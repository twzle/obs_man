package obsman

import (
	hcore "git.miem.hse.ru/hubman/hubman-lib/core"
)

// Represenation of OBS configuration entity
type ObsConf struct {
	HostPort            string `yaml:"host_port" json:"host_port"`
	Password            string `yaml:"password" json:"password"`
	HealthCheckInterval int    `yaml:"health_check_interval" json:"health_check_interval"`
}

// Representation of user configuration entity
type Conf struct {
	ObsConf ObsConf `yaml:"obs" json:"obs"`
}

var _ hcore.Configuration = &Conf{}

// Function validating user configuration
func (o Conf) Validate() error {
	return nil
}
