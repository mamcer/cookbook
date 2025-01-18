package repositories

import (
	"context"
	"database/sql"

	"github.com/mamcer/cookbook/internal/entities"
)

type UnitRepository struct {
	Conn *sql.DB
}

func NewUnitRepository(conn *sql.DB) *UnitRepository {
	return &UnitRepository{conn}
}

func (rr *UnitRepository) GetByName(c context.Context, n string) (r *entities.Unit, err error) {

	// var u UnitDto
	// err := db.QueryRow("SELECT id, name FROM unit WHERE lower(name) = lower(?)", name).Scan(&u.ID, &u.Name)
	// fmt.Printf("insert unit: %d", err)
	// if err == sql.ErrNoRows {
	// 	statement, _ := db.Prepare("INSERT INTO unit (name) VALUES (?)")
	// 	res, _ := statement.Exec(name)
	// 	u.ID, _ = res.LastInsertId()
	// 	u.Name = name
	// }

	// return &u

	return &entities.Unit{}, nil
}
