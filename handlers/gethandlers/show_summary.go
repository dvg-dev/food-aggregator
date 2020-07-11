package gethandlers

import (
	"encoding/json"

	"github.com/dvg-dev/food-aggregator/helpers"
	"github.com/labstack/echo"
)

//ShowSummary returns all the items inside the in memory data cache
func (g *GetHandler) ShowSummary(ctx echo.Context) error {
	ctx.Response().WriteHeader(200)
	return json.NewEncoder(ctx.Response()).Encode(helpers.CachedItems)
}
