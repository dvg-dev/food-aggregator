package gethandlers

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/dvg-dev/food-aggregator/helpers"
	"github.com/dvg-dev/food-aggregator/model"
	"github.com/labstack/echo"
)

//BuyByNameQty to fetch item by name and quantity
func (g *GetHandler) BuyByNameQty(ctx echo.Context) error {
	itemName := ctx.Param("name")
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	quantity, qtyErr := strconv.Atoi(ctx.Param("quantity"))
	if qtyErr != nil {
		return helpers.CommonResponseHandler(400, "Invalid Quantity", ctx)
	}

	items, err := helpers.GetItems()
	if err != nil {
		log.Println("Error in getting Items: ", err)
		return helpers.CommonResponseHandler(404, "NOT_FOUND", ctx)
	}

	item := findItembyQty(items, itemName, quantity)
	if item == (model.Item{}) {
		return helpers.CommonResponseHandler(404, "NOT_FOUND", ctx)
	}
	ctx.Response().WriteHeader(200)
	return json.NewEncoder(ctx.Response()).Encode(item)
}

//findItembyQty to find the item from the list based on name and quantity
func findItembyQty(items []model.Item, name string, quantity int) model.Item {
	for _, item := range items {
		if strings.EqualFold(name, item.Name) && quantity <= item.Quantity {
			return item
		}
	}
	return model.Item{}
}
