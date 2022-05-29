package login_user

import (
	"net/http"

	login_user_query "salvegame197/golangApi/src/application/user/query/login_user"
	custom_error "salvegame197/golangApi/src/domain/common/custom-error"
	user_errors "salvegame197/golangApi/src/domain/user/errors"

	"github.com/gin-gonic/gin"
)

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (controller *Controller) login(c *gin.Context) {
	var input LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	query, err := login_user_query.NewUserLoginQuery(
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
	user, err := controller.queryBus.Handle(query)

	if err != nil {
		switch err.(type) {
		case *user_errors.IncorrectUserOrPasswordError:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		default:
			httpError := controller.util.Error.HTTP.InternalServerError()
			c.JSON(httpError.Code(), httpError.Message())
		}
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}
