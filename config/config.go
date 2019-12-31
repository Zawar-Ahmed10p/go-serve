package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// type PostgresDb struct {
// 	Db *gorm.DB
// }
//
func GetPostgres() *gorm.DB {
	Db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=***** dbname=******* password=******* sslmode=disable")
	if err != nil {
		panic(err)
	}
	//a.Db = Db
	return Db
}
