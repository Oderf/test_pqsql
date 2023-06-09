package main

import (
	"database.sql"
   "fmt"
   "log"
   "os"
   _ "github.com/lib/pq"
   "github.com/gofiber/fiber/v2"
)

func main() {
   app := fiber.New()
 
   app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
	return deleteHandler(c, db)
	})

   port := os.Getenv("PORT")
   if port == "" {
       port = "3000"
   }
   log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

   connStr := "postgresql://checkmyconf:Hypocrisy2Scorebook2Alabaster2Clubbed2Progeny62Roamer@192.168.1.1/todos?sslmode=disable
   "
	  // Connect to database
	  db, err := sql.Open("postgres", connStr)
	  if err != nil {
		  log.Fatal(err)
	  }
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string
	rows, err := db.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
 }
 
 type todo struct {
	Item string
 }

 func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTodo := todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT into todos VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}
}
 
 func putHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
 }
 
 func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
 }