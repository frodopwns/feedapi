package actions

import (
	"encoding/json"
	"io/ioutil"

	"github.com/frodopwns/feedapi/models"
	"github.com/markbates/willie"
)

// Test_Subscriptions_Index makes sure the subscription index endpoint returns the correct response
func (as *ActionSuite) Test_Subscriptions_Index() {
	user := &models.User{Name: "TestUser", Email: "test@user.com"}
	res := as.JSON("/users").Post(user)
	as.Equal(201, res.Code)

	feed := &models.Feed{Name: "Tests"}
	res = as.JSON("/feeds").Post(feed)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(user)
	as.NoError(err)
	as.NotZero(user.ID)

	err = as.DB.First(feed)
	as.NoError(err)
	as.NotZero(feed.ID)

	subscription := &models.Subscription{
		FeedID: feed.ID,
		UserID: user.ID,
	}
	res = as.JSON("/subscriptions").Post(subscription)
	as.Equal(201, res.Code)

	res = as.JSON("/subscriptions").Get()

	b, err := ioutil.ReadAll(res.Body)
	as.NoError(err)

	subscriptions := models.Subscriptions{}
	err = json.Unmarshal(b, &subscriptions)
	as.NoError(err)

	as.Equal(feed.ID, subscriptions[0].FeedID)
	as.Equal(user.ID, subscriptions[0].UserID)
}

// Test_Subscriptions_Get makes sure we can get the subscription by id
func (as *ActionSuite) Test_Subscriptions_Get() {
	user := &models.User{Name: "TestUser", Email: "test@user.com"}
	res := as.JSON("/users").Post(user)
	as.Equal(201, res.Code)

	feed := &models.Feed{Name: "Tests"}
	res = as.JSON("/feeds").Post(feed)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(user)
	as.NoError(err)
	as.NotZero(user.ID)

	err = as.DB.First(feed)
	as.NoError(err)
	as.NotZero(feed.ID)

	subscription := &models.Subscription{
		FeedID: feed.ID,
		UserID: user.ID,
	}
	res = as.JSON("/subscriptions").Post(subscription)
	as.Equal(201, res.Code)

	err = as.DB.First(subscription)
	as.NoError(err)
	as.NotZero(subscription.ID)

	res = as.JSON("/subscriptions/" + subscription.ID.String()).Get()

	b, err := ioutil.ReadAll(res.Body)
	as.NoError(err)

	newSubscription := models.Subscription{}
	err = json.Unmarshal(b, &newSubscription)
	as.NoError(err)

	as.Equal(newSubscription.UserID, user.ID)
	as.Equal(newSubscription.FeedID, feed.ID)
}

// Test_Subscriptions_Create makes sure we can post new subscriptions
func (as *ActionSuite) Test_Subscriptions_Create() {
	user := &models.User{Name: "TestUser", Email: "test@user.com"}
	res := as.JSON("/users").Post(user)
	as.Equal(201, res.Code)

	feed := &models.Feed{Name: "Tests"}
	res = as.JSON("/feeds").Post(feed)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(user)
	as.NoError(err)
	as.NotZero(user.ID)

	err = as.DB.First(feed)
	as.NoError(err)
	as.NotZero(feed.ID)

	subscription := &models.Subscription{
		FeedID: feed.ID,
		UserID: user.ID,
	}
	res = as.JSON("/subscriptions").Post(subscription)
	as.Equal(201, res.Code)

	err = as.DB.First(subscription)
	as.NoError(err)
	as.NotZero(subscription.ID)

	as.Equal(user.ID, subscription.UserID)
	as.Equal(feed.ID, subscription.FeedID)
}

// Test_Subscriptions_Delete makes sure a DELETE requests is handles properly
func (as *ActionSuite) Test_Subscriptions_Delete() {
	user := &models.User{Name: "TestUser", Email: "test@user.com"}
	res := as.JSON("/users").Post(user)
	as.Equal(201, res.Code)

	feed := &models.Feed{Name: "Tests"}
	res = as.JSON("/feeds").Post(feed)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(user)
	as.NoError(err)
	as.NotZero(user.ID)

	err = as.DB.First(feed)
	as.NoError(err)
	as.NotZero(feed.ID)

	subscription := &models.Subscription{
		FeedID: feed.ID,
		UserID: user.ID,
	}
	res = as.JSON("/subscriptions").Post(subscription)
	as.Equal(201, res.Code)

	count, err = as.DB.Count("subscriptions")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(subscription)
	as.NoError(err)
	as.NotZero(subscription.ID)

	w := willie.New(as.App)
	res2 := w.Request("/subscriptions/" + subscription.ID.String()).Delete()
	as.Equal(200, res2.Code)

	count, err = as.DB.Count("subscriptions")
	as.NoError(err)
	as.Equal(0, count)
}
