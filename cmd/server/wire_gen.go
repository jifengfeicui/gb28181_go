// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gowvp/gb28181/internal/conf"
	"github.com/gowvp/gb28181/internal/data"
	"github.com/gowvp/gb28181/internal/web/api"
	"log/slog"
	"net/http"
)

// Injectors from wire.go:

func wireApp(bc *conf.Bootstrap, log *slog.Logger) (http.Handler, func(), error) {
	db, err := data.SetupDB(bc, log)
	if err != nil {
		return nil, nil, err
	}
	core := api.NewVersion(db)
	versionAPI := api.NewVersionAPI(core)
	smsCore := api.NewSMSCore(db, bc)
	smsAPI := api.NewSmsAPI(smsCore)
	uniqueidCore := api.NewUniqueID(db)
	mediaCore := api.NewMediaCore(db, uniqueidCore)
	webHookAPI := api.NewWebHookAPI(smsCore, mediaCore, bc)
	mediaAPI := api.NewMediaAPI(mediaCore, smsCore, bc)
	usecase := &api.Usecase{
		Conf:       bc,
		DB:         db,
		Version:    versionAPI,
		SMSAPI:     smsAPI,
		WebHookAPI: webHookAPI,
		UniqueID:   uniqueidCore,
		MediaAPI:   mediaAPI,
	}
	handler := api.NewHTTPHandler(usecase)
	return handler, func() {
	}, nil
}
