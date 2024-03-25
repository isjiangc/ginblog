//go:build wireinject
// +build wireinject

package main

import (
	"ginblog/internal/handler"
	"ginblog/internal/middleware"
	"ginblog/internal/repository"
	"ginblog/internal/server"
	"ginblog/internal/service"
	"ginblog/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var RepositorySet = wire.NewSet(
	repository.NewDb,
	repository.NewRepository,
	repository.NewUserRepository,
	repository.NewArticleRepository,
	repository.NewCategoryRepository,
	repository.NewCommentRepository,
	repository.NewProfileRepository,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewArticleService,
	service.NewCategoryService,
	service.NewCommentService,
	service.NewProfileService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewArticleHandler,
	handler.NewCategoryHandler,
	handler.NewCommentHandler,
	handler.NewProfileHandler,
	handler.NewLoginHandler,
)

func newApp(*viper.Viper, *log.Logger, *middleware.JWT) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	))
}
