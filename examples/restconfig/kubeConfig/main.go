package main

import (
	"flag"
	"fmt"
	"log"

	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// load the path to the kubeconfig file
	kubeconfigPath := flag.String("kubeconfig", "", "Path to your kubeconfig file to load")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfigPath)
	if (err != nil) {
		log.Fatalf(err.Error())
	}

	// pretty print config
	fmt.Printf("%#v\n", config)
}


