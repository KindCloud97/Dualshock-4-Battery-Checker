package bcheck

import (
	"io/ioutil"
	"strings"
)

type Controller struct {
	path string
}

func FindControllers() ([]Controller, error) {
	const path = "/sys/class/power_supply"
	var final []Controller

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "sony_controller_battery") {
			ctrl := Controller{path: path + "/" + file.Name()}
			final = append(final, ctrl)
		}
	}

	return final, nil
}

func (c *Controller) readFile(filename string) (string, error) {
	info, err := ioutil.ReadFile(c.path + "/" + filename)
	if err != nil {
		return "", err
	}

	return string(info), nil
}

//Expects absolute path to controller
func (c *Controller) Capacity() (string, error) {
	return c.readFile("capacity")
}

//Expects absolute path to controller
func (c *Controller) Status() (string, error) {
	return c.readFile("status")
}
