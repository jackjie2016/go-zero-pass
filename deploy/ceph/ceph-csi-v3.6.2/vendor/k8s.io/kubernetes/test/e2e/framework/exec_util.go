/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	"bytes"
	"context"
	"io"
	"net/url"
	"strings"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"

	"github.com/onsi/gomega"
)

// ExecOptions passed to ExecWithOptions
type ExecOptions struct {
	Command       []string
	Namespace     string
	PodName       string
	ContainerName string
	Stdin         io.Reader
	CaptureStdout bool
	CaptureStderr bool
	// If false, whitespace in std{err,out} will be removed.
	PreserveWhitespace bool
	Quiet              bool
}

// ExecWithOptions executes a command in the specified container,
// returning stdout, stderr and error. `options` allowed for
// additional parameters to be passed.
func (f *Framework) ExecWithOptions(options ExecOptions) (string, string, error) {
	if !options.Quiet {
		Logf("ExecWithOptions %+v", options)
	}
	config, err := LoadConfig()
	ExpectNoError(err, "failed to load restclient config")

	const tty = false

	Logf("ExecWithOptions: Clientset creation")
	req := f.ClientSet.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(options.PodName).
		Namespace(options.Namespace).
		SubResource("exec").
		Param("container", options.ContainerName)
	req.VersionedParams(&v1.PodExecOptions{
		Container: options.ContainerName,
		Command:   options.Command,
		Stdin:     options.Stdin != nil,
		Stdout:    options.CaptureStdout,
		Stderr:    options.CaptureStderr,
		TTY:       tty,
	}, scheme.ParameterCodec)

	var stdout, stderr bytes.Buffer
	Logf("ExecWithOptions: execute(POST %s %s)", req.URL())
	err = execute("POST", req.URL(), config, options.Stdin, &stdout, &stderr, tty)
	if options.PreserveWhitespace {
		return stdout.String(), stderr.String(), err
	}
	return strings.TrimSpace(stdout.String()), strings.TrimSpace(stderr.String()), err
}

// ExecCommandInContainerWithFullOutput executes a command in the
// specified container and return stdout, stderr and error
func (f *Framework) ExecCommandInContainerWithFullOutput(podName, containerName string, cmd ...string) (string, string, error) {
	return f.ExecWithOptions(ExecOptions{
		Command:            cmd,
		Namespace:          f.Namespace.Name,
		PodName:            podName,
		ContainerName:      containerName,
		Stdin:              nil,
		CaptureStdout:      true,
		CaptureStderr:      true,
		PreserveWhitespace: false,
	})
}

// ExecCommandInContainer executes a command in the specified container.
func (f *Framework) ExecCommandInContainer(podName, containerName string, cmd ...string) string {
	stdout, stderr, err := f.ExecCommandInContainerWithFullOutput(podName, containerName, cmd...)
	Logf("Exec stderr: %q", stderr)
	ExpectNoError(err,
		"failed to execute command in pod %v, container %v: %v",
		podName, containerName, err)
	return stdout
}

// ExecShellInContainer executes the specified command on the pod's container.
func (f *Framework) ExecShellInContainer(podName, containerName string, cmd string) string {
	return f.ExecCommandInContainer(podName, containerName, "/bin/sh", "-c", cmd)
}

func (f *Framework) execCommandInPod(podName string, cmd ...string) string {
	pod, err := f.PodClient().Get(context.TODO(), podName, metav1.GetOptions{})
	ExpectNoError(err, "failed to get pod %v", podName)
	gomega.Expect(pod.Spec.Containers).NotTo(gomega.BeEmpty())
	return f.ExecCommandInContainer(podName, pod.Spec.Containers[0].Name, cmd...)
}

func (f *Framework) execCommandInPodWithFullOutput(podName string, cmd ...string) (string, string, error) {
	pod, err := f.PodClient().Get(context.TODO(), podName, metav1.GetOptions{})
	ExpectNoError(err, "failed to get pod %v", podName)
	gomega.Expect(pod.Spec.Containers).NotTo(gomega.BeEmpty())
	return f.ExecCommandInContainerWithFullOutput(podName, pod.Spec.Containers[0].Name, cmd...)
}

// ExecShellInPod executes the specified command on the pod.
func (f *Framework) ExecShellInPod(podName string, cmd string) string {
	return f.execCommandInPod(podName, "/bin/sh", "-c", cmd)
}

// ExecShellInPodWithFullOutput executes the specified command on the Pod and returns stdout, stderr and error.
func (f *Framework) ExecShellInPodWithFullOutput(podName string, cmd string) (string, string, error) {
	return f.execCommandInPodWithFullOutput(podName, "/bin/sh", "-c", cmd)
}

func execute(method string, url *url.URL, config *restclient.Config, stdin io.Reader, stdout, stderr io.Writer, tty bool) error {
	exec, err := remotecommand.NewSPDYExecutor(config, method, url)
	if err != nil {
		return err
	}
	return exec.Stream(remotecommand.StreamOptions{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
		Tty:    tty,
	})
}
