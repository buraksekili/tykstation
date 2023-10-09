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

func (c *Client) GetCR(ctx context.Context, ns, name string, gvr schema.GroupVersionResource) (interface{}, error) {
	if c.clientSet == nil {
		return nil, nil
	}

	cr, err := c.dynamicClient.Resource(gvr).Namespace(ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return cr, nil
}

func (c *Client) GetCRs(ctx context.Context, ns string, gvr schema.GroupVersionResource) (interface{}, error) {
	if c.clientSet == nil {
		return nil, nil
	}

	crs, err := c.dynamicClient.Resource(gvr).Namespace(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return crs, nil
}
