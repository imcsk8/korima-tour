package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/imcsk8/korima-tour/korima_hishow/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
