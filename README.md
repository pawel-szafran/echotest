# echotest - fluent tests for RESTful APIs built with [Echo](https://github.com/labstack/echo)

[![wercker status](https://app.wercker.com/status/4634f21a9a01c7adac6e9d70ef4607d1/s/master "wercker status")](https://app.wercker.com/project/byKey/4634f21a9a01c7adac6e9d70ef4607d1)

### Example

```go
import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	. "github.com/pawel-szafran/echotest"
)

func TestYourEndpoint(t *testing.T) {
	e := registerYourEndpoint(echo.New())
	Get(e, "/card/4111111111111111/balance").AssertResp(t).
		StatusIs(http.StatusOK).
		BodyAsJsonIs(`
			{
			  "amount" : 10000,
			  "ccy"    : "USD"
			}`)
}
```
