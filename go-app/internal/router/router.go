package router

import (
	"next-go-task/internal/config"
	"next-go-task/internal/handler"
	"next-go-task/internal/service"
	"next-go-task/pkg/graph/generated"
	"next-go-task/pkg/graph/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Setup configures and returns the Echo router with REST and GraphQL support
func Setup(cfg *config.Config, taskService *service.TaskService) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{cfg.FrontURL},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// REST API Handlers
	healthHandler := handler.NewHealthHandler()
	taskHandler := handler.NewTaskHandler()

	// REST API Routes
	api := e.Group("/api")
	{
		api.GET("/health", healthHandler.Check)
		api.GET("/tasks", taskHandler.GetAll)
		api.POST("/tasks", taskHandler.Create)
	}

	// GraphQL Setup
	gqlResolver := resolver.NewResolver(taskService)
	gqlServer := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: gqlResolver,
	}))

	// GraphQL Routes
	e.POST("/query", func(c echo.Context) error {
		gqlServer.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// GraphQL Playground (development only)
	playgroundHandler := playground.Handler("GraphQL Playground", "/query")
	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// Legacy REST endpoints (for backward compatibility)
	e.GET("/health", healthHandler.Check)
	e.GET("/tasks", taskHandler.GetAll)
	e.POST("/tasks", taskHandler.Create)

	return e
}
