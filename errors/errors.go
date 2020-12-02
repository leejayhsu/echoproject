package errors

// BError implements Error
type BError struct {
	Message string
	Code    int
	Status  string
	Kind    string
	Extra   map[string](string) `json:"-"`
}

func (e *BError) Error() string {
	return e.Message
}
