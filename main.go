package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// Initialize Kubernetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error creating in-cluster config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// Query pod state
	namespace := "default" // Change this to the namespace you want to query
	pods, err := getPods(clientset, namespace)
	if err != nil {
		log.Fatalf("Error getting pods: %v", err)
	}

	// Print pod states
	for _, pod := range pods.Items {
		fmt.Printf("Pod Name: %s, Phase: %s\n", pod.Name, pod.Status.Phase)
	}
}

func getPods(clientset *kubernetes.Clientset, namespace string) (*v1.PodList, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods, nil
}
