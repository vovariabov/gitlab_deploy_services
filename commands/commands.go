package commands

import (
	"os/exec"
	"bytes"
	"github.com/pkg/errors"
	"regexp"
)

func Initialize() Collection {
	return &Command{}
}

type Collection interface {
	Clone(repo string, path string) error
	Branch(path string) (branches []string, err error)
	CheckOut(path string, targetBranch string) error
	Merge(path string, sourceBranch string, targetBranch string) error
}

type Command struct{}

func (c *Command) Clone(repo, path string) (err error) {
	var errb bytes.Buffer
	cmd := exec.Command("git", "clone", repo, path)
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, errb.String())
	}
	return
}

func (c *Command) Branch(path string) (branches []string, err error) {
	var errb, outb bytes.Buffer
	cmd := exec.Command("git", "branch")
	cmd.Dir = path
	cmd.Stderr = &errb
	cmd.Stdout = &outb
	err = cmd.Run()
	if err != nil {
		return nil, errors.Wrap(err, errb.String())
	}
	if err != nil {
		return nil, err
	}
	branches = regexp.MustCompile("[^\\s|*]+").FindAllString(outb.String(), -1)
	return
}

func (c *Command) CheckOut(path, targetBranch string) (err error) {
	var errb, outb bytes.Buffer
	cmd := exec.Command("git", "checkout", targetBranch)
	cmd.Dir = path
	cmd.Stderr = &errb
	cmd.Stdout = &outb
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, errb.String())
	}
	return
}

func (c *Command) Merge(path string, sourceBranch string, targetBranch string) (err error) {
	err = c.CheckOut(path, targetBranch)
	var errb, outb bytes.Buffer
	cmd := exec.Command("git", "merge", sourceBranch)
	cmd.Dir = path
	cmd.Stderr = &errb
	cmd.Stdout = &outb
	err = cmd.Run()
	if err != nil {
		return errors.Wrap(err, errb.String())
	}
	return
}