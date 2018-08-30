package main

import (
	"github.com/vovariabov/gitlab_deploy_services/importer"
	"github.com/vovariabov/gitlab_deploy_services"
	"fmt"
)



func main() {

//	usage := `Naval Fate.
//
//Usage:go get github.com/docopt/docopt-go
//  naval_fate ship new <name>...
//  naval_fate ship <name> move <x> <y> [--speed=<kn>]
//  naval_fate ship shoot <x> <y>
//  naval_fate mine (set|remove) <x> <y> [--moored|--drifting]
//  naval_fate -h | --help
//  naval_fate --version
//
//Options:
//  -h --help     Show this screen.
//  --version     Show version.
//  --speed=<kn>  Speed in knots [default: 10].
//  --moored      Moored (anchored) mine.
//  --drifting    Drifting mine.`

//	usage := `GitLab Deploy Services
//Usage:  deploy_util [--version] [--branch=<branch>] <command> [<args>...]
//
//
//		load <file>
//		deploy_util list [--version]
//		deploy_util service [--version] [--branch=<branch>]
//		deploy_util deploy (<service> <branch>)...
//		deploy_util -h | --help
//`
//	arguments, _ := docopt.ParseDoc(usage)
//	fmt.Println(arguments)

	//parser := &docopt.Parser{OptionsFirst: true}
	//args, err := parser.ParseArgs(usage, nil, "huy")
	tgmsDeploy, err := importer.Import(importer.DOMAIN, importer.GROUP, importer.TGMSDEPLOY)
	if err != nil {
		panic(err)
	}
	s, err := gitlab_deploy_services.FetchServices(tgmsDeploy)
	for _, item := range s {
		fmt.Println(item.GetPath())
	}
}


