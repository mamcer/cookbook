package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/mamcer/cookbook/internal/database"
	"github.com/mamcer/cookbook/internal/models"
)

// RecipeService handles business logic for recipes
type RecipeService struct {
	recipeRepo           database.RecipeRepository
	ingredientRepo       database.IngredientRepository
	unitRepo             database.UnitRepository
	recipeIngredientRepo database.RecipeIngredientRepository
}

// NewRecipeService creates a new recipe service
func NewRecipeService(
	recipeRepo database.RecipeRepository,
	ingredientRepo database.IngredientRepository,
	unitRepo database.UnitRepository,
	recipeIngredientRepo database.RecipeIngredientRepository,
) *RecipeService {
	return &RecipeService{
		recipeRepo:           recipeRepo,
		ingredientRepo:       ingredientRepo,
		unitRepo:             unitRepo,
		recipeIngredientRepo: recipeIngredientRepo,
	}
}

// CreateRecipe creates a new recipe with its ingredients
func (s *RecipeService) CreateRecipe(ctx context.Context, recipe *models.Recipe) error {
	// Create the recipe
	if err := s.recipeRepo.Create(ctx, recipe); err != nil {
		return fmt.Errorf("failed to create recipe: %w", err)
	}

	// Create ingredients and recipe-ingredient relationships
	for _, ingredient := range recipe.Ingredients {
		// Get or create ingredient
		ingredientDto, err := s.ingredientRepo.GetOrCreate(ctx, ingredient.Name)
		if err != nil {
			return fmt.Errorf("failed to get or create ingredient: %w", err)
		}

		// Get or create unit
		unitDto, err := s.unitRepo.GetOrCreate(ctx, ingredient.Unit)
		if err != nil {
			return fmt.Errorf("failed to get or create unit: %w", err)
		}

		// Create recipe-ingredient relationship
		if err := s.recipeIngredientRepo.Create(ctx, recipe.ID, ingredientDto.ID, unitDto.ID, ingredient.Quantity, ingredient.Note); err != nil {
			return fmt.Errorf("failed to create recipe ingredient: %w", err)
		}
	}

	return nil
}

// GetRecipe retrieves a recipe by ID with its ingredients
func (s *RecipeService) GetRecipe(ctx context.Context, id int64) (*models.Recipe, error) {
	recipe, err := s.recipeRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get recipe: %w", err)
	}

	if recipe == nil {
		return nil, nil
	}

	// Get ingredients for the recipe
	ingredients, err := s.recipeIngredientRepo.GetByRecipeID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get recipe ingredients: %w", err)
	}

	recipe.Ingredients = ingredients
	return recipe, nil
}

// GetAllRecipes retrieves all recipes with their ingredients
func (s *RecipeService) GetAllRecipes(ctx context.Context) ([]models.Recipe, error) {
	recipes, err := s.recipeRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all recipes: %w", err)
	}

	// Get ingredients for each recipe
	for i := range recipes {
		ingredients, err := s.recipeIngredientRepo.GetByRecipeID(ctx, recipes[i].ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get ingredients for recipe %d: %w", recipes[i].ID, err)
		}
		recipes[i].Ingredients = ingredients
	}

	return recipes, nil
}

// SearchRecipes searches for recipes by name and ingredients
func (s *RecipeService) SearchRecipes(ctx context.Context, query string, ingredients []string) ([]models.Recipe, error) {
	// Filter out empty ingredients
	filteredIngredients := make([]string, 0)
	for _, ing := range ingredients {
		if strings.TrimSpace(ing) != "" {
			filteredIngredients = append(filteredIngredients, strings.TrimSpace(ing))
		}
	}

	recipes, err := s.recipeRepo.Search(ctx, query, filteredIngredients)
	if err != nil {
		return nil, fmt.Errorf("failed to search recipes: %w", err)
	}

	// Get ingredients for each recipe
	for i := range recipes {
		ingredients, err := s.recipeIngredientRepo.GetByRecipeID(ctx, recipes[i].ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get ingredients for recipe %d: %w", recipes[i].ID, err)
		}
		recipes[i].Ingredients = ingredients
	}

	return recipes, nil
}

// GetRecipeCount returns the total number of recipes
func (s *RecipeService) GetRecipeCount(ctx context.Context) (int, error) {
	count, err := s.recipeRepo.Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get recipe count: %w", err)
	}

	return count, nil
} 