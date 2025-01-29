package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"

	config "github.com/KaDingMeaw/godb/config"
	models "github.com/KaDingMeaw/godb/models"
	module "github.com/KaDingMeaw/godb/modules"
)

func main() {
	var wg sync.WaitGroup
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
	err := config.LoadEnv(appEnv)
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

	// Insert the Cover object into the database
	for i := 0; i < 10; i++ {
		wg.Add(1)
		newCover := &models.Cover{
			Item: "Sample Item " + strconv.Itoa(i),
		}

		err = module.InsertItem(db, &wg, newCover)

		if err != nil {
			log.Fatalf("Error inserting item: %s", err)
		} else {
			fmt.Printf("Insert Completed %s\n", newCover.Item)
		}

	}

	wg.Wait()

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
