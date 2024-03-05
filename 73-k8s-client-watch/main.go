package main

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	config, _    = clientcmd.BuildConfigFromFlags("", "~/.kube/config")
	clientset, _ = kubernetes.NewForConfig(config)
)

func watchPods() {
	timeOut := int64(60)
	watcher, _ := clientset.CoreV1().
		Pods("dev").
		Watch(context.Background(), metav1.ListOptions{TimeoutSeconds: &timeOut})

	for event := range watcher.ResultChan() {
		item := event.Object.(*corev1.Pod)

		switch event.Type {
		case watch.Modified:
		case watch.Bookmark:
		case watch.Error:
		case watch.Deleted:
		case watch.Added:
			process(item.GetName())
		}
	}
}

func process(name string) {
	log.Info("Some processing for newly created pod: ", name)
}

func main() {
	var wg sync.WaitGroup
	go watchPods()
	wg.Add(1)
	wg.Wait()
}
