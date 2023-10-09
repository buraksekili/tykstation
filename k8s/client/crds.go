package client

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (c *Client) GetCRDs(ctx context.Context) (interface{}, error) {
	crds, err := c.crdClientSet.
		ApiextensionsV1().
		CustomResourceDefinitions().
		List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return crds, err
}

func (c *Client) GetCRs(ctx context.Context, ns, name, group, version, resource string) (interface{}, error) {
	if c.clientSet == nil {
		return nil, nil
	}

	var gvr = schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}

	cr, err := c.dynamicClient.Resource(gvr).Namespace(ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return cr, nil
}
