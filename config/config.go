package config

import (
    "encoding/json"
    "fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"
	"path/filepath"
	"regexp"
	"strings"
    "github.com/shen100/golang123/utils"
)

var jsonData map[string]interface{}

func initJSON() {
    bytes, err := ioutil.ReadFile("./config.json")
    if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		os.Exit(-1)
    }

	configStr := string(bytes[:])
	reg       := regexp.MustCompile(`/\*.*\*/`)

	configStr  = reg.ReplaceAllString(configStr, "")
	bytes      = []byte(configStr)

    if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
    }
}

type dBConfig struct {
	Dialect       string
	Database      string
	User          string
	Password      string
	Host          string
	Port          int
	Charset       string
	URL           string
	MaxIdleConns  int    
	MaxOpenConns  int 
}

// DBConfig 数据库相关配置
var DBConfig dBConfig

func initDB() {
	utils.SetStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", 
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)
	DBConfig.URL = url
}

type redisConfig struct {
	Host         string
	Port         int
	URL          string
}

// RedisConfig redis相关配置
var RedisConfig redisConfig

func initRedis() {
	utils.SetStructByJSON(&RedisConfig, jsonData["redis"].(map[string]interface{}))
	url := fmt.Sprintf("%s:%d", RedisConfig.Host, RedisConfig.Port)
	RedisConfig.URL = url
}

type serverConfig struct {
	APIPoweredBy        string
	SiteName            string
	Host                string
	Env                 string
	APIPrefix           string
	UploadImgDir        string
	ImgPath             string
	Port                int
	SessionID           string
	SessionTimeout      int
	PassSalt            string
	MailUser            string  //域名邮箱账号
	MailPass            string  //域名邮箱密码
	MailHost            string  //smtp邮箱域名
	MailPort            int     //smtp邮箱端口
	MailFrom            string  //邮件来源
	Github              string
}

// ServerConfig 服务器相关配置
var ServerConfig serverConfig

func initServer() {
	utils.SetStructByJSON(&ServerConfig, jsonData["go"].(map[string]interface{}))
	
	if ServerConfig.UploadImgDir == "" {
		sep := string(os.PathSeparator)
		execPath := filepath.Dir(os.Args[0])
		pathArr  := []string{"website", "static", "upload", "img"}
		length   := utf8.RuneCountInString(execPath)
		lastChar := execPath[length - 1:]
		if lastChar != sep {
			execPath = execPath + sep	
		}
		execPath = execPath + strings.Join(pathArr, sep)
		fmt.Println(execPath)
		ServerConfig.UploadImgDir = execPath
	}
}

func init() {
	initJSON()
	initDB()
	initRedis()
	initServer()
}