package service

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/merjn/furniripper-server/config"
	"github.com/merjn/furniripper-server/furni"
)

type Furni struct {
	Config config.Config
	Adder  furni.Adder
}

// AddFurni adds the furniture to the hotel.
func (f *Furni) AddFurni(swfName, swfContent, iconName, iconContent, x, y, z string) error {
	swfContentDecoded, err := base64.StdEncoding.DecodeString(swfContent)
	if err != nil {
		return err
	}

	iconContentDecoded, err := base64.StdEncoding.DecodeString(iconContent)
	if err != nil {
		return err
	}

	furniFile := fmt.Sprintf("%s\\%s", f.Config.FurniLocation, swfName)
	err = ioutil.WriteFile(furniFile, swfContentDecoded, os.ModePerm)
	if err != nil {
		return err
	}

	iconFile := fmt.Sprintf("%s\\%s", f.Config.IconLocation, iconName)
	err = ioutil.WriteFile(iconFile, iconContentDecoded, os.ModePerm)
	if err != nil {
		return err
	}

	width, err := strconv.Atoi(x)
	if err != nil {
		return err
	}

	length, err := strconv.Atoi(y)
	if err != nil {
		return err
	}

	height, err := strconv.ParseFloat(z, 32)
	if err != nil {
		return err
	}

	entity := furni.Furni{
		Name:   strings.Split(swfName, ".")[0], // remove file extension from swfName
		Width:  width,
		Length: length,
		Height: height,
	}

	err = f.Adder.Add(entity)
	if err != nil {
		return err
	}

	return nil
}
