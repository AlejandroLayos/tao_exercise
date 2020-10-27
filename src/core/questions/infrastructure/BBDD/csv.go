package BBDD

import (
	"../../domain"
	"../../domain/model"
	"github.com/tushar2708/altcsv"
	"os"
)

type questionCSVRepository struct {
}

/*
* Repository implementation for CSV files
*
 */

func (q questionCSVRepository) RetrieveListQuestion() (model.QuestionList, error) {
	csvfile, err := os.Open("src/core/questions/infrastructure/BBDD/questions.csv")
	if err != nil {
		return nil, err
	}
	defer csvfile.Close()

	//reader settings
	reader := altcsv.NewReader(csvfile)
	reader.FieldsPerRecord = -1
	reader.Comma = ','
	reader.LazyQuotes = true

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var list model.QuestionList

	//to avoid the titles
	firstLine := true

	//transform data in to our model
	for _, record := range rawCSVdata {
		if !firstLine {
			var question model.Question
			question.Text = record[0]
			question.CreatedAt = record[1]

			question.Choices = append(question.Choices, model.Choices{
				Text: record[2],
			})
			question.Choices = append(question.Choices, model.Choices{
				Text: record[3],
			})
			question.Choices = append(question.Choices, model.Choices{
				Text: record[4],
			})

			list = append(list, question)
		} else {
			firstLine = false
		}

	}
	return list, nil

}
func (q questionCSVRepository) CreateNewQuestion(question model.Question) error {
	csvfile, err := os.OpenFile("src/core/questions/infrastructure/BBDD/questions.csv", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer csvfile.Close()
	writer := altcsv.NewWriter(csvfile)
	writer.UseCRLF = true
	defer writer.Flush()
	var record = []string{question.Text,
		question.CreatedAt,
		question.Choices[0].Text,
		question.Choices[1].Text,
		question.Choices[2].Text}
	writer.AllQuotes = true
	writer.Write(record)
	return nil
}

func NewQuestionCSVRepository() domain.QuestionRepository {
	return &questionCSVRepository{}
}
