package scraper

import (
	"strings"

	"github.com/frukas/scraperiago/internal/models"
)

func toArticle(text []string, category string) *models.Article {
	var temp models.Article

	for i, res := range text {

		if strings.Contains(res, "/ja/articles/") {
			temp.Address = strings.Split(res, ":")[1]
		}

		if strings.Contains(res, "Card_card__title") {
			temp.Title = strings.Trim(strings.Split(text[i+1], ":")[1], "}]")
		}
	}

	temp.Category = category

	return &temp
}

func GetArticleList(text string, category string) []models.Article {
	var ArticleList []models.Article
	resultParcial := strings.Split(text, "<script>")

	for i, res := range resultParcial {
		if i <= 0 {
			continue
		}

		if strings.Contains(res, "Card_card__title") {
			ArticleList = append(ArticleList, *toArticle(strings.Split(res, ","), category))
		}
	}

	return ArticleList
}
