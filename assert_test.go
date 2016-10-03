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
	t.Run("same status", func(t *testing.T) {
		respAsserter.StatusIs(status1)
		errorer.AssertNoError()
	})
	t.Run("different status", func(t *testing.T) {
		respAsserter.StatusIs(status2)
		errorer.AssertError()
	})
	t.Run("fluent", func(t *testing.T) {
		assertFluent(t, respAsserter.StatusIs(status1))
	})
}

func TestBodyAsJsonIs(t *testing.T) {
	const (
		json1                = `{"foo":"bar"}`
		json1withExtraSpaces = `{ "foo" : "bar" }`
		json2                = `{"foo":"baz"}`
	)
	errorer := RegisterNewErrorer(t)
	resp.Body = bytes.NewBufferString(json1)
	t.Run("same JSON", func(t *testing.T) {
		respAsserter.BodyAsJsonIs(json1)
		errorer.AssertNoError()
	})
	t.Run("same JSON with extra spaces", func(t *testing.T) {
		respAsserter.BodyAsJsonIs(json1withExtraSpaces)
		errorer.AssertNoError()
	})
	t.Run("different JSON", func(t *testing.T) {
		respAsserter.BodyAsJsonIs(json2)
		errorer.AssertError()
	})
	t.Run("fluent", func(t *testing.T) {
		assertFluent(t, respAsserter.BodyAsJsonIs(json1))
	})
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

func (e *fakeErrorer) AssertError() {
	if e.error == false {
		e.t.Errorf("Want error")
	}
	e.error = false
}

func (e *fakeErrorer) AssertNoError() {
	if e.error == true {
		e.t.Errorf("Want no error")
	}
	e.error = false
}

func assertFluent(t *testing.T, gotAsserter RespAsserter) {
	if gotAsserter != respAsserter {
		t.Errorf("Method not fluent")
	}
}
