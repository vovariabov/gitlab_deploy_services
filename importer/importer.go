package importer

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/vovariabov/gitlab_deploy_services/commands"
)

const (
	DOMAIN       = "gitlab.qarea.org"
	GROUP        = "tgms"
	TGMSDEPLOY   = "tgms-deploy"
	gitRepoPath  = "git@%v:%v/%v.git"
	existsErrMsg = "already exists"
)

func InitImporter() Importer {
	return &GitLabPackage{}
}

type Importer interface {
	Import() error
	Branch() ([]string, error)
}

var c = commands.Initialize(DOMAIN, GROUP)

type GitLabPackage struct {
	Name     string
	Domain   string
	Group    string
	imported bool
}

func (g *GitLabPackage) Import() (err error) {
	err = c.Clone(g.Name)
	if err != nil && !cloneExistsErr(err) {
		//handle cloneExistsErr -> git fetch
		return
	}
	g.imported = true
	return nil
}

func (g *GitLabPackage) Branch() (branches []string, err error) {
	if !g.imported {
		return nil, errors.New("not imported")
	}
	branches, err = c.Branch(g.Name)
	return
}

//func (g *GitLabPackage) Merge(targetBranch, sourceBranch string) (err error) {
//	if !g.imported {
//		return errors.New("not imported")
//	}
//	return c.Merge(g.Name, sourceBranch, targetBranch)
//}

func (g *GitLabPackage) GetPath() string {
	return commands.GoPathSrc() + g.Domain + "/" + g.Group + "/" + g.Name
}

func Import(domain, group, name string) (*GitLabPackage, error) {
	var g = GitLabPackage{
		Name:   name,
		Group:  group,
		Domain: domain,
	}
	err := g.Import()
	return &g, err
}

func cloneExistsErr(err error) bool {
	return strings.Contains(err.Error(), existsErrMsg)
}
