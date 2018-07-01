package actions

import (
	"encoding/json"
	"io/ioutil"

	"github.com/frodopwns/feedapi/models"
	"github.com/markbates/willie"
)

// Test_Users_Create tests that we can POST users
func (as *ActionSuite) Test_Users_Create() {
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	user := &models.User{Name: "TestUser", Email: "test@user.com"}
	//res := as.HTML("/users").Post(user)
	res := as.JSON("/users").Post(user)
	as.Equal(201, res.Code)

	err = as.DB.First(user)
	as.NoError(err)

	as.NotZero(user.ID)
	as.Equal("TestUser", user.Name)
	as.Equal("test@user.com", user.Email)

}

// Test_Users_Delete tests that we can DELETE users
func (as *ActionSuite) Test_Users_Delete() {
	user := &models.User{Name: "TestUser2", Email: "test2@user.com"}
	//res := as.HTML("/users").Post(user)
	res := as.JSON("/users").Post(user)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(user)
	as.NoError(err)

	as.NotZero(user.ID)
	as.NotNil(user.ID)

	w := willie.New(as.App)
	res2 := w.Request("/users/" + user.ID.String()).Delete()
	as.Equal(200, res2.Code)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)
}

// Test_Users_Index tests that the user indes page returns items properly
func (as *ActionSuite) Test_Users_Index() {
	as.LoadFixture("test data")
	res := as.JSON("/users").Get()

	body := res.Body.String()

	as.Contains(body, "user #1")
	as.Contains(body, "user1@example.com")
	as.Contains(body, "user #2")
	as.Contains(body, "user2@example.com")
}

// Test_Users_Get tests that we can get a user by id
func (as *ActionSuite) Test_Users_Get() {
	user := &models.User{Name: "TestUser", Email: "test@user.com"}

	res := as.JSON("/users").Post(user)
	as.Equal(201, res.Code)

	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)

	err = as.DB.First(user)
	as.NoError(err)

	res = as.JSON("/users/" + user.ID.String()).Get()

	b, err := ioutil.ReadAll(res.Body)
	as.NoError(err)

	newUser := models.User{}
	err = json.Unmarshal(b, &newUser)
	as.NoError(err)

	as.Equal(newUser.Name, user.Name)
	as.Equal(newUser.Email, user.Email)
}

// Test_Users_Get_Feeds tests that getting a user's feeds returns the correct results
func (as *ActionSuite) Test_Users_Get_Feeds() {
	as.LoadFixture("test data")
	res := as.JSON("/users").Get()

	users := models.Users{}
	err := json.Unmarshal(res.Body.Bytes(), &users)
	as.NoError(err)

	res = as.JSON("/users/" + users[0].ID.String() + "/feeds").Get()
	as.Equal(200, res.Code)
	body := res.Body.String()

	as.Contains(body, "Feed1")
	as.NotContains(body, "Feed2")

	res = as.JSON("/users/" + users[1].ID.String() + "/feeds").Get()
	as.Equal(200, res.Code)
	body = res.Body.String()

	as.Contains(body, "Feed2")
	as.NotContains(body, "Feed1")

	res = as.JSON("/users/" + users[2].ID.String() + "/feeds").Get()
	as.Equal(200, res.Code)
	body = res.Body.String()

	as.Contains(body, "Feed2")
	as.Contains(body, "Feed1")
	as.NotContains(body, "Feed3")
}

// Test_Users_Get_Articles tests that gettings a user's articles returns the correct articles
func (as *ActionSuite) Test_Users_Get_Articles() {
	as.LoadFixture("test data")
	res := as.JSON("/users").Get()

	users := models.Users{}
	err := json.Unmarshal(res.Body.Bytes(), &users)
	as.NoError(err)

	res = as.JSON("/users/" + users[0].ID.String() + "/articles").Get()
	as.Equal(200, res.Code)
	body := res.Body.String()

	as.Contains(body, "Election Results")
	as.NotContains(body, "Game Results")
	as.NotContains(body, "Russia Cheats At Olympics")

	res = as.JSON("/users/" + users[1].ID.String() + "/articles").Get()
	as.Equal(200, res.Code)
	body = res.Body.String()

	as.NotContains(body, "Election Results")
	as.Contains(body, "Game Results")
	as.Contains(body, "Russia Cheats At Olympics")

	res = as.JSON("/users/" + users[2].ID.String() + "/articles").Get()
	as.Equal(200, res.Code)
	body = res.Body.String()

	as.Contains(body, "Election Results")
	as.Contains(body, "Game Results")
	as.Contains(body, "Russia Cheats At Olympics")
}
