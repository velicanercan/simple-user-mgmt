package domain

type User struct {
	ID        int    `json:"id" gorm:"primaryKey" bson:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique" binding:"required" validate:"email"`
	BirthDate string `json:"birth_date" binding:"required"`
}
