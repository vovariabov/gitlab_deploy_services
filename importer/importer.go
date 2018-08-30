package importer

import (
	"fmt"
	"os/exec"
	"os"
	"bytes"
	"github.com/pkg/errors"
)


const gitRepoPath = "git@%v:%v/%v.git"
type Importer interface {
	Import(branch string) error
}

type GitLabPackage struct {
	Name    string
	Domain  string
	Group   string
	path    string
	gitRepo string
}

func (g *GitLabPackage) Import(branch string) error{
	g.gitRepo = fmt.Sprintf(gitRepoPath, g.Domain, g.Group, g.Name)
	g.path    = gopathSrc()+g.Domain+"/"+g.Group+"/"+g.Name
	var errb bytes.Buffer
	cmd := exec.Command("git", "clone", g.gitRepo, "-b", "dev", g.path)
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		return errors.Wrap(err, errb.String())
	}
	fmt.Println("err:", errb.String())
	return nil
}

func (g *GitLabPackage) GetPath() string{
	return g.path
}

func Import(domain, group, name, branch string) (*GitLabPackage, error) {
	var g = GitLabPackage{
		Name: name,
		Group: group,
		Domain: domain,
	}
	err := g.Import(branch)
	return &g, err
}

func gopathSrc() string {
	return  os.Getenv("GOPATH")+"/src/"
}