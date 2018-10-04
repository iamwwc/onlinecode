package main

import (
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	log.SetLevel(log.DebugLevel)
	file := prepareLogFile("runlog.log")
	multiWriter := io.MultiWriter(file,os.Stdout)
	log.SetOutput(multiWriter)
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

func prepareLogFile(name string)*os.File{
	workDir := os.Getenv("CODE_WORK_DIR")
	if workDir == ""{
		panic("cannot get env CODE_WORK_DIR")
	}
	currentDir := filepath.Join(workDir,name)
	file, err := os.OpenFile(currentDir,os.O_CREATE | os.O_WRONLY | os.O_APPEND, 755)
	if err != nil{
		panic(err)
	}
	return file
}

