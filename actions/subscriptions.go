package actions

import (
	"encoding/json"
	"io/ioutil"

	"github.com/frodopwns/feedapi/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// SubscriptionsIndex default implementation.
func SubscriptionsIndex(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	subscriptions := models.Subscriptions{}
	err := tx.All(&subscriptions)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(subscriptions))
}

// SubscriptionsGet default implementation.
func SubscriptionsGet(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	subscription := models.Subscription{}
	err := tx.Find(&subscription, c.Param("subscription_id"))
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	return c.Render(200, r.JSON(subscription))
}

// SubscriptionsCreate default implementation.
func SubscriptionsCreate(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	req := c.Request()

	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	subscription := models.Subscription{}
	err = json.Unmarshal(b, &subscription)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	verrs, err := tx.ValidateAndCreate(&subscription)
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	if verrs.Count() != 0 {
		return c.Render(400, r.JSON(map[string]string{"error": verrs.String()}))
	}

	return c.Render(201, r.JSON(subscription))
}

// SubscriptionsDelete default implementation.
func SubscriptionsDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	subscription := models.Subscription{}
	err := tx.Find(&subscription, c.Param("subscription_id"))
	if err != nil {
		return c.Render(500, r.JSON(map[string]string{"error": err.Error()}))
	}

	err = tx.Destroy(&subscription)
	if err != nil {
		return c.Render(500, r.JSON(subscription))
	}

	return c.Render(200, r.JSON(subscription))
}
