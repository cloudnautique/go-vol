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
	err = vols.DeleteOrphans(false)
	if err != nil {
		log.Fatal("%v", err)
	}
}
