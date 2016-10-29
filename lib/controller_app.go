package lib

import (
	"github.com/ncodes/go-base/extend"
)

type AppController struct {
}

// Create a new controller instance
func NewAppController() *AppController {
	return &AppController{}
}

func (ac *AppController) Index(c *extend.Context) error {
	return c.String(200, "Hello World!")
}
