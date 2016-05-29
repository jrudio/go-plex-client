package plex

import "encoding/xml"

// Plex contains fields that are required to make
// an api call to your plex server
type Plex struct {
	URL   string
	token string
}

// SearchResults a list of media returned when searching
// for media via your plex server
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

type devicesResponse struct {
	ID         int    `json:"id"`
	LastSeenAt string `json:"lastSeenAt"`
	Name       string `json:"name"`
	Product    string `json:"product"`
	Version    string `json:"version"`
}

type friends struct {
	ID                        int    `xml:"id,attr"`
	Title                     string `xml:"title,attr"`
	Thumb                     string `xml:"thumb,attr"`
	Protected                 string `xml:"protected,attr"`
	Home                      string `xml:"home,attr"`
	AllowSync                 string `xml:"allowSync,attr"`
	AllowCameraUpload         string `xml:"allowCameraUpload,attr"`
	AllowChannels             string `xml:"allowChannels,attr"`
	FilterAll                 string `xml:"filterAll,attr"`
	FilterMovies              string `xml:"filterMovies,attr"`
	FilterMusic               string `xml:"filterMusic,attr"`
	FilterPhotos              string `xml:"filterPhotos,attr"`
	FilterTelevision          string `xml:"filterTelevision,attr"`
	Restricted                string `xml:"restricted,attr"`
	Username                  string `xml:"username,attr"`
	Email                     string `xml:"email,attr"`
	RecommendationsPlaylistID string `xml:"recommendationsPlaylistId,attr"`
	Server                    struct {
		ID                string `xml:"id,attr"`
		ServerID          string `xml:"serverId,attr"`
		MachineIdentifier string `xml:"machineIdentifier,attr"`
		Name              string `xml:"name,attr"`
		LastSeenAt        string `xml:"lastSeenAt,attr"`
		NumLibraries      string `xml:"numLibraries,attr"`
		AllLibraries      string `xml:"allLibraries,attr"`
		Owned             string `xml:"owned,attr"`
		Pending           string `xml:"pending,attr"`
	} `xml:"Server"`
}

type friendsResponse struct {
	XMLName           xml.Name  `xml:"MediaContainer"`
	FriendlyName      string    `xml:"friendlyName,attr"`
	Identifier        string    `xml:"identifier,attr"`
	MachineIdentifier string    `xml:"machineIdentifier,attr"`
	TotalSize         string    `xml:"totalSize,attr"`
	Size              int       `xml:"size,attr"`
	User              []friends `xml:"User"`
}

type resultResponse struct {
	XMLName xml.Name `xml:"Response"`
	Code    int      `xml:"code,attr"`
	Status  string   `xml:"status,attr"`
}

type inviteFriendResponse struct {
	XMLName           xml.Name `xml:"MediaContainer"`
	FriendlyName      string   `xml:"friendlyName,attr"`
	Identifier        string   `xml:"identifier,attr"`
	MachineIdentifier string   `xml:"machineIdentifier,attr"`
	Size              string   `xml:"size,attr"`
	SharedServer      struct {
		ID                string `xml:"id,attr"`
		Username          string `xml:"username,attr"`
		Email             string `xml:"email,attr"`
		UserID            string `xml:"userID,attr"`
		AccessToken       string `xml:"accessToken,attr"`
		Name              string `xml:"name,attr"`
		AcceptedAt        string `xml:"acceptedAt,attr"`
		InvitedAt         string `xml:"invitedAt,attr"`
		AllowSync         string `xml:"allowSync,attr"`
		AllowCameraUpload string `xml:"allowCameraUpload,attr"`
		AllowChannels     string `xml:"allowChannels,attr"`
		Owned             string `xml:"owned,attr"`
		Section           struct {
			ID     string `xml:"id,attr"`
			Key    string `xml:"key,attr"`
			Title  string `xml:"title,attr"`
			Type   string `xml:"type,attr"`
			Shared string `xml:"shared,attr"`
		} `xml:"Section"`
	} `xml:"SharedServer"`
}

// InviteFriendParams are the params to invite a friend
type InviteFriendParams struct {
	UsernameOrEmail string
	MachineID       string
	Label           string
	LibraryIDs      []int
}

// UpdateFriendParams optional parameters to update your friends access to your server
type UpdateFriendParams struct {
	AllowSync         string
	AllowCameraUpload string
	AllowChannels     string
	FilterMovies      string
	FilterTelevision  string
	FilterMusic       string
	FilterPhotos      string
}
type inviteFriendBody struct {
	ServerID        string                      `json:"server_id"`
	SharedServer    inviteFriendSharedServer    `json:"shared_server"`
	SharingSettings inviteFriendSharingSettings `json:"sharing_settings"`
}

