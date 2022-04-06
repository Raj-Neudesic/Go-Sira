package main

import (
	"time"
)

type Metadata struct {
	Name string `json:"name"`

	// Namespace where service is created
	Namespace string `json:"namespace"`

	// Time when the service was created
	CreationTimestamp time.Time `json:"creationTimestamp"`

	// List of all labels associated with the service
	Labels map[string]string `json:"labels"`

	// List of all annotations associated with the service
	//Annotations map[string]string `json:"annotations"`
}
type ServiceDto struct {
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
}

type Spec struct {
	Ports []Ports `json:"ports"`
	// Label selector of the service.
	Selector   map[string]string `json:"selector"`
	ClusterIP  string            `json:"clusterIP"`
	Type       string            `json:"type"`
	ExternalIP string            `json:"externalIP"`
}

type Ports struct {
	Protocol string `json:"protocol"`
	Port     int    `json:"port"`
	//TargetPort string `json:"targetPort"`
}
