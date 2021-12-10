package oracle

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGorm() *gorm.DB {
	conn, err := sql.Open("oracle", "oracle://OT:yourpassword@localhost/ORCLPDB1.localdomain")
	if err != nil {
		panic(err)
	}
	config := Config{
		DriverName: "ora",
		Conn:       conn,
	}
	dialector := New(config)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // Slow SQL threshold
			LogLevel:      logger.Error, // Log level
			Colorful:      false,        // Disable color
		},
	)
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
