package sender

import (
	"fmt"
	"github.com/AlekseyPorandaykin/crypto_loader/pkg/kucoin/response"
)

type ExternalError struct {
	HttpCode      int
	ExchangeError response.ExchangeError

	Message string
	Err     error
}

func (e ExternalError) Error() string {
	errStr := e.Message
	if e.ExchangeError.Code != "" && e.ExchangeError.Message != "" {
		errStr += fmt.Sprintf("; %s: %s; ", e.ExchangeError.Code, e.ExchangeError.Message)
	}
	if e.Err != nil {
		errStr += fmt.Sprintf(" err=%s;", e.Err.Error())
	}
	return errStr
}

func (e ExternalError) Detail() string {
	return fmt.Sprintf(
		"%s (httpCode=%d; errorCode=%d (%s); msg=%s; err=%s)",
		e.Message,
		e.HttpCode,
		e.ExchangeError.Code,
		parseHttpErrorCode(e.HttpCode),
		e.ExchangeError.Message,
		e.Err.Error(),
	)
}

func parseHttpErrorCode(code int) string {
	switch code {
	case 400:
		return "Bad Request -- Invalid request format."
	case 401:
		return "Unauthorized -- Invalid API Key."
	case 403:
		return "Forbidden or Too Many Requests -- The request is forbidden or Access limit breached."
	case 404:
		return "Not Found -- The specified resource could not be found."
	case 405:
		return "Method Not Allowed -- You tried to access the resource with an invalid method."
	case 415:
		return "Unsupported Media Type. You need to use: application/json."
	case 500:
		return "Internal Server Error -- We had a problem with our server. Try again later."
	case 503:
		return "Service Unavailable -- We're temporarily offline for maintenance. Please try again later."
	default:
		return ""
	}
}
