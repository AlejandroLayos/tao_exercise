package BBDD

import (
	"../../domain"
	"../../domain/model"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type questionJsonRepository struct {
}
/*
* Repository implementation for JSON files
*
 */
func (q questionJsonRepository) RetrieveListQuestion() (model.QuestionList, error) {

	jsonFile, err := os.Open("src/core/questions/infrastructure/BBDD/questions.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	var questionsList model.QuestionList
	err = json.Unmarshal(byteValue, &questionsList)
	if err != nil {
		return nil, err
	}
	return questionsList, nil
}

func (q questionJsonRepository) CreateNewQuestion(question model.Question) error {
	//retrieve the list
	list, err := q.RetrieveListQuestion()
	if err != nil {
		return err
	}
	//append the new question
	list = append(list, question)
	questionJson, err := json.Marshal(list)
	if err != nil {
		return err
	}
	//open the file to overwrite
	f, err := os.OpenFile("src/core/questions/infrastructure/BBDD/questions.json", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	//Empty the file to overwrite with the new json
	f.Truncate(0)
	io.WriteString(f, string(questionJson))

	return nil
}

func NewQuestionJsonRepository() domain.QuestionRepository {
	return &questionJsonRepository{}
}
