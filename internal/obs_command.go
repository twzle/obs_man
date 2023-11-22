package obsman

import (
	"errors"
)

type AgentConf struct {
	ExecutorPort uint16 `yaml:"server_port"`
	RedisUrl     string `yaml:"redis_url"`
}

type ObsConf struct {
	HostPort string `yaml:"host_port"`
	Password string `yaml:"password"`
}

type Conf struct {
	Agent   AgentConf `yaml:"agent"`
	ObsConf ObsConf   `yaml:"obs"`
}

func (o ObsConf) Validate() error {
	if o.HostPort == "" {
		return errors.New("empty host port")
	}
	return nil
}

type SetScene struct {
	SceneName string `hubman:"scene_name"`
}

func (s SetScene) Code() string {
	return "SetScene"
}

func (s SetScene) Description() string {
	return "Sets scene with given name active in obs"
}

type StartRecord struct {
}

func (s StartRecord) Code() string {
	return "StartRecord"
}

func (s StartRecord) Description() string {
	return "Toggles record, if it is started - is noop. Similar to start record button"
}

type StopRecord struct {
}

func (s StopRecord) Code() string {
	return "StopRecord"
}

func (s StopRecord) Description() string {
	return "Toggles off record, if it is off - is noop. Similar to stop record button"
}
