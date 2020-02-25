package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NavenduDuari/go-blockchain/bcgql/mutation"
	"github.com/NavenduDuari/go-blockchain/bcgql/query"
	"github.com/graphql-go/graphql"
)

type reqBody struct {
	Query string `json:"query"`
}

func GqlHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "No query data", 400)
			return
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		fmt.Fprintf(w, "%s", processQuery(rBody.Query))
		// query := r.URL.Query().Get("query")
		// json.NewEncoder(w).Encode(processQuery(query))
	})
}

// func processQuery(query string) *graphql.Result {
func processQuery(query string) string {
	result := graphql.Do(graphql.Params{
		Schema:        gqlSchema(),
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Println("Wrong result, unexpected errors: ", result.Errors)
	}

	jsonResult, _ := json.Marshal(result)
	return fmt.Sprintf("%s", jsonResult)

	// return result
}

func gqlSchema() graphql.Schema {
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(query.RootQuery),
		Mutation: graphql.NewObject(mutation.RootMutation),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		fmt.Println("Failed to create schema : ", err)
	}
	return schema
}
