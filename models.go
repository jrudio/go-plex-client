package plex

import (
	"encoding/xml"
	"net/http"
)

// Plex contains fields that are required to make
// an api call to your plex server
type Plex struct {
	URL              string
	Token            string
	ClientIdentifier string
	HTTPClient       http.Client
}

// SearchResults a list of media returned when searching
// for media via your plex server

type Provider struct {
	Key   string `json:"key"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

type SearchMediaContainer struct {
	MediaContainer
	Provider []Provider
}

type SearchResults struct {
	MediaContainer SearchMediaContainer `json:"MediaContainer"`
}

// Metadata ...
type Metadata struct {
	Player                Player       `json:"Player"`
	Session               Session      `json:"Session"`
	User                  User         `json:"User"`
	AddedAt               int          `json:"addedAt"`
	Art                   string       `json:"art"`
	ContentRating         string       `json:"contentRating"`
	Duration              int          `json:"duration"`
	GrandparentArt        string       `json:"grandparentArt"`
	GrandparentKey        string       `json:"grandparentKey"`
	GrandparentRatingKey  string       `json:"grandparentRatingKey"`
	GrandparentTheme      string       `json:"grandparentTheme"`
	GrandparentThumb      string       `json:"grandparentThumb"`
	GrandparentTitle      string       `json:"grandparentTitle"`
	GUID                  string       `json:"guid"`
	Index                 int64        `json:"index"`
	Key                   string       `json:"key"`
	LastViewedAt          int          `json:"lastViewedAt"`
	LibrarySectionID      int          `json:"librarySectionID"`
	LibrarySectionKey     string       `json:"librarySectionKey"`
	LibrarySectionTitle   string       `json:"librarySectionTitle"`
	OriginallyAvailableAt string       `json:"originallyAvailableAt"`
	ParentIndex           int64        `json:"parentIndex"`
	ParentKey             string       `json:"parentKey"`
	ParentRatingKey       string       `json:"parentRatingKey"`
	ParentThumb           string       `json:"parentThumb"`
	ParentTitle           string       `json:"parentTitle"`
	RatingCount           string       `json:"ratingCount"`
	Rating                float64      `json:"rating"`
	RatingKey             string       `json:"ratingKey"`
	SessionKey            string       `json:"sessionKey"`
	Summary               string       `json:"summary"`
	Thumb                 string       `json:"thumb"`
	Media                 []Media      `json:"Media"`
	Title                 string       `json:"title"`
	TitleSort             string       `json:"titleSort"`
	Type                  string       `json:"type"`
	UpdatedAt             int          `json:"updatedAt"`
	ViewCount             int          `json:"viewCount"`
	ViewOffset            int          `json:"viewOffset"`
	Year                  int          `json:"year"`
	Director              []TaggedData `json:"Director"`
	Writer                []TaggedData `json:"Writer"`
}

type MetadataV1 struct {
	Metadata
	Index            int64     `json:"index,string"`
	ParentIndex      int64     `json:"parentIndex,string"`
	AddedAt          string    `json:"addedAt"`
	Duration         string    `json:"duration"`
	LastViewedAt     string    `json:"lastViewedAt"`
	LibrarySectionID string    `json:"librarySectionID"`
	Media            []MediaV1 `json:"Media"`
	Rating           string    `json:"rating"`
	UpdatedAt        string    `json:"updatedAt"`
	ViewOffset       string    `json:"viewOffset"`
	Year             string    `json:"year"`
}

type Media struct {
	AspectRatio           float32 `json:"aspectRatio"`
	AudioChannels         int     `json:"audioChannels"`
	AudioCodec            string  `json:"audioCodec"`
	AudioProfile          string  `json:"audioProfile"`
	Bitrate               int     `json:"bitrate"`
	Container             string  `json:"container"`
	Duration              int     `json:"duration"`
	Has64bitOffsets       bool    `json:"has64bitOffsets"`
	Height                int     `json:"height"`
	ID                    int     `json:"id"`
	OptimizedForStreaming int     `json:"optimizedForStreaming"`
	Selected              bool    `json:"selected"`
	VideoCodec            string  `json:"videoCodec"`
	VideoFrameRate        string  `json:"videoFrameRate"`
	VideoProfile          string  `json:"videoProfile"`
	VideoResolution       string  `json:"videoResolution"`
	Width                 int     `json:"width"`
	Part                  []Part  `json:"Part"`
}

type MediaV1 struct {
	Media
	Part                  []PartV1 `json:"Part"`
	AudioChannels         float32  `json:"audioChannels,string"`
	AspectRatio           float32  `json:"aspectRatio,string"`
	Bitrate               int      `json:"bitrate,string"`
	Duration              int      `json:"duration,string"`
	Has64bitOffsets       string   `json:"has64bitOffsets"`
	Height                int      `json:"height,string"`
	ID                    int      `json:"id,string"`
	OptimizedForStreaming int      `json:"optimizedForStreaming,string"`
	Width                 int      `json:"width,string"`
}

type MediaContainer struct {
	Metadata            []Metadata `json:"Metadata"`
	AllowSync           bool       `json:"allowSync"`
	Identifier          string     `json:"identifier"`
	LibrarySectionID    int        `json:"librarySectionID"`
	LibrarySectionTitle string     `json:"librarySectionTitle"`
	LibrarySectionUUID  string     `json:"librarySectionUUID"`
	MediaTagPrefix      string     `json:"mediaTagPrefix"`
	MediaTagVersion     int        `json:"mediaTagVersion"`
	Size                int        `json:"size"`
}

// MediaMetadata ...
type MediaMetadata struct {
	MediaContainer MediaContainer `json:"MediaContainer"`
}

// Location is the path of a plex server directory
type Location struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
}

// Directory shows plex directory metadata
type Directory struct {
	Location   []Location `json:"Location"`
	Agent      string     `json:"agent"`
	AllowSync  bool       `json:"allowSync"`
	Art        string     `json:"art"`
	Composite  string     `json:"composite"`
	CreatedAt  int        `json:"createdAt"`
	Filter     bool       `json:"filters"`
	Key        string     `json:"key"`
	Language   string     `json:"language"`
	Refreshing bool       `json:"refreshing"`
	Scanner    string     `json:"scanner"`
	Thumb      string     `json:"thumb"`
	Title      string     `json:"title"`
	Type       string     `json:"type"`
	UpdatedAt  int        `json:"updatedAt"`
	UUID       string     `json:"uuid"`
}

// LibrarySections metadata of your library contents
type LibrarySections struct {
	MediaContainer struct {
		Directory []Directory `json:"Directory"`
	} `json:"MediaContainer"`
}

// TaggedData ...
type TaggedData struct {
	Tag    string `json:"tag"`
	Filter string `json:"filter"`
	id     string `json:"id"`
}

type Role struct {
	TaggedData
	Role  string `json:"role"`
	Thumb string `json:"thumb"`
}

// MetadataChildren returns metadata about a piece of media (tv show, movie, music, etc)
type MetadataChildren struct {
	MediaContainer MediaContainer `json:"MediaContainer"`
}

// SearchResultsEpisode contains metadata about an episode
type SearchResultsEpisode struct {
	MediaContainer MediaContainer `json:"MediaContainer"`
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

// CreateLibraryParams params required to create a library
type CreateLibraryParams struct {
	Name        string
	Location    string
	LibraryType string
	Agent       string
	Scanner     string
	Language    string
}

// DevicesResponse  metadata of a device that has connected to your server
type DevicesResponse struct {
	ID         int    `json:"id"`
	LastSeenAt string `json:"lastSeenAt"`
	Name       string `json:"name"`
	Product    string `json:"product"`
	Version    string `json:"version"`
}

// Friends are the plex accounts that have access to your server
type Friends struct {
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
	User              []Friends `xml:"User"`
}

type resultResponse struct {
	XMLName  xml.Name `xml:"Response"`
	Response struct {
		Code   int    `xml:"code,attr"`
		Status string `xml:"status,attr"`
	} `xml:"Response"`
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
		UserID            int    `xml:"userID,attr"`
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
	Device  []PMSDevices `xml:"Device"`
}

type terminateSessionResponse struct {
	XMLName xml.Name `xml:"MediaContainer"`
	Size    int      `xml:"size,attr"`
}

// PMSDevices is the result of the https://plex.tv/pms/resources endpoint
type PMSDevices struct {
	Name                 string       `json:"name" xml:"name,attr"`
	Product              string       `json:"product" xml:"product,attr"`
	ProductVersion       string       `json:"productVersion" xml:"productVersion,attr"`
	Platform             string       `json:"platform" xml:"platform,attr"`
	PlatformVersion      string       `json:"platformVersion" xml:"platformVersion,attr"`
	Device               string       `json:"device" xml:"device,attr"`
	ClientIdentifier     string       `json:"clientIdentifier" xml:"clientIdentifier,attr"`
	CreatedAt            string       `json:"createdAt" xml:"createdAt,attr"`
	LastSeenAt           string       `json:"lastSeenAt" xml:"lastSeenAt,attr"`
	Provides             string       `json:"provides" xml:"provides,attr"`
	Owned                string       `json:"owned" xml:"owned,attr"`
	AccessToken          string       `json:"accessToken" xml:"accessToken,attr"`
	HTTPSRequired        int          `json:"httpsRequired" xml:"httpsRequired,attr"`
	Synced               string       `json:"synced" xml:"synced,attr"`
	Relay                int          `json:"relay" xml:"relay,attr"`
	PublicAddressMatches string       `json:"publicAddressMatches" xml:"publicAddressMatches,attr"`
	PublicAddress        string       `json:"publicAddress" xml:"publicAddress,attr"`
	Presence             string       `json:"presence" xml:"presence,attr"`
	Connection           []Connection `json:"connection" xml:"Connection"`
}

// Connection lists options to connect to a device
type Connection struct {
	Protocol string `json:"protocol" xml:"protocol,attr"`
	Address  string `json:"address" xml:"address,attr"`
	Port     string `json:"port" xml:"port,attr"`
	URI      string `json:"uri" xml:"uri,attr"`
	Local    int    `json:"local" xml:"local,attr"`
}

// BaseAPIResponse info about the Plex Media Server
type BaseAPIResponse struct {
	MediaContainer struct {
		Directory []struct {
			Count int64  `json:"count"`
			Key   string `json:"key"`
			Title string `json:"title"`
		} `json:"Directory"`
		AllowCameraUpload             bool   `json:"allowCameraUpload"`
		AllowChannelAccess            bool   `json:"allowChannelAccess"`
		AllowSharing                  bool   `json:"allowSharing"`
		AllowSync                     bool   `json:"allowSync"`
		BackgroundProcessing          bool   `json:"backgroundProcessing"`
		Certificate                   bool   `json:"certificate"`
		CompanionProxy                bool   `json:"companionProxy"`
		CountryCode                   string `json:"countryCode"`
		Diagnostics                   string `json:"diagnostics"`
		EventStream                   bool   `json:"eventStream"`
		FriendlyName                  string `json:"friendlyName"`
		HubSearch                     bool   `json:"hubSearch"`
		ItemClusters                  bool   `json:"itemClusters"`
		Livetv                        int64  `json:"livetv"`
		MachineIdentifier             string `json:"machineIdentifier"`
		MediaProviders                bool   `json:"mediaProviders"`
		Multiuser                     bool   `json:"multiuser"`
		MyPlex                        bool   `json:"myPlex"`
		MyPlexMappingState            string `json:"myPlexMappingState"`
		MyPlexSigninState             string `json:"myPlexSigninState"`
		MyPlexSubscription            bool   `json:"myPlexSubscription"`
		MyPlexUsername                string `json:"myPlexUsername"`
		OwnerFeatures                 string `json:"ownerFeatures"`
		PhotoAutoTag                  bool   `json:"photoAutoTag"`
		Platform                      string `json:"platform"`
		PlatformVersion               string `json:"platformVersion"`
		PluginHost                    bool   `json:"pluginHost"`
		ReadOnlyLibraries             bool   `json:"readOnlyLibraries"`
		RequestParametersInCookie     bool   `json:"requestParametersInCookie"`
		Size                          int64  `json:"size"`
		StreamingBrainABRVersion      int64  `json:"streamingBrainABRVersion"`
		StreamingBrainVersion         int64  `json:"streamingBrainVersion"`
		Sync                          bool   `json:"sync"`
		TranscoderActiveVideoSessions int64  `json:"transcoderActiveVideoSessions"`
		TranscoderAudio               bool   `json:"transcoderAudio"`
		TranscoderLyrics              bool   `json:"transcoderLyrics"`
		TranscoderPhoto               bool   `json:"transcoderPhoto"`
		TranscoderSubtitles           bool   `json:"transcoderSubtitles"`
		TranscoderVideo               bool   `json:"transcoderVideo"`
		TranscoderVideoBitrates       string `json:"transcoderVideoBitrates"`
		TranscoderVideoQualities      string `json:"transcoderVideoQualities"`
		TranscoderVideoResolutions    string `json:"transcoderVideoResolutions"`
		UpdatedAt                     int64  `json:"updatedAt"`
		Updater                       bool   `json:"updater"`
		Version                       string `json:"version"`
		VoiceSearch                   bool   `json:"voiceSearch"`
	} `json:"MediaContainer"`
}

// UserPlexTV plex.tv user. should be used when interacting with plex.tv as the id is an int
type UserPlexTV struct {
	// ID is an int when signing in to Plex.tv but a string when access own server
	ID                  int    `json:"id"`
	UUID                string `json:"uuid"`
	Email               string `json:"email"`
	JoinedAt            string `json:"joined_at"`
	Username            string `json:"username"`
	Thumb               string `json:"thumb"`
	HasPassword         bool   `json:"hasPassword"`
	AuthToken           string `json:"authToken"`
	AuthenticationToken string `json:"authenticationToken"`
	Subscription        struct {
		Active   bool     `json:"active"`
		Status   string   `json:"Active"`
		Plan     string   `json:"lifetime"`
		Features []string `json:"features"`
	} `json:"subscription"`
	Roles struct {
		Roles []string `json:"roles"`
	} `json:"roles"`
	Entitlements []string `json:"entitlements"`
	ConfirmedAt  string   `json:"confirmedAt"`
	ForumID      string   `json:"forumId"`
	RememberMe   bool     `json:"rememberMe"`
	Title        string   `json:"title"`
}

// User plex server user. only difference is id is a string
type User struct {
	// ID is an int when signing in to Plex.tv but a string when access own server
	ID                  string `json:"id"`
	UUID                string `json:"uuid"`
	Email               string `json:"email"`
	JoinedAt            string `json:"joined_at"`
	Username            string `json:"username"`
	Thumb               string `json:"thumb"`
	HasPassword         bool   `json:"hasPassword"`
	AuthToken           string `json:"authToken"`
	AuthenticationToken string `json:"authenticationToken"`
	Subscription        struct {
		Active   bool     `json:"active"`
		Status   string   `json:"Active"`
		Plan     string   `json:"lifetime"`
		Features []string `json:"features"`
	} `json:"subscription"`
	Roles struct {
		Roles []string `json:"roles"`
	} `json:"roles"`
	Entitlements []string `json:"entitlements"`
	ConfirmedAt  string   `json:"confirmedAt"`
	ForumID      string   `json:"forumId"`
	RememberMe   bool     `json:"rememberMe"`
	Title        string   `json:"title"`
}

// SignInResponse ...
type SignInResponse struct {
	User UserPlexTV `json:"user"`
}

// ServerInfo is the result of the https://plex.tv/api/servers endpoint
type ServerInfo struct {
	XMLName           xml.Name `xml:"MediaContainer"`
	FriendlyName      string   `xml:"friendlyName,attr"`
	Identifier        string   `xml:"identifier,attr"`
	MachineIdentifier string   `xml:"machineIdentifier,attr"`
	Size              int      `xml:"size,attr"`
	Server            []struct {
		AccessToken       string `xml:"accessToken,attr"`
		Name              string `xml:"name,attr"`
		Address           string `xml:"address,attr"`
		Port              string `xml:"port,attr"`
		Version           string `xml:"version,attr"`
		Scheme            string `xml:"scheme,attr"`
		Host              string `xml:"host,attr"`
		LocalAddresses    string `xml:"localAddresses,attr"`
		MachineIdentifier string `xml:"machineIdentifier,attr"`
		CreatedAt         string `xml:"createdAt,attr"`
		UpdatedAt         string `xml:"updatedAt,attr"`
		Owned             string `xml:"owned,attr"`
		Synced            string `xml:"synced,attr"`
	} `xml:"Server"`
}

// SectionIDResponse the section id (or library id) of your server
// useful when inviting a user to the server
type SectionIDResponse struct {
	XMLName           xml.Name `xml:"MediaContainer"`
	FriendlyName      string   `xml:"friendlyName,attr"`
	Identifier        string   `xml:"identifier,attr"`
	MachineIdentifier string   `xml:"machineIdentifier,attr"`
	Size              int      `xml:"size,attr"`
	Server            []struct {
		Name              string           `xml:"name,attr"`
		Address           string           `xml:"address,attr"`
		Port              string           `xml:"port,attr"`
		Version           string           `xml:"version,attr"`
		Scheme            string           `xml:"scheme,attr"`
		Host              string           `xml:"host,attr"`
		LocalAddresses    string           `xml:"localAddresses,attr"`
		MachineIdentifier string           `xml:"machineIdentifier,attr"`
		CreatedAt         int              `xml:"createdAt,attr"`
		UpdatedAt         int              `xml:"updatedAt,attr"`
		Owned             int              `xml:"owned,attr"`
		Synced            string           `xml:"synced,attr"`
		Section           []ServerSections `xml:"Section"`
	} `xml:"Server"`
}

// ServerSections contains information of your library sections
type ServerSections struct {
	ID    int    `xml:"id,attr"`
	Key   string `xml:"key,attr"`
	Type  string `xml:"type,attr"`
	Title string `xml:"title,attr"`
}

// LibraryLabels are the existing labels set on your server
type LibraryLabels struct {
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
	Platform               string
	PlatformVersion        string
	Provides               string
	Product                string
	Version                string
	Device                 string
	ContainerSize          string
	ContainerStart         string
	Token                  string
	Accept                 string
	ContentType            string
	ClientIdentifier       string
	TargetClientIdentifier string
}

type request struct {
	headers
}

// Sessions

// TranscodeSessionsResponse is the result for transcode session endpoint /transcode/sessions
type TranscodeSessionsResponse struct {
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

// Stream ...
type Stream struct {
	AlbumGain          string  `json:"albumGain"`
	AlbumPeak          string  `json:"albumPeak"`
	AlbumRange         string  `json:"albumRange"`
	Anamorphic         bool    `json:"anamorphic"`
	AudioChannelLayout string  `json:"audioChannelLayout"`
	BitDepth           int     `json:"bitDepth"`
	Bitrate            int     `json:"bitrate"`
	BitrateMode        string  `json:"bitrateMode"`
	Cabac              string  `json:"cabac"`
	Channels           int     `json:"channels"`
	ChromaLocation     string  `json:"chromaLocation"`
	ChromaSubsampling  string  `json:"chromaSubsampling"`
	Codec              string  `json:"codec"`
	CodecID            string  `json:"codecID"`
	ColorRange         string  `json:"colorRange"`
	ColorSpace         string  `json:"colorSpace"`
	Default            bool    `json:"default"`
	DisplayTitle       string  `json:"displayTitle"`
	Duration           string  `json:"duration"`
	FrameRate          float64 `json:"frameRate"`
	FrameRateMode      string  `json:"frameRateMode"`
	Gain               string  `json:"gain"`
	HasScalingMatrix   bool    `json:"hasScalingMatrix"`
	Height             int     `json:"height"`
	ID                 int     `json:"id"`
	Index              int     `json:"index"`
	Language           string  `json:"language"`
	LanguageCode       string  `json:"languageCode"`
	Level              int     `json:"level"`
	Location           string  `json:"location"`
	Loudness           string  `json:"loudness"`
	Lra                string  `json:"lra"`
	Peak               string  `json:"peak"`
	PixelAspectRatio   string  `json:"pixelAspectRatio"`
	PixelFormat        string  `json:"pixelFormat"`
	Profile            string  `json:"profile"`
	RefFrames          int     `json:"refFrames"`
	SamplingRate       int     `json:"samplingRate"`
	ScanType           string  `json:"scanType"`
	Selected           bool    `json:"selected"`
	StreamIdentifier   string  `json:"streamIdentifier"`
	StreamType         int     `json:"streamType"`
	Width              int     `json:"width"`
}

type StreamV1 struct {
	Stream
	BitDepth         int     `json:"bitDepth,string"`
	Default          string  `json:"default"`
	Bitrate          int     `json:"bitrate,string"`
	FrameRate        float64 `json:"frameRate,string"`
	HasScalingMatrix string  `json:"hasScalingMatrix"`
	Height           int     `json:"height,string"`
	Width            int     `json:"width,string"`
	ID               int     `json:"id,string"`
	Index            int     `json:"index,string"`
	Level            int     `json:"level,string"`
	RefFrames        int     `json:"refFrames,string"`
	StreamType       int     `json:"streamType,string"`
	Channels         int     `json:"channels,string"`
	SamplingRate     int     `json:"samplingRate,string"`
	Selected         string  `json:"selected"`
}

// Part ...
type Part struct {
	AudioProfile          string   `json:"audioProfile"`
	Container             string   `json:"container"`
	Decision              string   `json:"decision"`
	Duration              int      `json:"duration"`
	File                  string   `json:"file"`
	Has64bitOffsets       bool     `json:"has64bitOffsets"`
	HasThumbnail          string   `json:"hasThumbnail"`
	ID                    int      `json:"id"`
	Key                   string   `json:"key"`
	OptimizedForStreaming bool     `json:"optimizedForStreaming"`
	Selected              bool     `json:"selected"`
	Size                  int      `json:"size"`
	Stream                []Stream `json:"Stream"`
	VideoProfile          string   `json:"videoProfile"`
}

type PartV1 struct {
	Part
	Duration              int        `json:"duration,string"`
	Has64bitOffsets       string     `json:"has64bitOffsets"`
	ID                    int        `json:"id,string"`
	OptimizedForStreaming string     `json:"optimizedForStreaming"`
	Size                  int        `json:"size,string"`
	Stream                []StreamV1 `json:"Stream"`
}

// Player ...
type Player struct {
	Address             string `json:"address"`
	Device              string `json:"device"`
	Local               bool   `json:"local"`
	MachineIdentifier   string `json:"machineIdentifier"`
	Model               string `json:"model"`
	Platform            string `json:"platform"`
	PlatformVersion     string `json:"platformVersion"`
	Product             string `json:"product"`
	Profile             string `json:"profile"`
	RemotePublicAddress string `json:"remotePublicAddress"`
	State               string `json:"state"`
	Title               string `json:"title"`
	UserID              int    `json:"userID"`
	Vendor              string `json:"vendor"`
	Version             string `json:"version"`
}

// Session ...
type Session struct {
	Bandwidth int    `json:"bandwidth"`
	ID        string `json:"id"`
	Location  string `json:"location"`
}

// CurrentSessions metadata of users consuming media
type CurrentSessions struct {
	MediaContainer struct {
		Metadata []MetadataV1 `json:"Metadata"`
		Size     int          `json:"size"`
	} `json:"MediaContainer"`
}
