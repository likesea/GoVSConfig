package main

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	MachineName string
	Environment string
}

type Serverslice struct {
	Servers []Server
}

func GetEnvironment(dir string) string {
	var s Serverslice
	envMap := s.getSettings(dir)
	osName, _ := os.Hostname()
	var env string
	if v, ok := envMap[osName]; ok {
		env = v
	}
	if v, ok := envMap["default"]; ok {
		env = v
	}
	return env
}
func (s *Serverslice) getSettings(dir string) map[string]string {
	//fmt.Println(dir + "\\settings.json")
	fi, err := os.Open(dir + "\\settings.json")
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	json.Unmarshal(fd, &s)
	envMap := make(map[string]string)
	for _, v := range s.Servers {
		envMap[v.MachineName] = v.Environment
	}
	return envMap
}
