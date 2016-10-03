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
	e.Any(url, func(c echo.Context) error {
		if c.Request() != req {
			t.Errorf("Doesn't pass request to the handler")
		}
		return c.String(status, body)
	})
	resp := Call(e, req)
	assertStatusAndBody(t, resp)
}

func TestGet(t *testing.T) {
	e := echo.New()
	e.Get(url, func(c echo.Context) error {
		assertReqBody(t, c.Request(), nil)
		return c.String(status, body)
	})
	resp := Get(e, url)
	assertStatusAndBody(t, resp)
}

func TestPostJson(t *testing.T) {
	e := echo.New()
	e.Post(url, func(c echo.Context) error {
		assertReqBody(t, c.Request(), strings.NewReader(body))
		if ct := c.Request().Header().Get(contentType); ct != applicationJSON {
			t.Errorf("Want %s %q, got %q", contentType, applicationJSON, ct)
		}
		return c.String(status, body)
	})
	resp := PostJson(e, url, body)
	assertStatusAndBody(t, resp)
}

func TestAssertResp(t *testing.T) {
	e := echo.New()
	resp := Get(e, url)
	gotAsserter := resp.AssertResp(t)
	wantAsserter := RespAsserter{t, resp.ResponseRecorder}
	if gotAsserter != wantAsserter {
		t.Errorf("Want asserter %+v, got %+v", wantAsserter, gotAsserter)
	}
}

func assertStatusAndBody(t *testing.T, resp Resp) {
	if resp.Status() != status {
		t.Errorf("Want status %d, got %d", status, resp.Status())
	}
	if resp.Body.String() != body {
		t.Errorf("Want body %q, got %q", body, resp.Body.String())
	}
}

func assertReqBody(t *testing.T, r engine.Request, body io.Reader) {
	switch {
	case body != nil && r.Body() == nil:
		t.Errorf("Want non-nil body reader")
	case body == nil && r.Body() != nil:
		t.Errorf("Want nil body reader")
	case body != nil && r.Body() != nil:
		gotBody, wantBody := bodyAsString(r.Body()), bodyAsString(body)
		if gotBody != wantBody {
			t.Errorf("Want body %q, got %q", wantBody, gotBody)
		}
	}
}

func bodyAsString(body io.Reader) string {
	bs, _ := ioutil.ReadAll(body)
	return string(bs)
}
