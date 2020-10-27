package handlers

import (
	"../../application"
	"../../domain/model"
	"encoding/json"
	"net/http"
)

type QuestionRestController interface {
	GetQuestionList(w http.ResponseWriter, r *http.Request)
	SaveQuestion(w http.ResponseWriter, r *http.Request)
}

type questionRestController struct {
	service application.QuestionService
}

func NewQuestionRestHandler(service application.QuestionService) QuestionRestController {
	return &questionRestController{
		service,
	}
}

func (q questionRestController) GetQuestionList(w http.ResponseWriter, r *http.Request) {

	key := r.FormValue("lang")

	//getting the list
	list, err := q.service.GetQuestionList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := q.service.EncodeQuestion(list)

	//this does not work
	q.service.TranslateQuestions(key,&list)

	//create the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(response); err != nil {
		http.Error(w, "Unable create the response ", http.StatusInternalServerError)
		return
	}
}

func (q questionRestController) SaveQuestion(w http.ResponseWriter, r *http.Request) {
	var request model.Question
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad format for request", http.StatusBadRequest)
		return
	}
	err = q.service.SaveQuestion(request)
	if err != nil {
		http.Error(w, "Unable save the campaign", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if _, err := w.Write([]byte("Question saved")); err != nil {
		http.Error(w, "Unable create the response ", http.StatusInternalServerError)
		return

	}
}
