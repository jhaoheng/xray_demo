package apis

import (
	"app/utility"
	"encoding/json"
	"log"
	"net/http"
)

/*
SQLByXrayWithSuccess ...
*/
type SQLByXrayWithSuccess struct{}

func (h *SQLByXrayWithSuccess) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("SQLByXrayWithSuccess")

	obj, err := utility.SQLByXrayWithSuccess(request.Context())
	if err != nil {
		writer.WriteHeader(http.StatusBadGateway)
		writer.Write([]byte(err.Error()))
		return
	}
	b, _ := json.Marshal(obj)
	writer.Write(b)
	writer.WriteHeader(http.StatusOK)
}
