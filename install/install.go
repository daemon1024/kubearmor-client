// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeArmor

package install

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/kubearmor/kubearmor-client/k8s"
	"golang.org/x/mod/semver"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func K8sInstaller(c *k8s.Client) error {
	env := autoDetectEnvironment(c)
	if env == "none" {
		return errors.New("unsupported environment or cluster not configured correctly")
	}
	fmt.Printf("Auto Detected Environment : %s\n", env)
	fmt.Printf("CRD %s ...\n", kspName)
	if _, err := CreateCustomResourceDefinition(c, kspName); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
		fmt.Printf("CRD %s already exists ...\n", kspName)
	}
	fmt.Printf("CRD %s ...\n", hspName)
	if _, err := CreateCustomResourceDefinition(c, hspName); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
		fmt.Printf("CRD %s already exists ...\n", hspName)
	}
	fmt.Print("Service Account ...\n")
	if _, err := c.K8sClientset.CoreV1().ServiceAccounts("kube-system").Create(context.Background(), serviceAccount, metav1.CreateOptions{}); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
		fmt.Print("Service Account already exists ...\n")
	}
	fmt.Print("Cluster Role Bindings ...\n")
	if _, err := c.K8sClientset.RbacV1().ClusterRoleBindings().Create(context.Background(), clusterRoleBinding, metav1.CreateOptions{}); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
		fmt.Print("Cluster Role Bindings already exists ...\n")
	}
	fmt.Print("KubeArmor Relay Service ...\n")
	if _, err := c.K8sClientset.CoreV1().Services("kube-system").Create(context.Background(), relayService, metav1.CreateOptions{}); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
		fmt.Print("KubeArmor Relay Service already exists ...\n")
	}
	fmt.Print("KubeArmor Relay Deployment ...\n")
	if _, err := c.K8sClientset.AppsV1().Deployments("kube-system").Create(context.Background(), relayDeployment, metav1.CreateOptions{}); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
		fmt.Print("KubeArmor Relay Deployment already exists ...\n")
	}
	fmt.Print("KubeArmor DaemonSet ...\n")
	if _, err := c.K8sClientset.AppsV1().DaemonSets("kube-system").Create(context.Background(), generateDaemonSet(env), metav1.CreateOptions{}); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
		fmt.Print("KubeArmor DaemonSet already exists ...\n")
	}
	fmt.Print("KubeArmor Policy Manager Service ...\n")
	if _, err := c.K8sClientset.CoreV1().Services("kube-system").Create(context.Background(), policyManagerService, metav1.CreateOptions{}); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
		fmt.Print("KubeArmor Policy Manager Service already exists ...\n")
	}
	fmt.Print("KubeArmor Policy Manager Deployment ...\n")
	if _, err := c.K8sClientset.AppsV1().Deployments("kube-system").Create(context.Background(), policyManagerDeployment, metav1.CreateOptions{}); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
		fmt.Print("KubeArmor Policy Manager Deployment already exists ...\n")
	}
	fmt.Print("KubeArmor Host Policy Manager Service ...\n")
	if _, err := c.K8sClientset.CoreV1().Services("kube-system").Create(context.Background(), hostPolicyManagerService, metav1.CreateOptions{}); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
		fmt.Print("KubeArmor Host Policy Manager Service already exists ...\n")
	}
	fmt.Print("KubeArmor Host Policy Manager Deployment ...\n")
	if _, err := c.K8sClientset.AppsV1().Deployments("kube-system").Create(context.Background(), hostPolicyManagerDeployment, metav1.CreateOptions{}); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			return err
		}
		fmt.Print("KubeArmor Host Policy Manager Deployment already exists ...\n")
	}
	return nil
}

