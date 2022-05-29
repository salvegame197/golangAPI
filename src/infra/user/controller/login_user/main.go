package login_user

import (
	query "salvegame197/golangApi/src/application/user/query/login_user"
	"salvegame197/golangApi/src/infra/common/util"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	route    *gin.Engine
	queryBus *query.Bus
	util     *util.Util
}

// New create profile controller instance
func New(route *gin.Engine, queryBus *query.Bus, util *util.Util) *Controller {
	controller := &Controller{
		route:    route,
		queryBus: queryBus,
		util:     util,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup profile router
func (controller *Controller) SetupRoutes() {
	controller.route.POST("/login", controller.login)
}
