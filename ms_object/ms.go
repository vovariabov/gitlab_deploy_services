package ms_object

import (
	"github.com/vovariabov/gitlab_deploy_services/importer"
)

func InitMsObj(mss []importer.GitLabPackage) MsObj {
	mapms := make(map[string]importer.Importer)
	for _, val := range mss {
		mapms[val.Name] = &val
	}
	return MsObj{
		Mss: mapms,
	}
}

type MsObj struct {
	Mss map[string]importer.Importer
}
