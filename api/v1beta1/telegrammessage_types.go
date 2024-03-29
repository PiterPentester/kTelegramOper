/*
Copyright 2020 Piter_Pentester.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TelegramMessageSpec defines the desired state of TelegramMessage
type TelegramMessageSpec struct {
	MessageToDeliver string `json:"messagetodeliver,omitempty"`
	Token            string `json:"token,omitempty"`
	ChatID           string `json:"chatid,omitempty"`
}

// TelegramMessageStatus defines the observed state of TelegramMessage
type TelegramMessageStatus struct {
	Delivered        string `json:"delivered,omitempty"`
	MessageDelivered string `json:"messagedelivered,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TelegramMessage is the Schema for the telegrammessages API
// +k8s:openapi-gen=true
type TelegramMessage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TelegramMessageSpec   `json:"spec,omitempty"`
	Status TelegramMessageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TelegramMessageList contains a list of TelegramMessage
type TelegramMessageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TelegramMessage `json:"items"`
}

func init() {

	SchemeBuilder.Register(&TelegramMessage{}, &TelegramMessageList{})
}
