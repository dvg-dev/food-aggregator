package helpers

import (
	"encoding/json"

	"github.com/dvg-dev/food-aggregator/model"
	"github.com/labstack/echo"
)

//CommonResponseHandler to return common responses based on http status code and message
func CommonResponseHandler(code int, message string, ctx echo.Context) error {
	response := ResponseMapper(code, message)
	ctx.Response().WriteHeader(code)
	return json.NewEncoder(ctx.Response()).Encode(response)
}

//ResponseMapper maps http code and message to response struct
func ResponseMapper(code int, message string) model.APIResponse {
	response := model.APIResponse{
		Code:    code,
		Message: message,
	}
	return response
}
