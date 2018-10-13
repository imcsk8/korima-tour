package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/imcsk8/korima-tour/korima/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
