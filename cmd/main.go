package main

import (
	"flag"
	"fmt"
	"github.com/erfgypO/dbcreator"
	"log"
	"os"
)

func main() {
	var project string
	flag.StringVar(&project, "project", "", "Project name")
	var namespace string
	flag.StringVar(&namespace, "namespace", "default", "Namespace")
	var host string
	flag.StringVar(&host, "host", "", "Host")
	var user string
	flag.StringVar(&user, "user", "", "User")
	var key string
	flag.StringVar(&key, "key", "", "Key (default \"\")")
	var keyPassword string
	flag.StringVar(&keyPassword, "key-password", "", "Key password")

	flag.Parse()

	if project == "" {
		log.Println("Project name is required")
		os.Exit(1)
	}
	if host == "" {
		log.Println("Host is required")
		os.Exit(1)
	}
	if user == "" {
		log.Println("User is required")
		os.Exit(1)
	}
	if key == "" {
		log.Println("Key is required")
		os.Exit(1)
	}

	volumeDir := fmt.Sprintf("/k8s-data/%s", project)
	err := dbcreator.CreateVolumeDir(host, user, key, keyPassword, volumeDir)
	if err != nil {
		panic(err)
	}

	err = dbcreator.CreateVolume(project)
	if err != nil {
		log.Panic(err)
	}

	err = dbcreator.CreateNamespace(namespace)
	if err != nil {
		log.Panic(err)
	}

	err = dbcreator.CreateStatefulSet(project, namespace)
	if err != nil {
		log.Panic(err)
	}

	err = dbcreator.CreateService(project, namespace)
	if err != nil {
		log.Panic(err)
	}
}
