package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ttacon/libphonenumber"
)

type bank struct {
	RoutingNumber string `json:"routing_number"`
	Bank          string `json:"bank"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
	Phone         string `json:"phone"`
}

var banks = []bank{}

func getBanks(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, banks)
}

func getBankByRoutingNumber(c *gin.Context) {
	routingNumber := c.Param("routing")
	for _, a := range banks {
		if a.RoutingNumber == routingNumber {
			num, _ := libphonenumber.Parse(string(a.Phone), "US")
			a.Phone = libphonenumber.Format(num, libphonenumber.NATIONAL)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bank not found"})

}

func main() {

	file, _ := ioutil.ReadFile("data/banks.json")

	err := json.Unmarshal(file, &banks)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	router := gin.Default()
	router.GET("/banks", getBanks)
	router.GET("/banks/:routing", getBankByRoutingNumber)

	router.Run("localhost:8080")
}
