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
// +kubebuilder:printcolumn:name="Count",type="integer",JSONPath=".status.dailyCount",description="Total number of button pushes today"
// +kubebuilder:printcolumn:name="Last Pushed",type="string",JSONPath=".status.lastPushed",description="ISO8601 timestamp of the last push"
// +kubebuilder:printcolumn:name="Nagging",type="boolean",JSONPath=".status.isNagging",description="Indicates if the user needs to be nagged"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

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
