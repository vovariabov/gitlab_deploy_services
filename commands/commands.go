package commands

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/pkg/errors"
)

const (
	gitRepoPath = "git@%v:%v/%v.git"
	master      = "master"
	dev         = "dev"
	staging     = "staging"
)

func Initialize(Domain, Group string) Collection {
	return &Command{
		Domain:   Domain,
		Group:    Group,
		basePath: GoPathSrc() + Domain + "/" + Group + "/",
		baseRepo: fmt.Sprintf(gitRepoPath, Domain, Group, "%v"),
	}
}

type Collection interface {
	Clone(string) error
	Branch(string) (branches []string, err error)
}

type Command struct {
	Domain   string
	Group    string
	basePath string
	baseRepo string
}

func (c *Command) Clone(msName string) (err error) {
	return execute(exec.Command("git", "clone", fmt.Sprintf(c.baseRepo, msName), c.basePath+msName))
}

func (c *Command) Branch(msName string) (branches []string, err error) {
	var path = c.basePath + msName
	var errb, outb bytes.Buffer
	cmd := exec.Command("git", "branch")
	cmd.Dir = path
	cmd.Stderr = &errb
	cmd.Stdout = &outb
	err = cmd.Run()
	if err != nil {
		return nil, errors.Wrap(err, errb.String())
	}

	branches = regexp.MustCompile("[^\\s|*]+").FindAllString(outb.String(), -1)
	return
}

func (c *Command) checkOut(msName, targetBranch string) (err error) {
	return execute(exec.Command("git", "checkout", targetBranch), c.basePath+msName)
}

func (c *Command) merge(path string, sourceBranch string, targetBranch string) (err error) {
	err = c.checkOut(path, targetBranch)
	if err != nil {
		return
	}
	return execute(exec.Command("git", "merge", sourceBranch), path)
}

func (c *Command) DeployToStaging(msName string) (err error) {
	var path = c.basePath + msName
	err = c.merge(path, dev, staging)
	if err != nil {
		return err
	}
	return c.pushChanges(path)
}

func (c *Command) DeployToProduction(msName string) (err error) {
	var path = c.basePath + msName
	err = c.merge(path, staging, master)
	if err != nil {
		return err
	}
	return c.pushChanges(path)
}

func (c *Command) pushChanges(path string) (err error) {
	return execute(exec.Command("git", "push", "origin", getCurrentBranch(path)), path)

}

//git branch | grep \* | cut -d ' ' -f2
func getCurrentBranch(path string) string {
	var errb, outb bytes.Buffer
	cmd := exec.Command("bash", "-c", "git branch | grep \\* | cut -d ' ' -f2")
	cmd.Dir = path
	cmd.Stderr = &errb
	cmd.Stdout = &outb
	cmd.Run()
	return outb.String()
}

func GoPathSrc() string {
	return os.Getenv("GOPATH") + "/src/"
}

func execute(cmd *exec.Cmd, t ...string) (err error) {
	if len(t) != 0 {
		cmd.Dir = t[0]
	}
	var errb, outb bytes.Buffer
	cmd.Stderr = &errb
	cmd.Stdout = &outb
	err = cmd.Run()
	fmt.Println(outb)
	if err != nil {
		return errors.Wrap(err, errb.String())
	}
	return
}
