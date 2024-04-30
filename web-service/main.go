package main

import (
	"fmt"
	"os"

	"database/sql"
	"net/http"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// connection string variables
var (
	host = os.Getenv("CLIMATE_DB_HOST")
	port = 5432
	user = os.Getenv("CLIMATE_DB_USER")
	password = os.Getenv("CLIMATE_DB_PASSWORD")
	dbname = "climate_data"
)

// getCoords: address -> lat, lng
func getCoords(c *gin.Context){
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

//wrapper for application-scoped DB connection
func get_vector_tile(db *sql.DB) gin.HandlerFunc {
    fn := func(c * gin.Context){
		rows, err := db.Query(`SELECT ST_SkewX(rast) from "netcdf:test_data";`)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{"error": "error"})
		} else {
			fmt.Println(rows)
			c.JSON(http.StatusOK, gin.H{"rows": rows})
		}
	} 
	return gin.HandlerFunc(fn)
}

func main() {
	// connect to db
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/coords/:address", getCoords)
	router.GET("/tile/", get_vector_tile(db))


	// start gin server
	fmt.Println("\n\nStarting server on localhost:8080")
	router.Run("localhost:8080")
}