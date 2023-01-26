package model

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
)

type kubernetesWatcher struct {
	ctx       context.Context
	config    *rest.Config
	clientSet *kubernetes.Clientset
}

func NewKubernetesWatcher() *kubernetesWatcher {
	ctx := context.Background()
	config := ctrl.GetConfigOrDie()
	clientSet := kubernetes.NewForConfigOrDie(config)


	return &kubernetesWatcher{ctx: ctx, config: config, clientSet: clientSet}
}

func (k kubernetesWatcher) GetPodsForNamespace(namespace string) []PodResume {
	var podResume []PodResume
	pods, _ := k.clientSet.CoreV1().Pods(namespace).List(k.ctx, metav1.ListOptions{})
	for _, pod := range pods.Items {
		podResume = append(podResume, PodResume{pod.Name, string(pod.Status.Phase), *pod.Status.StartTime.DeepCopy()})
	}

	return podResume
}
