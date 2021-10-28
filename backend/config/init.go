package config

import "fmt"

type ENV string

const (
	Dev ENV = "dev"
	Prod ENV = "prod"
	Local ENV = "local"
)

var env ENV
var isLocal bool
var mysqlPassword string




func Init(setEnv ENV) {
	initialize(setEnv)
	setMysqlPassword(env)
}

func initialize(setEnv ENV) {

	switch setEnv {
	case Dev :
		env = Dev
	case Prod :
		env = Prod
	default:
		env = setEnv
		isLocal = true
	}

}

func setMysqlPassword(env ENV) {

	// AWS 시크릿매니저, 로컬 , 어디선가 가져와서 변수에 대입
	fetchFromAWSSecretManger := fmt.Sprintf("myungsworld/:%s/mysql",env)

	mysqlPassword = fetchFromAWSSecretManger
}

func MySqlPassword() string {
	return mysqlPassword
}

func IsLocal() bool {
	return isLocal
}