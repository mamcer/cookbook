package models

// RecipeDto represents a recipe for database operations
type RecipeDto struct {
	ID          int64  `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Direction   string `db:"direction"`
}

// IngredientDto represents an ingredient for database operations
type IngredientDto struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

// UnitDto represents a unit for database operations
type UnitDto struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

// RecipeIngredientDto represents a recipe ingredient for database operations
type RecipeIngredientDto struct {
	RecipeID     int64   `db:"recipe_id"`
	IngredientID int64   `db:"ingredient_id"`
	UnitID       int64   `db:"unit_id"`
	Quantity     float64 `db:"quantity"`
	Note         string  `db:"note"`
} 