package commands

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"

	"os"

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
	DeployToProduction(msName string) (err error)
	DeployToStaging(msName string) (err error)
	GitFetch(path string)
}

type Command struct {
	Domain   string
	Group    string
	basePath string
	baseRepo string
}

func GoPathSrc() string {
	return os.Getenv("GOPATH") + "/src/"
}

func (c *Command) Clone(msName string) (err error) {
	return execute(exec.Command("git", "clone", fmt.Sprintf(c.baseRepo, msName), c.basePath+msName), &bytes.Buffer{})
}

func (c *Command) Branch(msName string) (branches []string, err error) {
	var outb bytes.Buffer
	err = execute(exec.Command("git", "branch"), &outb, c.basePath+msName)
	if err != nil {
		return nil, err
	}
	branches = regexp.MustCompile("[^\\s|*]+").FindAllString(outb.String(), -1)
	return
}

func (c *Command) DeployToStaging(msName string) (err error) {
	var path = c.basePath + msName
	c.GitFetch(path)
	err = c.merge(path, dev, staging)
	if err != nil {
		return err
	}
	return c.pushChanges(path)
}

func (c *Command) DeployToProduction(msName string) (err error) {
	var path = c.basePath + msName
	c.GitFetch(path)
	err = c.merge(path, staging, master)
	if err != nil {
		return err
	}
	return c.pushChanges(path)
}

func (c *Command) GitFetch(path string) {
	execute(exec.Command("git", "fetch"), &bytes.Buffer{}, path)
}

func (c *Command) checkOut(path, targetBranch string) (err error) {
	return execute(exec.Command("git", "checkout", targetBranch), &bytes.Buffer{}, path)
}

func (c *Command) merge(path string, sourceBranch string, targetBranch string) (err error) {
	err = c.checkOut(path, targetBranch)
	if err != nil {
		return
	}
	return execute(exec.Command("git", "merge", sourceBranch), &bytes.Buffer{}, path)
}

func (c *Command) pushChanges(path string) error {
	return execute(exec.Command("git", "push", "origin", getCurrentBranch(path)), &bytes.Buffer{}, path)

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

func execute(cmd *exec.Cmd, outb *bytes.Buffer, t ...string) (err error) {
	if len(t) != 0 {
		cmd.Dir = t[0]
	}
	var errb bytes.Buffer
	cmd.Stderr = &errb
	cmd.Stdout = outb
	err = cmd.Run()
	fmt.Println(cmd.Args)
	//fmt.Println(outb)
	if err != nil {
		return errors.Wrap(err, errb.String())
	}
	return
}
