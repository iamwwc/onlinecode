package global

import (
	"github.com/docker/docker/client"
)

const (
	CompilePath = "/api/compile"
	SharePath = "/api/share"
	FromSharePath = "/share/"
	IndexPath = "/"
	VolumeName = "onlinecode_onlinecodetemp"
	VolumePath = "/onlinecodetemp/"
)

var DockerClient , err = client.NewClientWithOpts(client.WithVersion("1.38"))


type Request struct{
	Version string
	Files map[string]map[string]string
	Lang string
}




//Response is
type Response struct{
	Error string
	Events []Event
}

type Event struct{
	Kind string
	Message string
}




