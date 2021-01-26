package service

import "fmt"

type Furni struct {
}

// AddFurni adds the furniture to the hotel.
func (f *Furni) AddFurni(swfName, swfContent, iconLocation, x, y, z string) error {
	fmt.Println(swfName)
	fmt.Println(swfContent)
	fmt.Println(iconLocation)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	return nil
}