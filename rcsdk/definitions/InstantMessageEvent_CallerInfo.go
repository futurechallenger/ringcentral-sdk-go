package definitions

type InstantMessageEvent_CallerInfo struct {
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	ExtensionNumber string `json:"extensionNumber,omitempty"`
	Location        string `json:"location,omitempty"`
	Name            string `json:"name,omitempty"`
}
