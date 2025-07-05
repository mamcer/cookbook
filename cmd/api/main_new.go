package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mamcer/cookbook/internal/config"
	"github.com/mamcer/cookbook/internal/database"
	"github.com/mamcer/cookbook/internal/handlers"
	"github.com/mamcer/cookbook/internal/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	db, err := database.NewConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	recipeRepo := database.NewMySQLRecipeRepository(db)
	ingredientRepo := database.NewMySQLIngredientRepository(db)
	unitRepo := database.NewMySQLUnitRepository(db)
	recipeIngredientRepo := database.NewMySQLRecipeIngredientRepository(db)

	// Initialize service
	recipeService := handlers.NewRecipeService(recipeRepo, ingredientRepo, unitRepo, recipeIngredientRepo)

	// Initialize handlers
	recipeHandler := handlers.NewRecipeHandler(recipeService)

	// Set up Gin router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Add middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.ErrorLoggingMiddleware())

	// API routes
	api := router.Group("/api/v1")
	{
		api.GET("/ping", recipeHandler.Ping)
		api.GET("/search", recipeHandler.Search)
		api.GET("/recipes", recipeHandler.GetRecipes)
		api.GET("/recipes/:id", recipeHandler.GetRecipe)
		api.GET("/recipes/count", recipeHandler.GetRecipeCount)
		api.POST("/recipes", recipeHandler.CreateRecipe)
	}

	// Legacy routes for backward compatibility
	router.GET("/ping", recipeHandler.Ping)
	router.GET("/search", recipeHandler.Search)
	router.GET("/recipes/", recipeHandler.GetRecipes)
	router.GET("/recipes/:id", recipeHandler.GetRecipe)
	router.GET("/recipes/count", recipeHandler.GetRecipeCount)
	router.POST("/recipes", recipeHandler.CreateRecipe)

	// Start web server in a goroutine
	go func() {
		webServer := &http.Server{
			Addr:    ":" + cfg.Server.WebPort,
			Handler: http.FileServer(http.Dir(".")),
		}

		log.Printf("Starting web server on port %s", cfg.Server.WebPort)
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start web server: %v", err)
		}
	}()

	// Start API server
	apiServer := &http.Server{
		Addr:    ":" + cfg.Server.APIPort,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting API server on port %s", cfg.Server.APIPort)
		if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start API server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down servers...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown API server
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Fatal("API server forced to shutdown:", err)
	}

	log.Println("Servers exited")
} 