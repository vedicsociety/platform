package placeholder

import (
	"fmt"
	"io"
	"net/http"

	"github.com/vedicsociety/platform/http/actionresults"
	"github.com/vedicsociety/platform/http/handling"
	"github.com/vedicsociety/platform/http/handling/params"
	"github.com/vedicsociety/platform/logging"
	"github.com/vedicsociety/platform/validation"
)

type UploadHandler struct {
	logging.Logger
	handling.URLGenerator
	validation.Validator
}

func (handler UploadHandler) GetUpload() actionresults.ActionResult {
	return actionresults.NewTemplateAction("upload.html", struct {
		Title string
	}{

		Title: "IngestCEX title",
	})
}

func (n UploadHandler) PostUpload(params params.InputParams, files params.Files) actionresults.ActionResult {
	n.Logger.Debugf("PostName method invoked with params %v, files %v ", params, files)

	return n.redirectOrError(UploadHandler.GetUpload)
}

// func (n NameHandler) GetUpload() actionresults.ActionResult {
// 	n.Logger.Debug("GetNames method invoked")
// 	return actionresults.NewTemplateAction("simple_message.html", names)
// }

func HandleMultipartForm(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10000000)
	fmt.Fprintf(writer, "Name: %v, City: %v\n",
		request.MultipartForm.Value["name"][0],
		request.MultipartForm.Value["city"][0])
	fmt.Fprintln(writer, "------")

	for _, header := range request.MultipartForm.File["files"] {
		fmt.Fprintf(writer, "Name: %v, Size: %v\n", header.Filename, header.Size)
		file, err := header.Open()
		if err == nil {
			defer file.Close()
			fmt.Fprintln(writer, "------")
			io.Copy(writer, file)
		} else {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (n UploadHandler) redirectOrError(handler interface{},
	data ...interface{}) actionresults.ActionResult {
	url, err := n.GenerateUrl(handler)
	if err == nil {
		return actionresults.NewRedirectAction(url)
	} else {
		return actionresults.NewErrorAction(err)
	}
}
