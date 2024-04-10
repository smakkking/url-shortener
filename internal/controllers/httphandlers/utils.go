package httphandlers

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Url    string `json:"url,omitempty"`
}

type Request struct {
	URL string `json:"url"`
}

func OK(url string) Response {
	return Response{
		Status: StatusOK,
		Url:    url,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}
