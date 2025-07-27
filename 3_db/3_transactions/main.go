package main

import (
	"context"
	"database/sql"
)

type User struct {
	ID   string
	Name string
}

func main() {
	// Example usage
	ctx := context.WithValue(context.Background(), "tenantID", "tenant123")
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	user := User{ID: "user123", Name: "John Doe"}
	err = updateUser(ctx, db, user)
	if err != nil {
		panic(err)
	}

	println("User updated:", user.ID, user.Name)
}

// update user with context and transaction
func updateUser(ctx context.Context, db *sql.DB, user User) error {
	tenantID := ctx.Value("tenantID").(string)

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE users SET name = ? WHERE id = ? AND tenant_id = ?", user.Name, user.ID, tenantID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
