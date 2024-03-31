package users

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
