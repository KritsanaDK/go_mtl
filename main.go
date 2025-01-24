package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

func loadEnv(env string) error {
	fileName := fmt.Sprintf(".env.%s", env)
	return godotenv.Load(fileName)
}

func main() {

	// Get the environment variable to determine which .env file to load
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		// Default to "dev" if APP_ENV is not set
		appEnv = "dev"
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

	products, err := selectProductsByPrice(db)
	if err != nil {
		log.Printf("Error %s when selecting product by price", err)
		return
	}
	for _, product := range products {
		log.Printf("Item: %s ID: %d", product.item, product.id)
	}

	ValueTest := 8

	cover, err := getItem(db, ValueTest)

	if err != nil {
		log.Fatalf("Error fetching item: %s", err)
	}
	log.Printf("Fetched Cover: %+v", cover)

	// Create a new Cover object
	newCover := &Cover{
		item: "Sample Item",
	}

	// Insert the Cover object into the database
	err = insertItem(db, newCover)
	if err != nil {
		log.Fatalf("Error inserting item: %s", err)
	}

	fmt.Println(cover.id, cover.item)

	updatedCover := &Cover{
		id:   ValueTest,
		item: "Updated Item",
	}

	// Update the Cover object in the database
	err = updateItem(db, updatedCover)
	if err != nil {
		log.Fatalf("Error updating item: %s", err)
	}

	err = delItem(db, ValueTest)
	if err != nil {
		log.Printf("Error deleting item: %s", err)
		return
	}

	log.Println("Item deleted successfully.")

}

func selectProductsByPrice(db *sql.DB) ([]Cover, error) {
	query := `SELECT id, item FROM dbo.itemData;`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return []Cover{}, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return []Cover{}, err
	}
	defer rows.Close()
	var covers = []Cover{}
	for rows.Next() {
		var prd Cover
		if err := rows.Scan(&prd.id, &prd.item); err != nil {
			return []Cover{}, err
		}
		covers = append(covers, prd)
	}
	if err := rows.Err(); err != nil {
		return []Cover{}, err
	}
	return covers, nil
}

func getItem(db *sql.DB, id int) (*Cover, error) {

	query := "SELECT  id, item  FROM dbo.itemData where id = $id"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, sql.Named("id", id))

	if err != nil {
		return nil, err
	}

	var cover = Cover{}

	err = row.Scan(&cover.id, &cover.item)

	if err != nil {
		return nil, err
	}
	return &cover, nil

}

func delItem(db *sql.DB, id int) error {
	// Correct query with named parameter
	query := "DELETE FROM dbo.itemData WHERE id = $id"

	// Create a context with a timeout
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	// Execute the DELETE statement with the parameter
	result, err := stmt.ExecContext(ctx, sql.Named("id", id))
	if err != nil {
		log.Printf("Error %s when executing SQL statement", err)
		return err
	}

	// Check the number of affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error %s when checking affected rows", err)
		return err
	}

	// Handle no rows affected (record not found)
	if rowsAffected == 0 {
		log.Printf("No record found with id %d", id)
		return fmt.Errorf("no record found with id %d", id)
	}

	// Log the successful deletion
	log.Printf("Successfully deleted %d record(s) with id %d", rowsAffected, id)
	return nil
}

func insertItem(db *sql.DB, cover *Cover) error {
	query := "INSERT INTO dbo.itemData (item) VALUES ($item)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	// Execute the INSERT statement
	_, err = stmt.ExecContext(ctx,
		sql.Named("item", cover.item),
	)
	if err != nil {
		log.Printf("Error %s when executing SQL statement", err)
		return err
	}

	log.Printf("Successfully inserted item: %+v", cover)
	return nil
}

func updateItem(db *sql.DB, cover *Cover) error {
	// Update query
	query := "UPDATE dbo.itemData SET item = $item WHERE id = $id"

	// Create a context with timeout for query execution
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	// Execute the UPDATE statement with named parameters
	result, err := stmt.ExecContext(ctx,
		sql.Named("id", cover.id),     // Named parameter for id
		sql.Named("item", cover.item), // Named parameter for item
	)
	if err != nil {
		log.Printf("Error %s when executing SQL statement", err)
		return err
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error %s when checking affected rows", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("No record found with id %d", cover.id)
		return fmt.Errorf("no record found with id %d", cover.id)
	}

	log.Printf("Successfully updated %d record(s) with id %d", rowsAffected, cover.id)
	return nil
}

type Cover struct {
	id   int
	item string
}
