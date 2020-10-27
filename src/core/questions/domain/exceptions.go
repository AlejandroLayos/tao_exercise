package domain

import (
	"errors"
)

/*
* Exception class for question model
*
 */

type QuestionsExceptions interface {
	ErrorFindingQuestionsList(string) error
	ErrorSavingQuestion(string) error
}

type questionsExceptions struct {
}

func (e questionsExceptions) ErrorFindingQuestionsList(err string) error {
	return errors.New("error retrieving list from file: " + err)
}

func (e questionsExceptions) ErrorSavingQuestion(err string) error {
	return errors.New("error saving new question in to file: " + err)
}

func NewQuestionExceptions() QuestionsExceptions {
	return &questionsExceptions{}
}
