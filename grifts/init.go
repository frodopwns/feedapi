package grifts

import (
	"github.com/frodopwns/feedapi/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
