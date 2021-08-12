package grifts

import (
	"github.com/diszy10/gorengan/mendoan/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
