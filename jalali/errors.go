package jalali

import "fmt"

// ErrorNilReference is happening when a pointer is nil.
type ErrorNilReference struct{}

// ErrorInvalidYear is happening when year() passed is is in proper range.
type ErrorInvalidYear struct {
	year int
}

func (e *ErrorNilReference) Error() string {
	return "jalali: reference is nil"
}

func (e *ErrorInvalidYear) Error() string {
	return fmt.Sprintf("jalali: %v is invalid year()", e.year)
}
