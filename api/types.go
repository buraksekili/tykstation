package api

var (
	coreV1Types = []string{
		"pods", "services", "secrets", "configmaps", "endpoints",
		"events", "namespaces", "nodes", "pvcs", "pvs", "replicacontrollers", "serviceaccounts",
	}

	appsV1Types = []string{"deploys", "daemonsets", "statefulsets", "replicasets"}
)
