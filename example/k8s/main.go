package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var err error
	var config *rest.Config
	// inCluster(Pod).kuberconfig(kubectl)
	var kubeconfig *string

	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(可选)kubeconfig 文件的绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 文件的绝对路径")
	}
	flag.Parse()
	// 首先使用 inCluster 模式(需要区配置对应的RBAC 权限,默认的sa是default-->是没有获取deployment的List权限)
	if config, err = rest.InClusterConfig(); err != nil {
		// 使用Kubeonfig文件配置集群配置Config对象
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
		fmt.Println(config.Host)
	}
	// 已经获得了rest.Config对象
	// 创建Clientset对象
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 就可以使用Clientset对象获取资源对象,进行增删改查
	// 想要获取default命令空间下面的Deployment的列表
	deployment, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for idx, deploy := range deployment.Items {
		fmt.Printf("%d-%s\n", idx, deploy.Name)
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") //windows
}
