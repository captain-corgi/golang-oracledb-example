package oracle

import (
	"database/sql"

	_ "github.com/sijms/go-ora/v2"
)

func NewConnection() {
	db, err := sql.Open("oracle", "oracle://OT:yourpassword@localhost/ORCLPDB1.localdomain")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
