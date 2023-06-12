package handler_test

import (
	"testing"

	"github.com/ericprd/articles-be/handlers"
)

func TestGetAllArticles(t *testing.T) {
	aList := handlers.GetAllArticles()

	if len(aList) != len(handlers.Articles) {
		t.Fail()
	}

	for i, v := range aList {
		if v.Content != handlers.Articles[i].Content ||
			v.ID != handlers.Articles[i].ID ||
			v.Title != handlers.Articles[i].Title {
				t.Fail()
				break
		}
	}
}
