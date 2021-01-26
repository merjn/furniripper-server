package handlers

import (
	"fmt"
	"net/http"

	"github.com/merjn/furniripper-server/service"
)

var ErrSwfNameNotFound = []byte("swf name not found")
var ErrIconLocationNotFound = []byte("icon_location not found")
var ErrHeightNotFound = []byte("height not found")
var ErrWidthNotFound = []byte("width not found")
var ErrLengthNotFound = []byte("length not found")
var ErrSwfContentNotFound = []byte("swf content not found")

type AddFurniHandler struct {
	FurniService *service.Furni
}

// Handle gets all data from the request and passes it to the furni service facade.
func (a AddFurniHandler) Handle(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	swfName := req.FormValue("swf_name")
	if swfName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrSwfNameNotFound)
		return
	}

	swfContent := req.FormValue("swf_content")
	if swfContent == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrSwfContentNotFound)
		return
	}

	swfIcon := req.FormValue("icon_location")
	if swfIcon == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrIconLocationNotFound)
		return
	}

	furniHeight := req.FormValue("furni_height")
	if furniHeight == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrHeightNotFound)
		return
	}

	furniWidth := req.FormValue("furni_width")
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

	err := a.FurniService.AddFurni(swfName, swfContent, swfIcon, furniWidth, furniLength, furniHeight)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
