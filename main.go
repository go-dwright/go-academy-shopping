package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-academy-shopping/internal/customer"
)

func main() {

	// create 1st customer
	cust := customer.New()
	cust.SetEmail("damon@wright.com")
	cust.SetFirstName("damon")
	cust.SetLastName("wright")

	// create 2nd customer
	cust2 := customer.New()
	cust2.SetEmail("joe@bloggs.com")
	cust2.SetFirstName("joe")
	cust2.SetLastName("bloggs")

	router := gin.Default()
	router.GET("/test", test)
	router.GET("/customers/all", GetAllCustomers)
	router.POST("/customer", createCustomer)

	router.Run("localhost:8080")
}

func GetAllCustomers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, customer.GetAll())
}

func createCustomer(c *gin.Context) {
}

func test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is a test")
}
