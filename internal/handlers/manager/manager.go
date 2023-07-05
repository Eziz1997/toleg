package handlermanager

import (
	"toleg/internal/client/admin"
	"toleg/pkg/client/postgresql"
	"toleg/pkg/logging"

	admindb "toleg/internal/client/admin/db"

	"github.com/gorilla/mux"
)

const (
	settingsURL = "/api/v1/settings"
	adminURL    = "/api/altynasyr"
	userURL     = "/api/v1/user"
)

func Manager(client postgresql.Client, logger *logging.Logger) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// Settings
	//settingsRouterManager := router.PathPrefix(settingsURL).Subrouter()
	//settingsRouterRepository := settingsdb.NewRepository(client, logger)
	//settingsRouterHandler := settings.NewHandler(settingsRouterRepository, logger)
	//settingsRouterHandler.Register(settingsRouterManager)

	//Admin
	adminRouterManager := router.PathPrefix(adminURL).Subrouter()
	adminRouterRepository := admindb.NewRepository(client, logger)
	adminRouterHandler := admin.NewHandler(adminRouterRepository, logger)
	adminRouterHandler.Register(adminRouterManager)

	// user
	//userRouterManager := router.PathPrefix(userURL).Subrouter()
	//userRouterRepository := userdb.NewRepository(client, logger)
	//userRouterHandler := user.NewHandler(userRouterRepository, logger)
	//userRouterHandler.Register(userRouterManager)

	return router
}