type inviteFriendSharedServer struct {
	InvitedEmail      string `json:"invited_email"`
	LibrarySectionIDs []int  `json:"library_section_ids"`
}
type inviteFriendSharingSettings struct {
	FilterMovies     string `json:"filterMovies"`
	FilterTelevision string `json:"filterTelevision"`
}

type resourcesResponse struct {
	XMLName xml.Name     `xml:"MediaContainer"`
	Size    int          `xml:"size,attr"`
	Device  []pmsDevices `xml:"Device"`
}

type pmsDevices struct {
	Name                 string `xml:"name,attr"`
	Product              string `xml:"product,attr"`
	ProductVersion       string `xml:"productVersion,attr"`
	Platform             string `xml:"platform,attr"`
	PlatformVersion      string `xml:"platformVersion,attr"`
	Device               int    `xml:"device,attr"`
	ClientIdentifier     string `xml:"clientIdentifier,attr"`
	CreatedAt            string `xml:"createdAt,attr"`
	LastSeenAt           string `xml:"lastSeenAt,attr"`
	Provides             string `xml:"provides,attr"`
	Owned                string `xml:"owned,attr"`
	AccessToken          string `xml:"accessToken,attr"`
	HTTPSRequired        string `xml:"httpsRequired,attr"`
	Synced               string `xml:"synced,attr"`
	PublicAddressMatches string `xml:"publicAddressMatches,attr"`
	Presence             string `xml:"presence,attr"`
	Connection           []struct {
		Protocol string `xml:"protocol,attr"`
		Address  string `xml:"address,attr"`
		Port     string `xml:"port,attr"`
		URI      string `xml:"uri,attr"`
		Local    int    `xml:"local,attr"`
	} `xml:"Connection"`
}

// Library
type sectionIDResponse struct {
	XMLName           xml.Name `xml:"MediaContainer"`
	FriendlyName      string   `xml:"friendlyName,attr"`
	Identifier        string   `xml:"identifier,attr"`
	MachineIdentifier string   `xml:"machineIdentifier,attr"`
	Size              int      `xml:"size,attr"`
	Server            []struct {
		Name              string `xml:"name,attr"`
		Address           string `xml:"address,attr"`
		Port              string `xml:"port,attr"`
		Version           string `xml:"version,attr"`
		Scheme            string `xml:"scheme,attr"`
		Host              string `xml:"host,attr"`
		LocalAddresses    string `xml:"localAddresses,attr"`
		MachineIdentifier string `xml:"machineIdentifier,attr"`
		CreatedAt         int    `xml:"createdAt,attr"`
		UpdatedAt         int    `xml:"updatedAt,attr"`
		Owned             int    `xml:"owned,attr"`
		Synced            string `xml:"synced,attr"`
		Section           []struct {
			ID    int    `xml:"id,attr"`
			Key   string `xml:"key,attr"`
			Type  string `xml:"type,attr"`
			Title string `xml:"title,attr"`
		} `xml:"Section"`
	} `xml:"Server"`
}

type librarySections struct {
	ElementType     string `json:"_elementType"`
	AllowSync       string `json:"allowSync"`
	Identifier      string `json:"identifier"`
	MediaTagPrefix  string `json:"mediaTagPrefix"`
	MediaTagVersion string `json:"mediaTagVersion"`
	Title1          string `json:"title1"`
	Children        []struct {
		Children []struct {
			ElementType string `json:"_elementType"`
			ID          int    `json:"id"`
			Path        string `json:"path"`
		} `json:"_children"`
		ElementType string `json:"_elementType"`
		Agent       string `json:"agent"`
		AllowSync   string `json:"allowSync"`
		Art         string `json:"art"`
		Composite   string `json:"composite"`
		CreatedAt   string `json:"createdAt"`
		Filters     string `json:"filters"`
		Key         string `json:"key"`
		Language    string `json:"language"`
		Refreshing  string `json:"refreshing"`
		Scanner     string `json:"scanner"`
		Thumb       string `json:"thumb"`
		Title       string `json:"title"`
		Type        string `json:"type"`
		UpdatedAt   string `json:"updatedAt"`
		UUID        string `json:"uuid"`
	} `json:"_children"`
}

