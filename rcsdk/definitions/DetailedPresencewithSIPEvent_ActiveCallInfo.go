package definitions

type DetailedPresencewithSIPEvent_ActiveCallInfo struct {
	From            string  `json:"from,omitempty"`
	To              string  `json:"to,omitempty"`
	TelephonyStatus string  `json:"telephonyStatus,omitempty"`
	SessionId       string  `json:"sessionId,omitempty"`
	SipData         SIPData `json:"sipData,omitempty"`
	Id              string  `json:"id,omitempty"`
	Direction       string  `json:"direction,omitempty"`
}
