package mssql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/KaDingMeaw/godb/models"
)

func GetAllItem(db *sql.DB) ([]models.Cover, error) {
	query := `SELECT id, item FROM dbo.itemData;`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return []models.Cover{}, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return []models.Cover{}, err
	}
	defer rows.Close()
	var covers = []models.Cover{}
	for rows.Next() {
		var prd models.Cover
		if err := rows.Scan(&prd.Id, &prd.Item); err != nil {
			return []models.Cover{}, err
		}
		covers = append(covers, prd)
	}
	if err := rows.Err(); err != nil {
		return []models.Cover{}, err
	}
	return covers, nil
}

func GetItem(db *sql.DB, id int) (*models.Cover, error) {

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

	var cover = models.Cover{}

	err = row.Scan(&cover.Id, &cover.Item)

	if err != nil {
		return nil, err
	}
	return &cover, nil

}

func DelItem(db *sql.DB, id int) error {
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

func InsertItem(db *sql.DB, cover *models.Cover) error {
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
		sql.Named("item", cover.Item),
	)
	if err != nil {
		log.Printf("Error %s when executing SQL statement", err)
		return err
	}

	log.Printf("Successfully inserted item: %+v", cover)
	return nil
}

func UpdateItem(db *sql.DB, cover *models.Cover) error {
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
		sql.Named("id", cover.Id),     // Named parameter for id
		sql.Named("item", cover.Item), // Named parameter for item
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
		log.Printf("No record found with id %d", cover.Id)
		return fmt.Errorf("no record found with id %d", cover.Id)
	}

	log.Printf("Successfully updated %d record(s) with id %d", rowsAffected, cover.Id)
	return nil
}

// type Cover struct {
// 	id   int
// 	item string
// }
