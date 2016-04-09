package env

import (
	"os"
	"path/filepath"
)

func GetProjectRoot() string{
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if(err != nil){
		return ""
	}
	return dir
}

func GetEnvValue( envName string) string{
	return os.Getenv(envName)
}

func SetEnvValue( key string, value string) error{
	err := os.Setenv(key, value)
	if(err != nil){
		return err
	}
	return nil
}

func GetAllEnv() []string  {
	return os.Environ()
}