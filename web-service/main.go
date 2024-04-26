package main

import (
	"fmt"
	"net/http"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/gin-gonic/gin"
)

const test_addr = "2426 ontario rd NW"

func geolocate(address string) (location string) {
	// geolocate user co-ords from address
	var geocoder = openstreetmap.Geocoder()
	coords, _ := geocoder.Geocode(test_addr)
	if (coords != nil) {
		fmt.Printf("%s location is (%.6f, %.6f)\n", test_addr, coords.Lat, coords.Lng)
	}
	
	coord := "test"
	return coord
}

func getTemp(c *gin.Context) {
	// var user_address = c.Param("address")
	temp := "temp_test"
	c.IndentedJSON(http.StatusOK, temp)
}

func main() {
	router := gin.Default()
	router.GET("/temp/:address", getTemp)
	geolocate(test_addr)

	// start gin server
	fmt.Println("Starting server on localhost:8080")
	router.Run("localhost:8080")
}