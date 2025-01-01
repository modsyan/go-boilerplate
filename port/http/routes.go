package http

import (
	"company-name/port/http/handlers"
	"company-name/port/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine               *gin.Engine
	authHandler          *handlers.AuthHandler
	contentBlocksHandler *handlers.ContentBlocksHandler
	userHandler          *handlers.UserHandler
	fileHandler          *handlers.FileHandler
}

func NewRouter(
	engine *gin.Engine,
	authHandler *handlers.AuthHandler,
	contentBlocksHandler *handlers.ContentBlocksHandler,
	userHandler *handlers.UserHandler,
	fileHandler *handlers.FileHandler,
) *Router {
	return &Router{
		engine:               engine,
		authHandler:          authHandler,
		contentBlocksHandler: contentBlocksHandler,
		userHandler:          userHandler,
		fileHandler:          fileHandler,
	}
}

func (r *Router) RegisterRoutes() error {
	api := r.engine.Group("/api/v1")
	api.Use(middleware.LocalizationMiddleware)
	r.registerAuthRoutes(api)
	r.registerContentBlocksRoutes(api)
	r.registerUsersRoutes(api)
	r.registerPaymentRoutes(api)
	r.registerFilesRoutes(api)

	return nil
}

func (r *Router) registerAuthRoutes(api *gin.RouterGroup) {
	authRoutes := api.Group("/auth")

	authRoutes.POST("/login", r.authHandler.Login)
	authRoutes.POST("/register", r.authHandler.Register)
	authRoutes.GET("/verify-email", r.authHandler.VerifyEmail)
}

func (r *Router) registerContentBlocksRoutes(api *gin.RouterGroup) {
	blocksRoutes := api.Group("/blocks")
	blocksRoutes.GET("/page/:name", r.contentBlocksHandler.GetPageContentBlocks)
	blocksRoutes.GET("/", r.contentBlocksHandler.GetContentBlock)
	blocksRoutes.POST("/", r.contentBlocksHandler.CreateContentBlock)
	blocksRoutes.PUT("/", r.contentBlocksHandler.UpdateContentBlock)
	blocksRoutes.DELETE("/", r.contentBlocksHandler.DeleteContentBlock)
}

func (r *Router) registerUsersRoutes(api *gin.RouterGroup) {
	userRoutes := api.Group("/users")
	userRoutes.GET("/", r.userHandler.GetAllUsers)
	userRoutes.GET("/:id", r.userHandler.GetDetailsUserByID)
	userRoutes.POST("/", r.userHandler.CreateUser)
	userRoutes.PUT("/:id", r.userHandler.UpdateUser)
	userRoutes.DELETE("/:id", r.userHandler.DeleteUser)
}

func (r *Router) registerFilesRoutes(api *gin.RouterGroup) {
	fileRoutes := api.Group("/files")
	fileRoutes.POST("/", r.fileHandler.UploadFile)
}

func (r *Router) registerPaymentRoutes(api *gin.RouterGroup) {
}
