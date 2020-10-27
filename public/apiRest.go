package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"../config"

	"../src/services/logger"
	"../src/services/server"

	"../src/core/questions/application"
	"../src/core/questions/domain"

	"../src/core/questions/infrastructure/BBDD"
	"../src/core/questions/infrastructure/handlers"
	"../src/core/questions/infrastructure/loggers"
)

func main() {

	//Main configuration
	configParser := config.NewConfig()
	configParser.ConfigYaml()

	//start the logger
	logger := logger.NewLogrusLogger()

	var questionRepo domain.QuestionRepository

	//selecting the type of data we are going to use
	switch os := os.Getenv("DATATYPE"); os {
	case "json":
		questionRepo = BBDD.NewQuestionJsonRepository()
	case "csv":
		questionRepo = BBDD.NewQuestionCSVRepository()

	}
	//selecting components
	questionLogs := loggers.NewLogrusQuestionLogger(logger.ReturnLogrus())
	questionService := application.NewQuestionService(questionRepo, questionLogs)
	questionRestHandler := handlers.NewQuestionRestHandler(questionService)
	
	//starting the router
	muxRouter := server.NewMuxRouter(questionRestHandler)
	muxRouter.AssignQuestionRoutes()

	http.ListenAndServe(os.Getenv("REST_PORT"), muxRouter.ReturnRouter())

	chErr := make(chan error)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-quit:
			goto end
		case err := <-chErr:
			fmt.Print(err)
		}
	}
end:
}
