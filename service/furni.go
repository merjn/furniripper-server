package service

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/merjn/furniripper-server/config"
	"github.com/merjn/furniripper-server/furni"
)

type Furni struct {
	Config config.Config
	Adder furni.Adder
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

	// Add to catalogue & furnidata
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	return nil
}