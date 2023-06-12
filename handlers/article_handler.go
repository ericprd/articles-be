package handlers

import "github.com/ericprd/articles-be/models"

var Articles = []models.Article {
	{ ID: 1, Title: "Article 1", Content: "Article 1 body" },
	{ ID: 2, Title: "Article 2", Content: "Article 2 body" },
}

func GetAllArticles() []models.Article {
	return Articles
}
