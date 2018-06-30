package actions

import (
	"encoding/json"
	"io/ioutil"

	"github.com/frodopwns/feedapi/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// UsersCreate default implementation.
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

	return c.Render(200, r.JSON(user))
}

// UsersDelete default implementation.
func UsersDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user := models.User{}
	err := tx.Find(&user, c.Param("user_id"))
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	err = tx.Destroy(&user)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(map[string]string{"message": "ok"}))
}

// UsersIndex default implementation.
func UsersIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	users := models.Users{}
	err := tx.All(&users)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(users))
}

// UserGet return a sinfle user with the given id
func UserGet(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	user := models.User{}
	err := tx.Find(&user, c.Param("user_id"))
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(user))
}
