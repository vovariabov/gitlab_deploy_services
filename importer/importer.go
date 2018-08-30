package importer

import "fmt"


const gitRepoPath = "git@%v:%v/%v.git"
type Importer interface {
	Import(branch string) error
}

type GitLabPackage struct {
	Name    string
	Domain  string
	Group   string
	gitRepo string
}

func (g *GitLabPackage) Import(branch string) error{
	g.repo()
	return nil
}

func (g *GitLabPackage) repo() {
	g.gitRepo = fmt.Sprintf(gitRepoPath, g.Domain, g.Group, g.Name)
}

func (g *GitLabPackage) GetPath() string{
	return g.Domain+"/"+g.Group+"/"+g.Name
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