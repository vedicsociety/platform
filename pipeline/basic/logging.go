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

	"github.com/tsiparinda/platform/logging"
	"github.com/tsiparinda/platform/pipeline"
	//"platform/services"
)

type LoggingResponseWriter struct {
	statusCode int
	http.ResponseWriter
}

func (w *LoggingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *LoggingResponseWriter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = http.StatusOK
	}
	return w.ResponseWriter.Write(b)
}

type LoggingComponent struct{}

func (lc *LoggingComponent) ImplementsProcessRequestWithServices() {}

func (lc *LoggingComponent) Init() {}

func (lc *LoggingComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	logger logging.Logger) {

	// var logger logging.Logger
	// err := services.GetServiceForContext(ctx.Request.Context(), &logger)
	// if (err != nil) {
	//     ctx.Error(err)
	//     return
	// }

	loggingWriter := LoggingResponseWriter{0, ctx.ResponseWriter}
	ctx.ResponseWriter = &loggingWriter

	logger.Infof("REQ --- %v - %v", ctx.Request.Method, ctx.Request.URL)
	next(ctx)
	logger.Infof("RSP %v %v", loggingWriter.statusCode, ctx.Request.URL)
}
