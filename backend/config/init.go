package config

import (
	"os"
	"toTheMoon/backend/constants"
)

var env ENV
var isLocal bool

func Init(setEnv ENV) {
	initialize(setEnv)
	setMysqlPassword()
}

func initialize(setEnv ENV) {

	var err error
	// value stored in eb environment
	_env := os.Getenv(constants.ENV)

	switch _env {
	case string(Dev):
		env = Dev
	case string(Prod):
		env = Prod
	default:

		isLocal = true
		env = setEnv
	}

	// read config by using viper

	v, err = readConfig(string(env))
	if err != nil {
		panic(err)
	}
}

func IsLocal() bool {
	return isLocal
}

var mysqlPassword string

func setMysqlPassword() {
	if IsLocal() {
		mysqlPassword = ""
	} else {

		// use it when deployed

		//secretName := fmt.Sprintf(mysqlSecretNameFmt, env)
		//
		//secretStr, _ := secretsmanager.GetSecret(secretName)
		//
		//var secret struct {
		//	Password string `json:"password"`
		//}
		//if err := json.Unmarshal([]byte(secretStr), &secret); err != nil {
		//	panic(cerror.UnMarshallError(err))
		//}
		//
		//mysqlPassword = secret.Password
	}
}

func MySqlPassword() string {
	return mysqlPassword
}
