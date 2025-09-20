package providers

import (
	"blog/config"
	authController "blog/modules/auth/controller"
	authRepo "blog/modules/auth/repository"
	authService "blog/modules/auth/service"
	userController "blog/modules/user/controller"
	userRepo "blog/modules/user/repository"
	userService "blog/modules/user/service"
	"blog/pkg/constants"

	"gorm.io/gorm"
	"github.com/samber/do"
)

func InDatabase(injector *do.Injector) {
	do.ProvideNamed(injector, constants.DB, func (i *do.Injector) (*gorm.DB, error) {
		return config.SetUpDatabaseConnection(), nil
	})
}

func RegisterDependencies(injector *do.Injector) {
	InDatabase(injector)

	do.ProvideNamed(injector, constants.JWTService, func(i *do.Injector) (authService.JWTService, error) {
		return authService.NewJWTService(), nil
	})

	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)
	jwtService := do.MustInvokeNamed[authService.JWTService](injector, constants.JWTService)

	userRepository := userRepo.NewUserRepository(db)
	refreshTokenRepository := authRepo.NewRefreshTokenRepository(db)

	userService := userService.NewUserService(userRepository, db)
	authService := authService.NewAuthService(userRepository, refreshTokenRepository, jwtService, db)

	do.Provide(
		injector, func(i *do.Injector) (userController.UserController, error) {
			return userController.NewUserController(i, userService), nil
		},
	)

	do.Provide(
		injector, func(i *do.Injector) (authController.AuthController, error) {
			return authController.NewAuthController(i, authService), nil
		},
	)
}