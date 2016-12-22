package goutil

import (
	"errors"
	"io/ioutil"
	"strconv"

	"gopkg.in/yaml.v2"
)

// IConfig interface for checking version and general config validation
type IConfig interface {
	GetVer() int
	Check() error
}

// NewConfig for parsing and validating YAML config
func NewConfig(file string, c IConfig, vers []int) error {
	if buf, err := ioutil.ReadFile(file); err != nil {
		msg := StrCat("Failed to open config: ", err.Error())
		return errors.New(msg)
	} else {
		if err = yaml.Unmarshal(buf, c); err != nil {
			msg := StrCat("Invalid config format: ", err.Error())
			return errors.New(msg)
		}
	}
	if !MemberOfSlice(c.GetVer(), vers) {
		msg := StrCat("Unsupported config version ", strconv.Itoa(c.GetVer()), ": need ", JoinInt(vers, ", "))
		return errors.New(msg)
	}

	return c.Check()
}
