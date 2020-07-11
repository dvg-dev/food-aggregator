package gethandlers

import (
	"encoding/json"
	"log"

	"github.com/dvg-dev/food-aggregator/helpers"
	"github.com/dvg-dev/food-aggregator/model"
	"github.com/labstack/echo"
)

//BuyByNameHandler returns the item by name
func (g *GetHandler) BuyByNameHandler(ctx echo.Context) error {
	//To get the path parameter
	itemName := ctx.Param("name")
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	//Get all Items
	items, err := helpers.GetItems()
	if err != nil {
		log.Println("Error in getting Items: ", err)
		return helpers.CommonResponseHandler(404, "NOT_FOUND", ctx)
	}

	item := findItembyName(items, itemName)
	if item == (model.Item{}) {
		return helpers.CommonResponseHandler(404, "NOT_FOUND", ctx)
	}

	ctx.Response().WriteHeader(200)
	return json.NewEncoder(ctx.Response()).Encode(item)
}

//findItembyName to find the item from the list of items fetched from the URLs
func findItembyName(items []model.Item, itemName string) model.Item {
	for _, item := range items {
		if itemName == item.Name {
			return item
		}
	}
	return model.Item{}
}
