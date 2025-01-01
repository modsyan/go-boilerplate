package api

import (
	"company-name/configs"
	"company-name/constants"
	"company-name/internal/auth"
	"company-name/internal/content-blocks"
	"company-name/internal/user"
	"company-name/pkg/database"
	"company-name/pkg/email"
	"company-name/pkg/file"
	"company-name/pkg/validators"
	"company-name/port/http"
	"company-name/port/http/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

type APIServer struct {
	engine       *gin.Engine
	config       *configs.Config
	validator    validators.IValidator
	emailService email.IEmailService
	db           database.IDatabase
}

func NewAPIServer(
	db database.IDatabase,
	cfg *configs.Config,
	emailService email.IEmailService,
	validator validators.IValidator,
) *APIServer {
	return &APIServer{
		engine:       gin.Default(),
		db:           db,
		emailService: emailService,
		config:       cfg,
		validator:    validator,
	}
}

func (s APIServer) Run() error {
	_ = s.engine.SetTrustedProxies(nil)
	if s.config.App.Environment == constants.EnvironmentProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := s.RegisterRoutes(); err != nil {
		log.Fatalf("Error registering routes: %v", err)
	}

	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := s.engine.Run(":" + s.config.App.Port); err != nil {
		log.Fatalf("Running Http Server: %v", err)
	}

	return nil
}

func (s APIServer) GetEngine() *gin.Engine {
	return s.engine
}

func (s APIServer) RegisterRoutes() error {
	// Initialize repositories
	authRepo := auth.NewAuthRepository(s.db)
	userRepo := user.NewUserRepository(s.db)
	contentRepo := blocks.NewContentBlockRepository(s.db)
	userRepo = user.NewUserRepository(s.db)

	// Initialize services
	authService := auth.NewAuthService(authRepo, s.config, s.validator, s.emailService)
	contentBlocksService := blocks.NewContentBlocksService(contentRepo, s.validator)
	userService := user.NewUserService(userRepo, s.validator)
	fileService := file.NewFileService(s.config.FileStorage.Directory)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService, s.validator)
	contentBlocksHandler := handlers.NewContentBlocksHandler(contentBlocksService, s.validator)
	userHandler := handlers.NewUserHandler(userService, s.validator)
	fileHandler := handlers.NewFileHandler(fileService)

	router := http.NewRouter(
		s.engine,
		authHandler,
		contentBlocksHandler,
		userHandler,
		fileHandler,
	)

	router.RegisterRoutes()

	return nil
}
