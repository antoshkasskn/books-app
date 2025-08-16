package entity

type ErrNotFound struct{}

func (e ErrNotFound) Error() string {
	return "not found"
}
