package scraper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// For test purpose only. reads the examplePage.html to avoid to many requests to the server.
func ReadHTML() string {
	content, err := os.ReadFile("examplePage.html")
	if err != nil {
		panic("Error reading file") // Use log.Fatalf for critical errors
	}

	return string(content)
}

func GetPageContents(pageAddress string) string {

	resp, err := http.Get(pageAddress)
	if err != nil {
		fmt.Printf("Response Status Code: %d\n", resp.StatusCode)
		panic(fmt.Sprintf("Not possible to get the page %v", pageAddress))
	}

	defer resp.Body.Close()

	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Not possible to get the body of the page %v", pageAddress))
	}

	return string(body)
}

func GetPageContentsWithPassword(pageAddress string, username string, password string) string {

	client := &http.Client{}

	req, err := http.NewRequest("GET", pageAddress, nil)
	if err != nil {
		log.Fatal("Error preparing the request", err)
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("error get the page:", err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return string(body)
}

func MultiPageSearch(pageAddress string, username string, password string) []string {

	var listPage []string
	IsLastPage := false

	for i := 1; !IsLastPage; i++ {
		page := fmt.Sprintf("%s&page=%d", pageAddress, i)
		fmt.Println(page)
		res := GetPageContentsWithPassword(page, username, password)
		listPage = append(listPage, res)
		if strings.Contains(res, "検索結果が見つかりませんでした") {
			fmt.Println("last page count: ", i)
			IsLastPage = true
		}
	}

	return listPage
}
