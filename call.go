package echotest

import (
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	. "github.com/labstack/echo/test"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
)

type Resp struct {
	*ResponseRecorder
}

func Call(e *echo.Echo, req engine.Request) Resp {
	resp := NewResponseRecorder()
	e.ServeHTTP(req, resp)
	return Resp{resp}
}

func Get(e *echo.Echo, url string) Resp {
	req := NewRequest("GET", url, nil)
	return Call(e, req)
}

func PostJson(e *echo.Echo, url, body string) Resp {
	req := NewRequest("POST", url, strings.NewReader(body))
	req.Header().Set(contentType, applicationJSON)
	return Call(e, req)
}

func (r Resp) AssertResp(t *testing.T) RespAsserter {
	return RespAsserter{t, r.ResponseRecorder}
}
