# Go parameters
API_REST=public/apiRest.go
NAME_API=applicationRest


apiRest:
	go build -o $(NAME_API) -race $(API_REST)
	./$(NAME_SERVER_PATH)

