package main

import (
	"fmt"

	"github.com/frukas/scraperiago/internal/ia"
	"github.com/frukas/scraperiago/internal/models"
	"github.com/frukas/scraperiago/internal/repository"
	"github.com/frukas/scraperiago/internal/scraper"
	"github.com/frukas/scraperiago/internal/util"
)

func main() {
	fmt.Println("Program started")
	defer fmt.Println("Program ended")

	//pages := scraper.MultiPageSearch("https://test.brasilnippou.com/ja/search?query=経済指標", "gnw", "123")

	var pages []string

	pa := scraper.ReadHTML()

	pages = append(pages, pa)

	work := util.WorkFactory(2)

	work.RunWorker()
	defer work.Wait()

	repo, _ := repository.NewRepository()
	repo.Migration(models.Article{})

	go ia.QuestionFactoryString(pages, *repo, work.TaskChan, "自動車")

}
