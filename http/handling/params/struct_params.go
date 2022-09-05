package params

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
)

func getFileFromMultipartForm(funcType reflect.Type, form *http.Request) (params []reflect.Value, err error) {
	err = form.ParseMultipartForm(20 << 20)
	file, handler, err := form.FormFile("file")
	if err != nil {
		err = errors.New("file_not_found")
		return
	} else if filepath.Ext(handler.Filename) != ".cex" {
		err = errors.New("bad_file_ext")
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	datafile := []string{string(data)}

	if len(datafile) == funcType.NumIn()-1 {
		params = make([]reflect.Value, funcType.NumIn()-1)
		for i := 0; i < len(datafile); i++ {
			params[i], err = parseValueToType(funcType.In(i+1), datafile[i])
			if err != nil {
				return
			}
		}
	} else {
		err = errors.New("Parameter number mismatch")
	}

	return
}

func populateStructFromForm(structVal reflect.Value,
	formVals map[string][]string) (err error) {
	for i := 0; i < structVal.Elem().Type().NumField(); i++ {
		field := structVal.Elem().Type().Field(i)
		for key, vals := range formVals {
			if strings.EqualFold(key, field.Name) && len(vals) > 0 {
				valField := structVal.Elem().Field(i)
				if valField.CanSet() {
					valToSet, convErr := parseValueToType(valField.Type(), vals[0])
					if convErr == nil {
						valField.Set(valToSet)
					} else {
						err = convErr
					}
				}
			}
		}
	}
	return
}

func populateStructFromJSON(structVal reflect.Value,
	reader io.ReadCloser) (err error) {
	return json.NewDecoder(reader).Decode(structVal.Interface())
}
