package handlers

import (
	"net/http"

	"github.com/merjn/furniripper-server/service"
)

var ErrSwfLocationNotFound = []byte("swf_location not found")
var ErrIconLocationNotFound = []byte("icon_location not found")
var ErrHeightNotFound = []byte("height not found")
var ErrWidthNotFound = []byte("width not found")
var ErrLengthNotFound = []byte("length not found")

type AddFurniHandler struct {
	FurniService *service.Furni
}

// Handle gets all data from the request and passes it to the furni service facade.
func (a AddFurniHandler) Handle(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	swfLocation := req.Form.Get("swf_location")
	if swfLocation == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrSwfLocationNotFound)
		return
	}

	swfIcon := req.Form.Get("icon_location")
	if swfIcon == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrIconLocationNotFound)
		return
	}

	furniHeight := req.Form.Get("furni_height")
	if furniHeight == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrHeightNotFound)
		return
	}

	furniWidth := req.Form.Get("furni_width")
	if furniWidth == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrWidthNotFound)
		return
	}

	furniLength := req.Form.Get("furni_length")
	if furniLength == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrLengthNotFound)
		return
	}

	err := a.FurniService.AddFurni(swfLocation, swfIcon, furniWidth, furniLength, furniHeight)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
