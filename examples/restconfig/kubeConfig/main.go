package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir" // cross platform consistent way to get user home directory
)

func main() {
	// load the path to the kubeconfig file
	defaultKubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
	kubeconfigPath := flag.String("kubeconfig", defaultKubeconfigPath, "Path to your kubeconfig file to load")
	useProtobuf := flag.Bool("protobuf", false, "Use protobuf as the content type")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfigPath)
	if (err != nil) {
		log.Fatalf(err.Error())
	}

	// modify config based on flags
	if *useProtobuf {
		config.ContentType = "application/protobuf"
	}

	// pretty print config
	fmt.Printf("%#v\n", config)
}


