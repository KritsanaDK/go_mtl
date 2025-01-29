package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"

	models "github.com/KaDingMeaw/godb/models"
	module "github.com/KaDingMeaw/godb/modules"
)

func loadEnv(env string) error {
	fileName := fmt.Sprintf(".env.%s", env)
	return godotenv.Load(fileName)
}

func main() {

	// Get all arguments
	args := os.Args
	appEnv := "dev"

	// Print arguments if available
	if len(args) > 1 {
		appEnv = args[1]
		fmt.Println("Arguments:", args[1:])
	} else {
		fmt.Println("No arguments provided")
	}

	// Load the appropriate .env file
	err := loadEnv(appEnv)
	if err != nil {
		log.Fatalf("Error loading .env.%s file: %v", appEnv, err)
	}

	// Port := os.Getenv("APP_PORT")
	DbHost := os.Getenv("DB_HOST")
	// DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", DbHost, DbUser, DbPassword, DbName)

	fmt.Println(connString)

	db, err := sql.Open("mssql", connString)
	// db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	defer db.Close()

	if err != nil {
		log.Fatal("Error executing query:", err)
	}

	items, err := module.GetAllItem(db)
	if err != nil {
		log.Printf("Error %s when selecting product by price", err)
		return
	}
	for _, item := range items {
		log.Printf("Item: %s ID: %d", item.Item, item.Id)
	}

	ValueTest := items[0].Id

	cover, err := module.GetItem(db, ValueTest)

	if err != nil {
		log.Fatalf("Error fetching item: %s", err)
	}
	log.Printf("Fetched Cover: %+v", cover)

	// Create a new Cover object
	newCover := &models.Cover{
		Item: "Sample Item",
	}

	// Insert the Cover object into the database
	for i := 0; i < 1; i++ {
		err = module.InsertItem(db, newCover)
		// Sleep for 2 seconds
		time.Sleep(200 * time.Millisecond)
		if err != nil {
			log.Fatalf("Error inserting item: %s", err)
		}
	}

	fmt.Println(cover.Id, cover.Item)

	updatedCover := &models.Cover{
		Id:   items[1].Id,
		Item: "Updated Item",
	}

	// Update the Cover object in the database
	err = module.UpdateItem(db, updatedCover)
	if err != nil {
		log.Fatalf("Error updating item: %s", err)
	}

	err = module.DelItem(db, ValueTest)
	if err != nil {
		log.Printf("Error deleting item: %s", err)
		return
	}

	log.Println("Item deleted successfully.")

}
