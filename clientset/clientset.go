package clientset

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfig *string

func NewClientSet() (*kubernetes.Clientset, error) {

	kubeconfig := "/Users/vijaysamanthapuri/Downloads/karpentertest-config.yaml"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("error while loading kubeconfig")
	}
	return kubernetes.NewForConfig(config)
}
