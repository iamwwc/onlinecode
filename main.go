package main

import (
	"chaochaogege.com/onlinecode/global"
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
)

func main() {
	log.SetLevel(log.DebugLevel)
	file := prepareLogFile("runlog.log")
	multiWriter := io.MultiWriter(file,os.Stdout)
	log.SetOutput(multiWriter)
	hook := filename.NewHook()
	c := global.DecodeConfigJson("./")
	s, err := newServer(func(s *server) error {
		str := c["max_run_container_numbers"]
		if counts , err:= strconv.Atoi(str); err != nil{
			panic("bad config, error in config.json max_run_container_numbers, only allow int")
		}else {
			num := int32(counts)
			atomic.StoreInt32(&s.maxCount,num)
		}
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
	file, err := os.OpenFile(name,os.O_CREATE | os.O_WRONLY | os.O_APPEND, 755)
	if err != nil{
		panic(err)
	}
	return file
}

