package helpers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/dvg-dev/food-aggregator/model"
)

var urls model.URLs

//InitURLs to load the URLs from the json file
func InitURLs() {
	urlFile, err := os.Open("urls.json")
	if err != nil {
		log.Fatal("URL file not found")
	}
	defer urlFile.Close()

	urlByteValue, _ := ioutil.ReadAll(urlFile)
	json.Unmarshal([]byte(urlByteValue), &urls)
}

//GetItemsParallelly fetches all items from different urls concurrently without waiting for any URL
func GetItemsParallelly() ([]model.Item, error) {
	var wg sync.WaitGroup
	var items []model.Item
	var genItem model.Item
	var fruitItems []model.FruitItem
	var vegItems []model.VegItem
	var grainItems []model.GrainItem
	var getFruits, getVegetables, getGrains []byte
	var getFruitsErr, getVegErr, getGrainErr error

	//Using goroutines with waitgroups for fetching data parallelly
	wg.Add(3)
	go func(getFruits *[]byte, getFruitsErr *error, wg *sync.WaitGroup) {
		defer wg.Done()
		*getFruits, *getFruitsErr = FetchItemsFromURL(urls.FruitsURL)
	}(&getFruits, &getFruitsErr, &wg)
	go func(getVegetables *[]byte, getVegErr *error, wg *sync.WaitGroup) {
		defer wg.Done()
		*getVegetables, *getVegErr = FetchItemsFromURL(urls.VegetablesURL)
	}(&getVegetables, &getVegErr, &wg)
	go func(getGrains *[]byte, getGrainErr *error, wg *sync.WaitGroup) {
		defer wg.Done()
		*getGrains, *getGrainErr = FetchItemsFromURL(urls.GrainsURL)
	}(&getGrains, &getGrainErr, &wg)
	wg.Wait()

	if getGrainErr != nil || getVegErr != nil || getFruitsErr != nil {
		return nil, errors.New("Unable to get data from API")
	}

	json.Unmarshal(getFruits, &fruitItems)
	json.Unmarshal(getVegetables, &vegItems)
	json.Unmarshal(getGrains, &grainItems)

	for _, item := range fruitItems {
		genItem = model.Item(item)
		items = append(items, genItem)
	}
	for _, item := range vegItems {
		genItem = model.Item(item)
		items = append(items, genItem)
	}
	for _, item := range grainItems {
		genItem = model.Item(item)
		items = append(items, genItem)
	}
	return items, nil
}

//GetItems to fetch all items sequentially
func GetItems() ([]model.Item, error) {
	var items []model.Item
	var genItem model.Item
	var fruitItems []model.FruitItem
	var vegItems []model.VegItem
	var grainItems []model.GrainItem
	getFruits, getFruitsErr := FetchItemsFromURL(urls.FruitsURL)
	getVegetables, getVegErr := FetchItemsFromURL(urls.VegetablesURL)
	getGrains, getGrainErr := FetchItemsFromURL(urls.GrainsURL)
	if getGrainErr != nil || getVegErr != nil || getFruitsErr != nil {
		return nil, errors.New("Unable to get data from API")
	}
	json.Unmarshal(getFruits, &fruitItems)
	json.Unmarshal(getVegetables, &vegItems)
	json.Unmarshal(getGrains, &grainItems)
	for _, item := range fruitItems {
		genItem = model.Item(item)
		items = append(items, genItem)
	}
	for _, item := range vegItems {
		genItem = model.Item(item)
		items = append(items, genItem)
	}
	for _, item := range grainItems {
		genItem = model.Item(item)
		items = append(items, genItem)
	}
	return items, nil
}

//FetchItemsFromURL to fetch items from the URL specified
func FetchItemsFromURL(url string) ([]byte, error) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("Error in fetching items: ", err)
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error in fetching items: ", err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error in fetching items: ", err)
		return nil, err
	}
	return body, nil
}
