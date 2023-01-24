package actionresults

import (
	"encoding/json"
	"net/http"
)

func NewErrorJsonAction(data interface{}) ActionResult {
	return &ErrorJsonActionResult{data: data}
}

type ErrorJsonActionResult struct {
	data interface{}
}

func (action *ErrorJsonActionResult) Execute(ctx *ActionContext) error {
	ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(ctx.ResponseWriter)
	return encoder.Encode(action.data)
}
