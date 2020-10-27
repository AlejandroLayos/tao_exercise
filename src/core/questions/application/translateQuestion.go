package application

import (
	"../domain/model"
)

func (q questionService) TranslateQuestions(lang string, list *model.QuestionList) {

	/*ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		//TODO
			}
	defer client.Close()

	req := &translatepb.TranslateTextRequest{
		SourceLanguageCode: "EN",
		TargetLanguageCode: lang,
		MimeType:           "text/plain", // Mime types: "text/plain", "text/html"
		Contents:           []string{"this is a test"},
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		//TODO
		}*/


}
