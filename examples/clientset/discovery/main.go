package main

import (
	"fmt"
	"log"

	"k8s.io/client-go/kubernetes"

	"github.com/sandeshbhatjr/kubernetes-examples/examples/restconfig"
)

// The discovery options is very useful to retrieve information about existing API groups and their versions
//
// For example, the kubectl CLI uses the discovery to fetch the shortnames
// in order to make it easier for the user query API resources with their shortnames
// Say you did kubectl get pods, kubectl first fetches all the api discovery information
// so it knows you mean core/v1.Pods

func printAllAPIGroups(clientset kubernetes.Interface) {
	discoveryClient := clientset.Discovery()
	apiGroups, err := discoveryClient.ServerGroups()
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, apiGroup := range(apiGroups.Groups) {
		fmt.Printf("%#v: ", apiGroup.Name)
		for i, groupVersion := range(apiGroup.Versions) {
			if i >= 1 {
				fmt.Printf(", ")
			}
			fmt.Printf("%#v", groupVersion.Version)
		}
		fmt.Printf("\n")
	}
}

func main() {
	config := restconfig.ConfigFromFlags()
	clientset, err := kubernetes.NewForConfig(&config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	printAllAPIGroups(clientset)
	
}