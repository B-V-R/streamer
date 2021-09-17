package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type ErrorDetails struct {
	Message string
}

type Router struct {
	Log *log.Logger
}

func (router *Router)Stream(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		{
			if HasContentType(request, "application/json") {
				buf := bytes.Buffer{}
				io.Copy(&buf, request.Body)
				data := buf.Bytes()

				if json.Valid(data) {
					errCh := make(chan error, 1)
					defer close(errCh)

					file, stdout := New(router.Log, data)
					go file.Write()
					go stdout.Write()

					errorDetails := ErrorDetails{}

					if fileErr := <-file.ErrorCh(); fileErr != nil {
						errorDetails.Message = fileErr.Error()
						data, _ := json.MarshalIndent(errorDetails," ", "")
						writer.Write(data)
						return
					}

					if stdoutErr := <-stdout.ErrorCh(); stdoutErr != nil {
						errorDetails.Message = stdoutErr.Error()
						data, _ := json.MarshalIndent(errorDetails," ", "")
						writer.Write(data)
						return
					}

					writer.WriteHeader(http.StatusOK)
					writer.Write([]byte(`{"Message": "Completed"}`))
				} else {
					fmt.Println("JSON is not valid")
					writer.WriteHeader(http.StatusBadRequest)
					writer.Write([]byte(`{"Message": "JSON is Not Valid"}`))
				}
			} else {
				router.Log.Error("Content Type is not application/json")
				writer.WriteHeader(http.StatusUnsupportedMediaType)
			}
		}

	default:
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte(`{"Error": "Method Not allowed"}`))
	}
}
