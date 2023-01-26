package model

type NamespaceStatus struct {
	Namespace    string
	Pods         []PodResume
	HasOldCanary bool
}
