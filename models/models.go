package models 

// import "gorm.io/gorm"

type Player struct{
	ID 		uint            `gorm:"primarykey"`
	Name	string			`gorm:"<-"` 
	Country string			`gorm:"<-:create"`
	Score	uint			`gorm:"<-"` 
}