package main

import (
	"github.com/NYTimes/logrotate"
	logging "github.com/op/go-logging"
	"os"
)

var (
	log    = logging.MustGetLogger("example")
	format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ [%{level:.5s}] %{id:03x}%{color:reset} %{message}`)
)

func main() {
	logfile, err := logrotate.NewFile("sample.log")
	if err != nil {
		log.Fatal(err)
	}

	//w := bufio.NewWriter(logfile)
	backend1 := logging.SetBackend(logging.NewLogBackend(logfile, "", 0))
	backend2 := logging.NewLogBackend(os.Stdout, "", 0)

	logging.SetFormatter(format)
	logging.SetBackend(backend1, backend2)

	log.Debug("Hello Debug Log!!!")

}
