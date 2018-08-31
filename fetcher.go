package gitlab_deploy_services

import (
	"github.com/vovariabov/gitlab_deploy_services/importer"
	"io/ioutil"
	"github.com/go-yaml/yaml"
)

const (
	serviceListFile = "/service_list/all"
	nameSuffics     = "ms"
)
type AllServices struct {
	Services []string `yaml:"all_services"`
}


func FetchServices(tgmsDeploy *importer.GitLabPackage) (services []importer.GitLabPackage, err error) {
	var (
		path = tgmsDeploy.GetPath()
		list   AllServices
	)
	file, err := ioutil.ReadFile(path + serviceListFile)
	if err != nil {
		return
	}
	err = yaml.Unmarshal([]byte(file), &list)
	if err != nil {
		return
	}
	s := convertNames(list.Services)
	for _, ms := range s {
		service := importer.GitLabPackage{Name:ms, Domain: tgmsDeploy.Domain, Group:tgmsDeploy.Group}
		services = append(services, service)
	}
	return
}
func convertNames(names []string) []string {
	var resNames []string
	for _, n := range names {
		resNames = append(resNames, n + nameSuffics)
	}
	return resNames
}
