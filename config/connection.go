package config

import (
	"encoding/json"
	"io/ioutil"
)

func getConfig() (*Connection, error) {
	filename := "./defaultEnvVariables.json"
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	conJSON := Connection{}
	err = json.Unmarshal(fileContent, &conJSON)
	if err != nil {
		return nil, err
	}
	return &conJSON, nil
}
type Connection struct{
	Environment string
	Database string
	Port string
	Server string
	Userid string
	Password string
}

func GetConnection() string{
	c,err := getConfig()
	if (err != nil){
		panic(err)
	}
	return "server="+c.Server+";port="+c.Port+";database="+c.Database+";user id="+c.Userid+";password="+c.Password
}