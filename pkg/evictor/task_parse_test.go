package evictor

import "testing"

func TestParseEvictVolumeTask(t *testing.T) {
	tests := []struct {
		name        string
		task        string
		wantVolume  string
		wantNode    string
		expectError bool
	}{
		{name: "valid", task: "volume-a/node-a", wantVolume: "volume-a", wantNode: "node-a"},
		{name: "missing slash", task: "volume-a", expectError: true},
		{name: "missing volume", task: "/node-a", expectError: true},
		{name: "missing node", task: "volume-a/", expectError: true},
		{name: "too many parts", task: "volume-a/node-a/extra", expectError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVolume, gotNode, err := parseEvictVolumeTask(tt.task)
			if (err != nil) != tt.expectError {
				t.Fatalf("parseEvictVolumeTask() error = %v, expectError %v", err, tt.expectError)
			}
			if gotVolume != tt.wantVolume {
				t.Fatalf("parseEvictVolumeTask() volume = %q, want %q", gotVolume, tt.wantVolume)
			}
			if gotNode != tt.wantNode {
				t.Fatalf("parseEvictVolumeTask() node = %q, want %q", gotNode, tt.wantNode)
			}
		})
	}
}

func TestParseEvictPodTask(t *testing.T) {
	tests := []struct {
		name          string
		task          string
		wantNamespace string
		wantPod       string
		expectError   bool
	}{
		{name: "valid", task: "default/pod-a", wantNamespace: "default", wantPod: "pod-a"},
		{name: "missing slash", task: "pod-a", expectError: true},
		{name: "missing namespace", task: "/pod-a", expectError: true},
		{name: "missing pod", task: "default/", expectError: true},
		{name: "too many parts", task: "default/pod-a/extra", expectError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNamespace, gotPod, err := parseEvictPodTask(tt.task)
			if (err != nil) != tt.expectError {
				t.Fatalf("parseEvictPodTask() error = %v, expectError %v", err, tt.expectError)
			}
			if gotNamespace != tt.wantNamespace {
				t.Fatalf("parseEvictPodTask() namespace = %q, want %q", gotNamespace, tt.wantNamespace)
			}
			if gotPod != tt.wantPod {
				t.Fatalf("parseEvictPodTask() pod = %q, want %q", gotPod, tt.wantPod)
			}
		})
	}
}
