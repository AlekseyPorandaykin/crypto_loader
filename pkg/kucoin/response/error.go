package response

type Error interface {
	error
	Detail() string
}

type ExchangeError struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}
