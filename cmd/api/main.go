package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mamcer/cookbook/internal/handlers"
	"github.com/mamcer/cookbook/internal/services"
)

func getDB(driver, source string) *sql.DB {
	var err error
	db, err := sql.Open(driver, source)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, struct{}{})
}

func main() {
	apiPort := os.Getenv("SPELLBOOK_API_PORT")
	webPort := os.Getenv("SPELLBOOK_WEB_PORT")
	dbDriver := os.Getenv("SPELLBOOK_DB_DRIVER")
	dbUser := os.Getenv("SPELLBOOK_DB_USER")
	dbPass := os.Getenv("SPELLBOOK_DB_PASS")
	dbHost := os.Getenv("SPELLBOOK_DB_HOST")
	dbPort := os.Getenv("SPELLBOOK_DB_PORT")
	dbName := os.Getenv("SPELLBOOK_DB_NAME")

	db := getDB(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))

	err := db.Ping()
	if err != nil {
		fmt.Printf("there is an error in the db connection %s", err)
		panic(err)
	}

	g := gin.Default()

	handlers.NewPingHandler(g, &services.PingService{Message: "pong"})

	go func() {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}

		http.Handle("/",
			http.StripPrefix("/",
				http.FileServer(http.Dir(filepath.Dir(ex)))))
		log.Fatal(http.ListenAndServe(":"+webPort, nil))
	}()

	g.Run(":" + apiPort)
}

// func ping(c *gin.Context) {
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

// 	c.JSON(200, gin.H{
// 		"message": "pong",
// 	})
// }

// func search(c *gin.Context) {
// 	query := c.DefaultQuery("q", "Default")
// 	ing := c.Query("ingredient")
// 	fmt.Printf("ing: '%v'\n", ing)
// 	i := strings.Split(ing, ",")
// 	fmt.Printf("i '%v' len:'%v'\n", i, len(i))
// 	ingredients := []string{"$$$", "$$$", "$$$"}
// 	fmt.Printf("ingredients: '%v'\n", ingredients)
// 	for k, j := range i {
// 		if j != "" {
// 			ingredients[k] = j
// 		}
// 	}
// 	fmt.Printf("ingredients: '%v'\n", ingredients)

// 	var recipes []Recipe
// 	db := getDB()
// 	defer db.Close()
// 	rows, _ := db.Query(`SELECT r.id, r.name, r.description
// 	FROM recipe as r
// 	WHERE (lower(r.name) like ? or ? = '%')
// 	AND (? = '$$$' or ? in (select lower(i.name) from ingredient as i, recipe_ingredient as ri where ri.recipe_id = r.id and i.id = ri.ingredient_id))
// 	AND (? = '$$$' or ? in (select lower(i.name) from ingredient as i, recipe_ingredient as ri where ri.recipe_id = r.id and i.id = ri.ingredient_id))
// 	AND (? = '$$$' or ? in (select lower(i.name) from ingredient as i, recipe_ingredient as ri where ri.recipe_id = r.id and i.id = ri.ingredient_id))
// 	GROUP BY r.id`,
// 		strings.ToLower(query)+"%", strings.ToLower(query)+"%", strings.ToLower(ingredients[0]), strings.ToLower(ingredients[0]), strings.ToLower(ingredients[1]), strings.ToLower(ingredients[1]), strings.ToLower(ingredients[2]), strings.ToLower(ingredients[2]))

// 	for rows.Next() {
// 		var r Recipe
// 		rows.Scan(&r.ID, &r.Name, &r.Description)
// 		recipes = append(recipes, r)
// 	}

// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
// 	c.JSON(http.StatusOK, gin.H{
// 		"query":   query,
// 		"recipes": recipes,
// 	})
// }

// func recipesCount(c *gin.Context) {
// 	count := 0
// 	db := getDB()
// 	defer db.Close()
// 	db.QueryRow("SELECT count(id) from recipe").Scan(&count)

// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
// 	c.JSON(http.StatusOK, gin.H{
// 		"count": count,
// 	})
// }

