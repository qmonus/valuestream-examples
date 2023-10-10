package schema

type Todo struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	Comment string `json:"comment" binding:"required"`
	State   int    `json:"state"`
}
