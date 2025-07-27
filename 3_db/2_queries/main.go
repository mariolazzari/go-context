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

	user, err := queryUser(ctx, db, "user123")
	if err != nil {
		panic(err)
	}

	println("User ID:", user.ID, "Name:", user.Name)
}

// query user with context
func queryUser(ctx context.Context, db *sql.DB, userID string) (User, error) {
	tenantID := ctx.Value("tenantID").(string)
	var user User
	err := db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = ? AND tenant_id = ?", userID, tenantID).Scan(&user.ID, &user.Name)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
