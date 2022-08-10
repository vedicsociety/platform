/*
The project doesn’t contain any middleware components that generate responses, which would typically be defined as part of the application.
For the moment, however,this component will generate simple responses as I develop other features.
This component produces a simple text response, which is just enough to ensure that the pipeline works as expected.
Next, create the platform/placeholder/files folder and add to it a file named hello.json with the content shown
	{
		"message": "Hello from the JSON file"
	}
To set the location from which static files will be read, add the setting to the config.json file in the platform folder:
	},
	"files":
	{
		"path": "placeholder/files"
	}
*/
package placeholder

import (
	//"io"
	//"errors"
	"github.com/tsiparinda/platform/config"
	"github.com/tsiparinda/platform/pipeline"

	//"platform/services"
	"github.com/tsiparinda/platform/templates"
)

type SimpleMessageComponent struct {
	Message string
	config.Configuration
}

func (lc *SimpleMessageComponent) ImplementsProcessRequestWithServices() {}

func (c *SimpleMessageComponent) Init() {
	c.Message = c.Configuration.GetStringDefault("main:message",
		"Default Message")
}

func (c *SimpleMessageComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	executor templates.TemplateExecutor) {
	err := executor.ExecTemplate(ctx.ResponseWriter,
		"simple_message.html", c.Message)
	if err != nil {
		ctx.Error(err)
	} else {
		next(ctx)
	}
}
