package params

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"
)

// A FileHeader describes a file part of a multipart request.
type FileHeader struct {
	Filename string
	Size     int64
	Content  []byte
}

type Files struct {
	File []FileHeader
}

func (f *Files) AddFile(name string, size int64, data []byte) []FileHeader {
	f.File = append(f.File, FileHeader{Filename: name, Size: size, Content: data})
	return f.File
}

type InputParam struct {
	NameParam string
	Value     []string
}

type InputParams struct {
	InputParam []InputParam
}

func (p *InputParams) AddParam(name string, value []string) []InputParam {
	p.InputParam = append(p.InputParam, InputParam{NameParam: name, Value: value})
	return p.InputParam
}

func getStructFromMultipartForm(funcType reflect.Type, request *http.Request) (params []reflect.Value, err error) {
	err = request.ParseMultipartForm(20 << 20)
	if err != nil {
		err = errors.New("getStructFromMultipartForm: error ParseMultipartForm")
		return
	}

	/////////////////////////////

	//	fmt.Printf("params", params)
	//	Value map[string][]string
	var inpar InputParams
	for name, value := range request.MultipartForm.Value {
		//	fmt.Printf(" value: %v", value)
		inpar.AddParam(name, value)
	}

	var filestruct Files
	var file multipart.File
	var filedata []byte
	for _, header := range request.MultipartForm.File["files"] {

		file, err = header.Open()

		filedata, err = ioutil.ReadAll(file)
		if err == nil {
			defer file.Close()
			filestruct.AddFile(header.Filename, header.Size, filedata)
		} else {
			return
		}
	}
	params = make([]reflect.Value, 2)
	params[0] = reflect.ValueOf(inpar)

	// params = make([]reflect.Value, funcType.NumIn()-1)
	// for i := 0; i < len(datafile); i++ {
	// 	params[i], err = parseValueToType(funcType.In(i+1), datafile[i])
	// 	if err != nil {
	// 		return
	// 	}
	// }
	params[1] = reflect.ValueOf(filestruct)
	//fmt.Printf("params", params)
	return

	////////////////////////////////
	// file, handler, err := request.FormFile("file")
	// if err != nil {
	// 	err = errors.New("getFileFromMultipartForm: file_not_found")
	// 	return
	// } else if filepath.Ext(handler.Filename) != ".cex" {
	// 	err = errors.New("getFileFromMultipartForm: bad_file_extention")
	// 	return
	// }
	// defer file.Close()

	// data, err := ioutil.ReadAll(file)
	// datafile := []string{string(data)}
	// if len(datafile) == funcType.NumIn()-1 {
	// 	params = make([]reflect.Value, funcType.NumIn()-1)
	// 	for i := 0; i < len(datafile); i++ {
	// 		params[i], err = parseValueToType(funcType.In(i+1), datafile[i])
	// 		if err != nil {
	// 			return
	// 		}
	// 	}
	// } else {
	// 	err = errors.New("getFileFromMultipartForm: Parameter number mismatch")
	// }
	// return
}
