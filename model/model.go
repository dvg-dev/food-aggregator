package model

//APIResponse to encode http status code and message
type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//FruitItem struct to store fruit items
type FruitItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	QtyPrice
}

//VegItem struct to store vegetable items
type VegItem struct {
	ID   string `json:"productId"`
	Name string `json:"productName"`
	QtyPrice
}

//GrainItem struct to store Grain items
type GrainItem struct {
	ID   string `json:"itemId"`
	Name string `json:"itemName"`
	QtyPrice
}

//Item is a general structure for items
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	QtyPrice
}

//QtyPrice is a part of composite struct that holds common data for all the above items
type QtyPrice struct {
	Quantity int    `json:"quantity"`
	Price    string `json:"price"`
}

//URLs to store the URLs of suppliers
type URLs struct {
	FruitsURL     string `json:"fruitsURL"`
	VegetablesURL string `json:"vegetablesURL"`
	GrainsURL     string `json:"grainsURL"`
}
