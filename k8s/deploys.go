package k8s

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Client) Deploys(ctx context.Context, namespace string, lo metav1.ListOptions) (*appsv1.DeploymentList, error) {
	if c.ClientSet == nil {
		return nil, nil
	}

	return c.ClientSet.AppsV1().Deployments(namespace).List(ctx, lo)
}
