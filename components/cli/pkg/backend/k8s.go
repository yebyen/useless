package backend

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

type K8sBackend struct {
	client client.Client
}

func NewK8sBackend() (*K8sBackend, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	cl, err := client.New(cfg, client.Options{})
	if err != nil {
		return nil, err
	}
	return &K8sBackend{client: cl}, nil
}

var gvk = schema.GroupVersionKind{
	Group:   "mecris.io",
	Version: "v1alpha1",
	Kind:    "UselessMachine",
}

func (b *K8sBackend) GetStatus(ctx context.Context) (*MachineStatus, error) {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	
	err := b.client.Get(ctx, client.ObjectKey{Name: "global", Namespace: "default"}, u)
	if err != nil {
		return nil, err
	}

	status, found, err := unstructured.NestedMap(u.Object, "status")
	if err != nil || !found {
		return nil, fmt.Errorf("status not found")
	}

	dailyCount, _, _ := unstructured.NestedInt64(status, "dailyCount")
	lastPushed, _, _ := unstructured.NestedString(status, "lastPushed")
	isNagging, _, _ := unstructured.NestedBool(status, "isNagging")

	return &MachineStatus{
		DailyCount: int32(dailyCount),
		LastPushed: lastPushed,
		IsNagging:  isNagging,
	}, nil
}

func (b *K8sBackend) PushButton(ctx context.Context) (string, error) {
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	u.SetName("global")
	u.SetNamespace("default")

	patch := client.RawPatch(client.Merge.Type(), []byte(`{"spec":{"action":"push"}}`))
	err := b.client.Patch(ctx, u, patch)
	if err != nil {
		return "", err
	}

	return "Push action recorded in Kubernetes.", nil
}
