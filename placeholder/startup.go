package placeholder

import (
	"platform/authorization"
	"platform/http"
	"platform/http/handling"
	"platform/pipeline"
	"platform/pipeline/basic"
	"platform/services"
	"platform/sessions"
	"sync"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
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
			//handling.HandlerEntry{ "",  CounterHandler{}},
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
