package client

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func (c *Client) ListCRD(ctx context.Context) (interface{}, error) {
	crds, err := c.crdClientSet.
		ApiextensionsV1().
		CustomResourceDefinitions().
		List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return crds, err
}

// GetCRD gets the detail of CustomResourceDefinition specified by given name. The name MUST be in the format
// <.spec.name>.<.spec.group>. See following for the doc:
//
//	https://github.com/kubernetes/apiextensions-apiserver/blob/ee7666a3e09f4241647a3cafa0c64fbc5ee8dc99/pkg/apis/apiextensions/v1/types.go#L359-L360
func (c *Client) GetCRD(ctx context.Context, name string) (interface{}, error) {
	crd, err := c.crdClientSet.
		ApiextensionsV1().
		CustomResourceDefinitions().
		Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return crd, err
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
