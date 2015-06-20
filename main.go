package main

import (
	log "github.com/Sirupsen/logrus"
)

const (
	dockerDir    = "/var/lib/docker"
	volumeDir    = "/var/lib/docker/volumes"
	dockerSocket = "unix:///var/run/docker.sock"
)

func main() {
	vols := &Volumes{}
	err := vols.GetVolumes(volumeDir)
	if err != nil {
		log.Fatal("%v", err)
	}

	for k, v := range *vols {
		log.Infof("Key: %v", k)
		log.Infof("Attached: %v", v.Attached)
	}
}
