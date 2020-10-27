package application

import (
	"../domain/model"
	"encoding/json"
	"io"
)
/*
* Class that decode json to request to our model
*
* input - body io.ReadCloser
* @return - model model.Question
 */

func (q questionService) Decode(body io.ReadCloser) model.Question {
	var request model.Question
	decoder := json.NewDecoder(body)
	decoder.Decode(&request)
	q.logger.QuestionInfo("Question Decoded")
	return request
}
