# ScraperIAGO

ScraperIAGO is a Go application designed to scrape news articles from the brasilnippou.com website. It leverages Google's Gemini AI to intelligently verify if the scraped articles are correctly categorized based on their titles and persists the results in a SQLite database.  - work in progress -

## Overview
The primary function of this project is to automate the process of data collection and content verification. It fetches paginated search results for specific keywords, extracts article information, and uses a concurrent worker pool to process them. For each new article, it queries the Gemini AI to confirm category relevance before saving the data, thus preventing redundant processing.

## Features
**Web Scraping:** Fetches and parses article data (title, URL) from the brasilnippou.com news portal.

**AI-Powered Verification:** Integrates with the Google Gemini API to validate if an article's title aligns with its assigned category.

**Concurrent Processing:** Employs a worker pool to efficiently process multiple articles in parallel, significantly speeding up the scraping 
process.

**Data Persistence:** Uses a local SQLite database (artigos.db) to store article information and verification results, ensuring data is not re-processed on subsequent runs.

**Basic Authentication:** Supports fetching content from pages protected by basic HTTP authentication.

## Architecture
The project is structured into several distinct packages, each with a specific responsibility:

**main.go:** The application's entry point. It orchestrates the scraper, the worker pool, and the database repository.

**/scraper:** Contains the logic for fetching raw HTML from web pages and extracting article information.

**/ia:** Manages all interactions with the Google Gemini API, including prompt engineering and API requests.

**/repository:** Implements the data access layer using GORM, handling all communication with the SQLite database.

**/models:** Defines the core data structures used throughout the application, most notably the Article struct.

**/util:** Provides a generic, concurrency-safe worker pool for executing tasks.

## How It Works
The application begins by fetching HTML content from the search result pages of brasilnippou.com for a predefined keyword.
The scraper parses the raw HTML to extract a list of articles, including their titles and URLs.
These articles are grouped and dispatched as tasks to a concurrent worker pool.
Each worker first checks the SQLite database to see if an article has already been processed.
For new articles, the worker constructs a specific prompt and sends it to the Gemini AI, requesting validation of the article's category based on its title.
The application parses the JSON response from the AI, which contains the validation result (IsCorrect).
Finally, the new article data, along with the AI's verification status, is saved to the SQLite database.

# Getting Started
## Prerequisites
Go (version 1.22 or later)
A Google Gemini API Key

### Installation
Clone the repository to your local machine:

```
git clone https://github.com/frukas/scraperiago.git
cd scraperiago
```

### Install the required dependencies:

```
go mod tidy
```

### Configuration
Before running the application, you must add your Google Gemini API key.

Open the file internal/ia/geminiAI.go.
Locate the apikey variable and replace the empty string with your actual API key:
// internal/ia/geminiAI.go

```
func AskGemini(question string) string {
    apikey := "YOUR_GEMINI_API_KEY"
    // ...
}
```

Usage
To run the application, execute the following command from the root directory:

```
go run main.go
```

The program will start, perform the scraping and AI verification process, and store the results in the artigos.db SQLite file.
