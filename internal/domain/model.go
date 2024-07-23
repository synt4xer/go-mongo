package domain

import "time"

type User struct {
	ID        string     `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string     `json:"name" validate:"required" bson:"name"`
	Email     string     `json:"email" validate:"required,email" bson:"email"`
	Password  string     `validate:"required" bson:"password"`
	CreatedAt *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	// Address   *Address   `json:"address" bson:"address"`
}

// type Address struct {
// 	Street  string `json:"street" bson:"street"`
// 	City    string `json:"city" bson:"city"`
// 	State   string `json:"state" bson:"state"`
// 	Country string `json:"country" bson:"country"`
// }
