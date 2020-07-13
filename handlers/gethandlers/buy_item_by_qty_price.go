package gethandlers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/dvg-dev/food-aggregator/helpers"
	"github.com/dvg-dev/food-aggregator/model"
	"github.com/labstack/echo"
)

//BuyByNameQtyPrice to fetch item based on given name, quantity and price
func (g *GetHandler) BuyByNameQtyPrice(ctx echo.Context) error {
	itemName := ctx.Param("name")
	quantity, qtyErr := strconv.Atoi(ctx.Param("quantity"))
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	price, priceErr := strconv.ParseFloat(ctx.Param("price"), 64)
	if qtyErr != nil || priceErr != nil {
		return helpers.CommonResponseHandler(400, "Invalid Quantity or price", ctx)
	}

	//Check for the item from the in memory data structure
	item := findItembyQtyPrice(helpers.CachedItems, itemName, quantity, price)
	if item == (model.Item{}) {
		//Find the item from APIs if not found in memory
		items, err := helpers.GetItems()
		if err != nil {
			return helpers.CommonResponseHandler(404, "NOT_FOUND", ctx)
		}
		helpers.CachedItems = append(helpers.CachedItems, items...)
		item = findItembyQtyPrice(items, itemName, quantity, price)
		if item == (model.Item{}) {
			return helpers.CommonResponseHandler(404, "NOT_FOUND", ctx)
		}
		ctx.Response().WriteHeader(200)
		return json.NewEncoder(ctx.Response()).Encode(item)
	}
	ctx.Response().WriteHeader(200)
	return json.NewEncoder(ctx.Response()).Encode(item)
}

//findItembyQtyPrice to find the item based on name, quantity and price
func findItembyQtyPrice(items []model.Item, itemName string, quantity int, itemPrice float64) model.Item {
	for _, item := range items {
		price, _ := strconv.ParseFloat(item.Price[1:], 64)
		if strings.EqualFold(itemName, item.Name) && quantity <= item.Quantity && ((itemPrice / float64(quantity)) <= (price / float64(item.Quantity))) {
			return item
		}
	}
	return model.Item{}
}
