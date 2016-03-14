package echotest

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
	. "github.com/labstack/echo/test"
)

const (
	url    = "/foo/123"
	status = http.StatusAccepted
	body   = "bar"
)

func TestCall(t *testing.T) {
	req := NewRequest("GET", url, nil)
	e := echo.New()
	e.Any(url, echo.HandlerFunc(func(c echo.Context) error {
		if c.Request() != req {
			t.Errorf("Doesn't pass request to the handler")
		}
		return c.String(status, body)
	}))
	resp := Call(e, req)
	assertStatusAndBody(t, resp)
}

func TestGet(t *testing.T) {
	e := echo.New()
	e.Get(url, echo.HandlerFunc(func(c echo.Context) error {
		assertReqBody(t, c.Request(), nil)
		return c.String(status, body)
	}))
	resp := Get(e, url)
	assertStatusAndBody(t, resp)
}

func TestPostJson(t *testing.T) {
	e := echo.New()
	e.Post(url, echo.HandlerFunc(func(c echo.Context) error {
		assertReqBody(t, c.Request(), strings.NewReader(body))
		if ct := c.Request().Header().Get(ContentType); ct != ApplicationJSON {
			t.Errorf("Want %s '%s', got '%s'", ContentType, ApplicationJSON, ct)
		}
		return c.String(status, body)
	}))
	resp := PostJson(e, url, body)
	assertStatusAndBody(t, resp)
}

func TestAssertResp(t *testing.T) {
	e := echo.New()
	resp := Get(e, url)
	respAsserter := resp.AssertResp(t)
	if respAsserter.errorer != t {
		t.Errorf("Doesn't pass Errorer to asserter")
	}
	if respAsserter.resp != resp.ResponseRecorder {
		t.Errorf("Doesn't pass ResponseRecorder to asserter")
	}
}

func assertStatusAndBody(t *testing.T, resp Resp) {
	if resp.Status() != status {
		t.Errorf("Want status %d, got %d", status, resp.Status())
	}
	if resp.Body.String() != body {
		t.Errorf("Want body '%s', got '%s'", body, resp.Body.String())
	}
}

func assertReqBody(t *testing.T, r engine.Request, body io.Reader) {
	if r.Body() == nil {
		if body != nil {
			t.Errorf("Want not-nil body reader")
		}
	} else {
		if body == nil {
			t.Errorf("Want nil body reader")
		} else {
			gotBody, wantBody := bodyAsString(r.Body()), bodyAsString(body)
			if gotBody != wantBody {
				t.Errorf("Want body '%s', got '%s'", wantBody, gotBody)
			}
		}
	}
}

func bodyAsString(body io.Reader) string {
	bs, _ := ioutil.ReadAll(body)
	return string(bs)
}
