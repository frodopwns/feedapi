package actions

import (
	"encoding/json"
	"io/ioutil"

	"github.com/frodopwns/feedapi/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// FeedsIndex default implementation.
func FeedsIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	feeds := models.Feeds{}
	err := tx.All(&feeds)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(feeds))
}

// FeedsGet default implementation.
func FeedsGet(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	feed := models.Feed{}
	err := tx.Find(&feed, c.Param("feed_id"))
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(feed))
}

// FeedsCreate default implementation.
func FeedsCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	req := c.Request()

	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	feed := models.Feed{}
	err = json.Unmarshal(b, &feed)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	verrs, err := tx.ValidateAndCreate(&feed)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	if verrs.Count() != 0 {
		return c.Render(400, r.JSON(map[string]string{"error": verrs.String()}))
	}

	return c.Render(201, r.JSON(feed))
}

// FeedsDelete default implementation.
func FeedsDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	feed := models.Feed{}
	err := tx.Find(&feed, c.Param("feed_id"))
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	err = tx.Destroy(&feed)
	if err != nil {
		return c.Render(500, r.JSON(feed))
	}

	return c.Render(200, r.JSON(feed))
}
