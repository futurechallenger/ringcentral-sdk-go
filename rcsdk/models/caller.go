package models

type CallerInfo struct {
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	ExtensionNumber string `json:"extensionNumber,omitempty"`
	Location        string `json:"location,omitempty"`
	Name            string `json:"name,omitempty"`
}
