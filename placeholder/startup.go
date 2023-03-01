/*
configure the pipeline required by the application and use it to configure and start the HTTP server.
This is a task that will be performed by the application once I start development in Chapter 35.
The createPipeline function creates a pipeline with the middleware components created previously.
The Start function calls createPipeline and uses the result to configure and start the HTTP server.
*/
package placeholder

import (
	"sync"

	"github.com/vedicsociety/platform/authorization"
	"github.com/vedicsociety/platform/http"
	"github.com/vedicsociety/platform/http/handling"
	"github.com/vedicsociety/platform/pipeline"
	"github.com/vedicsociety/platform/pipeline/basic"
	"github.com/vedicsociety/platform/services"
	"github.com/vedicsociety/platform/sessions"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.AuthComponent{},
		&basic.StaticFileComponent{},
		&sessions.SessionComponent{},

		//&SimpleMessageComponent{},

		// The changes create a branch of the pipeline that has the /protected prefix,
		// which is restricted to users who have been assigned to the Administrator role.
		// The CounterHandler, defined earlier, is the only handler on the branch.
		// The AuthenticationHandler is added to the main branch of the pipeline.
		authorization.NewAuthComponent(
			"protected",
			authorization.NewRoleCondition("Administrator"),
			CounterHandler{},
		),
		handling.NewRouter(
			handling.HandlerEntry{"", NameHandler{}},
			handling.HandlerEntry{"", DayHandler{}},
			handling.HandlerEntry{"", UploadHandler{}},
			handling.HandlerEntry{"", AuthenticationHandler{}},
		).AddMethodAlias("/", NameHandler.GetNames),
	)
}

func Start() {
	sessions.RegisterSessionService()
	authorization.RegisterDefaultSignInService()
	authorization.RegisterDefaultUserService()
	RegisterPlaceholderUserStore()
	results, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
