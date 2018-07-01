package actions

import (
	"encoding/json"
	"io/ioutil"

	"github.com/frodopwns/feedapi/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// ArticlesDelete default implementation.
func ArticlesDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	article := models.Article{}
	err := tx.Find(&article, c.Param("article_id"))
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	err = tx.Destroy(&article)
	if err != nil {
		return c.Render(500, r.JSON(article))
	}

	return c.Render(200, r.JSON(article))
}

// ArticlesCreate default implementation.
func ArticlesCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	req := c.Request()

	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	article := models.Article{}
	err = json.Unmarshal(b, &article)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	verrs, err := tx.ValidateAndCreate(&article)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	if verrs.Count() != 0 {
		return c.Render(400, r.JSON(map[string]string{"error": verrs.String()}))
	}

	return c.Render(201, r.JSON(article))
}

// ArticlesIndex default implementation.
func ArticlesIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	articles := models.Articles{}
	err := tx.All(&articles)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(articles))
}

// ArticlesGet default implementation.
func ArticlesGet(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	article := models.Article{}
	err := tx.Find(&article, c.Param("article_id"))
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(article))
}
