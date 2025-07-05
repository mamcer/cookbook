package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mamcer/cookbook/internal/models"
)

// RecipeHandler handles HTTP requests for recipes
type RecipeHandler struct {
	service *RecipeService
}

// NewRecipeHandler creates a new recipe handler
func NewRecipeHandler(service *RecipeService) *RecipeHandler {
	return &RecipeHandler{service: service}
}

// Ping handles the ping endpoint
func (h *RecipeHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Search handles recipe search requests
func (h *RecipeHandler) Search(c *gin.Context) {
	query := c.DefaultQuery("q", "")
	ingredientParam := c.Query("ingredient")

	// Parse ingredients from comma-separated string
	var ingredients []string
	if ingredientParam != "" {
		ingredients = strings.Split(ingredientParam, ",")
	}

	recipes, err := h.service.SearchRecipes(c.Request.Context(), query, ingredients)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to search recipes",
		})
		return
	}

	response := models.RecipeSearchResponse{
		Query:   query,
		Recipes: recipes,
		Count:   len(recipes),
	}

	c.JSON(http.StatusOK, response)
}

// GetRecipes handles getting all recipes
func (h *RecipeHandler) GetRecipes(c *gin.Context) {
	recipes, err := h.service.GetAllRecipes(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get recipes",
		})
		return
	}

	c.JSON(http.StatusOK, recipes)
}

// GetRecipe handles getting a single recipe by ID
func (h *RecipeHandler) GetRecipe(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Recipe ID is required",
		})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid recipe ID",
		})
		return
	}

	recipe, err := h.service.GetRecipe(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get recipe",
		})
		return
	}

	if recipe == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found",
		})
		return
	}

	c.JSON(http.StatusOK, recipe)
}

// CreateRecipe handles creating a new recipe
func (h *RecipeHandler) CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Basic validation
	if recipe.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Recipe name is required",
		})
		return
	}

	if recipe.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Recipe description is required",
		})
		return
	}

	if recipe.Direction == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Recipe direction is required",
		})
		return
	}

	err := h.service.CreateRecipe(c.Request.Context(), &recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create recipe",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          recipe.ID,
		"name":        recipe.Name,
		"description": recipe.Description,
		"direction":   recipe.Direction,
	})
}

// GetRecipeCount handles getting the total number of recipes
func (h *RecipeHandler) GetRecipeCount(c *gin.Context) {
	count, err := h.service.GetRecipeCount(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get recipe count",
		})
		return
	}

	response := models.RecipeCountResponse{Count: count}
	c.JSON(http.StatusOK, response)
} 