package client

import (
	"context"
	"errors"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func (c *Client) GetAppsV1(
	ctx context.Context,
	namespace, name, resource string,
	opts metav1.GetOptions,
) (interface{}, error) {
	if c.ClientSet == nil {
		return nil, nil
	}

	switch resource {
	case "deploys":
		return c.ClientSet.AppsV1().Deployments(namespace).Get(ctx, name, opts)
	case "daemonsets":
		return c.ClientSet.AppsV1().DaemonSets(namespace).Get(ctx, name, opts)
	case "statefulsets":
		return c.ClientSet.AppsV1().StatefulSets(namespace).Get(ctx, name, opts)
	case "replicasets":
		return c.ClientSet.AppsV1().ReplicaSets(namespace).Get(ctx, name, opts)
	default:
		return nil, errors.New("invalid operation provided")
	}

}

func (c *Client) WatchAppsV1(
	ctx context.Context,
	namespace, resource string,
	opts metav1.ListOptions,
) (watch.Interface, error) {
	if c.ClientSet == nil {
		return nil, nil
	}

	defer func() { fmt.Println("watch interface worked") }()

	switch resource {
	case "deploys":
		return c.ClientSet.AppsV1().Deployments(namespace).Watch(ctx, opts)
	case "daemonsets":
		return c.ClientSet.AppsV1().DaemonSets(namespace).Watch(ctx, opts)
	case "statefulsets":
		return c.ClientSet.AppsV1().StatefulSets(namespace).Watch(ctx, opts)
	case "replicasets":
		return c.ClientSet.AppsV1().ReplicaSets(namespace).Watch(ctx, opts)
	default:
		return nil, errors.New("invalid operation provided")
	}

}

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
