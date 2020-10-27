package application

import (
	"../domain/model"
	"encoding/json"
)

/*
* Class that encode the model to json ([]byte)
*
* input - list model.QuestionList
* @return - response []byte
 */

func (q questionService) EncodeQuestion(list model.QuestionList) []byte {
	response, _ := json.Marshal(list)
	q.logger.QuestionInfo("Question encoded")
	return response
}
