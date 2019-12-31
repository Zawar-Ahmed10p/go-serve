package repositories

import (
	"fmt"

	"github.com/Zawar-Ahmed10p/go_postgres/config"
	"github.com/Zawar-Ahmed10p/go_postgres/types"

	//justify

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//
// const (
// 	db = config.PostgresDb.GetPostgres
// )

func AddUser(user types.User) {
	db := config.GetPostgres()
	fmt.Println("--------")
	fmt.Println(user)
	db.Create(user)
	//db.Save(user)
	db.Close()
}
