/*
Almost all web applications require support for serving static files, even if it is just for CSS stylesheets.
The standard library contains built-in support for serving files, which is helpful because it is a task that is fraught with potential problems.
But fortunately, it is a simple matter to integrate the standard library features into the request pipeline in the example project.
This handler uses the Init method to read the configuration settings that specify the prefix used for file requests
and the directory from which to serve files and uses the handlers provided by the net/http package to serve files.
*/
package basic

import (
	"net/http"
	"platform/config"
	"platform/pipeline"

	//"platform/services"
	"strings"
)

type StaticFileComponent struct {
	urlPrefix     string
	stdLibHandler http.Handler
	Config        config.Configuration
}

func (sfc *StaticFileComponent) Init() {
	// var cfg config.Configuration
	// services.GetService(&cfg)
	sfc.urlPrefix = sfc.Config.GetStringDefault("files:urlprefix", "/files/")
	path, ok := sfc.Config.GetString("files:path")
	if ok {
		sfc.stdLibHandler = http.StripPrefix(sfc.urlPrefix,
			http.FileServer(http.Dir(path)))
	} else {
		panic("Cannot load file configuration settings")
	}
}

func (sfc *StaticFileComponent) ProcessRequest(ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext)) {

	if !strings.EqualFold(ctx.Request.URL.Path, sfc.urlPrefix) &&
		strings.HasPrefix(ctx.Request.URL.Path, sfc.urlPrefix) {
		sfc.stdLibHandler.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	} else {
		next(ctx)
	}
}
