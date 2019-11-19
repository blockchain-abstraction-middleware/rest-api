package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// SwaggerResource implements server.Resource interface
type SwaggerResource struct {
	path string
}

// SwaggerResource constructor func
func (r *SwaggerResource) NewResource(path string) *SwaggerResource {
	return &SwaggerResource{
		path: path,
	}
}

// Path returns the Resource base path
func (r *SwaggerResource) Path() string {
	return r.path
}

// Routes bootstraps health routes
func (r *SwaggerResource) Routes() http.Handler {
	router := chi.NewRouter()

	router.Handle("/docs/*", r.serverSwagger())

	return router
}

func (r *SwaggerResource) serverSwagger() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		prefix := "/api/v1/swagger"

		hfs := http.FileServer(http.Dir("docs"))
		fs := http.StripPrefix(prefix+"/docs/", hfs)
		fs.ServeHTTP(res, req)
	}
}
