package graphqlClient

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/machinebox/graphql"
)

// CreateClient initializes and returns a new GraphQL client
func CreateClient() *graphql.Client {
	if os.Getenv("GRAPHQL_ENDPOINT") == "" {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file and GRAPHQL_ENDPOINT environment variable not set")
		}
	}

	endpoint := []byte(os.Getenv("GRAPHQL_ENDPOINT"))
	if len(endpoint) == 0 {
		panic("GRAPHQL_ENDPOINT environment variable not set")
	}

	return graphql.NewClient(string(endpoint))
}

// QueryHasura executes a GraphQL query and returns the response data
func QueryHasura(client *graphql.Client, query string, variables map[string]interface{}) (map[string]interface{}, error) {
	req := graphql.NewRequest(query)

	// First, we'll set any variables required for the query
	for key, value := range variables {
		req.Var(key, value)
	}

	// Define a context for the request
	ctx := context.Background()

	// Then, we'll define a variable to capture the response
	var respData map[string]interface{}

	// Finally, execute the request
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Println("Error executing query:", err)

		return nil, err
	}

	return respData, nil
}
