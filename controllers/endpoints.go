package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetEndpoint(w http.ResponseWriter, req *http.Request) {
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
	} // Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Use the content
	bodyString := string(bodyBytes)
	fmt.Printf("DATA %s\n", bodyString)
	json.NewEncoder(w).Encode(req)
}

func CreateEndpoint(w http.ResponseWriter, req *http.Request) {
	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(req.Body)
	} // Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Use the content
	bodyString := string(bodyBytes)
	fmt.Printf("DATA %s\n", bodyString)

	json.NewEncoder(w).Encode(req)
}
