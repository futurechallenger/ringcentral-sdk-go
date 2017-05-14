package definitions

type FullCountryInfo struct {
	Uri              string `json:"uri,omitempty"`
	CallingCode      string `json:"callingCode,omitempty"`
	EmergencyCalling bool   `json:"emergencyCalling,omitempty"`
	IsoCode          string `json:"isoCode,omitempty"`
	Name             string `json:"name,omitempty"`
	NumberSelling    bool   `json:"numberSelling,omitempty"`
	LoginAllowed     bool   `json:"loginAllowed,omitempty"`
	Id               string `json:"id,omitempty"`
}
