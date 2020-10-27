package domain

import (
	"../domain/model"
)

/*
* Repository Interface for question model
*
*	RetrieveListQuestion() (model.QuestionList, error) --> get the complete list of questions
*	CreateNewQuestion(question model.Question) error --> save a question in the file
*
 */
type QuestionRepository interface {
	RetrieveListQuestion() (model.QuestionList, error)
	CreateNewQuestion(question model.Question) error
}
