package models

import "gorm.io/gorm"

// Ovo je iz dokumentacije sa gorma i prikazuje praksu pisanja tj.
//Gorm ima svoj ugradjen model koji dolazi sa uintm, createdAt,updatedAt,deletedAt

//type User struct {
//	gorm.Model
//	Name string
//}
//// equals
//type User struct {
//	ID        uint           `gorm:"primaryKey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//	Name string
//}

type Post struct {
	gorm.Model
	Title string
	Body  string
}
