package bcgql

import (
	"fmt"
	"net/http"

	"github.com/NavenduDuari/go-blockchain/bcgql/httphandler"
)

func Run() error {

	http.Handle("/", httphandler.GqlHandler())
	fmt.Println("Server running...")
	err := http.ListenAndServe(":8080", nil)
	return err
}
