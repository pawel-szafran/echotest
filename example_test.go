package echotest

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
)

func Example() {
	_ = func(t *testing.T) {
		e := echo.New()
		// registerYourEndpoint(e)
		Get(e, "/card/4111111111111111/balance").AssertResp(t).
			StatusIs(http.StatusOK).
			BodyAsJsonIs(`
				{
				  "amount" : 10000,
				  "ccy"    : "USD"
				}`)
	}
}
