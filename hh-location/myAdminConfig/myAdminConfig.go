package myAdminConfig

import (
	"beacon/hh-location/provider"
	"github.com/qor/admin"
	"beacon/hh-location/models"
	"net/http"
)

func Init() *http.ServeMux {
	DB := provider.GetProvider().GetDB()

	DB.AutoMigrate(&models.Beacon{}, &models.Device{}, &models.Position{})

	Admin := admin.New(&admin.AdminConfig{DB: DB})
	Admin.AddResource(&models.Beacon{})
	Admin.AddResource(&models.Device{})

	m := http.NewServeMux()
	Admin.MountTo("/admin", m)

	return m
}
