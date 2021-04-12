package myAdminConfig

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"hh-location-web/hh-location/configurator"
	"hh-location-web/hh-location/models"
	"net/http"
)

func Init() *http.ServeMux {

	dbString := configurator.GetConfiguration().DbString
	DB, _ := gorm.Open("mysql", dbString)
	DB.AutoMigrate(&models.Device{}, &models.Beacon{}, &models.DevicesPositions{})

	Admin := admin.New(&admin.AdminConfig{DB: DB})

	deviceMeta := Admin.AddResource(&models.Device{}, &admin.Config{
		Menu: []string{"Тестовые устройства"},
		Name: "Устройства",
	})
	deviceMeta.Meta(&admin.Meta{Name: "Name", Label: "Заводское название"})
	deviceMeta.Meta(&admin.Meta{Name: "Nickname", Label: "Народное название"})
	deviceMeta.Meta(&admin.Meta{Name: "Private", Label: "Личный"})
	deviceMeta.Meta(&admin.Meta{Name: "OS", Valuer: func(record interface{}, context *qor.Context) interface{} {
		if p, ok := record.(*models.Device); ok {
			return p.OSName + " " + p.OSVersion
		}
		return ""
	}})
	deviceMeta.Meta(&admin.Meta{Name: "OS", Label: "Платформа"})

	deviceMeta.IndexAttrs("-Id", "-Uid", "-OSName", "-OSVersion", "-Shell", "-Resolution", "-Type", "-Comment", "-TokenUid", "-OwnerId", "-Private")
	deviceMeta.EditAttrs("-OS", "-Uid", "-OwnerId")
	deviceMeta.NewAttrs("-OS", "-Uid", "-OwnerId")
	deviceMeta.Meta(&admin.Meta{Name: "OSName", Config: &admin.SelectOneConfig{Collection: []string{"iOS", "Android", "Другое"}}})
	deviceMeta.Meta(&admin.Meta{Name: "Type", Config: &admin.SelectOneConfig{Collection: []string{"Phone", "Tablet"}}})

	Admin.AddResource(&models.Beacon{}, &admin.Config{
		Menu: []string{"Тестовые устройства"},
		Name: "Метки",
	})

	m := http.NewServeMux()
	Admin.MountTo("/admin", m)

	return m
}
