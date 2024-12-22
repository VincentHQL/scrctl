package operator

import (
	"net/http"

	apiv1 "github.com/VincentHQL/scrctl/api/operator/apiv1"
	"github.com/gorilla/mux"
)

// Creates a router with handlers for the following endpoints:
// GET  /infra_config
func CreateHttpHandlers(
	pool *DevicePool,
	polledSet *PolledSet,
	config apiv1.InfraConfig) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		connect(w, config, pool, polledSet)
	}).Methods("GET")
	return router
}

func connect(w http.ResponseWriter, config apiv1.InfraConfig,
	pool *DevicePool, polledSet *PolledSet) {
	// Implement the function logic here
}
