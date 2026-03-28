package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// UselessMachineSpec defines the desired state of UselessMachine
type UselessMachineSpec struct {
	// Action is the requested operation (e.g., "push")
	// +optional
	Action string `json:"action,omitempty"`
}

// UselessMachineStatus defines the observed state of UselessMachine
type UselessMachineStatus struct {
	// DailyCount is the total number of button pushes today
	DailyCount int32 `json:"dailyCount"`
	// LastPushed is the ISO8601 timestamp of the last push
	LastPushed string `json:"lastPushed"`
	// IsNagging indicates if the user needs to be nagged
	IsNagging bool `json:"isNagging"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// UselessMachine is the Schema for the uselessmachines API
type UselessMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UselessMachineSpec   `json:"spec,omitempty"`
	Status UselessMachineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UselessMachineList contains a list of UselessMachine
type UselessMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UselessMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UselessMachine{}, &UselessMachineList{})
}
