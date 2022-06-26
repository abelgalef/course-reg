package utils

type Error struct {
	Code     int    `json:"code"`
	E        string `json:"e"`
	Overview string `json:"message"`
}

func NewError(code int, e error, Overview string) *Error {
	return &Error{code, e.Error(), Overview}
}
