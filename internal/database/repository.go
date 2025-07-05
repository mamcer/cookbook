package database

import (
	"context"
	"github.com/mamcer/cookbook/internal/models"
)

// RecipeRepository defines the interface for recipe database operations
type RecipeRepository interface {
	Create(ctx context.Context, recipe *models.Recipe) error
	GetByID(ctx context.Context, id int64) (*models.Recipe, error)
	GetAll(ctx context.Context) ([]models.Recipe, error)
	Search(ctx context.Context, query string, ingredients []string) ([]models.Recipe, error)
	Count(ctx context.Context) (int, error)
}

// IngredientRepository defines the interface for ingredient database operations
type IngredientRepository interface {
	GetOrCreate(ctx context.Context, name string) (*models.IngredientDto, error)
}

// UnitRepository defines the interface for unit database operations
type UnitRepository interface {
	GetOrCreate(ctx context.Context, name string) (*models.UnitDto, error)
}

// RecipeIngredientRepository defines the interface for recipe ingredient database operations
type RecipeIngredientRepository interface {
	Create(ctx context.Context, recipeID int64, ingredientID int64, unitID int64, quantity float64, note string) error
	GetByRecipeID(ctx context.Context, recipeID int64) ([]models.RecipeIngredient, error)
} 