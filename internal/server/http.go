package server

import (
	"ginblog/internal/handler"
	"ginblog/internal/middleware"
	"ginblog/pkg/helper/resp"
	"ginblog/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewServerHTTP(
	logger *log.Logger,
	userHandler handler.UserHandler,
	articleHandler handler.ArticleHandler,
	categoryHandler handler.CategoryHandler,
	commentHandler handler.CommentHandler,
	profileHandler handler.ProfileHandler,
	loginHandler handler.LoginHandler,
) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.GET("/", func(ctx *gin.Context) {
		resp.HandleSuccess(ctx, map[string]interface{}{
			"say": "Hi Nunu!",
		})
	})
	r.GET("/user", userHandler.GetUserById)

	auth := r.Group("/api/v1")

	// 后台管理路由接口
	var conf *viper.Viper
	auth.Use(middleware.JwtToken(conf))
	{
		// 用户模块的路由接口
		auth.GET("admin/users", userHandler.GetUsers)
		auth.PUT("user/:id", userHandler.EditUser)
		auth.DELETE("user/:id", userHandler.DeleteUser)
		// 修改密码
		auth.PUT("admin/changepw/:id", userHandler.ChangeUserPassword)

		// 分类模块的路由接口
		auth.GET("admin/category", categoryHandler.GetCate)
		auth.POST("category/add", categoryHandler.AddCategory)
		auth.PUT("category/:id", categoryHandler.EditCate)
		auth.DELETE("category/:id", categoryHandler.DeleteCate)

		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", articleHandler.GetArtInfo)
		auth.GET("admin/article", articleHandler.GetArt)
		auth.POST("article/add", articleHandler.AddArticle)
		auth.PUT("article/:id", articleHandler.EditArt)
		auth.DELETE("article/:id", articleHandler.DeleteArt)

		// 更新个人设置
		auth.GET("admin/profile/:id", profileHandler.GetProfile)
		auth.PUT("profile/:id", profileHandler.UpdateProfile)

		// 评论模块
		auth.GET("comment/list", commentHandler.GetCommentList)
		auth.DELETE("delcomment/:id", commentHandler.DeleteComment)
		auth.PUT("checkcomment/:id", commentHandler.CheckComment)
		auth.PUT("uncheckcomment/:id", commentHandler.UncheckComment)

	}

	router := r.Group("/api/v1")
	{

		// 用户信息模块
		router.POST("user/add", userHandler.AddUser)
		router.GET("user/:id", userHandler.GetUserInfo)
		router.GET("users", userHandler.GetUsers)

		// 文章分类信息模块
		router.GET("category", categoryHandler.GetCate)
		router.GET("category/:id", categoryHandler.GetCateInfo)

		// 文章模块路由接口
		router.GET("article", articleHandler.GetArt)
		router.GET("article/list/:id", articleHandler.GetCateArt)
		router.GET("article/info/:id", articleHandler.GetArtInfo)

		// 登录控制
		router.POST("login", loginHandler.Login)
		router.POST("loginfront", loginHandler.LoginFront)

		// 评论模块
		router.POST("addcomment", commentHandler.AddComment)
		router.GET("comment/info/:id", commentHandler.GetComment)
		router.GET("commentfront/:id", commentHandler.GetCommentListFront)
		router.GET("commentcount/:id", commentHandler.GetCommentCont)
	}

	return r
}
