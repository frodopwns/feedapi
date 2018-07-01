package actions

import (
	"encoding/json"
	"io/ioutil"

	"github.com/frodopwns/feedapi/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// UsersCreate creates a user from json in request body
func UsersCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	req := c.Request()

	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	user := models.User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	verrs, err := tx.ValidateAndCreate(&user)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	if verrs.Count() != 0 {
		return c.Render(400, r.JSON(map[string]string{"error": verrs.String()}))
	}

	return c.Render(201, r.JSON(user))
}

// UsersDelete deletes the user that has the same id as the user_id parameter
func UsersDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user := models.User{}
	err := tx.Find(&user, c.Param("user_id"))
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	err = tx.Destroy(&user)
	if err != nil {
		return c.Render(500, r.JSON(user))
	}

	return c.Render(200, r.JSON(user))
}

// UsersIndex returns all users
func UsersIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	users := models.Users{}
	err := tx.All(&users)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(users))
}

// UserGet return a single user with the given id
func UserGet(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user := models.User{}
	err := tx.Find(&user, c.Param("user_id"))
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(user))
}

// UserGetFeeds returns all feeds a user is subscribed to
func UserGetFeeds(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	feeds := models.Feeds{}
	allFeeds := tx.Where("feeds.name IS NOT NULL")
	query := allFeeds.LeftJoin("subscriptions", "subscriptions.feed_id=feeds.id").
		Where(`subscriptions.user_id=?`, c.Param("user_id"))

	sql, args := query.ToSQL(&pop.Model{Value: models.Feeds{}}, "feeds.*")
	err := tx.RawQuery(sql, args...).All(&feeds)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(feeds))
}

// UserGetArticles returns all articles for feeds a user is subscribed to
func UserGetArticles(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	articles := models.Articles{}
	allArticles := tx.Where("articles.name IS NOT NULL")
	query := allArticles.LeftJoin("subscriptions", "subscriptions.feed_id=articles.feed_id").
		Where(`subscriptions.user_id=?`, c.Param("user_id"))

	sql, args := query.ToSQL(&pop.Model{Value: models.Articles{}}, "articles.*")
	err := tx.RawQuery(sql, args...).All(&articles)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(articles))
}
