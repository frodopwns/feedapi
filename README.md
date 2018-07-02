# FeedAPI

An API for creating Users, Feeds, Articles, and Subscriptions.

## API Doc

### Paths

METHOD | PATH                             |  DESCRIPTIOM
------ | ----                             |  -------
GET    | /articles                        | Get all articles
POST   | /articles                        | Create a new article
DELETE | /articles/{article_id}           | Delete an article
GET    | /articles/{article_id}           | Get an article by id
POST   | /feeds                           | Create a feed
GET    | /feeds                           | Get all feeds
GET    | /feeds/{feed_id}                 | Get a feed by id
DELETE | /feeds/{feed_id}                 | Delete a feed
POST   | /subscriptions                   | Create a subscription
GET    | /subscriptions                   | Get all subscriptions
GET    | /subscriptions/{subscription_id} | Get a subscription by id
DELETE | /subscriptions/{subscription_id} | Delete a subscription
GET    | /users                           | Get all users
POST   | /users                           | Create a user
GET    | /users/{user_id}                 | Get a user by id
DELETE | /users/{user_id}                 | Delete a user
GET    | /users/{user_id}/articles        | Get articles from feeds a user is subscribed to
GET    | /users/{user_id}/feeds           | Get feeds a user is subscribed to

### Entities

#### User

Name and Email are required. Email is a unique field.

```
{
	"id": "59d31403-926a-409b-91fb-8582a933abed",
	"created_at": "2018-06-29T22:30:36.0863Z",
	"updated_at": "2018-06-29T22:30:36.086301Z",
	"name": "Bob Dole",
	"email": "dole@home.place"
}
```

#### Feed

Name is a unique field.

```
{
	"id": "2874cbfd-ec23-48e5-a475-5f3bd4225620",
	"created_at": "2018-06-29T22:30:36.208254Z",
	"updated_at": "2018-06-29T22:30:36.208256Z",
	"name": "Turtles"
}
```

#### Article

Name is a unique field. Feed_id is the ID of the feed the article is attached to.

```
{
	"id": "14e30d66-df0e-405e-8a25-86cdc8889cc3",
	"feed_id": "2874cbfd-ec23-48e5-a475-5f3bd4225620",
	"created_at": "2018-06-29T22:30:36.314212Z",
	"updated_at": "2018-06-29T22:30:36.314213Z",
	"name": "What I Love About Turtles",
	"content": "Turtles are the most exciting ninja warri..."
}
```

#### Subscription

This entity binds a User to a Feed.

```
{
	"id": "14e30d66-df0e-405e-8a25-86cdc8889cc3",
	"feed_id": "2874cbfd-ec23-48e5-a475-5f3bd4225620",
	"user_id": "59d31403-926a-409b-91fb-8582a933abed"
	"created_at": "2018-06-29T22:30:36.314212Z",
	"updated_at": "2018-06-29T22:30:36.314213Z"
}
```

## Prereqs

- Golang
- GoBuffalo (`go get -u -v github.com/gobuffalo/buffalo/buffalo`)
- Dep (https://github.com/golang/dep)

## Running

If you already have access to a running Postgres instance then you can skip this part.

### Deploy a POstgres instance

```
helm install stable/postgresql
```

Note: In order to persist data through restarts of Postgres you will need a PVC or a host volume mount comvined with pod affinity.

### Update database.yaml

To tell the app where to find its database, open database.yaml and configure the `development` section with your postgresql instance data.

eg. If you are running POstgresql on localhost at the default port (5432) with the default user/apss your config would look like this:

```
development:
  dialect: postgres
  database: feedapi_development
  user: postgres
  password: postgres
  host: 127.0.0.1
  pool: 5
```

### Build the binary

```
dep ensure
buffalo build
```

### Test

From the root of the application, run:

```
./bin/feedapi test
```

This will run the test suite. All the happy parths for the HTTP endpoints will be exercised.

Eventually, you should see something like this:

```
dropped database feedapi_test
created database feedapi_test
loaded schema for feedapi_test
INFO[0002] go test -p 1 -tags development github.com/frodopwns/feedapi github.com/frodopwns/feedapi/actions github.com/frodopwns/feedapi/grifts github.com/frodopwns/feedapi/models 
?   	github.com/frodopwns/feedapi	[no test files]
ok  	github.com/frodopwns/feedapi/actions	8.371s
?   	github.com/frodopwns/feedapi/grifts	[no test files]
ok  	github.com/frodopwns/feedapi/models	0.021s
```

### Run

To interact with the API yourself check out the examples in the actions/ test files and the python notebook in examples/. 

To preseed the db with some Users, Articles, and Feeds, run:

```
./bin/feedapi task db:seed
```

Then when you are ready to rock, run:

```
./bin/feedapi migrate up
./bin/feedapi dev
```

This will start the dev server and tell you where to hit the API.