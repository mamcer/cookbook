package repositories

import (
	"context"
	"database/sql"

	"github.com/mamcer/cookbook/internal/entities"
)

type IngredientRepository struct {
	Conn *sql.DB
}

func NewIngredientRepository(conn *sql.DB) *IngredientRepository {
	return &IngredientRepository{conn}
}

func (rr *IngredientRepository) create(c context.Context, r *entities.Ingredient) (err error) {
	// var i IngredientDto
	// err := db.QueryRow("SELECT id, name FROM ingredient WHERE lower(name) = lower(?)", name).Scan(&i.ID, &i.Name)
	// fmt.Printf("insert ingredient: %d", err)
	// if err == sql.ErrNoRows {
	// 	statement, _ := db.Prepare("INSERT INTO ingredient (name) VALUES (?)")
	// 	res, _ := statement.Exec(name)
	// 	i.ID, _ = res.LastInsertId()
	// 	i.Name = name
	// }

	// return &i

	// return &ri
	return
}

func (rr *IngredientRepository) getByID(c context.Context, id int) (i *entities.Ingredient, err error) {
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

	return &entities.Ingredient{1, "papa"}, nil
}
