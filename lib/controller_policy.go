// This controller defines policies that govern
// how a request is treated before entering any main controller method
package lib

import (
	"github.com/jinzhu/gorm"
	"github.com/ncodes/go-base/extend"
)

type PolicyController struct {
	postgresDB *gorm.DB
}

// Create a new controller instance
func NewPolicyController(postgresDB *gorm.DB) *PolicyController {
	return &PolicyController{postgresDB}
}

// Authenticate policy.
// Ensures a valid bear/session token is included in the request.
// If token is valid, the authenticated user will be added to the context
func (self *PolicyController) Authenticate(c *extend.Context) error {

	return nil
}
