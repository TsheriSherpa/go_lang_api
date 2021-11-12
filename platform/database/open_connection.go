package database

import "github.com/tsheri/go-fiber/app/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries // load queries from User model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries: &queries.UserQueries{DB: db}, // from User model
	}, nil
}
