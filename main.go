package main

import (
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetLevel(log.DebugLevel)
	hook := filename.NewHook()
	s, err := newServer(func(s *server) error {
		//do nothing, for future
		return nil
	})

	if err != nil {
		panic(err)
	}
	port := "8086"
	log.AddHook(hook)
	log.Infof("Http Server Listen on %v\n", port)
	log.Errorf("Error when HTTP Listen on %v : %v", port, http.ListenAndServe(":"+port, s))

}

