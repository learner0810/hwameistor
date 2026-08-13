package hwameistor

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	apisv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	hwameistorapi "github.com/hwameistor/hwameistor/pkg/apiserver/api"
)

func newEmptyDiskController(t *testing.T) *LocalStorageNodeController {
	t.Helper()
	scheme := runtime.NewScheme()
	if err := apisv1alpha1.AddToScheme(scheme); err != nil {
		t.Fatalf("AddToScheme() error = %v", err)
	}
	cli := fake.NewClientBuilder().WithScheme(scheme).Build()
	return NewLocalStorageNodeController(cli, nil, record.NewFakeRecorder(1))
}

func TestReserveStorageNodeDiskReturnsErrorWhenDiskIsMissing(t *testing.T) {
	controller := newEmptyDiskController(t)
	_, err := controller.ReserveStorageNodeDisk(hwameistorapi.QueryPage{
		NodeName:        "node-a",
		DeviceShortPath: "sda",
	})
	if err == nil {
		t.Fatal("ReserveStorageNodeDisk() expected an error for a missing disk")
	}
}

func TestRemoveReserveStorageNodeDiskReturnsErrorWhenDiskIsMissing(t *testing.T) {
	controller := newEmptyDiskController(t)
	_, err := controller.RemoveReserveStorageNodeDisk(hwameistorapi.QueryPage{
		NodeName:        "node-a",
		DeviceShortPath: "sda",
	})
	if err == nil {
		t.Fatal("RemoveReserveStorageNodeDisk() expected an error for a missing disk")
	}
}
