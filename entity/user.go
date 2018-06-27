package entity

//entity
type User struct {
	ID     uint   `json:"id" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Age    uint   `json:"age" gorm:"column:age"`
	Gender string `json:"gender" gorm:"column:gender"`
	Email  string `json:"email" gorm:"column:email"`
}
