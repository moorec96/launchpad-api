package http_res

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type HttpResponse struct {
	StatusCode int
	Messages   []interface{}
}

func GenerateHttpResponse(code int, messages ...interface{}) *HttpResponse {
	var messagesArray []interface{}
	for _, msg := range messages {
		messagesArray = append(messagesArray, msg)
	}
	return &HttpResponse{
		StatusCode: code,
		Messages:   messagesArray,
	}
}

func SendHttpResponse(writer http.ResponseWriter, res *HttpResponse) {
	if res.StatusCode < 300 {
		sendHttpErrorWithError(writer, res.StatusCode, res.Messages...)
	} else {
		var okMessage interface{} = ""
		if len(res.Messages) >= 1 {
			okMessage = res.Messages[0]
		}
		sendHttpOk(writer, res.StatusCode, okMessage)
	}
}

func sendHttpErrorWithError(writer http.ResponseWriter, code int, errors ...interface{}) {
	var errorStrings []string
	for _, err := range errors {
		errString := err.(error).Error()
		errString = strings.ToUpper(string(errString[0])) + errString[1:]
		errorStrings = append(errorStrings, errString)
	}
	sendHttpError(writer, code, errorStrings...)
}

func sendHttpError(writer http.ResponseWriter, code int, errors ...string) {
	output := make(map[string]interface{})

	if len(errors) == 0 {
		_ = append(errors, "Could not process request")
	}

	output["errors"] = errors

	writer.Header().Add("Content-Type", "application/json")
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Request-Method", "*")
	writer.Header().Add("Access-Control-Allow-Headers", "*")
	writer.Header().Add("Cache-Control", "no-cache")
	writer.Header().Add("Connection", "keep-alive")

	writer.WriteHeader(code)
	marsh, err := json.Marshal(output)

	if err != nil {
		log.Printf("Could not marshal HTTP error message with code %d", code)
	}
	_, err = writer.Write(marsh)
	if err != nil {
		log.Printf("Could not writer to http request with code %d", code)
	}
}

func sendHttpOk(writer http.ResponseWriter, code int, out interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Header().Add("Access-Control-Request-Method", "*")
	writer.Header().Add("Access-Control-Allow-Headers", "*")
	writer.Header().Add("Cache-Control", "no-cache")
	writer.Header().Add("Connection", "keep-alive")

	output, err := json.Marshal(out)

	if err != nil {
		log.Printf("Could not marshal HTTP error message with code %d", code)
	}
	_, err = writer.Write(output)
	if err != nil {
		log.Printf("Could not writer to http request with code %d", code)
	}
}
