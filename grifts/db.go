package grifts

import (
	"fmt"

	"github.com/frodopwns/feedapi/models"
	"github.com/gobuffalo/uuid"
	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		fmt.Println("seeding!")

		users := models.Users{
			{Name: "Bob Dole", Email: "dole@home.place"},
			{Name: "Al Gore", Email: "al@green.peace"},
			{Name: "Kermit", Email: "frog@hops.alot"},
		}
		err := models.DB.Create(&users)
		if err != nil {
			return err
		}

		turtleID, _ := uuid.NewV4()
		cowID, _ := uuid.NewV4()

		feeds := models.Feeds{
			{Name: "Turtles", ID: turtleID},
			{Name: "Cows", ID: cowID},
		}
		err = models.DB.Create(&feeds)
		if err != nil {
			return err
		}

		articles := models.Articles{
			{
				Name:    "What I Love About Turtles",
				Content: "Turtles are the most exciting ninja warri...",
				FeedID:  turtleID,
			},
			{
				Name:    "Turtle Soup",
				Content: "yadayada..",
				FeedID:  turtleID,
			},
			{
				Name:    "Cows Make Great Pets",
				Content: "They play fetch, they trim the lawn, they fetilize the garden...",
				FeedID:  cowID,
			},
			{
				Name:    "Cows Have Feelings Too",
				Content: "fun times",
				FeedID:  cowID,
			},
		}
		err = models.DB.Create(&articles)
		if err != nil {
			return err
		}

		return nil
	})

})