func K8sUninstaller(c *k8s.Client) error {
	fmt.Print("Service Account ...\n")
	if err := c.K8sClientset.CoreV1().ServiceAccounts("kube-system").Delete(context.Background(), serviceAccountName, metav1.DeleteOptions{}); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		fmt.Print("Service Account not found ...\n")
	}
	fmt.Print("Cluster Role Bindings ...\n")
	if err := c.K8sClientset.RbacV1().ClusterRoleBindings().Delete(context.Background(), clusterRoleBindingName, metav1.DeleteOptions{}); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		fmt.Print("Cluster Role Bindings not found ...\n")
	}
	fmt.Print("KubeArmor Relay Service ...\n")
	if err := c.K8sClientset.CoreV1().Services("kube-system").Delete(context.Background(), relayServiceName, metav1.DeleteOptions{}); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		fmt.Print("KubeArmor Relay Service not found ...\n")
	}
	fmt.Print("KubeArmor Relay Deployment ...\n")
	if err := c.K8sClientset.AppsV1().Deployments("kube-system").Delete(context.Background(), relayDeploymentName, metav1.DeleteOptions{}); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		fmt.Print("KubeArmor Relay Deployment not found ...\n")
	}
	fmt.Print("KubeArmor DaemonSet ...\n")
	if err := c.K8sClientset.AppsV1().DaemonSets("kube-system").Delete(context.Background(), "kubearmor", metav1.DeleteOptions{}); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		fmt.Print("KubeArmor DaemonSet not found ...\n")
	}
	fmt.Print("KubeArmor Policy Manager Service ...\n")
	if err := c.K8sClientset.CoreV1().Services("kube-system").Delete(context.Background(), policyManagerServiceName, metav1.DeleteOptions{}); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		fmt.Print("KubeArmor Policy Manager Service not found ...\n")
	}
	fmt.Print("KubeArmor Policy Manager Deployment ...\n")
	if err := c.K8sClientset.AppsV1().Deployments("kube-system").Delete(context.Background(), policyManagerDeploymentName, metav1.DeleteOptions{}); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		fmt.Print("KubeArmor Policy Manager Deployment not found ...\n")
	}
	fmt.Print("KubeArmor Host Policy Manager Service ...\n")
	if err := c.K8sClientset.CoreV1().Services("kube-system").Delete(context.Background(), hostPolicyManagerServiceName, metav1.DeleteOptions{}); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		fmt.Print("KubeArmor Host Policy Manager Service not found ...\n")
	}
	fmt.Print("KubeArmor Host Policy Manager Deployment ...\n")
	if err := c.K8sClientset.AppsV1().Deployments("kube-system").Delete(context.Background(), hostPolicyManagerDeploymentName, metav1.DeleteOptions{}); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		fmt.Print("KubeArmor Host Policy Manager Deployment not found ...\n")
	}
	fmt.Printf("CRD %s ...\n", kspName)
	if err := c.APIextClientset.ApiextensionsV1().CustomResourceDefinitions().Delete(context.Background(), kspName, metav1.DeleteOptions{}); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		fmt.Printf("CRD %s not found ...\n", kspName)
	}
	fmt.Printf("CRD %s ...\n", hspName)
	if err := c.APIextClientset.ApiextensionsV1().CustomResourceDefinitions().Delete(context.Background(), hspName, metav1.DeleteOptions{}); err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
		fmt.Printf("CRD %s not found ...\n", hspName)
	}
	return nil
}

func autoDetectEnvironment(c *k8s.Client) (name string) {
	var env = "none"
	contextName := c.RawConfig.CurrentContext
	clusterContext, exists := c.RawConfig.Contexts[contextName]
	if !exists {
		return env
	}
	clusterName := clusterContext.Cluster

	// Detecting Environment based on cluster name and context
	if clusterName == "minikube" || contextName == "minikube" {
		env = "minikube"
		return env
	}

	if strings.HasPrefix(clusterName, "microk8s-") || contextName == "microk8s" {
		env = "microk8s"
		return env
	}

	if strings.HasPrefix(clusterName, "gke_") {
		env = "gke"
		return env
	}

	// Environment is Self Managed K8s, checking container runtime and it's version

	nodes, _ := c.K8sClientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	containerRuntime := nodes.Items[0].Status.NodeInfo.ContainerRuntimeVersion

	s := strings.Split(containerRuntime, "://")
	runtime := s[0]
	version := "v" + s[1]
	if runtime == "docker" && semver.Compare(version, "v18.9") >= 0 {
		env = "docker"
		return env
	}
	if (runtime == "docker" && semver.Compare(version, "v19.3") >= 0) || runtime == "containerd" {
		env = "generic"
		return env
	}
	return env
}
