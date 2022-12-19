// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"net/http"

	"github.com/chokey2nv/blockchain/app/services/node/handlers/v1/private"
	"github.com/chokey2nv/blockchain/app/services/node/handlers/v1/public"
	"github.com/chokey2nv/blockchain/foundation/web"
	"go.uber.org/zap"
)

const version = "v1"

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log *zap.SugaredLogger
}

// PublicRoutes binds all the version 1 public routes.
func PublicRoutes(app *web.App, cfg Config) {
	pbl := public.Handlers{
		Log: cfg.Log,
	}

	app.Handle(http.MethodGet, version, "/sample", pbl.SampleHander)
}

// PrivateRoutes binds all the version 1 private routes.
func PrivateRoutes(app *web.App, cfg Config) {
	prv := private.Handlers{
		Log: cfg.Log,
	}

	app.Handle(http.MethodPost, version, "/node/sample", prv.SampleHander)
}
