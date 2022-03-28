package controllers

import (
	"encoding/json"
	"main/connections"
	"main/models"
	"net/http"
)

type BodyData struct {
	UserId  int             `json:"userId"`
	Answers []models.Answer `json:"answers"`
}

type ResponseItem struct {
	Questions  []QuestionItem
	TotalPoint int
	TotalRank  int
}

type QuestionItem struct {
	QuestionId int  `json:"questionId"`
	IsSuccess  bool `json:"isSuccess"`
}

var checkToken bool

func GetResult(writer http.ResponseWriter, request *http.Request) {
	bodyData := BodyData{}
	responseData := ResponseItem{}
	checkToken = false

	json.NewDecoder(request.Body).Decode(&bodyData)

	responseData, _ = getQuestions(writer, bodyData)

	if checkToken == true {
		json, _ := json.Marshal(responseData)
		connections.SendReponse(writer, http.StatusOK, json)
	}

}

func getQuestions(writer http.ResponseWriter, bodyData BodyData) (ResponseItem, error) {
	user := models.User{}
	questionItem := QuestionItem{}
	responseData := ResponseItem{}

	db := connections.GetConnection()
	defer db.Close()

	userResult := db.First(&user, "id = ?", bodyData.UserId)
	if userResult.RowsAffected > 0 {
		if user.Token != "" {
			for _, element := range bodyData.Answers {
				question := models.Question{}

				result := db.First(&element, "(question_id,user_id) = (?,?)", element.QuestionId, bodyData.UserId)
				questionResult := db.First(&question, "id = ?", element.QuestionId)

				if questionResult.RowsAffected > 0 {
					if result.RowsAffected > 0 {
						Message := []byte(`{"Error": "User already solved the question"}`)
						connections.ErrorMsg(writer, http.StatusBadRequest, Message)
					} else {
						checkToken = true
						element.UserId = bodyData.UserId

						db.Create(&element)

						if question.Answer == element.Chosen {
							questionItem.IsSuccess = true
							responseData.TotalPoint += 1
							db.Model(&user).Where("id = ?", bodyData.UserId).Update("score", responseData.TotalPoint)
						} else {
							questionItem.IsSuccess = false
						}

						questionItem.QuestionId = question.ID
						responseData.Questions = append(responseData.Questions, questionItem)
					}
				}
			}
		} else {
			Message := []byte(`{"Error": "Please login first "}`)
			connections.ErrorMsg(writer, http.StatusUnauthorized, Message)
			checkToken = false
		}
	}

	responseData.TotalRank = getTotalRank(bodyData)

	return responseData, nil
}

func getTotalRank(bodyData BodyData) int {
	var rank int
	var users []models.User
	db := connections.GetConnection()
	defer db.Close()
	result := db.Order("score desc").Where("score > ?", 0).Find(&users)
	if result.RowsAffected > 0 {
		for index, element := range users {
			if element.ID == bodyData.UserId {
				rank = 100 - (100 / len(users) * (index + 1))
			}
		}
	}
	return rank
}
