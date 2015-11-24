package goutil

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type IConfig interface {
	GetVer() int
}

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
		msg := fmt.Sprintf("Unsupported config version %s: need %s", c.GetVer(), JoinInt(vers, ", "))
		return errors.New(msg)
	}
	return nil
}
