package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/ttacon/libphonenumber"

	_ "egov/routing-number-info/docs"
)

type bank struct {
	RoutingNumber string `json:"routing_number"`
	Bank          string `json:"bank"`
	Address       string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	Zip           string `json:"zip"`
	Phone         string `json:"phone"`
	Message       string `json:"message"`
}

var banks = []bank{}

// GetBanks godoc
// @Summary List all banks
// @Description Get a list of all banks with their routing numbers and information
// @Tags banks
// @Produce json
// @Success 200 {array} bank
// @Router /banks [get]
func getBanks(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	c.IndentedJSON(http.StatusOK, banks)
}

// GetBankByRoutingNumber godoc
// @Summary Get bank by routing number
// @Description Get bank information for a specific routing number
// @Tags banks
// @Produce json
// @Param routing path string true "Routing Number"
// @Success 200 {object} bank
// @Failure 404 {object} map[string]string
// @Router /banks/{routing} [get]
func getBankByRoutingNumber(c *gin.Context) {
	routingNumber := c.Param("routing")
	for _, a := range banks {
		if a.RoutingNumber == routingNumber {
			num, _ := libphonenumber.Parse(string(a.Phone), "US")
			a.Phone = libphonenumber.Format(num, libphonenumber.NATIONAL)
			a.Message = "OK"
			c.Header("Access-Control-Allow-Origin", "*")
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "bank not found"})

}

// Health godoc
// @Summary Health check
// @Description Check if the API is running
// @Tags health
// @Produce json
// @Success 200 {string} string "up"
// @Router /health [get]
func health(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusOK, "up")
}

// @title Routing Number API
// @version 1.0
// @description API for looking up bank information by routing number
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	gin.SetMode(gin.ReleaseMode)
	var filename string

	if len(os.Args) < 2 {
		filename = "data/banks.json"
	} else {
		filename = os.Args[1]
	}

	file, _ := ioutil.ReadFile(filename)

	err := json.Unmarshal(file, &banks)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	router := gin.Default()
	router.GET("/health", health)
	router.GET("/banks", getBanks)
	router.GET("/banks/:routing", getBankByRoutingNumber)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
