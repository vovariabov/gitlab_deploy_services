package ms_object

import (
	"gitlab_deploy_services/importer"
)

func InitMsObj(mss []importer.GitLabPackage) MsObj {
	mapms := make(map[string]importer.Importer)
	for _, val := range mss {
		v := val
		mapms[v.Name] = &v
	}
	return MsObj{
		Mss: mapms,
	}
}

type MsObj struct {
	Mss map[string]importer.Importer
}
