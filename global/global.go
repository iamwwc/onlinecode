package global

import (
	"encoding/json"
	"github.com/docker/docker/client"
	"io/ioutil"
	"path/filepath"
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


func DecodeConfigJson(dir string)map[string]string{
	config := filepath.Join(dir,"config.json")
	bs, err := ioutil.ReadFile(config)
	if err != nil{
		panic(err)
	}
	m := make(map[string]string)
	if err := json.Unmarshal(bs,&m); err != nil{
		panic(err)
	}
	return m
}


