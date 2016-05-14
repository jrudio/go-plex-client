package plex

// Plex ...
type Plex struct {
	URL   string
	token string
}

// SearchResults ...
type SearchResults struct {
	Children []struct {
		ElementType string `json:"_elementType"`
		Count       string `json:"count"`
		Key         string `json:"key"`
		Title       string `json:"title"`
	} `json:"_children"`
	ElementType                   string `json:"_elementType"`
	AllowCameraUpload             string `json:"allowCameraUpload"`
	AllowChannelAccess            string `json:"allowChannelAccess"`
	AllowSync                     string `json:"allowSync"`
	BackgroundProcessing          string `json:"backgroundProcessing"`
	Certificate                   string `json:"certificate"`
	CompanionProxy                string `json:"companionProxy"`
	FriendlyName                  string `json:"friendlyName"`
	HubSearch                     string `json:"hubSearch"`
	MachineIdentifier             string `json:"machineIdentifier"`
	Multiuser                     string `json:"multiuser"`
	MyPlex                        string `json:"myPlex"`
	MyPlexMappingState            string `json:"myPlexMappingState"`
	MyPlexSigninState             string `json:"myPlexSigninState"`
	MyPlexSubscription            string `json:"myPlexSubscription"`
	MyPlexUsername                string `json:"myPlexUsername"`
	Platform                      string `json:"platform"`
	PlatformVersion               string `json:"platformVersion"`
	RequestParametersInCookie     string `json:"requestParametersInCookie"`
	Sync                          string `json:"sync"`
	TranscoderActiveVideoSessions string `json:"transcoderActiveVideoSessions"`
	TranscoderAudio               string `json:"transcoderAudio"`
	TranscoderLyrics              string `json:"transcoderLyrics"`
	TranscoderPhoto               string `json:"transcoderPhoto"`
	TranscoderSubtitles           string `json:"transcoderSubtitles"`
	TranscoderVideo               string `json:"transcoderVideo"`
	TranscoderVideoBitrates       string `json:"transcoderVideoBitrates"`
	TranscoderVideoQualities      string `json:"transcoderVideoQualities"`
	TranscoderVideoResolutions    string `json:"transcoderVideoResolutions"`
	UpdatedAt                     string `json:"updatedAt"`
	Version                       string `json:"version"`
}

type plexResponse struct {
	Children []struct {
		ElementType string `json:"_elementType"`
		Count       string `json:"count"`
		Key         string `json:"key"`
		Title       string `json:"title"`
	} `json:"_children"`
	ElementType                   string `json:"_elementType"`
	AllowCameraUpload             string `json:"allowCameraUpload"`
	AllowChannelAccess            string `json:"allowChannelAccess"`
	AllowSync                     string `json:"allowSync"`
	BackgroundProcessing          string `json:"backgroundProcessing"`
	Certificate                   string `json:"certificate"`
	CompanionProxy                string `json:"companionProxy"`
	FriendlyName                  string `json:"friendlyName"`
	HubSearch                     string `json:"hubSearch"`
	MachineIdentifier             string `json:"machineIdentifier"`
	Multiuser                     string `json:"multiuser"`
	MyPlex                        string `json:"myPlex"`
	MyPlexMappingState            string `json:"myPlexMappingState"`
	MyPlexSigninState             string `json:"myPlexSigninState"`
	MyPlexSubscription            string `json:"myPlexSubscription"`
	MyPlexUsername                string `json:"myPlexUsername"`
	Platform                      string `json:"platform"`
	PlatformVersion               string `json:"platformVersion"`
	RequestParametersInCookie     string `json:"requestParametersInCookie"`
	Sync                          string `json:"sync"`
	TranscoderActiveVideoSessions string `json:"transcoderActiveVideoSessions"`
	TranscoderAudio               string `json:"transcoderAudio"`
	TranscoderLyrics              string `json:"transcoderLyrics"`
	TranscoderPhoto               string `json:"transcoderPhoto"`
	TranscoderSubtitles           string `json:"transcoderSubtitles"`
	TranscoderVideo               string `json:"transcoderVideo"`
	TranscoderVideoBitrates       string `json:"transcoderVideoBitrates"`
	TranscoderVideoQualities      string `json:"transcoderVideoQualities"`
	TranscoderVideoResolutions    string `json:"transcoderVideoResolutions"`
	UpdatedAt                     string `json:"updatedAt"`
	Version                       string `json:"version"`
}

type killTranscodeResponse struct {
	Children []struct {
		ElementType   string  `json:"_elementType"`
		AudioChannels int     `json:"audioChannels"`
		AudioCodec    string  `json:"audioCodec"`
		AudioDecision string  `json:"audioDecision"`
		Container     string  `json:"container"`
		Context       string  `json:"context"`
		Duration      int     `json:"duration"`
		Height        int     `json:"height"`
		Key           string  `json:"key"`
		Progress      float64 `json:"progress"`
		Protocol      string  `json:"protocol"`
		Remaining     int     `json:"remaining"`
		Speed         float64 `json:"speed"`
		Throttled     bool    `json:"throttled"`
		VideoCodec    string  `json:"videoCodec"`
		VideoDecision string  `json:"videoDecision"`
		Width         int     `json:"width"`
	} `json:"_children"`
	ElementType string `json:"_elementType"`
}

type transcodeSessionsResponse struct {
	Children []struct {
		ElementType   string  `json:"_elementType"`
		AudioChannels int     `json:"audioChannels"`
		AudioCodec    string  `json:"audioCodec"`
		AudioDecision string  `json:"audioDecision"`
		Container     string  `json:"container"`
		Context       string  `json:"context"`
		Duration      int     `json:"duration"`
		Height        int     `json:"height"`
		Key           string  `json:"key"`
		Progress      float64 `json:"progress"`
		Protocol      string  `json:"protocol"`
		Remaining     int     `json:"remaining"`
		Speed         float64 `json:"speed"`
		Throttled     bool    `json:"throttled"`
		VideoCodec    string  `json:"videoCodec"`
		VideoDecision string  `json:"videoDecision"`
		Width         int     `json:"width"`
	} `json:"_children"`
	ElementType string `json:"_elementType"`
}

type headers struct {
	Platform         string
	PlatformVersion  string
	Provides         string
	ClientIdentifier string
	Product          string
	Version          string
	Device           string
	ContainerSize    string
	ContainerStart   string
	Token            string
	Accept           string
}

type request struct {
	headers
}
