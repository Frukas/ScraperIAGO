package ia

import (
	"fmt"
	"log"

	"github.com/frukas/scraperiago/internal/models"
	"github.com/frukas/scraperiago/internal/repository"
	"github.com/frukas/scraperiago/internal/scraper"
	"github.com/frukas/scraperiago/internal/util"
)

type CategoryDoubleCheck struct {
	Articles []models.Article
	repo     repository.Repository
}

func (ca *CategoryDoubleCheck) Run() {
	var newArticles []models.Article

	for _, art := range ca.Articles {
		if !ca.repo.Exists(art.Address) {
			newArticles = append(newArticles, art)
		}
	}

	if len(newArticles) == 0 {
		fmt.Println("Articles already processed, ignoring")
		return
	}

	fmt.Printf("Sending %d new articles to IA\n", len(newArticles))
	answer := AskGemini(MakeTheQuestion(newArticles))

	newArticles, err := models.TextToArticle(answer)
	if err != nil {
		log.Fatal("Errors converting from text to Article", err)
	}

	fmt.Println("Saving on db: ")
	if err := ca.repo.SaveAll(newArticles); err != nil {
		fmt.Println("Error saving new article:", err)
	}
}

func QuestionFactory(Articles []models.Article, repo repository.Repository, taskChan chan util.Task) {

	defer close(taskChan)

	for i := 0; i < len(Articles); i += 5 {
		start := i
		end := i + 5

		if end > len(Articles) {
			end = len(Articles)
		}

		taskChan <- &CategoryDoubleCheck{

			Articles: Articles[start:end],
			repo:     repo,
		}
	}
}

func QuestionFactoryString(Articles []string, repo repository.Repository, taskChan chan util.Task, category string) {

	defer close(taskChan)

	for _, art := range Articles {
		list := scraper.GetArticleList(art, category)
		if len(list) > 0 {
			taskChan <- &CategoryDoubleCheck{
				Articles: list,
				repo:     repo,
			}
		}
	}
}
