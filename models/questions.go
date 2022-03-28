package models

type Question struct {
	ID       int    `json:"id" gorm:"primary_key;auto_increment"`
	Question string `json:"question"`
	ChoiceA  string `json:"choiceA"`
	ChoiceB  string `json:"choiceB"`
	ChoiceC  string `json:"choiceC"`
	ChoiceD  string `json:"choiceD"`
	Answer   string `json:"answer"`
}
