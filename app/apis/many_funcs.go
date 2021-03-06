package apis

import (
	"app/utility"
	"log"
	"net/http"
	"time"
)

/*
ManyFuncsHandler ...
*/
type ManyFuncsHandler struct{}

func (h *ManyFuncsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("many funcs ... ")
	//
	utility.XrayMiddle(request.Context(), "funcOne", func() error {
		funcOne()
		return nil
	})
	//
	utility.XrayMiddle(request.Context(), "funcTwo", func() error {
		funcTwo()
		return nil
	})
	writer.WriteHeader(http.StatusOK)
}

func funcOne() {
	log.Println("do func one")
	time.Sleep(500 * time.Microsecond)
}

func funcTwo() {
	log.Println("do func two")
	time.Sleep(1000 * time.Microsecond)
}
