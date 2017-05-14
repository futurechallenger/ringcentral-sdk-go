package requests

type AccountExtensionMeetingPostRequestBody struct {
	AudioOptions	[]string	`json:"audioOptions,omitempty"`
	Topic	string	`json:"topic,omitempty"`
	MeetingType	string	`json:"meetingType,omitempty"`
	Password	string	`json:"password,omitempty"`
	Schedule	MeetingScheduleInfo	`json:"schedule,omitempty"`
	AllowJoinBeforeHost	bool	`json:"allowJoinBeforeHost,omitempty"`
	StartHostVideo	bool	`json:"startHostVideo,omitempty"`
	StartParticipantsVideo	bool	`json:"startParticipantsVideo,omitempty"`
}