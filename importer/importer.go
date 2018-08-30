package importer

import (
	"fmt"
	"os"
	"github.com/vovariabov/gitlab_deploy_services/commands"
	"strings"
	"github.com/pkg/errors"
)


const (
	gitRepoPath  = "git@%v:%v/%v.git"
	existsErrMsg = "already exists"
)
type Importer interface {
	Import() error
	Branch() ([]string, error)
}

var c = commands.Initialize()

type GitLabPackage struct {
	Name     string
	Domain   string
	Group    string
	path     string
	gitRepo  string
	imported bool
}

func (g *GitLabPackage) Import() (err error) {
	g.gitRepo = fmt.Sprintf(gitRepoPath, g.Domain, g.Group, g.Name)
	g.path = gopathSrc()+g.Domain+"/"+g.Group+"/"+g.Name
	err = c.Clone(g.gitRepo, g.path)
	if err != nil && !cloneExistsErr(err) {
		return
	}
	g.imported = true
	return nil
}

func (g *GitLabPackage) Branch() (branches []string, err error) {
	if !g.imported {
		return nil, errors.New("not imported")
	}
	branches, err = c.Branch(g.path)
	return
}

func (g *GitLabPackage) Merge(targetBranch, sourceBranch string) (err error) {
	if !g.imported {
		return errors.New("not imported")
	}
	return c.Merge(g.path, sourceBranch, targetBranch)
}

func (g *GitLabPackage) GetPath() string{
	return g.path
}

func Import(domain, group, name string) (*GitLabPackage, error) {
	var g = GitLabPackage{
		Name: name,
		Group: group,
		Domain: domain,
	}
	err := g.Import()
	return &g, err
}

func gopathSrc() string {
	return  os.Getenv("GOPATH")+"/src/"
}

func cloneExistsErr(err error) bool {
	return strings.Contains(err.Error(), existsErrMsg) 
}