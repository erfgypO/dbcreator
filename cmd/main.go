package main

import (
	"fmt"
	"github.com/erfgypO/dbcreator"
	"log"
)

func main() {
	projectName := "sample-187"
	volumeDir := fmt.Sprintf("/k8s-data/%s", projectName)
	err := dbcreator.CreateVolumeDir("501st.tech", "root", "/Users/jhell/.ssh/id_ed25519", "", volumeDir)
	if err != nil {
		panic(err)
	}

	err = dbcreator.CreateVolume(projectName)
	if err != nil {
		log.Panic(err)
	}

	err = dbcreator.CreateNamespace("devdevdev")
	if err != nil {
		log.Panic(err)
	}

	err = dbcreator.CreateStatefulSet(projectName, "devdevdev")
	if err != nil {
		log.Panic(err)
	}

	err = dbcreator.CreateService(projectName, "devdevdev")
	if err != nil {
		log.Panic(err)
	}
}
