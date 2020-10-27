package server

import (
	"../../../src/core/questions/infrastructure/handlers"
	"github.com/gorilla/mux"
)

type muxRouter struct {
	mux              *mux.Router
	questionHandlers handlers.QuestionRestController
}

type Mux interface {
	AssignQuestionRoutes() error
	ReturnRouter() *mux.Router
}

func NewMuxRouter(restHandler handlers.QuestionRestController) Mux {
	return &muxRouter{
		mux.NewRouter(), restHandler,
	}
}

func (m muxRouter) AssignQuestionRoutes() error {

	m.mux.HandleFunc("/questions", m.questionHandlers.GetQuestionList).Methods("GET").Queries("lang", "{lang}")
	m.mux.HandleFunc("/questions", m.questionHandlers.SaveQuestion).Methods("POST")
	return nil
}

func (m muxRouter) ReturnRouter() *mux.Router {
	return m.mux
}
