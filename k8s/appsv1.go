package k8s

import (
	"context"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) ListAppsV1(ctx context.Context, namespace, resource string, listOptions metav1.ListOptions) (interface{}, error) {
	if c.ClientSet == nil {
		return nil, nil
	}

	switch resource {
	case "deploys":
		return c.ClientSet.AppsV1().Deployments(namespace).List(ctx, listOptions)
	case "daemonsets":
		return c.ClientSet.AppsV1().DaemonSets(namespace).List(ctx, listOptions)
	case "statefulsets":
		return c.ClientSet.AppsV1().StatefulSets(namespace).List(ctx, listOptions)
	case "replicasets":
		return c.ClientSet.AppsV1().ReplicaSets(namespace).List(ctx, listOptions)
	default:
		return nil, errors.New("invalid operation provided")
	}

}
