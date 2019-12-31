package main

import (
	"github.com/Zawar-Ahmed10p/go_postgres/app"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// type User struct {
// 	Email string `gorm:"column:username"`
// }

func main() {

	app.Run()
}
