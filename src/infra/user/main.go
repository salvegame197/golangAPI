package user

import (
	create_user_command "salvegame197/golangApi/src/application/user/command/create_user"
	login_user_query "salvegame197/golangApi/src/application/user/query/login_user"
	create_user_service "salvegame197/golangApi/src/domain/user/service/create_user"
	get_user_by_email_service "salvegame197/golangApi/src/domain/user/service/get_user_by_email"
	login_user_service "salvegame197/golangApi/src/domain/user/service/login_user"
	auth_service "salvegame197/golangApi/src/infra/common/services/auth"
	security_service "salvegame197/golangApi/src/infra/common/services/security"
	"salvegame197/golangApi/src/infra/common/util"
	create_user_controller "salvegame197/golangApi/src/infra/user/controller/create_user"
	login_user_controller "salvegame197/golangApi/src/infra/user/controller/login_user"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	user_repository "salvegame197/golangApi/src/infra/user/repository"

	"gorm.io/gorm"
)

// Initialize initialize profile module
func Initialize(
	engine *gin.Engine, util *util.Util, postgressClient *gorm.DB, redisClient *redis.Client,
) {
	redisService := auth_service.NewRedisService(redisClient)

	tokenService := auth_service.NewTokenService()

	securityService := security_service.New()

	//**********REGISTER USER*************//
	userRepository := user_repository.New(postgressClient)
	userServiceFinder := get_user_by_email_service.New(userRepository)
	createUserService := create_user_service.New(userRepository, userServiceFinder)
	createUserCommandBus := create_user_command.New(createUserService, securityService)
	create_user_controller.New(engine, createUserCommandBus, util)

	//**********LOGIN USER*************//
	loginUserService := login_user_service.New(userServiceFinder, securityService, redisService.Auth, tokenService)
	loginCreateUserQueryBus := login_user_query.New(loginUserService)
	login_user_controller.New(engine, loginCreateUserQueryBus, util)

}
