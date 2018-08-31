package main

import (
	"github.com/docopt/docopt-go"
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

	usage := `GitLab Deploy Services
	Usage:  deploy_util [-v] <command> [<args>...]
			deploy_util import (--all | <service>...)
			deploy_util deploy_to_staging (--all | <service>...)
			deploy_util deploy_to_production (--all | <service>...)
	`
	arguments, err := docopt.ParseDoc(usage)

	fmt.Sprintf("%+v", arguments)

	parser := &docopt.Parser{OptionsFirst: true}
	args, err := parser.ParseArgs(usage, nil, "huy")
	if err != nil {
		fmt.Println("ERR:", err)
	}
	fmt.Println(args)

	//tgmsDeploy, err := importer.Import(importer.DOMAIN, importer.GROUP, importer.TGMSDEPLOY)
	if err != nil {
		panic(err)
	}
	//s, err := gitlab_deploy_services.FetchServices(tgmsDeploy)
	//
	//msCollection := ms_object.InitMsObj(s)
	//msCollection.Mss["planningms"].Import()
}
