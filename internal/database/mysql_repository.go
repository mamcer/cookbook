package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/mamcer/cookbook/internal/models"
)

// MySQLRecipeRepository implements RecipeRepository for MySQL
type MySQLRecipeRepository struct {
	db *DB
}

// MySQLIngredientRepository implements IngredientRepository for MySQL
type MySQLIngredientRepository struct {
	db *DB
}

// MySQLUnitRepository implements UnitRepository for MySQL
type MySQLUnitRepository struct {
	db *DB
}

// MySQLRecipeIngredientRepository implements RecipeIngredientRepository for MySQL
type MySQLRecipeIngredientRepository struct {
	db *DB
}

// NewMySQLRecipeRepository creates a new MySQL recipe repository
func NewMySQLRecipeRepository(db *DB) *MySQLRecipeRepository {
	return &MySQLRecipeRepository{db: db}
}

// NewMySQLIngredientRepository creates a new MySQL ingredient repository
func NewMySQLIngredientRepository(db *DB) *MySQLIngredientRepository {
	return &MySQLIngredientRepository{db: db}
}

// NewMySQLUnitRepository creates a new MySQL unit repository
func NewMySQLUnitRepository(db *DB) *MySQLUnitRepository {
	return &MySQLUnitRepository{db: db}
}

// NewMySQLRecipeIngredientRepository creates a new MySQL recipe ingredient repository
func NewMySQLRecipeIngredientRepository(db *DB) *MySQLRecipeIngredientRepository {
	return &MySQLRecipeIngredientRepository{db: db}
}

// Create implements RecipeRepository.Create
func (r *MySQLRecipeRepository) Create(ctx context.Context, recipe *models.Recipe) error {
	query := "INSERT INTO recipe (name, description, direction) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, recipe.Name, recipe.Description, recipe.Direction)
	if err != nil {
		return fmt.Errorf("failed to create recipe: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	recipe.ID = id
	return nil
}

// GetByID implements RecipeRepository.GetByID
func (r *MySQLRecipeRepository) GetByID(ctx context.Context, id int64) (*models.Recipe, error) {
	query := "SELECT id, name, description, direction FROM recipe WHERE id = ?"
	
	var recipe models.Recipe
	err := r.db.QueryRowContext(ctx, query, id).Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Direction)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get recipe by id: %w", err)
	}

	return &recipe, nil
}

// GetAll implements RecipeRepository.GetAll
func (r *MySQLRecipeRepository) GetAll(ctx context.Context) ([]models.Recipe, error) {
	query := "SELECT id, name, description, direction FROM recipe ORDER BY id ASC"
	
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all recipes: %w", err)
	}
	defer rows.Close()

	var recipes []models.Recipe
	for rows.Next() {
		var recipe models.Recipe
		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description, &recipe.Direction); err != nil {
			return nil, fmt.Errorf("failed to scan recipe: %w", err)
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

// Search implements RecipeRepository.Search
func (r *MySQLRecipeRepository) Search(ctx context.Context, query string, ingredients []string) ([]models.Recipe, error) {
	// Pad ingredients array to 3 elements
	paddedIngredients := make([]string, 3)
	for i := range paddedIngredients {
		paddedIngredients[i] = "$$$"
	}
	for i, ing := range ingredients {
		if i < 3 && ing != "" {
			paddedIngredients[i] = strings.ToLower(ing)
		}
	}

	sqlQuery := `
		SELECT r.id, r.name, r.description 
		FROM recipe as r
		WHERE (lower(r.name) like ? or ? = '%') 
		AND (? = '$$$' or ? in (select lower(i.name) from ingredient as i, recipe_ingredient as ri where ri.recipe_id = r.id and i.id = ri.ingredient_id))
		AND (? = '$$$' or ? in (select lower(i.name) from ingredient as i, recipe_ingredient as ri where ri.recipe_id = r.id and i.id = ri.ingredient_id))
		AND (? = '$$$' or ? in (select lower(i.name) from ingredient as i, recipe_ingredient as ri where ri.recipe_id = r.id and i.id = ri.ingredient_id))								
		GROUP BY r.id`

	searchQuery := strings.ToLower(query) + "%"
	args := []interface{}{
		searchQuery, searchQuery,
		paddedIngredients[0], paddedIngredients[0],
		paddedIngredients[1], paddedIngredients[1],
		paddedIngredients[2], paddedIngredients[2],
	}

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search recipes: %w", err)
	}
	defer rows.Close()

	var recipes []models.Recipe
	for rows.Next() {
		var recipe models.Recipe
		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.Description); err != nil {
			return nil, fmt.Errorf("failed to scan recipe: %w", err)
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

// Count implements RecipeRepository.Count
func (r *MySQLRecipeRepository) Count(ctx context.Context) (int, error) {
	query := "SELECT COUNT(id) FROM recipe"
	
	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count recipes: %w", err)
	}

	return count, nil
}

// GetOrCreate implements IngredientRepository.GetOrCreate
func (r *MySQLIngredientRepository) GetOrCreate(ctx context.Context, name string) (*models.IngredientDto, error) {
	// Try to get existing ingredient
	var ingredient models.IngredientDto
	query := "SELECT id, name FROM ingredient WHERE lower(name) = lower(?)"
	
	err := r.db.QueryRowContext(ctx, query, name).Scan(&ingredient.ID, &ingredient.Name)
	if err == nil {
		return &ingredient, nil
	}
	
	if err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to query ingredient: %w", err)
	}

	// Create new ingredient
	insertQuery := "INSERT INTO ingredient (name) VALUES (?)"
	result, err := r.db.ExecContext(ctx, insertQuery, name)
	if err != nil {
		return nil, fmt.Errorf("failed to create ingredient: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id for ingredient: %w", err)
	}

	return &models.IngredientDto{ID: id, Name: name}, nil
}

