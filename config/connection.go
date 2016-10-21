package config

import (
	"encoding/json"
	"io/ioutil"
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"fmt"
)
type Connection struct{
	Environment string
	Database string
	Port string
	Server string
	Userid string
	Password string
	db *sql.DB
}
var _connection = Connection{}

func init(){
	fmt.Println("[connection]_connection",_connection)
	err := getConfig()
	fmt.Println("[connection]_connection",_connection)
	if (err != nil){
		panic(err)
	}

}

func getConfig() ( error) {
	filename := "./defaultEnvVariables.json"
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fileContent, &_connection)
	if err != nil {
		return err
	}
	return nil
}

func getConnectionString() string{
	fmt.Println("[getConnectionString]db",_connection)

	return "server="+_connection.Server+";port="+_connection.Port+";database="+_connection.Database+";user id="+_connection.Userid+";password="+_connection.Password
}




func Connect() *sql.DB{
	fmt.Println("[connection]Concect is called")
	fmt.Println("[connection]db",_connection.db)

	_connection.db, _ = sql.Open("mssql", getConnectionString())

	return _connection.db
}

func Disconnect(){
	fmt.Println("[connection]Disconnect is called")
	_connection.db.Close()

}