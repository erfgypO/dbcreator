package dbcreator

import (
	_ "embed"
	"fmt"
	"github.com/melbahja/goph"
	"log"
	"os"
	"os/exec"
	"strings"
)

//go:embed templates/stateful-set.template.yml
var statefulSetTemplate string

//go:embed templates/persistent-volume.template.yml
var persistentVolumeTemplate string

//go:embed templates/service.template.yml
var serviceTemplate string

func CreateVolumeDir(host string, user string, keyPath string, keyPassword string, volumeDir string) error {
	auth, err := goph.Key(keyPath, keyPassword)
	if err != nil {
		return err
	}

	client, err := goph.New(user, host, auth)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Run(fmt.Sprintf("mkdir -p %s", volumeDir))
	if err != nil {
		return err
	}

	return nil
}

func CreateVolume(projectName string) error {
	log.Println("Creating volume...")

	volumeConfig := strings.ReplaceAll(persistentVolumeTemplate, "{{PROJECT}}", projectName)

	ymlFile, err := os.Create("volume.yml")
	if err != nil {
		return err
	}
	defer func() {
		_ = ymlFile.Close()
		_ = os.Remove("volume.yml")
	}()
	_, err = ymlFile.Write([]byte(volumeConfig))
	if err != nil {
		return err
	}

	cmd := exec.Command("kubectl", "apply", "-f", "volume.yml")
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Error creating volume: %s", string(out))
		return err
	}

	log.Printf("%s", string(out))
	return nil
}

func CreateNamespace(namespace string) error {
	log.Println("Creating namespace...")
	cmd := exec.Command("kubectl", "create", "namespace", namespace)
	out, err := cmd.CombinedOutput()
	if err != nil {
		errorMessage := string(out)

		if !strings.Contains(errorMessage, "Error from server (AlreadyExists):") {
			log.Printf("Error creating namespace: %s", string(out))
			return err
		} else {
			log.Println("Namespace already exists")
			return nil
		}
	}

	log.Printf("%s", string(out))
	return nil
}

func CreateStatefulSet(projectName string, namespace string) error {
	log.Println("Creating stateful set...")
	statefulSetConfig := strings.ReplaceAll(statefulSetTemplate, "{{PROJECT}}", projectName)

	ymlFile, err := os.Create("stateful-set.yml")
	if err != nil {
		return err
	}
	defer func() {
		_ = ymlFile.Close()
		_ = os.Remove("stateful-set.yml")
	}()

	_, err = ymlFile.Write([]byte(statefulSetConfig))
	if err != nil {
		return err
	}

	cmd := exec.Command("kubectl", "apply", "-f", "stateful-set.yml", "-n", namespace)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Error creating stateful set: %s", string(out))
		return err
	}

	log.Printf("%s", string(out))
	return nil
}

func CreateService(projectName string, namespace string) error {
	log.Println("Creating service...")
	serviceConfig := strings.ReplaceAll(serviceTemplate, "{{PROJECT}}", projectName)

	ymlFile, err := os.Create("service.yml")
	if err != nil {
		return err
	}
	defer func() {
		_ = ymlFile.Close()
		_ = os.Remove("service.yml")
	}()

	_, err = ymlFile.Write([]byte(serviceConfig))
	if err != nil {
		return err
	}

	cmd := exec.Command("kubectl", "apply", "-f", "service.yml", "-n", namespace)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Error creating service: %s", string(out))
		return err
	}

	log.Printf("%s", string(out))
	return nil
}
