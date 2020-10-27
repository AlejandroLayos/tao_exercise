package application

import (
	"../domain"
	"../domain/model"
)

type QuestionService interface {
	GetQuestionList() (model.QuestionList, error)
	SaveQuestion(question model.Question) error
	EncodeQuestion(model.QuestionList) []byte
	TranslateQuestions(lang string, list *model.QuestionList)
}

type questionService struct {
	repo       domain.QuestionRepository
	logger     domain.QuestionLogger
	exceptions domain.QuestionsExceptions
}

func NewQuestionService(repo domain.QuestionRepository, logger domain.QuestionLogger) QuestionService {
	return &questionService{
		repo, logger, domain.NewQuestionExceptions(),
	}
}
