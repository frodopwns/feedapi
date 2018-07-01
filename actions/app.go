package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/frodopwns/feedapi/models"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_feedapi_session",
		})
		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Set the request content type to JSON
		app.Use(middleware.SetContentType("application/json"))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Wraps each request in a transaction.
		//  c.Value("tx").(*pop.PopTransaction)
		// Remove to disable this.
		app.Use(middleware.PopTransaction(models.DB))

		app.GET("/", HomeHandler)

		// Users CRUD routes
		app.GET("/users", UsersIndex)
		app.GET("/users/{user_id}", UserGet)
		app.GET("/users/{user_id}/feeds", UserGetFeeds)
		app.GET("/users/{user_id}/articles", UserGetArticles)
		app.POST("/users", UsersCreate)
		app.DELETE("/users/{user_id}", UsersDelete)

		// Articles CRUD routes
		app.GET("/articles", ArticlesIndex)
		app.GET("/articles/{article_id}", ArticlesGet)
		app.POST("/articles", ArticlesCreate)
		app.DELETE("/articles/{article_id}", ArticlesDelete)

		// Feeds CRUD
		app.GET("/feeds", FeedsIndex)
		app.GET("/feeds/{feed_id}", FeedsGet)
		app.POST("/feeds", FeedsCreate)
		app.DELETE("/feeds/{feed_id}", FeedsDelete)

		// Subscriptions CRUD
		app.GET("/subscriptions", SubscriptionsIndex)
		app.GET("/subscriptions/{subscription_id}", SubscriptionsGet)
		app.POST("/subscriptions", SubscriptionsCreate)
		app.DELETE("/subscriptions/{subscription_id}", SubscriptionsDelete)
	}

	return app
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return ssl.ForceSSL(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
