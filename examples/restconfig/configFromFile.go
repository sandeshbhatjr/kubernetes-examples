package restconfig

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func getDefaultKubeconfigPath() string {
	// kubeconfig is usually at ~/.kube/config
	// or the path is provided in the environment variable: KUBECONFIG
	if path := os.Getenv("KUBECONFIG"); path != "" {
		return path
	} else {
		return filepath.Join(homedir.HomeDir(), ".kube", "config")
	}
}

func ConfigFromFlags() rest.Config {
	defaultKubeconfigPath := getDefaultKubeconfigPath()
	kubeconfig := flag.String("kubeconfig", defaultKubeconfigPath, "Path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return *config
}