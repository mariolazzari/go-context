# Context Package In Go Lang

## Introduction

### Context package

- Go standard lib
- Helps manage request-scoped values and deadlines
- Individual requests or operations
- Deadlines
- Cancellation signals
- Request scoped values

#### Context type

Container for passing in formations between functions and go routines.

#### Context with deadlines

It allows you to set a time limit for specific operations and handle them accordingly.

#### Context with timeout

It sets a timeout for an operation and hanfle it accordingly.

#### Context with values

It adds additional values to specific request.

### Value passing context

Concurretly safe:
  
- tokens
- request IDs

```go
ctx := context.WithValue(parentContext, key, value)
val := ctx.Value(key)
```

### Context functions and methods

- Manage life cycles and cancellations
- Creates new context

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func(){
    // operations
    cancel()
}()

ctx, cancel := context.WithTimeout(context.Background())
time.Duration(time.Millisecond * 100)
defer cancel()
req = req.WithContext(ctx)
```

### Context cancellation

It terminates operations gracefully.

### Use cases

- Go routines
- Database transactions
- HTTP requests
- Testing

## Go routines

### Goroutine Overview

- Threads managed by go runtime
- Functions run indipendently
- Concurrency

```go
package main

import (
 "fmt"
 "time"
)

func main() {
    go sayHello(("Hello"))
    sayHello("World")
}

func sayHello(msg string) {
    for range 5 {
     fmt.Println(msg)
     time.Sleep(time.Second)
 }
}
```

### Managing context in go routines

Pass information without expliciting function paramenters.

```go
package main

import (
 "context"
 "time"
)

func main() {
 // create a context
 ctx := context.Background()
 // create cancelable context with a timeout
 ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
 defer cancel() // ensure resources are cleaned up
 myRoutine(ctx)
}

func myRoutine(ctx context.Context) {
 <-ctx.Done()
}
```

## Using Context in database operations

### Overview

- Connection pooling
- Cancellations

### Queries

### Transactions

```go
func queryUser(ctx context.Context, db *sql.DB, userID string) (User, error) {
 tenantID := ctx.Value("tenantID").(string)
 var user User
 err := db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = ? AND tenant_id = ?", userID, tenantID).Scan(&user.ID, &user.Name)
 if err != nil {
  return User{}, err
 }
 return user, nil
}
```

### Transactions

- Commit or rollback transaction

```go
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
```
