package models

import "time"

// Recipe represents a recipe in the system
type Recipe struct {
	ID          int64              `json:"id" db:"id"`
	Name        string             `json:"name" db:"name" validate:"required"`
	Description string             `json:"description" db:"description" validate:"required"`
	Direction   string             `json:"direction" db:"direction" validate:"required"`
	Ingredients []RecipeIngredient `json:"ingredients,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" db:"updated_at"`
}

// RecipeIngredient represents an ingredient in a recipe
type RecipeIngredient struct {
	Name     string  `json:"name" validate:"required"`
	Quantity float64 `json:"quantity" validate:"required,gt=0"`
	Unit     string  `json:"unit" validate:"required"`
	Note     string  `json:"note"`
}

// RecipeSearchRequest represents a search request for recipes
type RecipeSearchRequest struct {
	Query       string   `json:"query" form:"q"`
	Ingredients []string `json:"ingredients" form:"ingredient"`
}

// RecipeSearchResponse represents a search response
type RecipeSearchResponse struct {
	Query   string   `json:"query"`
	Recipes []Recipe `json:"recipes"`
	Count   int      `json:"count"`
}

// RecipeCountResponse represents a count response
type RecipeCountResponse struct {
	Count int `json:"count"`
} 