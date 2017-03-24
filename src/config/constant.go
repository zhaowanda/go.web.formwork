package config

import (
	"yaml"
	//"os/exec"
	"os"
)

var (
	execDirAbsPath, _ = os.Getwd()
	YamlConfig, Err = yaml.InitYamlConfig(execDirAbsPath + "/src/application.yaml")
)

