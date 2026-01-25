package ia

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

// TODO: change this funcion name to more generic one, latter prepare a interface to use more easly other IA
func AskGemini(question string) string {
	apikey := ""
	

	customSchema := &genai.Schema{
		Type: genai.TypeArray,
		Items: &genai.Schema{
			Type: genai.TypeObject,
			Properties: map[string]*genai.Schema{
				"Id": {
					Type:        genai.TypeInteger,
					Description: "The unique ID of the article",
				},
				"Title": {
					Type: genai.TypeString,
				},
				"Category": {
					Type: genai.TypeString,
				},
				"Address": {
					Type: genai.TypeString,
				},
				"IsCorrect": {
					Type:        genai.TypeBoolean,
					Description: "True if the article matches the assigned category",
				},
			},
			Required: []string{"Id", "Title", "Category", "Address", "IsCorrect"},
		},
	}
	ctx := context.Background()
	geminiClient, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apikey,
	})
	if err != nil {
		panic("Error on gemini")
	}

	resultGemini, err := geminiClient.Models.GenerateContent(
		ctx, "gemini-2.5-flash", genai.Text(question), &genai.GenerateContentConfig{
			ResponseMIMEType:   "application/json",
			ResponseJsonSchema: customSchema,
		})

	//Change the gemini client
	// resultGemini, err := geminiClient.Models.GenerateContent(
	// 	ctx, "gemini-2.5-flash-lite", genai.Text(question), &genai.GenerateContentConfig{
	// 		ResponseMIMEType:   "application/json",
	// 		ResponseJsonSchema: customSchema,
	// 	})

	if err != nil {
		fmt.Println(err)
		panic("Error on the gemini question")
	}
	return resultGemini.Text()

}
