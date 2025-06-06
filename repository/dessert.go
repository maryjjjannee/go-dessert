package repository

import (
	"context"
	"database/sql"
	"go-dessert/models"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllDesserts(db *sql.DB) ([]models.Dessert, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `SELECT id, name, description, price, image_url FROM sweet_dessert`)
	if err != nil {
		log.Printf("Error querying all desserts: %v", err)
		return nil, err
	}
	defer rows.Close()

	desserts := make([]models.Dessert, 0)
	for rows.Next() {
		var dessert models.Dessert
		if err := rows.Scan(&dessert.ID, &dessert.Name, &dessert.Description, &dessert.Price, &dessert.ImageURL); err != nil {
			log.Printf("Error scanning dessert row: %v", err)
			return nil, err
		}
		desserts = append(desserts, dessert)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating dessert rows: %v", err)
		return nil, err
	}

	return desserts, nil
}

func GetDessertByID(db *sql.DB, id int) (*models.Dessert, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dessert := &models.Dessert{}
	err := db.QueryRowContext(ctx, `SELECT id, name, description, price, image_url FROM sweet_dessert WHERE id = ?`, id).
		Scan(&dessert.ID, &dessert.Name, &dessert.Description, &dessert.Price, &dessert.ImageURL)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Printf("Error querying dessert by ID %d: %v", id, err)
		return nil, err
	}
	return dessert, nil
}

func InsertDessert(db *sql.DB, dessert models.Dessert) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := db.ExecContext(ctx, `INSERT INTO sweet_dessert (name, description, price, image_url) VALUES (?, ?, ?, ?)`,
		dessert.Name, dessert.Description, dessert.Price, dessert.ImageURL)
	if err != nil {
		log.Printf("Error inserting dessert: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return 0, err
	}
	return id, nil
}

func UpdateDessert(db *sql.DB, dessert models.Dessert) error {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    _, err := db.ExecContext(ctx, `UPDATE sweet_dessert SET name=?, description=?, price=?, image_url=? WHERE id=?`,
        dessert.Name, dessert.Description, dessert.Price, dessert.ImageURL, dessert.ID)
    if err != nil {
        log.Printf("Error updating dessert with ID %d: %v", dessert.ID, err)
        return err
    }
    return nil
}

func DeleteDessert(db *sql.DB, id int) error {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    _, err := db.ExecContext(ctx, `DELETE FROM sweet_dessert WHERE id = ?`, id)
    if err != nil {
        log.Printf("Error deleting dessert with ID %d: %v", id, err)
        return err
    }
    return nil
}