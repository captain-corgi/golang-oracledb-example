package main

import (
	"database/sql"
	"fmt"

	"github.com/captain-corgi/golang-oracledb-example/pkg/oracle"
)

func main() {
	// Greeting
	fmt.Println("Hello, Oracle DB!")

	// Create connection
	conn := oracle.NewConnection()
	defer conn.Close()

	// Create statement
	stmt, err := conn.Prepare(`SELECT
									EMPLOYEE_ID,
									FIRST_NAME,
									LAST_NAME,
									EMAIL,
									PHONE,
									HIRE_DATE,
									MANAGER_ID,
									JOB_TITLE
								FROM
									EMPLOYEES`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Execute statement
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	// Iterate over results
	for rows.Next() {
		// Define columns
		var EMPLOYEE_ID,
			FIRST_NAME,
			LAST_NAME,
			EMAIL,
			PHONE,
			HIRE_DATE,
			MANAGER_ID,
			JOB_TITLE sql.NullString

		// Scan columns
		err = rows.Scan(&EMPLOYEE_ID,
			&FIRST_NAME,
			&LAST_NAME,
			&EMAIL,
			&PHONE,
			&HIRE_DATE,
			&MANAGER_ID,
			&JOB_TITLE)
		if err != nil {
			panic(err)
		}

		// Print results
		fmt.Println(EMPLOYEE_ID.String,
			FIRST_NAME.String,
			LAST_NAME.String,
			EMAIL.String,
			PHONE.String,
			HIRE_DATE.String,
			MANAGER_ID.String,
			JOB_TITLE.String)
	}

}
