package config

import (
	"crypto/tls"
	"github.com/spf13/viper"
	"strings"
)

var v *viper.Viper

func readConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName(filename)
	var configPath string

	if IsLocal() {
		configPath = "./config"
	} else {
		configPath = "./backend/config"
	}
	v.AddConfigPath(configPath)
	err := v.ReadInConfig()
	return v, err
}

func All() interface{} {
	return v.AllSettings()
}

func Get(key string) interface{} {
	return v.Get(key)
}

func GetString(key string) string {
	return v.GetString(key)
}

func GetStringList(key string) []string {
	return v.GetStringSlice(key)
}

func GetInt(key string) int {
	return v.GetInt(key)
}

func GetBool(key string) bool {
	return v.GetBool(key)
}

func GetTLSCertificate(certKey, keyKey string) (certificate tls.Certificate, err error) {
	cert := GetString(certKey)
	key := GetString(keyKey)
	if strings.HasPrefix(cert, "-----") {
		certificate, err = tls.X509KeyPair([]byte(cert), []byte(key))
	} else {
		certificate, err = tls.LoadX509KeyPair(cert, key)
	}
	return
}

func Unmarshal(key string, rawVal interface{}) error {
	return v.UnmarshalKey(key, rawVal)
}

func Set(key string, value string) {
	v.Set(key, value)
}
