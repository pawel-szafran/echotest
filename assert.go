package echotest

import (
	"encoding/json"
	"reflect"

	. "github.com/labstack/echo/test"
)

type errorer interface {
	Errorf(format string, args ...interface{})
}

type RespAsserter struct {
	errorer errorer
	resp    *ResponseRecorder
}

func (r RespAsserter) StatusIs(status int) RespAsserter {
	if r.resp.Status() != status {
		r.errorer.Errorf("Want status %d, got %d", status, r.resp.Status())
	}
	return r
}

func (r RespAsserter) BodyAsJsonIs(body string) RespAsserter {
	gotBody := r.resp.Body.String()
	if !reflect.DeepEqual(asJson(gotBody), asJson(body)) {
		r.errorer.Errorf("Want JSON %q, got %q", body, gotBody)
	}
	return r
}

func asJson(text string) (jsonMap interface{}) {
	json.Unmarshal([]byte(text), &jsonMap)
	return
}
