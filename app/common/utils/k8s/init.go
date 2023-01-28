package k8s

import (
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func NewK8sClient() (*kubernetes.Clientset, error) {
	//下载 kubectl:https://kubernetes.io/docs/tasks/tools/#tabset-2
	//macos:
	// 1.curl -LO "https://dl.k8s.io/release/v1.21.0/bin/darwin/amd64/kubectl"
	// 2.chmod +x ./kubectl
	// 3.sudo mv ./kubectl /usr/local/bin/kubectl
	// 4.sudo chown root: /usr/local/bin/kubectl
	// 5.kubectl version --client
	// 6.集群模式下直接拷贝服务端~/.kube/config 文件到本机 ~/.kube/confg 中
	//   注意：- config中的域名要能解析正确
	//        - 生产环境可以创建另一个证书
	// 7.kubectl get ns 查看是否正常
	//
	//创建k8s连接
	//在集群外部使用
	//-v /Users/cap/.kube/config:/root/.kube/config
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "kubeconfig file 在当前系统中的地址")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig file 在当前系统中的地址")
	}
	flag.Parse()
	//创建 config 实例
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		logx.Error(err.Error())
		return nil, err
	}

	//在集群中使用
	//config , err := rest.InClusterConfig()
	//if err!=nil {
	//	panic(err.Error())
	//}

	//创建程序可操作的客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logx.Error(err.Error())
	}
	return clientset, err
}
