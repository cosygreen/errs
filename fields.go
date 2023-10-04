package errs

// WithFields adds additional fields to an error that will be added to the log statement.
func WithFields(err error, fields map[string]interface{}) error {
	if err == nil {
		return nil
	}

	return fieldsError{
		err:    withStack(err),
		fields: fields,
	}
}

// FieldsError defines an interface for errors that contain a map of fields.
type FieldsError interface {
	GetFields() map[string]interface{}
}

type fieldsError struct {
	err error

	fields map[string]interface{}
}

func (s fieldsError) Error() string {
	return s.err.Error()
}

func (s fieldsError) Unwrap() error {
	return s.err
}

func (s fieldsError) GetFields() map[string]interface{} {
	return s.fields
}

func (s fieldsError) HasErrInfo() {}
