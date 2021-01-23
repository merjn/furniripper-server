package handlers

import (
	"github.com/merjn/furniripper-server/furni"
	"net/http"
)

type AddFurniHandler struct {
	Adder furni.Adder
}

func (a AddFurniHandler) Handle(w http.ResponseWriter, req *http.Request) {

}

