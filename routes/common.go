package routes

var gpath string
var userDir string
var gUser *User
var linkQuery string = ""

var dbIp, dbPort string

func InitDB(ip string, port string) {
	dbIp = ip
	dbPort = port
}
