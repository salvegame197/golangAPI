package create_user

import (
	command "salvegame197/golangApi/src/application/user/command/create_user"
	"salvegame197/golangApi/src/infra/common/util"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	util       *util.Util
}

// New create profile controller instance
func New(
	route *gin.Engine,
	commandBus *command.Bus,
	util *util.Util,
) *Controller {
	controller := &Controller{
		route:      route,
		commandBus: commandBus,
		util:       util,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup profile router
func (controller *Controller) SetupRoutes() {
	controller.route.POST("/register", controller.create)
}
