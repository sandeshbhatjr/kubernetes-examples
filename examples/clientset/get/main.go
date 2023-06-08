package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/sandeshbhatjr/kubernetes-examples/examples/restconfig"	
)

func main() {
	config := restconfig.ConfigFromFlags()
	clientset, err := kubernetes.NewForConfig(&config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// context provides cancellable HTTP queries
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	for _, namespace := range(namespaces.Items) {
		fmt.Printf("%#v\n", namespace.Name)
	}
}