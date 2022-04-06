package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	_ "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"reflect"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	var kubeconfig *string
	// Connect to K8 server
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "/Users/rajaxceltron/.kube")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// List HPA
	hpas, err := clientset.AutoscalingV1().HorizontalPodAutoscalers("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	var hpavals []hpadto
	b1, _ := json.Marshal(hpas.Items)
	err = json.Unmarshal(b1, &hpavals)
	if err != nil {
		fmt.Println(err)
	}

	//List Service
	svc, err := clientset.CoreV1().Services("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println("%s", svc)
	var svcvals []ServiceDto
	b, _ := json.Marshal(svc.Items)
	//fmt.Println(string(b))
	err = json.Unmarshal(b, &svcvals)
	if err != nil {
		fmt.Println(err)
	}

	//Deployment list
	deploymentlist, err := clientset.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})
	var depvals []depdto
	if err != nil {
		fmt.Println(err)
	}
	b2, _ := json.Marshal(deploymentlist.Items)
	err = json.Unmarshal(b2, &depvals)
	for _, v := range svcvals {
		for _, v1 := range depvals {
			if reflect.DeepEqual(v.Spec.Selector, v1.Spec.Selector.MatchLabels) == true {
				for _, v2 := range hpavals {
					if v1.Metadata.Name == v2.Spec.ScaleTargetRef.Name {
						//fmt.Println(v2)
						b3, _ := json.Marshal(v2)
						fmt.Println(string(b3))
					}
				}
			}
		}
	}

}
