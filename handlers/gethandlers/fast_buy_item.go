package gethandlers

import (
	"encoding/json"

	"github.com/dvg-dev/food-aggregator/helpers"
	"github.com/dvg-dev/food-aggregator/model"
	"github.com/labstack/echo"
)

//FastBuyItem to fetch item by making parallel API calls
func (g *GetHandler) FastBuyItem(ctx echo.Context) error {
	itemName := ctx.Param("name")
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	items, err := helpers.GetItemsParallelly()
	if err != nil {
		return helpers.CommonResponseHandler(404, "NOT_FOUND", ctx)
	}

	item := findItembyName(items, itemName)
	if item == (model.Item{}) {
		return helpers.CommonResponseHandler(404, "NOT_FOUND", ctx)
	}

	ctx.Response().WriteHeader(200)
	return json.NewEncoder(ctx.Response()).Encode(item)
}
