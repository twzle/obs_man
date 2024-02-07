package obsman

import (
	hcore "git.miem.hse.ru/hubman/hubman-lib/core"
)

type ObsConf struct {
	HostPort            string `yaml:"host_port" json:"host_port"`
	Password            string `yaml:"password" json:"password"`
	HealthCheckInterval int    `yaml:"health_check_interval" json:"health_check_interval"`
}

type Conf struct {
	ObsConf ObsConf `yaml:"obs" json:"obs"`
}

var _ hcore.Configuration = &Conf{}

func (o Conf) Validate() error {
	return nil
}
