package echotest

import (
	"bytes"
	"net/http"
	"testing"

	. "github.com/labstack/echo/test"
)

var (
	resp         *ResponseRecorder
	respAsserter RespAsserter
)

func init() {
	resp = NewResponseRecorder()
	respAsserter = RespAsserter{resp: resp}
}

func TestStatusIs(t *testing.T) {
	const (
		status1 = http.StatusAccepted
		status2 = http.StatusNotFound
	)
	errorer := RegisterNewErrorer(t)
	resp.WriteHeader(status1)
	respAsserter.StatusIs(status1)
	errorer.AssertError(false, "same status")
	respAsserter.StatusIs(status2)
	errorer.AssertError(true, "different status")
	assertFluent(t, respAsserter.StatusIs(status1))
}

func TestBodyAsJsonIs(t *testing.T) {
	const (
		json1                = `{"foo":"bar"}`
		json1withExtraSpaces = `{ "foo" : "bar" }`
		json2                = `{"foo":"baz"}`
	)
	errorer := RegisterNewErrorer(t)
	resp.Body = bytes.NewBufferString(json1)
	respAsserter.BodyAsJsonIs(json1)
	errorer.AssertError(false, "same JSON")
	respAsserter.BodyAsJsonIs(json1withExtraSpaces)
	errorer.AssertError(false, "same JSON with extra spaces")
	respAsserter.BodyAsJsonIs(json2)
	errorer.AssertError(true, "different JSON")
	assertFluent(t, respAsserter.BodyAsJsonIs(json1))
}

func RegisterNewErrorer(t *testing.T) (errorer *fakeErrorer) {
	defer func() { respAsserter.errorer = errorer }()
	return &fakeErrorer{t: t}
}

type fakeErrorer struct {
	t     *testing.T
	error bool
}

func (e *fakeErrorer) Errorf(format string, args ...interface{}) {
	e.error = true
}

func (e *fakeErrorer) AssertError(error bool, comment string) {
	if e.error != error {
		e.t.Errorf("Want error=%v for %s", error, comment)
	}
	e.error = false
}

func assertFluent(t *testing.T, gotAsserter RespAsserter) {
	if gotAsserter != respAsserter {
		t.Errorf("Method not fluent")
	}
}
