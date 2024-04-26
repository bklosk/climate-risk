package main

import (
	"fmt"
	"net/http"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/gin-gonic/gin"
)

// getCoords: address -> lat, lng
func getCoords(c *gin.Context) {
	// spin up openstreetmap geocoder
	var geocoder = openstreetmap.Geocoder()
	var address = c.Param("address")

	// geocode address
	coded_address, _ := geocoder.Geocode(address)
	if (coded_address != nil) {
		fmt.Printf("%s location is (%.4f, %.4f)\n", address, coded_address.Lat, coded_address.Lng)
		// return coords as JSON
		c.JSON(http.StatusOK, gin.H{"lat": coded_address.Lat, "lng": coded_address.Lng})
	} else {
		fmt.Printf("Could not find coordinates for %s\n", address)
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find coordinates for " + address})
	}
}

func main() {
	router := gin.Default()
	router.GET("/coords/:address", getCoords)

	// start gin server
	fmt.Println("\n\nStarting server on localhost:8080")
	router.Run("localhost:8080")
}