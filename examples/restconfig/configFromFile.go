package restconfig

import (
	"flag"
	"log"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func ConfigFromFlags() rest.Config {
	defaultKubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
	kubeconfig := flag.String("kubeconfig", defaultKubeconfigPath, "Path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return *config
}