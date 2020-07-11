package main

import (
	"log"

	"github.com/dvg-dev/food-aggregator/helpers"
	"github.com/dvg-dev/food-aggregator/receivers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	log.Println("Starting the Food Aggregator Application at port 8080...")
	helpers.InitURLs()
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))
	g := e.Group("/food-aggregator")
	g.GET("/ping", receivers.Get.PingHandler)
	g.GET("/buy-item/:name", receivers.Get.BuyByNameHandler)
	g.GET("/buy-item-qty/:name/:quantity", receivers.Get.BuyByNameQty)
	g.GET("/buy-item-qty-price/:name/:quantity/:price", receivers.Get.BuyByNameQtyPrice)
	g.GET("/show-summary", receivers.Get.ShowSummary)
	g.GET("/fast-buy-item/:name", receivers.Get.FastBuyItem)
	e.Start(":8080")
}
