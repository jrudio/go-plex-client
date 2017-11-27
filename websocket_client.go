package plex

// TimelineEntry ...
type TimelineEntry struct {
	Identifier    string `json:"identifier"`
	ItemID        int64  `json:"itemID"`
	MetadataState string `json:"metadataState"`
	SectionID     int64  `json:"sectionID"`
	State         int64  `json:"state"`
	Title         string `json:"title"`
	Type          int64  `json:"type"`
	UpdatedAt     int64  `json:"updatedAt"`
}

// ActivityNotification ...
type ActivityNotification struct {
	Activity struct {
		Cancellable bool   `json:"cancellable"`
		Progress    int64  `json:"progress"`
		Subtitle    string `json:"subtitle"`
		Title       string `json:"title"`
		Type        string `json:"type"`
		UserID      int64  `json:"userID"`
		UUID        string `json:"uuid"`
	} `json:"Activity"`
	Event string `json:"event"`
	UUID  string `json:"uuid"`
}

// StatusNotification ...
type StatusNotification struct {
	Description      string `json:"description"`
	NotificationName string `json:"notificationName"`
	Title            string `json:"title"`
}

// PlaySessionStateNotification ...
type PlaySessionStateNotification []struct {
	GUID             string `json:"guid"`
	Key              string `json:"key"`
	PlayQueueItemID  int64  `json:"playQueueItemID"`
	RatingKey        string `json:"ratingKey"`
	SessionKey       string `json:"sessionKey"`
	State            string `json:"state"`
	URL              string `json:"url"`
	ViewOffset       int64  `json:"viewOffset"`
	TranscodeSession `json:"transcodeSession"`
}

// ReachabilityNotification ...
type ReachabilityNotification struct {
	Reachability bool `json:"reachability"`
}

// BackgroundProcessingQueueEventNotification ...
type BackgroundProcessingQueueEventNotification struct {
	Event   string `json:"event"`
	QueueID int64  `json:"queueID"`
}

// TranscodeSession ...
type TranscodeSession struct {
	AudioChannels        int64   `json:"audioChannels"`
	AudioCodec           string  `json:"audioCodec"`
	AudioDecision        string  `json:"audioDecision"`
	Complete             bool    `json:"complete"`
	Container            string  `json:"container"`
	Context              string  `json:"context"`
	Duration             int64   `json:"duration"`
	Key                  string  `json:"key"`
	Progress             int64   `json:"progress"`
	Protocol             string  `json:"protocol"`
	Remaining            int64   `json:"remaining"`
	SourceAudioCodec     string  `json:"sourceAudioCodec"`
	SourceVideoCodec     string  `json:"sourceVideoCodec"`
	Speed                float64 `json:"speed"`
	Throttled            bool    `json:"throttled"`
	TranscodeHwRequested bool    `json:"transcodeHwRequested"`
	VideoCodec           string  `json:"videoCodec"`
	VideoDecision        string  `json:"videoDecision"`
}

// Setting ...
type Setting struct {
	Advanced bool   `json:"advanced"`
	Default  string `json:"default"`
	Group    string `json:"group"`
	Hidden   bool   `json:"hidden"`
	ID       string `json:"id"`
	Label    string `json:"label"`
	Summary  string `json:"summary"`
	Type     string `json:"type"`
	Value    int64  `json:"value"`
}

// NotificationContainer read pms notifications
type NotificationContainer struct {
	TimelineEntry []TimelineEntry `json:"TimelineEntry"`

	ActivityNotification []ActivityNotification `json:"ActivityNotification"`

	StatusNotification []StatusNotification `json:"StatusNotification"`

	PlaySessionStateNotification []PlaySessionStateNotification `json:"PlaySessionStateNotification"`

	ReachabilityNotification []ReachabilityNotification `json:"ReachabilityNotification"`

	BackgroundProcessingQueueEventNotification []BackgroundProcessingQueueEventNotification `json:"BackgroundProcessingQueueEventNotification"`

	TranscodeSession []TranscodeSession `json:"TranscodeSession"`

	Setting []Setting `json:"Setting"`

	Size int64 `json:"size"`
	// Type can be one of: playing, reachability, transcode.end, preference
	Type string `json:"type"`
}

// WebsocketNotification websocket payload of notifications from a plex media server
type WebsocketNotification struct {
	NotificationContainer `json:"NotificationContainer"`
}

// func (p *Plex) SubscribeToNofications()