// GetOrCreate implements UnitRepository.GetOrCreate
func (r *MySQLUnitRepository) GetOrCreate(ctx context.Context, name string) (*models.UnitDto, error) {
	// Try to get existing unit
	var unit models.UnitDto
	query := "SELECT id, name FROM unit WHERE lower(name) = lower(?)"
	
	err := r.db.QueryRowContext(ctx, query, name).Scan(&unit.ID, &unit.Name)
	if err == nil {
		return &unit, nil
	}
	
	if err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to query unit: %w", err)
	}

	// Create new unit
	insertQuery := "INSERT INTO unit (name) VALUES (?)"
	result, err := r.db.ExecContext(ctx, insertQuery, name)
	if err != nil {
		return nil, fmt.Errorf("failed to create unit: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert id for unit: %w", err)
	}

	return &models.UnitDto{ID: id, Name: name}, nil
}

// Create implements RecipeIngredientRepository.Create
func (r *MySQLRecipeIngredientRepository) Create(ctx context.Context, recipeID int64, ingredientID int64, unitID int64, quantity float64, note string) error {
	query := "INSERT INTO recipe_ingredient (recipe_id, ingredient_id, unit_id, quantity, note) VALUES (?, ?, ?, ?, ?)"
	
	_, err := r.db.ExecContext(ctx, query, recipeID, ingredientID, unitID, quantity, note)
	if err != nil {
		return fmt.Errorf("failed to create recipe ingredient: %w", err)
	}

	return nil
}

// GetByRecipeID implements RecipeIngredientRepository.GetByRecipeID
func (r *MySQLRecipeIngredientRepository) GetByRecipeID(ctx context.Context, recipeID int64) ([]models.RecipeIngredient, error) {
	query := `
		SELECT i.name, ri.quantity, u.name, ri.note 
		FROM ingredient as i, unit as u, recipe_ingredient as ri 
		WHERE ri.recipe_id = ? and i.id = ri.ingredient_id and u.id = ri.unit_id 
		ORDER BY u.name DESC`

	rows, err := r.db.QueryContext(ctx, query, recipeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get recipe ingredients: %w", err)
	}
	defer rows.Close()

	var ingredients []models.RecipeIngredient
	for rows.Next() {
		var ingredient models.RecipeIngredient
		if err := rows.Scan(&ingredient.Name, &ingredient.Quantity, &ingredient.Unit, &ingredient.Note); err != nil {
			return nil, fmt.Errorf("failed to scan recipe ingredient: %w", err)
		}
		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
} 