package models

type Answer struct {
	ID         int    `json:"id" gorm:"primary_key;auto_increment"`
	QuestionId int    `json:"questionId"`
	UserId     int    `json:"userId"`
	Chosen     string `json:"chosen"`
}
