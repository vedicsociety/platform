package services

import (
	"github.com/vedicsociety/platform/config"
	"github.com/vedicsociety/platform/logging"
	"github.com/vedicsociety/platform/templates"
	"github.com/vedicsociety/platform/validation"
)

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

	err = AddSingleton(
		func() validation.Validator {
			return validation.NewDefaultValidator(validation.DefaultValidators())
		})
	if err != nil {
		panic(err)
	}

}
