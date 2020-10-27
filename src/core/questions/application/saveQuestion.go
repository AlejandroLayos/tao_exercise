package application

import (
	"../domain/model"
)

/*
* Class that save the requested question
*
* input - question model.Question
* @return - err error
 */

func (q questionService) SaveQuestion(question model.Question) error {

	err := q.repo.CreateNewQuestion(question)
	if err != nil {
		q.logger.QuestionError(err)
		return q.exceptions.ErrorSavingQuestion(err.Error())
	}
	return nil
}
