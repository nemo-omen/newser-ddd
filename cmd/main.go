package main

import (
	"fmt"
	"log"
	"newser/internal/domain/article"
	"newser/internal/domain/newsfeed"
	"newser/internal/domain/user"
	"newser/internal/domain/value"
	"time"
)

func main() {
	u1, err := user.NewUser("jeff@hello.whatever", "Jeff Caldwell")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("u1: %+v\n", u1)

	a1, err := article.NewArticle(article.ArticleProps{
		Title:           "The Title",
		Link:            "http://www.example.com",
		Links:           []string{"http://www.example.com"},
		Description:     "",
		Content:         "<p>This is the content. It's not very long</p>",
		Updated:         "2020-01-01T00:00:00Z",
		UpdatedParsed:   time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		Published:       "2020-01-01T00:00:00Z",
		PublishedParsed: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		GUID:            "123456",
		Image:           value.NewId().String(),
		FeedId:          value.NewId().String(),
	})
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("a1: %+v\n", a1)

	n1, err := newsfeed.NewNewsfeed(newsfeed.NewsfeedProps{
		Title:       "The Title",
		SiteUrl:     "http://www.example.com",
		FeedUrl:     "http://www.example.com/feed",
		Description: "The Description",
		Language:    "en",
		FeedType:    "rss",
	})

	if err != nil {
		log.Println(err.Error())
	}

	n1.AddArticle(a1.ID())
	fmt.Printf("n1: %+v\n", n1)

}
