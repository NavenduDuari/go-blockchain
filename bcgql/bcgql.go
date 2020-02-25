package bcgql

import (
	"fmt"
	"net/http"

	"github.com/NavenduDuari/go-blockchain/bcgql/httphandler"
	"github.com/friendsofgo/graphiql"
)

func Run() error {

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/")
	if err != nil {
		panic(err)
	}
	http.Handle("/graphiql", graphiqlHandler)

	http.Handle("/", httphandler.GqlHandler())
	fmt.Println("Server running...")
	err = http.ListenAndServe(":8080", nil)
	return err
}
