package main

import (
	"fmt"
	"time"

	"github.com/captain-corgi/golang-oracledb-example/pkg/oracle"
)

type Employees struct {
	EmployeeID int
	FirstName  string
	LastName   string
	Email      string
	Phone      string
	HireDate   time.Time
	ManagerID  int
	JobTitle   string
}

func main() {
	// Create connection
	db := oracle.NewGorm()

	// Read
	var employees []Employees
	db.Debug().
		Where("employee_id = ?", 100).
		Limit(5).
		Find(&employees)

	// Print
	for _, employee := range employees {
		fmt.Printf("%+v\n", employee)
	}
}
