package actions

import (
	"encoding/json"
	"io/ioutil"

	"github.com/frodopwns/feedapi/models"
	"github.com/markbates/willie"
)

// Test_Articles_Delete creates and then deletes an article making sure it is gone
func (as *ActionSuite) Test_Articles_Delete() {
	feed := &models.Feed{Name: "TestFeed"}
	res := as.JSON("/feeds").Post(feed)
	as.Equal(201, res.Code)

	err := as.DB.First(feed)
	as.NoError(err)

	article := &models.Article{
		Name:    "TestArticle",
		Content: "I am content!",
		FeedID:  feed.ID,
	}
	res = as.JSON("/articles").Post(article)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("articles")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(article)
	as.NoError(err)

	w := willie.New(as.App)
	res2 := w.Request("/articles/" + article.ID.String()).Delete()
	as.Equal(200, res2.Code)

	count, err = as.DB.Count("articles")
	as.NoError(err)
	as.Equal(0, count)
}

// Test_Articles_Create POSTs an article and then ensures it exists in the DB
func (as *ActionSuite) Test_Articles_Create() {
	feed := &models.Feed{Name: "TestFeed"}
	res := as.JSON("/feeds").Post(feed)
	as.Equal(201, res.Code)

	err := as.DB.First(feed)
	as.NoError(err)

	article := &models.Article{
		Name:    "TestArticle",
		Content: "I am content!",
		FeedID:  feed.ID,
	}
	res = as.JSON("/articles").Post(article)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("articles")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(article)
	as.NoError(err)

	as.Equal("TestArticle", article.Name)
	as.Equal("I am content!", article.Content)
	as.Equal(feed.ID, article.FeedID)
}

// Test_Articles_Index loads fixture data and then ensures the articles index endpoint returns all expected articles
func (as *ActionSuite) Test_Articles_Index() {
	count, err := as.DB.Count("articles")
	as.NoError(err)
	as.Equal(0, count)

	as.LoadFixture("test data")

	res := as.JSON("/articles").Get()
	body := res.Body.String()

	as.Contains(body, "Election Results")
	as.Contains(body, "Primary Election Results updated by the hour...")
	as.Contains(body, "Game Results")
	as.Contains(body, "World Cup Results updated by the hour...")
	as.Contains(body, "Russia Cheats At Olympics")
	as.Contains(body, "Russia Cheats At Olympics")
}

// Test_Articles_Get tests whether we can get and article by id
func (as *ActionSuite) Test_Articles_Get() {
	feed := &models.Feed{Name: "TestFeed"}
	res := as.JSON("/feeds").Post(feed)
	as.Equal(201, res.Code)

	err := as.DB.First(feed)
	as.NoError(err)

	article := &models.Article{
		Name:    "TestArticle",
		Content: "I am content!",
		FeedID:  feed.ID,
	}
	res = as.JSON("/articles").Post(article)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("articles")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(article)
	as.NoError(err)

	res = as.JSON("/articles/" + article.ID.String()).Get()

	b, err := ioutil.ReadAll(res.Body)
	as.NoError(err)

	newArticle := models.Article{}
	err = json.Unmarshal(b, &newArticle)
	as.NoError(err)

	as.Equal(newArticle.Name, article.Name)
	as.Equal(newArticle.Content, article.Content)
}
