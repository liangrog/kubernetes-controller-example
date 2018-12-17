package main

import (
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	fooClientset "github.com/liangrog/kube-crd-example/pkg/client/clientset/versioned"
	fooInformer "github.com/liangrog/kube-crd-example/pkg/client/informers/externalversions/foo/v1"
)

func main() {
	kubeConfigPath := os.Getenv("HOME") + "/.kube/config"

	// create the config from the path
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}
	// generate the client based off of the config
	/*client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}*/

	fooClient, err := fooClientset.NewForConfig(config)
	if err != nil {
		log.Fatalf("getClusterConfig: %v", err)
	}

	_ = fooInformer.NewHelloTypeInformer(
		fooClient,
		metav1.NamespaceAll,
		0,
		cache.Indexers{},
	)

}
