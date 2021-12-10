package oracle

import (
	"database/sql"
)

func NewConnection() *sql.DB {
	conn, err := sql.Open("oracle", "oracle://OT:yourpassword@localhost/ORCLPDB1.localdomain")
	if err != nil {
		panic(err)
	}

	return conn
}
