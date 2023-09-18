
# pgcontroller


`pgcontroller` is a Go package designed to simplify database management, migration, and CRUD operations for PostgreSQL databases. It provides a convenient and efficient way to interact with your PostgreSQL database within your Go applications.

## Features

- **Database Migration:** `pgcontroller` allows you to easily perform database migrations using a simple command-line interface.
- **CRUD Operations:** It provides a structured framework for performing Create, Read, Update, and Delete (CRUD) operations on PostgreSQL tables.
- **Query Parameter Handling:** The package includes utilities for handling query parameters in HTTP requests, making it easier to filter and sort database records.
- **Pagination Support:** You can handle pagination of query results effortlessly.
- **Error Handling:** Comprehensive error handling to ensure your application remains robust.

## Installation

To use `pgcontroller` in your Go project, you need to install it using `go get`:

```sh
go get github.com/mis-hashemi/pgcontroller

   ```

## Usage

Here is a basic example of how to use pgcontroller in your Go application:

```
package main

import (
	// Import necessary packages
	"github.com/mis-hashemi/pgcontroller"
	"github.com/mis-hashemi/pgcontroller/migrator"
	// ... other imports
)

func main() {
	// Initialize pgcontroller
	controller := pgcontroller.NewPgController(pgcontroller.Config{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "123",
		DBName:   "ex01",
	}, false)

	// Generate database schema if needed
	err := controller.Generate()
	if err != nil {
		panic(err)
	}

	// Initialize the database connection
	err = controller.Init()
	if err != nil {
		panic(err)
	}

	// Create a migrator
	mgr := migrator.New(controller.GetDataContext(), "./repository/migrations")

	// Perform database migration (up or down)
	migrateOperation("up", mgr)

	// Create an Echo instance for handling HTTP requests
	e := echo.New()

	// Define query parameter information

	// Define repository, service, and API routes

	// Start the server
	e.Start(":8080")
}

func migrateOperation(flag string, mg migrator.Migrator) {
	// Perform database migration (up or down)
	// ...
}

```

```
localhost:8080/users?first_name=in:fatemeh,ali,sara&last_name=nin:hashemi
```

For more advanced usage and integration details, please refer to the library's documentation or code examples in the "examples" directory.

## Contributing

Contributions to the "pgcontroller" library are welcome. If you have ideas for improvements, bug fixes, or new features, please feel free to open an issue or submit a pull request. Make sure to follow the project's code of conduct.