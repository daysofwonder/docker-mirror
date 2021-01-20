// +build !darwin

package main

import (
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
	log "github.com/sirupsen/logrus"
)

func getDockerCredentials(registry string) (*docker.AuthConfiguration, error) {
	dockerlog := log.WithField("registry", registry)

	// Creds from docker config file
	dockerlog.Debug("Get credentials from docker config file")
	authOptions, err := docker.NewAuthConfigurationsFromDockerCfg()
	if err != nil {
		dockerlog.Debug("No docker configuration found")
	} else {
		creds, ok := authOptions.Configs[registry]
		if ok {
			return &creds, nil
		}
		dockerlog.Debug("No creds found")
	}

	// Creds from docker credentials helpers
	dockerlog.Debug("Get credentials from docker credentials helpers")
	creds, err := docker.NewAuthConfigurationsFromCredsHelpers(registry)
	if err != nil {
		dockerlog.Debug(err)
	} else {
		return creds, nil
	}

	return nil, fmt.Errorf("No auth found for %s", registry)
}
