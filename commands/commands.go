package commands

import (
	"fmt"
	"os/exec"
)

func Initialize() Collection {
	return &Command{}
}

type Collection interface {
	Clone(msName string) (err error)
}

type Command struct{}

func (c *Command) Clone(msName string) (err error) {
	arg := fmt.Sprintf("git@gitlab.qarea.org:tgms/%s.git", msName)
	osCmd := exec.Command("git clone", arg)
	var out []byte
	out, err = osCmd.Output()
	if err != nil {
		return
	}
	fmt.Println(string(out))
	return
}
