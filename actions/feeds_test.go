package actions

import (
	"encoding/json"

	"github.com/frodopwns/feedapi/models"
	"github.com/markbates/willie"
)

// Test_Feeds_Index creates a feed and ensures it is returned in the list of feeds at /feeds
func (as *ActionSuite) Test_Feeds_Index() {
	count, err := as.DB.Count("feeds")
	as.NoError(err)
	as.Equal(0, count)

	as.LoadFixture("test data")

	res := as.JSON("/feeds").Get()
	as.Equal(200, res.Code)
	body := res.Body.String()

	as.Contains(body, "Feed1")
	as.Contains(body, "Feed2")
	as.Contains(body, "Feed3")
}

// Test_Feeds_Get makes sure we can get feeds by id
func (as *ActionSuite) Test_Feeds_Get() {
	feed := &models.Feed{Name: "TestFeed"}
	res := as.JSON("/feeds").Post(feed)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("feeds")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(feed)
	as.NoError(err)

	res = as.JSON("/feeds/" + feed.ID.String()).Get()
	as.Equal(200, res.Code)

	newFeed := models.Feed{}
	err = json.Unmarshal(res.Body.Bytes(), &newFeed)
	as.NoError(err)

	as.Equal(newFeed.Name, feed.Name)
}

// Test_Feeds_Create makes sure we can POST new feeds
func (as *ActionSuite) Test_Feeds_Create() {
	feed := &models.Feed{Name: "TestFeed"}
	res := as.JSON("/feeds").Post(feed)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("feeds")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(feed)
	as.NoError(err)

	as.Equal("TestFeed", feed.Name)
}

// Test_Feeds_Delete tests that DELETE queries work properly
func (as *ActionSuite) Test_Feeds_Delete() {
	feed := &models.Feed{Name: "TestFeed"}
	res := as.JSON("/feeds").Post(feed)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("feeds")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(feed)
	as.NoError(err)
	as.NotZero(feed.ID)

	w := willie.New(as.App)
	res2 := w.Request("/feeds/" + feed.ID.String()).Delete()
	as.Equal(200, res2.Code)

	count, err = as.DB.Count("feeds")
	as.NoError(err)
	as.Equal(0, count)
}
