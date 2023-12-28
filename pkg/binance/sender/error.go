package sender

import "fmt"

var (
	UnauthorizedError   = NewError("Unauthorized")
	InternalError       = NewError("the issue is on Binance's side")
	RateLimitError      = NewError("breaking a request rate limit")
	WAFLimitErr         = NewError("WAF Limit (Web Application Firewall) has been violated")
	CancelReplaceErr    = NewError("CancelReplace order partially succeeds")
	IPAutoBannedErr     = NewError("IP has been auto-banned for continuing to send requests after receiving 429 codes")
	MalformedRequestErr = NewError("malformed requests; the issue is on the sender's side code")
)

type SenderError interface {
	error
	Detail() string
	WithExternalError(err ExternalError)
}

type ExternalError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type Error struct {
	message       string
	httpCode      int
	externalError ExternalError
	cause         error
}

func WrapErr(err error, message string) SenderError {
	return &Error{message: message, cause: err}
}

func NewError(message string) *Error {
	return &Error{message: message}
}

func (e *Error) WithHttpCode(code int) *Error {
	e.httpCode = code
	return e
}
func (e *Error) WithExternalError(err ExternalError) {
	e.externalError = err
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Detail() string {
	var detail string = e.message

	if e.httpCode != 0 {
		detail = fmt.Sprintf("%s, http_code=%d", detail, e.httpCode)
	}
	if e.externalError.Code != 0 || e.externalError.Message != "" {
		detail = fmt.Sprintf("%s, ExchangeCode=%d ExternalMessage=%s", detail, e.externalError.Code, e.externalError.Message)
	}
	if e.cause != nil {
		detail = fmt.Sprintf("%s, cause=%s", detail, e.cause.Error())
	}

	return detail
}
