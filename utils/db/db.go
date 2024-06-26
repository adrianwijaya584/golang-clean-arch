package db

import (
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"clean_arch/domain"
	"clean_arch/utils/env"
)

func DbConn() *gorm.DB {
	conn := mysql.Open(os.Getenv(env.MYSQL_URL))

	db, err := gorm.Open(conn, &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatal(err.Error())

		return nil
	}

	if os.Getenv(env.GO_ENV) == "dev" {
		db = db.Debug()
	}

	return db
}

func DbMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)

		return nil, nil
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		CreateBatchSize: 3,
	})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

func DbInit() {
	conn := DbConn()

	conn.AutoMigrate(domain.Category{})
	conn.AutoMigrate(domain.User{})
	conn.AutoMigrate(domain.Cat{})
}
