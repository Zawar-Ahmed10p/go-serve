package types

type User struct {
	//ID    int64  `sql:",PK"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}
