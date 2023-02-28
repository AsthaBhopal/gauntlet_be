package globals

type HttpError struct {
	HttpErrorCode int
	Msg           string
	Err           error
}

func NewHttpError(httpCode int, msg string, err error) *HttpError {
	return &HttpError{
		HttpErrorCode: httpCode,
		Msg:           msg,
		Err:           err,
	}
}
