package main

import (
	"encoding/json"
	"io/ioutil"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/fsouza/go-dockerclient"
)

type Volume struct {
	ID          string "json:ID"
	Attached    bool   "json: Attached, omitempty"
	Path        string "json: Path"
	IsBindMount bool   "json: IsBindMount"
	Writable    bool   "json: Writeable"
}

type Volumes map[string]Volume

func (v *Volumes) GetVolumes(volumeDir string) error {
	// Get all Docker volumes from Disk.
	files, err := ioutil.ReadDir(volumeDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		log.Infof("Found volume: %v", f.Name())
		filePath := path.Join(volumeDir, f.Name(), "config.json")

		log.Infof("Reading Config: %s", filePath)
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Errorf("%v", err)
			return err
		}

		volume := &Volume{}
		err = json.Unmarshal(fileContent, &volume)
		if err != nil {
			log.Errorf("%v", err)
			return err
		}

		(*v)[volume.Path] = *volume
		log.Debugf("Volume path: %v", volume.Path)
	}

	err = v.setAttachedVolumes()
	if err != nil {
		return err
	}

	return nil
}

func (v *Volumes) setAttachedVolumes() error {
	client, _ := docker.NewClient(dockerSocket)

	existingContainers, err := client.ListContainers(
		docker.ListContainersOptions{
			All: true,
		})

	if err != nil {
		return err
	}

	// loop over existing containers
	for _, container := range existingContainers {
		containerInfo, _ := client.InspectContainer(container.ID)
		for _, val := range containerInfo.Volumes {
			if _, exists := (*v)[val]; exists {
				log.Infof("Attached: %v", val)
				volume := (*v)[val]
				volume.Attached = true
				(*v)[val] = volume
			}
		}
	}

	return nil
}
