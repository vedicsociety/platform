/*
This component logs basic details of the request and response using the Logger service created in Chapter 32.
The ResponseWriter interface doesn’t provide access to the status code sent in a response
and so a LoggingResponseWriter is created and passed to the next component in the pipeline.
This component performs actions before and after the next function is invoked, logging a message before passing on the request

	and logging another message that writes out the status code after the request has been processed.

	This component obtains a Logger service when it processes a request.
	We could obtain a Logger just once, but that works only because we know that the Logger has been registered as a singleton service.
	Instead, we prefer not to make assumptions about the Logger lifecycle, which means that we won’t get unexpected results
	 if the lifecycle is changed in the future.
*/
package basic

import (
	"net/http"

	"github.com/vedicsociety/platform/config"
	"github.com/vedicsociety/platform/logging"
	"github.com/vedicsociety/platform/pipeline"
)

type AuthComponent struct {
	Config              config.Configuration
	Logger              logging.Logger
	Is_BasicAuthEnabled bool
}

func (c *AuthComponent) Init() {
	c.Is_BasicAuthEnabled = c.Config.GetBoolDefault("auth:isenabled", false)

}

func (c *AuthComponent) ProcessRequest(ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext)) {

	c.Logger.Debugf("AuthComponent.Is_BasicAutenticated :", ctx.Is_BasicAutenticated, c.Is_BasicAuthEnabled)

	if c.Is_BasicAuthEnabled {
		
		if !ctx.Is_BasicAutenticated {

			ctx.ResponseWriter.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)

			user, pass, ok := ctx.Request.BasicAuth()
			c.Logger.Debugf("AuthComponent after", ok)

			if ok {
				osuser, _ := c.Config.GetString("auth:user")
				ospassw, _ := c.Config.GetString("auth:password")
				if osuser == user && ospassw == pass {
					c.Logger.Debugf("AuthComponent :", user, pass, ok)
					ctx.Is_BasicAutenticated = true
					ctx.ResponseWriter.WriteHeader(0)
					next(ctx)
				}
			}
		}

	} else {
		next(ctx)
	}

}
