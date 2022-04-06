package main

import "time"

//import (
//	"time"
//)

type specdep struct {
	Selector struct {
		MatchLabels map[string]string `json:"matchLabels"`
	} `json:"selector"`
}
type Metadatadep struct {
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
type depdto struct {
	Metadata Metadatadep `json:"metadata"`
	Spec     specdep     `json:"spec"`
}
