package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// employee represents data about a an employee record.
type employee struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Department string `json:"department"`
	Gendor     string `json:"gendor"`
}

// employees slice to seed record album data.
var employees = []employee{
	{ID: "P0P1", Role: "Associate", Name: "Aakash", Department: "R&D",Gendor: "Male"},
	{ID: "P0P2", Role: "Consultant", Name: "Divit", Department: "R&D", Gendor: "Male"},
	{ID: "P0P3", Role: "Manager", Name: "Dhriti", Department: "R&D", Gendor: "Female"},
}

// getEmployees responds with the list of all albums as JSON.
func GetEmployees(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, employees)
}

// postEmployees adds an album from JSON received in the request body.
func CreateEmployees(c *gin.Context) {
	var newemployee employee

	// Call BindJSON to bind the received JSON to
	// newEmployee.
	if err := c.BindJSON(&newemployee); err != nil {
		return
	}

	// Add the new employee to the slice.
	employees = append(employees,  newemployee)
	c.IndentedJSON(http.StatusCreated,  newemployee)
}

// getEmployeeByID locates the employee whose ID value matches the id
// parameter sent by the client, then returns that employee as a response.
func GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of employees, looking for
	// an employee whose ID value matches the parameter.
	for _, a := range employees {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee not found"})
}

// Delete a particular employee record by ID
func DeleteEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of employees, looking for
	// an employee whose ID value matches the parameter.
	for _, a := range employees {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "employee records deleted"})
}
