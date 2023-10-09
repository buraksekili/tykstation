package client

import (
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) GetCoreV1(ctx context.Context, namespace, name, resource string, opts metav1.GetOptions) (interface{}, error) {
	if c.ClientSet == nil {
		return nil, nil
	}

	switch resource {
	case "pods":
		return c.ClientSet.CoreV1().Pods(namespace).Get(ctx, name, opts)
	case "services":
		return c.ClientSet.CoreV1().Services(namespace).Get(ctx, name, opts)
	case "secrets":
		return c.ClientSet.CoreV1().Secrets(namespace).Get(ctx, name, opts)
	case "configmaps":
		return c.ClientSet.CoreV1().ConfigMaps(namespace).Get(ctx, name, opts)
	case "endpoints":
		return c.ClientSet.CoreV1().Endpoints(namespace).Get(ctx, name, opts)
	case "events":
		return c.ClientSet.CoreV1().Events(namespace).Get(ctx, name, opts)
	case "namespaces":
		return c.ClientSet.CoreV1().Namespaces().Get(ctx, name, opts)
	case "nodes":
		return c.ClientSet.CoreV1().Nodes().Get(ctx, name, opts)
	case "pvcs":
		return c.ClientSet.CoreV1().PersistentVolumeClaims(namespace).Get(ctx, name, opts)
	case "pvs":
		return c.ClientSet.CoreV1().PersistentVolumes().Get(ctx, name, opts)
	case "replicacontrollers":
		return c.ClientSet.CoreV1().ReplicationControllers(namespace).Get(ctx, name, opts)
	case "serviceaccounts":
		return c.ClientSet.CoreV1().ServiceAccounts(namespace).Get(ctx, name, opts)
	default:
		return nil, errors.New("invalid operation provided")
	}

}

func (c *Client) ListCoreV1(ctx context.Context, namespace, resource string, listOptions metav1.ListOptions) (interface{}, error) {
	if c.ClientSet == nil {
		return nil, nil
	}

	switch resource {
	case "pods":
		return c.ClientSet.CoreV1().Pods(namespace).List(ctx, listOptions)
	case "services":
		return c.ClientSet.CoreV1().Services(namespace).List(ctx, listOptions)
	case "secrets":
		return c.ClientSet.CoreV1().Secrets(namespace).List(ctx, listOptions)
	case "configmaps":
		return c.ClientSet.CoreV1().ConfigMaps(namespace).List(ctx, listOptions)
	case "endpoints":
		return c.ClientSet.CoreV1().Endpoints(namespace).List(ctx, listOptions)
	case "events":
		return c.ClientSet.CoreV1().Events(namespace).List(ctx, listOptions)
	case "namespaces":
		return c.ClientSet.CoreV1().Namespaces().List(ctx, listOptions)
	case "nodes":
		return c.ClientSet.CoreV1().Nodes().List(ctx, listOptions)
	case "pvcs":
		return c.ClientSet.CoreV1().PersistentVolumeClaims(namespace).List(ctx, listOptions)
	case "pvs":
		return c.ClientSet.CoreV1().PersistentVolumes().List(ctx, listOptions)
	case "replicacontrollers":
		return c.ClientSet.CoreV1().ReplicationControllers(namespace).List(ctx, listOptions)
	case "serviceaccounts":
		return c.ClientSet.CoreV1().ServiceAccounts(namespace).List(ctx, listOptions)
	default:
		return nil, errors.New("invalid operation provided")
	}

}
