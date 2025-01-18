package repositories

import (
	"context"
	"database/sql"

	"github.com/mamcer/cookbook/internal/entities"
)

type RecipeRepository struct {
	Conn *sql.DB
}

func NewRecipeRepository(conn *sql.DB) *RecipeRepository {
	return &RecipeRepository{conn}
}

func (rr *RecipeRepository) create(c context.Context, r *entities.Recipe) (err error) {
	// var r RecipeDto
	// err := db.QueryRow("SELECT id, name, description, direction FROM recipe WHERE lower(name) = lower(?)", name).Scan(&r.ID, &r.Name, &r.Description, &r.Direction)
	// fmt.Printf("insert recipe: %d", err)
	// if err == sql.ErrNoRows {
	// 	statement, _ := db.Prepare("INSERT INTO recipe (`name`, `description`, `direction`) VALUES (?, ?, ?)")
	// 	res, _ := statement.Exec(name, description, direction)
	// 	r.ID, _ = res.LastInsertId()
	// 	r.Name = name
	// 	r.Description = description
	// 	r.Direction = direction
	// }
	// return &r

	return
}
