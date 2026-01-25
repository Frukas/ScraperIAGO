package models

import (
	"encoding/json"
	"strings"
)

func TextToArticle(text string) ([]Article, error) {

	var resp []Article

	text = fixTextJason(text)

	if err := json.Unmarshal([]byte(text), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func fixTextJason(text string) string {
	text = strings.TrimSpace(text)

	if !strings.HasSuffix(text, "]") {
		if !strings.HasSuffix(text, "}") {
			text += "}]"
		} else {
			text += "]"
		}
	}

	return text
}
