package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Recipe container
type Recipe struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Direction   string             `json:"direction"`
	Ingredients []RecipeIngredient `json:"ingredients"`
}

// RecipeIngredient container
type RecipeIngredient struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
	Note     string  `json:"note"`
}

// RecipeDto container
type RecipeDto struct {
	ID          int64
	Name        string
	Description string
	Direction   string
}

// IngredientDto container
type IngredientDto struct {
	ID   int64
	Name string
}

// RecipeIngredientDto container
type RecipeIngredientDto struct {
	RecipeID     int64
	IngredientID int64
	UnitID       int64
	Quantity     float64
	Note         string
}

// UnitDto container
type UnitDto struct {
	ID   int64
	Name string
}

// Configuration container
type Configuration struct {
	ApiPort          string
	WebPort          string
	DBDriverName     string
	DBDataSourceName string
}

var config Configuration

func getDB() *sql.DB {
	var err error
	db, err := sql.Open(config.DBDriverName, config.DBDataSourceName)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func insertUnit(db *sql.DB, name string) *UnitDto {
	var u UnitDto
	err := db.QueryRow("SELECT id, name FROM unit WHERE lower(name) = lower(?)", name).Scan(&u.ID, &u.Name)
	fmt.Printf("insert unit: %d", err)
	if err == sql.ErrNoRows {
		statement, _ := db.Prepare("INSERT INTO unit (name) VALUES (?)")
		res, _ := statement.Exec(name)
		u.ID, _ = res.LastInsertId()
		u.Name = name
	}

	return &u
}

func insertIngredient(db *sql.DB, name string) *IngredientDto {
	var i IngredientDto
	err := db.QueryRow("SELECT id, name FROM ingredient WHERE lower(name) = lower(?)", name).Scan(&i.ID, &i.Name)
	fmt.Printf("insert ingredient: %d", err)
	if err == sql.ErrNoRows {
		statement, _ := db.Prepare("INSERT INTO ingredient (name) VALUES (?)")
		res, _ := statement.Exec(name)
		i.ID, _ = res.LastInsertId()
		i.Name = name
	}

	return &i
}

func insertRecipe(db *sql.DB, name string, description string, direction string) *RecipeDto {
	var r RecipeDto
	err := db.QueryRow("SELECT id, name, description, direction FROM recipe WHERE lower(name) = lower(?)", name).Scan(&r.ID, &r.Name, &r.Description, &r.Direction)
	fmt.Printf("insert recipe: %d", err)
	if err == sql.ErrNoRows {
		statement, _ := db.Prepare("INSERT INTO recipe (`name`, `description`, `direction`) VALUES (?, ?, ?)")
		res, _ := statement.Exec(name, description, direction)
		r.ID, _ = res.LastInsertId()
		r.Name = name
		r.Description = description
		r.Direction = direction
	}

	return &r
}

func insertRecipeIngredient(db *sql.DB, recipeID int64, ingredientID int64, unitID int64, quantity float64, note string) *RecipeIngredientDto {
	var ri RecipeIngredientDto
	err := db.QueryRow("SELECT recipe_id, ingredient_id, unit_id, quantity, note FROM recipe_ingredient WHERE recipe_id = ? and ingredient_id = ? and unit_id = ?", recipeID, ingredientID, unitID).Scan(&ri.RecipeID, &ri.IngredientID, &ri.UnitID, &ri.Quantity, &ri.Note)
	fmt.Printf("insert recipe  ingredient: %d", err)
	if err == sql.ErrNoRows {
		statement, _ := db.Prepare("INSERT INTO recipe_ingredient (recipe_id, ingredient_id, unit_id, quantity, note) VALUES (?, ?, ?, ?, ?)")
		statement.Exec(recipeID, ingredientID, unitID, quantity, note)
		ri.RecipeID = recipeID
		ri.IngredientID = ingredientID
		ri.UnitID = unitID
		ri.Quantity = quantity
		ri.Note = note
	}

	return &ri
}

func ping(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func search(c *gin.Context) {
	query := c.DefaultQuery("q", "Default")
	ing := c.Query("ingredient")
	fmt.Printf("ing: '%v'\n", ing)
	i := strings.Split(ing, ",")
	fmt.Printf("i '%v' len:'%v'\n", i, len(i))
	ingredients := []string{"$$$", "$$$", "$$$"}
	fmt.Printf("ingredients: '%v'\n", ingredients)
	for k, j := range i {
		if j != "" {
			ingredients[k] = j
		}
	}
	fmt.Printf("ingredients: '%v'\n", ingredients)

	var recipes []Recipe
	db := getDB()
	defer db.Close()
	rows, _ := db.Query(`SELECT r.id, r.name, r.description 
	FROM recipe as r
	WHERE (lower(r.name) like ? or ? = '%') 
	AND (? = '$$$' or ? in (select lower(i.name) from ingredient as i, recipe_ingredient as ri where ri.recipe_id = r.id and i.id = ri.ingredient_id))
	AND (? = '$$$' or ? in (select lower(i.name) from ingredient as i, recipe_ingredient as ri where ri.recipe_id = r.id and i.id = ri.ingredient_id))
	AND (? = '$$$' or ? in (select lower(i.name) from ingredient as i, recipe_ingredient as ri where ri.recipe_id = r.id and i.id = ri.ingredient_id))								
	GROUP BY r.id`,
		strings.ToLower(query)+"%", strings.ToLower(query)+"%", strings.ToLower(ingredients[0]), strings.ToLower(ingredients[0]), strings.ToLower(ingredients[1]), strings.ToLower(ingredients[1]), strings.ToLower(ingredients[2]), strings.ToLower(ingredients[2]))

	for rows.Next() {
		var r Recipe
		rows.Scan(&r.ID, &r.Name, &r.Description)
		recipes = append(recipes, r)
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{
		"query":   query,
		"recipes": recipes,
	})
}

func recipesCount(c *gin.Context) {
	count := 0
	db := getDB()
	defer db.Close()
	db.QueryRow("SELECT count(id) from recipe").Scan(&count)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{
		"count": count,
	})
}

func createRecipe(c *gin.Context) {
	var recipe Recipe
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.BindJSON(&recipe)

	db := getDB()
	defer db.Close()
	r := insertRecipe(db, recipe.Name, recipe.Description, recipe.Direction)
	for _, ingredient := range recipe.Ingredients {
		u := insertUnit(db, ingredient.Unit)
		i := insertIngredient(db, ingredient.Name)
		insertRecipeIngredient(db, r.ID, i.ID, u.ID, ingredient.Quantity, ingredient.Note)
	}

	c.JSON(http.StatusCreated, gin.H{
		"name":        recipe.Name,
		"description": recipe.Description,
		"direction":   recipe.Direction,
	})
}

func recipesController(c *gin.Context) {
	var recipes []Recipe
	db := getDB()
	defer db.Close()
	recipeRows, _ := db.Query("SELECT id, name, description, direction FROM recipe order by id asc")
	for recipeRows.Next() {
		var r Recipe
		recipeRows.Scan(&r.ID, &r.Name, &r.Description, &r.Direction)

		var ingredients []RecipeIngredient
		rows, _ := db.Query("SELECT i.name, ri.quantity, u.name, ri.note FROM ingredient as i, unit as u, recipe_ingredient as ri WHERE ri.recipe_id = ? and i.id = ri.ingredient_id and u.id = ri.unit_id order by u.name desc", r.ID)
		for rows.Next() {
			var ri RecipeIngredient
			rows.Scan(&ri.Name, &ri.Quantity, &ri.Unit, &ri.Note)
			ingredients = append(ingredients, ri)
		}
		r.Ingredients = ingredients

		recipes = append(recipes, r)
	}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, recipes)
}

