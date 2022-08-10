package services

import (
	"github.com/tsiparinda/platform/config"
	"github.com/tsiparinda/platform/logging"
	"github.com/tsiparinda/platform/templates"
)

// The RegisterDefaultServices creates Configuration and Logger services.
// These services are created using the AddSingleton function, which means that a single instance of the structs that implement each interface
// will be shared by the entire application.
func RegisterDefaultServices() {

	err := AddSingleton(func() (c config.Configuration) {
		c, loadErr := config.Load("config.json")
		if loadErr != nil {
			panic(loadErr)
		}
		return
	})

	err = AddSingleton(func(appconfig config.Configuration) logging.Logger {
		return logging.NewDefaultLogger(appconfig)
	})
	if err != nil {
		panic(err)
	}

	err = AddSingleton(
		func(c config.Configuration) templates.TemplateExecutor {
			templates.LoadTemplates(c)
			return &templates.LayoutTemplateProcessor{}
		})
	if err != nil {
		panic(err)
	}
}
