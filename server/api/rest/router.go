package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// InitRouter initialize the mux.Router.
func InitRouter(port string) {
	router := mux.NewRouter()

	log.Info(fmt.Sprintf("Server started on port %s", port))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