// func createRecipe(c *gin.Context) {
// 	var recipe Recipe
// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
// 	c.BindJSON(&recipe)

// 	db := getDB()
// 	defer db.Close()
// 	r := insertRecipe(db, recipe.Name, recipe.Description, recipe.Direction)
// 	for _, ingredient := range recipe.Ingredients {
// 		u := insertUnit(db, ingredient.Unit)
// 		i := insertIngredient(db, ingredient.Name)
// 		insertRecipeIngredient(db, r.ID, i.ID, u.ID, ingredient.Quantity, ingredient.Note)
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"name":        recipe.Name,
// 		"description": recipe.Description,
// 		"direction":   recipe.Direction,
// 	})
// }

// func recipesController(c *gin.Context) {
// 	var recipes []Recipe
// 	db := getDB()
// 	defer db.Close()
// 	recipeRows, _ := db.Query("SELECT id, name, description, direction FROM recipe order by id asc")
// 	for recipeRows.Next() {
// 		var r Recipe
// 		recipeRows.Scan(&r.ID, &r.Name, &r.Description, &r.Direction)

// 		var ingredients []RecipeIngredient
// 		rows, _ := db.Query("SELECT i.name, ri.quantity, u.name, ri.note FROM ingredient as i, unit as u, recipe_ingredient as ri WHERE ri.recipe_id = ? and i.id = ri.ingredient_id and u.id = ri.unit_id order by u.name desc", r.ID)
// 		for rows.Next() {
// 			var ri RecipeIngredient
// 			rows.Scan(&ri.Name, &ri.Quantity, &ri.Unit, &ri.Note)
// 			ingredients = append(ingredients, ri)
// 		}
// 		r.Ingredients = ingredients

// 		recipes = append(recipes, r)
// 	}

// 	c.Header("Access-Control-Allow-Origin", "*")
// 	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
// 	c.JSON(http.StatusOK, recipes)
// }

// func recipeController(c *gin.Context) {
// 	id := c.Param("id")

// 	var r Recipe
// 	if id != "" {
// 		db := getDB()
// 		defer db.Close()
// 		err := db.QueryRow("SELECT id, name, description, direction FROM recipe where id = ?", id).Scan(&r.ID, &r.Name, &r.Description, &r.Direction)
// 		if err != sql.ErrNoRows {
// 			var ingredients []RecipeIngredient
// 			rows, _ := db.Query("SELECT i.name, ri.quantity, u.name, ri.note FROM ingredient as i, unit as u, recipe_ingredient as ri WHERE ri.recipe_id = ? and i.id = ri.ingredient_id and u.id = ri.unit_id order by u.name desc", id)
// 			for rows.Next() {
// 				var ri RecipeIngredient
// 				rows.Scan(&ri.Name, &ri.Quantity, &ri.Unit, &ri.Note)
// 				ingredients = append(ingredients, ri)
// 			}
// 			r.Ingredients = ingredients

// 			c.Header("Access-Control-Allow-Origin", "*")
// 			c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
// 			c.JSON(http.StatusOK, r)
// 		} else {
// 			c.Header("Access-Control-Allow-Origin", "*")
// 			c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
// 			c.JSON(http.StatusNotFound, struct{}{})
// 		}
// 	}

// }

// g.GET("/ping", ping)
// g.OPTIONS("/ping", preflight)

// g.GET("/search", search)
// g.OPTIONS("/search", preflight)

// g.GET("/recipes/", recipesController)
// g.OPTIONS("/recipes/", preflight)

// g.GET("/recipes/:id", recipeController)
// g.OPTIONS("/recipes/:id", preflight)

// g.GET("/recipes/count", recipesCount)
// g.OPTIONS("/recipes/count", preflight)

// g.POST("/recipes", createRecipe)
// g.OPTIONS("/recipes", preflight)

// go func() {
// 	http.Handle("/",
// 		http.StripPrefix("/",
// 			http.FileServer(http.Dir("../../web"))))
// 	log.Fatal(http.ListenAndServe(":"+config.WebPort, nil))
// }()

// g.Run(":" + config.ApiPort)
