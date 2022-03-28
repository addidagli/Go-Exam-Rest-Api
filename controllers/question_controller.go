package controllers

import (
	"encoding/json"
	"log"
	"main/connections"
	"main/models"
	"net/http"

	"github.com/gorilla/mux"
)

func AddQuestion(writer http.ResponseWriter, request *http.Request) {
	question := models.Question{}

	db := connections.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&question)

	result := db.First(&question, "question = ?", question.Question)
	if result.RowsAffected > 0 {
		Message := []byte(`{"Error": "Question already exist"}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	} else {
		if error != nil {
			log.Fatal(error)
			connections.SendError(writer, http.StatusBadRequest)
			return
		}

		error = db.Create(&question).Error

		if error != nil {
			log.Fatal(error)
			connections.SendError(writer, http.StatusInternalServerError)
			return
		}

		json, _ := json.Marshal(question)

		connections.SendReponse(writer, http.StatusCreated, json)
	}
}

func GetQuestions(writer http.ResponseWriter, request *http.Request) {
	var questions []models.Question

	db := connections.GetConnection()
	defer db.Close()

	result := db.Find(&questions)
	if result.RowsAffected > 0 {
		json, _ := json.Marshal(questions)
		connections.SendReponse(writer, http.StatusOK, json)
	} else {
		Message := []byte(`{"Error": "Questions not found"}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	}
}

func GetAnswer(writer http.ResponseWriter, request *http.Request) {
	question := models.Question{}

	id := mux.Vars(request)["id"]

	db := connections.GetConnection()
	defer db.Close()

	db.Find(&question, "id = ?", id)

	if question.ID > 0 {
		json, _ := json.Marshal(question.Answer)
		connections.SendReponse(writer, http.StatusOK, json)
	} else {
		Message := []byte(`{"Error": "Invalid question id "}`)
		connections.ErrorMsg(writer, http.StatusBadRequest, Message)
	}

}
