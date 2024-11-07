package main

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"math/rand"
	"time"
)

type Dish struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func getDishNames() []string {
	rawJson, err := ioutil.ReadFile("data/dish_name.json")
	if err != nil {
		panic(err)
	}

	var dishes []string

	err = json.Unmarshal(rawJson, &dishes)
	if err != nil {
		panic(err)
	}

	return dishes
}

func getDishDescriptions() []string {
	rawJson, err := ioutil.ReadFile("data/dish_description.json")
	if err != nil {
		panic(err)
	}

	var dishesDesc []string

	err = json.Unmarshal(rawJson, &dishesDesc)
	if err != nil {
		panic(err)
	}

	return dishesDesc
}

func getRandomValueFromArray(array []string) string {
	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(array))

	return array[randomIndex]
}

func getRandomPrice() float64 {
	rand.Seed(time.Now().UnixNano())

	randomPrice := rand.Float64() * 20

	return math.Round(randomPrice*100) / 100
}

func saveToFile(data []byte) {
	err := ioutil.WriteFile("response.json", data, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	var dishes [1000]Dish

	dishNames := getDishNames()
	dishDescriptions := getDishDescriptions()

	for i := 0; i < len(dishes); i++ {
		dishes[i] = Dish{
			Id:          i,
			Name:        getRandomValueFromArray(dishNames),
			Description: getRandomValueFromArray(dishDescriptions),
			Price:       getRandomPrice(),
		}
	}

	dishesJson, err := json.Marshal(dishes)
	if err != nil {
		panic(err)
	}

	saveToFile(dishesJson)
}