func recipeController(c *gin.Context) {
	id := c.Param("id")

	var r Recipe
	if id != "" {
		db := getDB()
		defer db.Close()
		err := db.QueryRow("SELECT id, name, description, direction FROM recipe where id = ?", id).Scan(&r.ID, &r.Name, &r.Description, &r.Direction)
		if err != sql.ErrNoRows {
			var ingredients []RecipeIngredient
			rows, _ := db.Query("SELECT i.name, ri.quantity, u.name, ri.note FROM ingredient as i, unit as u, recipe_ingredient as ri WHERE ri.recipe_id = ? and i.id = ri.ingredient_id and u.id = ri.unit_id order by u.name desc", id)
			for rows.Next() {
				var ri RecipeIngredient
				rows.Scan(&ri.Name, &ri.Quantity, &ri.Unit, &ri.Note)
				ingredients = append(ingredients, ri)
			}
			r.Ingredients = ingredients

			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
			c.JSON(http.StatusOK, r)
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
			c.JSON(http.StatusNotFound, struct{}{})
		}
	}

}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}

func main() {
	f, err := os.Open("config.json")
	if err != nil {
		fmt.Printf("error opening config.json: %v", err)
		return
	}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Printf("error decoding config.json: %v", err)
		return
	}

	g := gin.Default()

	g.GET("/ping", ping)
	g.OPTIONS("/ping", preflight)

	g.GET("/search", search)
	g.OPTIONS("/search", preflight)

	g.GET("/recipes/", recipesController)
	g.OPTIONS("/recipes/", preflight)

	g.GET("/recipes/:id", recipeController)
	g.OPTIONS("/recipes/:id", preflight)

	g.GET("/recipes/count", recipesCount)
	g.OPTIONS("/recipes/count", preflight)

	g.POST("/recipes", createRecipe)
	g.OPTIONS("/recipes", preflight)

	go func() {
		http.Handle("/",
			http.StripPrefix("/",
				http.FileServer(http.Dir("./"))))
		log.Fatal(http.ListenAndServe(":"+config.WebPort, nil))
	}()

	g.Run(":" + config.ApiPort)
}
