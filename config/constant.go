package config

import (
	"github.com/zhaowanda/go.web.formwork/yaml"
	//"os/exec"
	"os"
)

var (
	execDirAbsPath, _ = os.Getwd()
	YamlConfig, Err   = yaml.InitYamlConfig(execDirAbsPath + "/src/application.yaml")
)
