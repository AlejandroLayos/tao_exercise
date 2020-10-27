package application

import (
	"../domain/model"
)

/*
* Class that retrieve the complete list of questions
*
* input -
* @return - list model.QuestionList, err error
 */

func (q questionService) GetQuestionList() (model.QuestionList, error) {
	listToReturn, err := q.repo.RetrieveListQuestion()
	if err != nil {
		q.logger.QuestionError(err)
		return nil, q.exceptions.ErrorFindingQuestionsList(err.Error())
	}
	q.logger.QuestionInfo("Question list retrieved")
	return listToReturn, nil
}
