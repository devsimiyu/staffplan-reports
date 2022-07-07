package errors

type NotFoundError string

func (n NotFoundError) Error() string {
	return string(n)
}
