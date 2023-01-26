package model

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodResume struct {
	Name      string
	Status    string
	StartTime metav1.Time
}

func (p PodResume) RuntimeSeconds() int64 {
	return time.Now().Unix() - p.StartTime.Unix()
}

func (p PodResume) RuntimeExceededSeconds(seconds int64) bool {
	return p.RuntimeSeconds() > seconds
}
