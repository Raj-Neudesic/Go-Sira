package main

import (
	"time"
)

type Metadatahpa struct {
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
type spechpa struct {
	ScaleTargetRef struct {
		Kind       string `json:"kind"`
		Name       string `json:"name"`
		ApiVersion string `json:"apiVersion"`
	} `json:"scaleTargetRef"`
	MinReplicas                    int `json:"minReplicas"`
	MaxReplicas                    int `json:"maxReplicas"`
	TargetCPUUtilizationPercentage int `json:"targetCPUUtilizationPercentage"`
}

type hpadto struct {
	Metadata Metadatahpa `json:"metadata"`
	Spec     spechpa     `json:"spec"`
}
