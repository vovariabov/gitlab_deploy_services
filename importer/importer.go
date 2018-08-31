package importer

import (
	"github.com/vovariabov/gitlab_deploy_services/commands"
	"strings"
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
	CloneRepo() (err error)
	DeployServiceToStaging() (err error)
	DeployServiceToProduction() (err error)
}

var c = commands.Initialize(DOMAIN, GROUP)

type GitLabPackage struct {
	Name     string
	Domain   string
	Group    string
	imported bool
}

func (g *GitLabPackage) CloneRepo() (err error) {
	err = c.Clone(g.Name)
	if err != nil && !cloneExistsErr(err) {
		return
	}
	c.PullOrigin(g.GetPath())
	g.imported = true
	return nil
}

func (g *GitLabPackage) DeployServiceToStaging() (err error) {
	err = g.CloneRepo()
	if err != nil {
		return
	}
	return c.DeployToStaging(g.Name)
}

func (g *GitLabPackage) DeployServiceToProduction() (err error) {
	err = g.CloneRepo()
	if err != nil {
		return
	}
	return c.DeployToProduction(g.Name)
}

func (g *GitLabPackage) GetPath() string {
	return commands.GoPathSrc() + g.Domain + "/" + g.Group + "/" + g.Name
}

func Import(domain, group, name string) (*GitLabPackage, error) {
	var g = GitLabPackage{
		Name:   name,
		Group:  group,
		Domain: domain,
	}
	err := g.CloneRepo()
	return &g, err
}

func cloneExistsErr(err error) bool {
	return strings.Contains(err.Error(), existsErrMsg)
}
