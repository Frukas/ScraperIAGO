package ia

import (
	"encoding/json"
	"fmt"

	"github.com/frukas/scraperiago/internal/models"
)

func MakeTheQuestion(Articles []models.Article) string {

	var inputList []models.Article

	for _, a := range Articles {
		inputList = append(inputList, models.Article{
			Id:       a.Id,
			Category: a.Category,
			Address:  a.Address,
			Title:    a.Title,
		})
	}

	jsonData, _ := json.MarshalIndent(inputList, "", " ")

	return fmt.Sprintf(`Please verify the following articles. 
Check if the "title" matches the "category" based on economic interpretation.
Output: > Return the updated JSON array with the IsCorrect field set to true or false wherever matches the category. Do not add any conversational text.

Articles Data:
%s`, string(jsonData))

}
