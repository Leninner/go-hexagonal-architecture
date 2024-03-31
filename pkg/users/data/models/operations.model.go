package users

type Operation struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ModuleID    int    `json:"module_id"`
}
