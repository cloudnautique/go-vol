package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/cloudnautique/go-vol/volumes"
)

const (
	volumeDir = "/var/lib/docker/volumes"
)

func main() {
	vols := &volumes.Volumes{}
	err := vols.GetVolumes(volumeDir)
	err = vols.DeleteAllOrphans(false)
	if err != nil {
		log.Fatal("%v", err)
	}
}
