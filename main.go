package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/kz3r/k8s-canarinho/model"
	"github.com/kz3r/k8s-canarinho/settings"
)

func CheckRuntimeExceed(pod model.PodResume, limitSeconds int64) error {
	if pod.RuntimeExceededSeconds(limitSeconds) {
		return errors.New("Pod exceed runtime limit")
	}

	return nil
}

func FindOldCanary(pods []model.PodResume, maxTime int32) bool {
	defaultCanarySuffix := "-canary"
	maxTimeInSeconds := int64(maxTime * 60)
	oldCanaryFound := false
	for _, pod := range pods {
		if strings.Contains(pod.Name, defaultCanarySuffix) {
			fmt.Println("Checking pod " + pod.Name)
			if pod.RuntimeExceededSeconds(maxTimeInSeconds) {
				oldCanaryFound = true
				break
			}
		}
	}

	return oldCanaryFound
}

func main() {

	conf, err := settings.ReadConf("conf.yaml")
	if err != nil {
		fmt.Println(err)
	}

	watchedNamespaces := make(map[string]model.NamespaceStatus)
	kubernetesWatcher := model.NewKubernetesWatcher()

	for {
		for _, namespace := range conf.Conf.Namespaces {
			pods := kubernetesWatcher.GetPodsForNamespace(namespace)
			hasOldCanary := FindOldCanary(pods, conf.Conf.CanaryMaxTimeInMinutes)
			status := model.NamespaceStatus{
				Namespace:    namespace,
				Pods:         pods,
				HasOldCanary: hasOldCanary,
			}

			watchedNamespaces[namespace] = status
		}

		for _, namespace := range watchedNamespaces {
			if namespace.HasOldCanary {
				beeep.Alert("Old canary running...", "Check namespace ["+namespace.Namespace+"]", "assets/canary-icon.svg")
			}
		}

		time.Sleep(time.Minute * 30)
	}
}
