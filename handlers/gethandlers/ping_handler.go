package gethandlers

import (
	"github.com/dvg-dev/food-aggregator/helpers"
	"github.com/labstack/echo"
)

//GetHandler is an empty struct for all methods of get handler
type GetHandler struct{}

//PingHandler is a Simple ping handler which returns 200 OK response
func (g *GetHandler) PingHandler(ctx echo.Context) error {
	//To set the response type
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return helpers.CommonResponseHandler(200, "OK", ctx)
}
