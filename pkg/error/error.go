package custom_error

type NotFound struct {}

func (e *NotFound) Error() string {
	return "Not Found"
}