package goutil

import (
	"errors"
	"fmt"
	"io/ioutil"

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
		msg := fmt.Sprintf("Failed to open config: %s", err)
		return errors.New(msg)
	} else {
		if err = yaml.Unmarshal(buf, c); err != nil {
			msg := fmt.Sprintf("Invalid config format: %s", err)
			return errors.New(msg)
		}
	}
	if !MemberOfSlice(c.GetVer(), vers) {
		msg := fmt.Sprintf("Unsupported config version %d: need %s", c.GetVer(), JoinInt(vers, ", "))
		return errors.New(msg)
	}

	return c.Check()
}
