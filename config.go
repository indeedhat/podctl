package main

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/pelletier/go-toml"
)

const (
	ConfigFileName        = ".podctl.toml"
	DefaultConfigPath     = ".config/podctl"
	DefaultEditorFallback = "vi"
)

var DefaultEditors = []string{"nvim", "vim", "vi", "nano"}

type Config struct {
	Pod  PodConfig
	Env  EnvConfig
	Logs LogConfig
}

type PodConfig struct {
	Name      string
	Namespace string `default:"default"`
}

type EnvConfig struct {
	ConfigDir string `toml:"config_dir"`
	Editor    string
}

type LogConfig struct {
	Prefix LogIndex `default:"index"`
}

// loadConfig will attempt to unmarshal the config file in the current dir
func loadConfig() (*Config, error) {
	data, err := ioutil.ReadFile(configPath())

	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = toml.Unmarshal(data, conf)

	if err != nil {
		return nil, err
	}

	if conf.Pod.Name == "" {
		return nil, errors.New("invalid config (pod.name not set)")
	}

	applyConfigDefaults(conf)

	return conf, nil
}

// configPath will build the path the the current dirs config file
func configPath() string {
	pwd, err := os.Getwd()

	if err != nil {
		return ConfigFileName
	}

	return path.Join(pwd, ConfigFileName)
}

// applyConfigDefaults will apply default values to any config param not set
func applyConfigDefaults(conf *Config) {
	userDir, _ := os.UserHomeDir()

	if conf.Env.ConfigDir == "" {
		conf.Env.ConfigDir = path.Join(userDir, DefaultConfigPath, conf.Pod.Name)
	} else if strings.HasPrefix(conf.Env.ConfigDir, "~") {
		conf.Env.ConfigDir = userDir + conf.Env.ConfigDir[1:]
	}

	if conf.Env.Editor == "" {
		conf.Env.Editor = os.Getenv("EDITOR")
	}

	if conf.Env.Editor == "" {
		for _, editor := range DefaultEditors {
			if _, err := exec.LookPath(editor); err != nil {
				conf.Env.Editor = editor
				break
			}
		}
	}

	if conf.Env.Editor == "" {
		conf.Env.Editor = DefaultEditorFallback
	}
}