type libraryLabels struct {
	ElementType     string `json:"_elementType"`
	AllowSync       string `json:"allowSync"`
	Art             string `json:"art"`
	Content         string `json:"content"`
	Identifier      string `json:"identifier"`
	MediaTagPrefix  string `json:"mediaTagPrefix"`
	MediaTagVersion string `json:"mediaTagVersion"`
	Thumb           string `json:"thumb"`
	Title1          string `json:"title1"`
	Title2          string `json:"title2"`
	ViewGroup       string `json:"viewGroup"`
	ViewMode        string `json:"viewMode"`
	Children        []struct {
		ElementType string `json:"_elementType"`
		FastKey     string `json:"fastKey"`
		Key         string `json:"key"`
		Title       string `json:"title"`
	} `json:"_children"`
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

// Sessions

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

// currentSessions is xml because plex returns a dynamic type (string or number) for the duration field
type currentSessions struct {
	XMLName xml.Name `xml:"MediaContainer"`
	Size    string   `xml:"size,attr"`
	Video   []struct {
		AddedAt               string `xml:"addedAt,attr"`
		Art                   string `xml:"art,attr"`
		ChapterSource         string `xml:"chapterSource,attr"`
		ContentRating         string `xml:"contentRating,attr"`
		Duration              int    `xml:"duration,attr"`
		GUID                  string `xml:"guid,attr"`
		Key                   string `xml:"key,attr"`
		LibrarySectionID      string `xml:"librarySectionID,attr"`
		OriginallyAvailableAt string `xml:"originallyAvailableAt,attr"`
		PrimaryExtraKey       string `xml:"primaryExtraKey,attr"`
		Rating                string `xml:"rating,attr"`
		RatingKey             string `xml:"ratingKey,attr"`
		SessionKey            string `xml:"sessionKey,attr"`
		Studio                string `xml:"studio,attr"`
		Summary               string `xml:"summary,attr"`
		Tagline               string `xml:"tagline,attr"`
		Thumb                 string `xml:"thumb,attr"`
		Title                 string `xml:"title,attr"`
		TitleSort             string `xml:"titleSort,attr"`
		Type                  string `xml:"type,attr"`
		UpdatedAt             string `xml:"updatedAt,attr"`
		ViewOffset            int    `xml:"viewOffset,attr"`
		Year                  string `xml:"year,attr"`
		Media                 struct {
			AspectRatio           string `xml:"aspectRatio,attr"`
			AudioChannels         string `xml:"audioChannels,attr"`
			AudioCodec            string `xml:"audioCodec,attr"`
			AudioProfile          string `xml:"audioProfile,attr"`
			Bitrate               string `xml:"bitrate,attr"`
			Container             string `xml:"container,attr"`
			Duration              string `xml:"duration,attr"`
			Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
			Height                string `xml:"height,attr"`
			ID                    string `xml:"id,attr"`
			OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
			VideoCodec            string `xml:"videoCodec,attr"`
			VideoFrameRate        string `xml:"videoFrameRate,attr"`
			VideoProfile          string `xml:"videoProfile,attr"`
			VideoResolution       string `xml:"videoResolution,attr"`
			Width                 string `xml:"width,attr"`
			Part                  struct {
				AudioProfile          string `xml:"audioProfile,attr"`
				Container             string `xml:"container,attr"`
				Duration              string `xml:"duration,attr"`
				File                  string `xml:"file,attr"`
				Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
				ID                    string `xml:"id,attr"`
				Indexes               string `xml:"indexes,attr"`
				Key                   string `xml:"key,attr"`
				OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
				Size                  string `xml:"size,attr"`
				VideoProfile          string `xml:"videoProfile,attr"`
				Stream                []struct {
					BitDepth           string `xml:"bitDepth,attr"`
					Bitrate            string `xml:"bitrate,attr"`
					Cabac              string `xml:"cabac,attr"`
					ChromaSubsampling  string `xml:"chromaSubsampling,attr"`
					Codec              string `xml:"codec,attr"`
					CodecID            string `xml:"codecID,attr"`
					ColorRange         string `xml:"colorRange,attr"`
					ColorSpace         string `xml:"colorSpace,attr"`
					Default            string `xml:"default,attr"`
					Duration           string `xml:"duration,attr"`
					FrameRate          string `xml:"frameRate,attr"`
					FrameRateMode      string `xml:"frameRateMode,attr"`
					HasScalingMatrix   string `xml:"hasScalingMatrix,attr"`
					Height             string `xml:"height,attr"`
					ID                 string `xml:"id,attr"`
					Index              string `xml:"index,attr"`
					Level              string `xml:"level,attr"`
					PixelFormat        string `xml:"pixelFormat,attr"`
					Profile            string `xml:"profile,attr"`
					RefFrames          string `xml:"refFrames,attr"`
					ScanType           string `xml:"scanType,attr"`
					StreamIdentifier   string `xml:"streamIdentifier,attr"`
					StreamType         string `xml:"streamType,attr"`
					Width              string `xml:"width,attr"`
					AudioChannelLayout string `xml:"audioChannelLayout,attr"`
					BitrateMode        string `xml:"bitrateMode,attr"`
					Channels           string `xml:"channels,attr"`
					Language           string `xml:"language,attr"`
					LanguageCode       string `xml:"languageCode,attr"`
					SamplingRate       string `xml:"samplingRate,attr"`
					Selected           string `xml:"selected,attr"`
					Format             string `xml:"format,attr"`
					Key                string `xml:"key,attr"`
				} `xml:"Stream"`
			} `xml:"Part"`
		} `xml:"Media"`
		Genre []struct {
			Count string `xml:"count,attr"`
			ID    string `xml:"id,attr"`
			Tag   string `xml:"tag,attr"`
		} `xml:"Genre"`
		Writer []struct {
			ID    string `xml:"id,attr"`
			Tag   string `xml:"tag,attr"`
			Count string `xml:"count,attr"`
		} `xml:"Writer"`
		Director struct {
			Count string `xml:"count,attr"`
			ID    string `xml:"id,attr"`
			Tag   string `xml:"tag,attr"`
		} `xml:"Director"`
		Producer []struct {
			Count string `xml:"count,attr"`
			ID    string `xml:"id,attr"`
			Tag   string `xml:"tag,attr"`
		} `xml:"Producer"`
		Country struct {
			Count string `xml:"count,attr"`
			ID    string `xml:"id,attr"`
			Tag   string `xml:"tag,attr"`
		} `xml:"Country"`
		Role []struct {
			Count string `xml:"count,attr"`
			ID    string `xml:"id,attr"`
			Role  string `xml:"role,attr"`
			Tag   string `xml:"tag,attr"`
		} `xml:"Role"`
		Collection struct {
			Count string `xml:"count,attr"`
			ID    string `xml:"id,attr"`
			Tag   string `xml:"tag,attr"`
		} `xml:"Collection"`
		Label struct {
			ID  string `xml:"id,attr"`
			Tag string `xml:"tag,attr"`
		} `xml:"Label"`
		Field struct {
			Locked string `xml:"locked,attr"`
			Name   string `xml:"name,attr"`
		} `xml:"Field"`
		User struct {
			ID    string `xml:"id,attr"`
			Title string `xml:"title,attr"`
			Thumb string `xml:"thumb,attr"`
		} `xml:"User"`
		Player struct {
			Address           string `xml:"address,attr"`
			Device            string `xml:"device,attr"`
			MachineIdentifier string `xml:"machineIdentifier,attr"`
			Model             string `xml:"model,attr"`
			Platform          string `xml:"platform,attr"`
			PlatformVersion   string `xml:"platformVersion,attr"`
			Product           string `xml:"product,attr"`
			Profile           string `xml:"profile,attr"`
			State             string `xml:"state,attr"`
			Title             string `xml:"title,attr"`
			Vendor            string `xml:"vendor,attr"`
			Version           string `xml:"version,attr"`
		} `xml:"Player"`
		GrandparentArt       string `xml:"grandparentArt,attr"`
		GrandparentKey       string `xml:"grandparentKey,attr"`
		GrandparentRatingKey string `xml:"grandparentRatingKey,attr"`
		GrandparentTheme     string `xml:"grandparentTheme,attr"`
		GrandparentThumb     string `xml:"grandparentThumb,attr"`
		GrandparentTitle     string `xml:"grandparentTitle,attr"`
		Index                string `xml:"index,attr"`
		LastViewedAt         string `xml:"lastViewedAt,attr"`
		ParentIndex          string `xml:"parentIndex,attr"`
		ParentKey            string `xml:"parentKey,attr"`
		ParentRatingKey      string `xml:"parentRatingKey,attr"`
		ParentThumb          string `xml:"parentThumb,attr"`
		ViewCount            string `xml:"viewCount,attr"`
		TranscodeSession     struct {
			Key           string `xml:"key,attr"`
			Throttled     string `xml:"throttled,attr"`
			Progress      string `xml:"progress,attr"`
			Speed         string `xml:"speed,attr"`
			Duration      string `xml:"duration,attr"`
			Remaining     string `xml:"remaining,attr"`
			Context       string `xml:"context,attr"`
			VideoDecision string `xml:"videoDecision,attr"`
			AudioDecision string `xml:"audioDecision,attr"`
			Protocol      string `xml:"protocol,attr"`
			Container     string `xml:"container,attr"`
			VideoCodec    string `xml:"videoCodec,attr"`
			AudioCodec    string `xml:"audioCodec,attr"`
			AudioChannels string `xml:"audioChannels,attr"`
			Width         string `xml:"width,attr"`
			Height        string `xml:"height,attr"`
		} `xml:"TranscodeSession"`
	} `xml:"Video"`
}
