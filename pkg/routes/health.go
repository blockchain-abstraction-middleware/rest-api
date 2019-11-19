package routes

import (
	"net/http"

	"github.com/blockchain-abstraction-middleware/rest-api/pkg/health"
	"github.com/go-chi/chi"
)

// HealthResource implements server.Resource interface
type HealthResource struct {
	path string
}

// NewResource constructor func
func (r *HealthResource) NewResource(path string) *HealthResource {
	return &HealthResource{
		path: path,
	}
}

// Path returns the Resource base path
func (r *HealthResource) Path() string {
	return r.path
}

// Routes bootstraps health routes
func (r *HealthResource) Routes() http.Handler {
	router := chi.NewRouter()

	router.Get("/status", r.healthCheck())

	return router
}

func (r *HealthResource) healthCheck() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		json, _ := health.Handler()

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(json)
	}
}
