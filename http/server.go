/*
create the HTTP server and use a pipeline to handle the requests it receives.
The Serve function uses the Configuration service to read the settings for HTTP and HTTPS
and uses the features provided by the standard library to receive requests and pass them to the pipeline for processing.
(I will enable HTTPS support in Chapter 38 when I prepare for deployment, but until then,
I will use the default settings that listen for HTTP requests on port 5000.)
*/
package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/vedicsociety/platform/config"
	"github.com/vedicsociety/platform/logging"
	"github.com/vedicsociety/platform/pipeline"
)

type pipelineAdaptor struct {
	pipeline.RequestPipeline
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}

func (p pipelineAdaptor) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
		enableCors(&writer)
		
	p.ProcessRequest(request, writer)
}

func Serve(pl pipeline.RequestPipeline, cfg config.Configuration, logger logging.Logger) *sync.WaitGroup {
	wg := sync.WaitGroup{}

	adaptor := pipelineAdaptor{RequestPipeline: pl}
	enableHttp := cfg.GetBoolDefault("http:enableHttp", true)
	if enableHttp {
		httpPort := cfg.GetIntDefault("http:port", 5000)
		// for compatability with heroku
		osport := os.Getenv("PORT")
		if osport != "" {
			httpPort, _ = strconv.Atoi(osport)
		}
		logger.Debugf("Starting HTTP server on port %v", httpPort)
		wg.Add(1)
		go func() {
			err := http.ListenAndServe(fmt.Sprintf(":%v", httpPort), adaptor)
			if err != nil {
				panic(err)
			}
		}()
	}
	enableHttps := cfg.GetBoolDefault("http:enableHttps", false)
	if enableHttps {
		httpsPort := cfg.GetIntDefault("http:httpsPort", 5500)
		certFile, cfok := cfg.GetString("http:httpsCert")
		keyFile, kfok := cfg.GetString("http:httpsKey")
		if cfok && kfok {
			logger.Debugf("Starting HTTPS server on port %v", httpsPort)
			wg.Add(1)
			go func() {
				err := http.ListenAndServeTLS(fmt.Sprintf(":%v", httpsPort),
					certFile, keyFile, adaptor)
				if err != nil {
					panic(err)
				}
			}()
		} else {
			panic("HTTPS certificate settings not found")
		}
	}
	return &wg
}
