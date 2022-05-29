package create_user

import (
	"net/http"

	create_user_command "salvegame197/golangApi/src/application/user/command/create_user"
	custom_error "salvegame197/golangApi/src/domain/common/custom-error"
	user_errors "salvegame197/golangApi/src/domain/user/errors"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func (controller *Controller) create(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	command, err := create_user_command.NewCreateUserCommand(
		input.FirstName,
		input.LastName,
		input.Email,
		input.Password,
	)

	if err != nil {
		switch err.(type) {
		case *custom_error.InvalidArgumentError:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		default:
			httpError := controller.util.Error.HTTP.InternalServerError()
			c.JSON(httpError.Code(), httpError.Message())
		}
		return
	}
	err = controller.commandBus.Handle(command)
	if err != nil {
		switch err.(type) {
		case *user_errors.AlreadyExistsError:
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
		default:
			httpError := controller.util.Error.HTTP.InternalServerError()
			c.JSON(httpError.Code(), httpError.Message())
		}
		return
	}
	c.JSON(http.StatusCreated, nil)
}
