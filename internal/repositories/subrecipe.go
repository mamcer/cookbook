package repositories

import (
	"context"
	"database/sql"

	"github.com/mamcer/cookbook/internal/entities"
)

type SubrecipeRepository struct {
	Conn *sql.DB
}

func NewSubrecipeRepository(conn *sql.DB) *SubrecipeRepository {
	return &SubrecipeRepository{conn}
}

func (rr *SubrecipeRepository) create(c context.Context, r *entities.Subrecipe) (err error) {
	// err := db.QueryRow("SELECT recipe_id, ingredient_id, unit_id, quantity, note FROM recipe_ingredient WHERE recipe_id = ? and ingredient_id = ? and unit_id = ?", r.ID, ingredientID, unitID).Scan(&ri.RecipeID, &ri.IngredientID, &ri.UnitID, &ri.Quantity, &ri.Note)
	// fmt.Printf("insert recipe  ingredient: %d", err)
	// if err == sql.ErrNoRows {
	// 	statement, _ := db.Prepare("INSERT INTO recipe_ingredient (recipe_id, ingredient_id, unit_id, quantity, note) VALUES (?, ?, ?, ?, ?)")
	// 	statement.Exec(recipeID, ingredientID, unitID, quantity, note)
	// 	ri.RecipeID = recipeID
	// 	ri.IngredientID = ingredientID
	// 	ri.UnitID = unitID
	// 	ri.Quantity = quantity
	// 	ri.Note = note
	// }

	// return &ri

	return
}
