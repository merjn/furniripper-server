package service

import (
	"fmt"
	"encoding/base64"
)

type Furni struct {
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

	fmt.Println(swfContentDecoded)
	fmt.Println(iconContentDecoded)

	// Move swf content to furnidata

	// Move icon content to icon data	

	// Add to catalogue & furnidata
	
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	return nil
}