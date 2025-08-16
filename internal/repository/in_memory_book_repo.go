package repository

import (
	"book-app/internal/entity"
	"book-app/pkg/logger"
	"context"
	"github.com/jackc/pgx"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	workerSize      = 10
	durationRestore = 5 * time.Minute
)

type inMemoryBookRepository struct {
	storage *sync.Map
	db      *pgx.ConnPool
}

func NewBookRepo(db *pgx.ConnPool) entity.BookRepo {
	r := &inMemoryBookRepository{
		storage: new(sync.Map),
		db:      db,
	}
	if err := r.storeFromDisk(); err != nil {
		log.Fatal(err)
	}
	r.taskWriteOnDisk(durationRestore)
	return r
}

func (r *inMemoryBookRepository) Create(ctx context.Context, book *entity.Book) error {
	r.storage.Store(book.Id, book)
	return nil
}

func (r *inMemoryBookRepository) GetById(ctx context.Context, s string) (*entity.Book, error) {
	book, ok := r.storage.Load(s)
	if !ok {
		return nil, entity.ErrNotFound{}
	}
	return book.(*entity.Book), nil
}

func (r *inMemoryBookRepository) GetList(ctx context.Context) (books []*entity.Book, totalCount int, err error) {
	r.storage.Range(func(k, v interface{}) bool {
		books = append(books, v.(*entity.Book))
		totalCount++
		return true
	})
	return
}

func (r *inMemoryBookRepository) DeleteById(ctx context.Context, s string) error {
	r.storage.Delete(s)
	return nil
}

func (r *inMemoryBookRepository) storeFromDisk() error {
	var (
		query = `select * from books`
		wg    = new(sync.WaitGroup)
		books = make(chan *entity.Book)
	)

	rows, err := r.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	wg.Add(workerSize)
	for range workerSize {
		go func() {
			defer wg.Done()

			for book := range books {
				r.storage.Store(book.Id, book)
			}
		}()
	}
	for rows.Next() {
		var book entity.Book
		if err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.PublicationDate, &book.Publisher, &book.Edition, &book.Location); err != nil {
			logger.Printf("Can't scan book: %v \n", err)
			continue
		}
		books <- &book
	}
	close(books)
	wg.Wait()
	return nil
}

func (r *inMemoryBookRepository) taskWriteOnDisk(duration time.Duration) {
	ticker := time.NewTicker(duration)
	notify := make(chan os.Signal, 1)

	signal.Notify(notify, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				logger.Println("Tick-tack - it's time to store on disk!!!")
				r.storeOnDisk()
			case <-notify:
				logger.Println("Oh no, you killed me. Enable evacuate data!")
				r.storeOnDisk()
				return
			}
		}
	}()
}

func (r *inMemoryBookRepository) storeOnDisk() {
	logger.Println("Start storeOnDisk")
	now := time.Now()
	booksChan := make(chan *entity.Book)
	wg := &sync.WaitGroup{}

	wg.Add(workerSize)
	for range workerSize {
		go func() {
			defer wg.Done()

			for book := range booksChan {
				query := `INSERT INTO books(id, title, author, publication_date, publisher, edition, location)` +
					` VALUES ($1, $2, $3, $4, $5, $6, $7)` +
					` ON CONFLICT (id) ` +
					` DO UPDATE SET 
  						  title = EXCLUDED.title,
						  author = EXCLUDED.author,
    					  publication_date = EXCLUDED.publication_date,
   						  publisher = EXCLUDED.publisher,
						  edition = EXCLUDED.edition,
     					  location = EXCLUDED.location;`
				_, err := r.db.Exec(query, book.Id, book.Title, book.Author, book.PublicationDate, book.Publisher, book.Edition, book.Location)
				if err != nil {
					logger.Printf("Ошибка при сохранении книги на диск: %v \n", err)
				}
			}
		}()
	}

	r.storage.Range(func(k, v interface{}) bool {
		book, ok := v.(*entity.Book)
		if ok {
			booksChan <- book
			return true
		}
		return false
	})
	close(booksChan)
	wg.Wait()
	logger.Printf("End storeOnDisk by %f seconds \n", time.Since(now).Seconds())
	return
}
