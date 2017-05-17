package plex

import (
	"encoding/xml"
	"net/http"
)

// Plex contains fields that are required to make
// an api call to your plex server
type Plex struct {
	URL        string
	Token      string
	HTTPClient http.Client
}

// SearchResults a list of media returned when searching
// for media via your plex server
type SearchResults struct {
	Children []struct {
		Children []struct {
			ElementType string `json:"_elementType"`
			Tag         string `json:"tag"`
		} `json:"_children"`
		ElementType           string `json:"_elementType"`
		AddedAt               int    `json:"addedAt"`
		AllowSync             string `json:"allowSync"`
		Art                   string `json:"art"`
		Banner                string `json:"banner"`
		ChildCount            int    `json:"childCount"`
		ContentRating         string `json:"contentRating"`
		Duration              int    `json:"duration"`
		Index                 int    `json:"index"`
		Key                   string `json:"key"`
		LastViewedAt          int    `json:"lastViewedAt"`
		LeafCount             int    `json:"leafCount"`
		LibrarySectionID      string `json:"librarySectionID"`
		LibrarySectionTitle   string `json:"librarySectionTitle"`
		LibrarySectionUUID    string `json:"librarySectionUUID"`
		OriginallyAvailableAt string `json:"originallyAvailableAt"`
		Personal              string `json:"personal"`
		Rating                string `json:"rating"`
		RatingKey             int    `json:"ratingKey"`
		SourceTitle           string `json:"sourceTitle"`
		Studio                string `json:"studio"`
		Summary               string `json:"summary"`
		Theme                 string `json:"theme"`
		Thumb                 string `json:"thumb"`
		Title                 string `json:"title"`
		Type                  string `json:"type"`
		UpdatedAt             int    `json:"updatedAt"`
		ViewCount             int    `json:"viewCount"`
		ViewedLeafCount       int    `json:"viewedLeafCount"`
		Year                  int    `json:"year"`
	} `json:"_children"`
	ElementType     string `json:"_elementType"`
	Identifier      string `json:"identifier"`
	MediaTagPrefix  string `json:"mediaTagPrefix"`
	MediaTagVersion string `json:"mediaTagVersion"`
}

// MediaMetadata using xml because plex kept spitting out different types when using json
type MediaMetadata struct {
	XMLName             xml.Name `json:"MediaContainer" xml:"MediaContainer"`
	Size                string   `json:"size" xml:"size,attr"`
	AllowSync           string   `json:"allowSync" xml:"allowSync,attr"`
	Identifier          string   `json:"identifier" xml:"identifier,attr"`
	LibrarySectionID    string   `json:"librarySectionID" xml:"librarySectionID,attr"`
	LibrarySectionTitle string   `json:"librarySectionTitle" xml:"librarySectionTitle,attr"`
	LibrarySectionUUID  string   `json:"librarySectionUUID" xml:"librarySectionUUID,attr"`
	MediaTagPrefix      string   `json:"mediaTagPrefix" xml:"mediaTagPrefix,attr"`
	MediaTagVersion     string   `json:"mediaTagVersion" xml:"mediaTagVersion,attr"`
	Video               struct {
		RatingKey            string `json:"ratingKey" xml:"ratingKey,attr"`
		Key                  string `json:"key" xml:"key,attr"`
		GrandparentTitle     string `json:"grandparentTitle" xml:"grandparentTitle,attr"`
		GrandparentRatingKey string `json:"grandparentRatingKey" xml:"grandparentRatingKey,attr"`
		ParentRatingKey      string `json:"parentRatingKey" xml:"parentRatingKey,attr"`
		ParentYear           string `json:"parentYear" xml:"parentYear,attr"`
		ParentTitle          string `json:"parentTitle" xml:"parentTitle,attr"`
		GUID                 string `json:"guid" xml:"guid,attr"`
		LibrarySectionID     string `json:"librarySectionID" xml:"librarySectionID,attr"`
		Type                 string `json:"type" xml:"type,attr"`
		Title                string `json:"title" xml:"title,attr"`
		Summary              string `json:"summary" xml:"summary,attr"`
		ViewCount            string `json:"viewCount" xml:"viewCount,attr"`
		LastViewedAt         string `json:"lastViewedAt" xml:"lastViewedAt,attr"`
		Year                 string `json:"year" xml:"year,attr"`
		Thumb                string `json:"thumb" xml:"thumb,attr"`
		Art                  string `json:"art" xml:"art,attr"`
		Duration             string `json:"duration" xml:"duration,attr"`
		AddedAt              string `json:"addedAt" xml:"addedAt,attr"`
		UpdatedAt            string `json:"updatedAt" xml:"updatedAt,attr"`
		ChapterSource        string `json:"chapterSource" xml:"chapterSource,attr"`
		Media                struct {
			VideoResolution string `json:"videoResolution" xml:"videoResolution,attr"`
			ID              string `json:"id" xml:"id,attr"`
			Duration        string `json:"duration" xml:"duration,attr"`
			Bitrate         string `json:"bitrate" xml:"bitrate,attr"`
			Width           string `json:"width" xml:"width,attr"`
			Height          string `json:"height" xml:"height,attr"`
			AspectRatio     string `json:"aspectRatio" xml:"aspectRatio,attr"`
			AudioChannels   string `json:"audioChannels" xml:"audioChannels,attr"`
			AudioCodec      string `json:"audioCodec" xml:"audioCodec,attr"`
			VideoCodec      string `json:"videoCodec" xml:"videoCodec,attr"`
			Container       string `json:"container" xml:"container,attr"`
			VideoFrameRate  string `json:"videoFrameRate" xml:"videoFrameRate,attr"`
			VideoProfile    string `json:"videoProfile" xml:"videoProfile,attr"`
			Part            struct {
				ID           string `json:"id" xml:"id,attr"`
				Key          string `json:"key" xml:"key,attr"`
				Duration     string `json:"duration" xml:"duration,attr"`
				File         string `json:"file" xml:"file,attr"`
				Size         string `json:"size" xml:"size,attr"`
				Container    string `json:"container" xml:"container,attr"`
				VideoProfile string `json:"videoProfile" xml:"videoProfile,attr"`
				Stream       []struct {
					ID                 string `json:"id" xml:"id,attr"`
					StreamType         string `json:"streamType" xml:"streamType,attr"`
					Default            string `json:"default" xml:"default,attr"`
					Codec              string `json:"codec" xml:"codec,attr"`
					Index              string `json:"index" xml:"index,attr"`
					Bitrate            string `json:"bitrate" xml:"bitrate,attr"`
					BitDepth           string `json:"bitDepth" xml:"bitDepth,attr"`
					Cabac              string `json:"cabac" xml:"cabac,attr"`
					ChromaSubsampling  string `json:"chromaSubsampling" xml:"chromaSubsampling,attr"`
					CodecID            string `json:"codecID" xml:"codecID,attr"`
					ColorRange         string `json:"colorRange" xml:"colorRange,attr"`
					ColorSpace         string `json:"colorSpace" xml:"colorSpace,attr"`
					Duration           string `json:"duration" xml:"duration,attr"`
					FrameRate          string `json:"frameRate" xml:"frameRate,attr"`
					FrameRateMode      string `json:"frameRateMode" xml:"frameRateMode,attr"`
					HasScalingMatrix   string `json:"hasScalingMatrix" xml:"hasScalingMatrix,attr"`
					HeaderStripping    string `json:"headerStripping" xml:"headerStripping,attr"`
					Height             string `json:"height" xml:"height,attr"`
					Level              string `json:"level" xml:"level,attr"`
					PixelFormat        string `json:"pixelFormat" xml:"pixelFormat,attr"`
					Profile            string `json:"profile" xml:"profile,attr"`
					RefFrames          string `json:"refFrames" xml:"refFrames,attr"`
					ScanType           string `json:"scanType" xml:"scanType,attr"`
					Width              string `json:"width" xml:"width,attr"`
					Selected           string `json:"selected" xml:"selected,attr"`
					Channels           string `json:"channels" xml:"channels,attr"`
					AudioChannelLayout string `json:"audioChannelLayout" xml:"audioChannelLayout,attr"`
					BitrateMode        string `json:"bitrateMode" xml:"bitrateMode,attr"`
					DialogNorm         string `json:"dialogNorm" xml:"dialogNorm,attr"`
					SamplingRate       string `json:"samplingRate" xml:"samplingRate,attr"`
				} `json:"stream" xml:"Stream"`
			} `json:"part" xml:"Part"`
		} `json:"media" xml:"Media"`
		Label struct {
			ID  string `json:"id" xml:"id,attr"`
			Tag string `json:"tag" xml:"tag,attr"`
		} `json:"label" xml:"Label"`
		Field []struct {
			Name   string `json:"name" xml:"name,attr"`
			Locked string `json:"locked" xml:"locked,attr"`
		} `json:"field" xml:"Field"`
	} `json:"video" xml:"Video"`
	Directory struct {
		RatingKey             string `json:"ratingKey" xml:"ratingKey,attr"`
		Key                   string `json:"key" xml:"key,attr"`
		GUID                  string `json:"guid" xml:"guid,attr"`
		LibrarySectionID      string `json:"librarySectionID" xml:"librarySectionID,attr"`
		ParentTitle           string `json:"parentTitle" xml:"parentTitle,attr"`
		ParentYear            string `json:"parentYear" xml:"parentYear,attr"`
		Studio                string `json:"studio" xml:"studio,attr"`
		Type                  string `json:"type" xml:"type,attr"`
		Title                 string `json:"title" xml:"title,attr"`
		TitleSort             string `json:"titleSort" xml:"titleSort,attr"`
		ContentRating         string `json:"contentRating" xml:"contentRating,attr"`
		Summary               string `json:"summary" xml:"summary,attr"`
		Index                 string `json:"index" xml:"index,attr"`
		Rating                string `json:"rating" xml:"rating,attr"`
		ViewCount             string `json:"viewCount" xml:"viewCount,attr"`
		LastViewedAt          string `json:"lastViewedAt" xml:"lastViewedAt,attr"`
		Year                  string `json:"year" xml:"year,attr"`
		Thumb                 string `json:"thumb" xml:"thumb,attr"`
		Art                   string `json:"art" xml:"art,attr"`
		Banner                string `json:"banner" xml:"banner,attr"`
		Theme                 string `json:"theme" xml:"theme,attr"`
		Duration              string `json:"duration" xml:"duration,attr"`
		OriginallyAvailableAt string `json:"originallyAvailableAt" xml:"originallyAvailableAt,attr"`
		LeafCount             string `json:"leafCount" xml:"leafCount,attr"`
		ViewedLeafCount       string `json:"viewedLeafCount" xml:"viewedLeafCount,attr"`
		ChildCount            string `json:"childCount" xml:"childCount,attr"`
		AddedAt               string `json:"addedAt" xml:"addedAt,attr"`
		UpdatedAt             string `json:"updatedAt" xml:"updatedAt,attr"`
		Genre                 []struct {
			ID  string `json:"id" xml:"id,attr"`
			Tag string `json:"tag" xml:"tag,attr"`
		} `json:"genre" xml:"Genre"`
		Role []struct {
			ID    string `json:"id" xml:"id,attr"`
			Tag   string `json:"tag" xml:"tag,attr"`
			Role  string `json:"role" xml:"role,attr"`
			Thumb string `json:"thumb" xml:"thumb,attr"`
		} `json:"role" xml:"Role"`
		Field struct {
			Name   string `json:"name" xml:"name,attr"`
			Locked string `json:"locked" xml:"locked,attr"`
		} `json:"field" xml:"Field"`
		Location string `json:"location" xml:"Location"`
	} `json:"directory" xml:"Directory"`
}

// MediaMetadataChildren returns metadata about a piece of media (tv show, movie, music, etc)
type MediaMetadataChildren struct {
	XMLName             xml.Name `json:"MediaContainer" xml:"MediaContainer"`
	Size                string   `json:"size" xml:"size,attr"`
	AllowSync           string   `json:"allowSync" xml:"allowSync,attr"`
	Identifier          string   `json:"identifier" xml:"identifier,attr"`
	LibrarySectionID    string   `json:"librarySectionID" xml:"librarySectionID,attr"`
	LibrarySectionTitle string   `json:"librarySectionTitle" xml:"librarySectionTitle,attr"`
	LibrarySectionUUID  string   `json:"librarySectionUUID" xml:"librarySectionUUID,attr"`
	MediaTagPrefix      string   `json:"mediaTagPrefix" xml:"mediaTagPrefix,attr"`
	MediaTagVersion     string   `json:"mediaTagVersion" xml:"mediaTagVersion,attr"`
	Key                 string   `json:"key" xml:"key,attr"`
	ParentYear          string   `json:"parentYear" xml:"parentYear,attr"`
	ParentTitle         string   `json:"parentTitle" xml:"parentTitle,attr"`
	ParentIndex         string   `json:"parentIndex" xml:"parentIndex,attr"`
	Video               []struct {
		RatingKey        string `json:"ratingKey" xml:"ratingKey,attr"`
		Key              string `json:"key" xml:"key,attr"`
		GUID             string `json:"guid" xml:"guid,attr"`
		LibrarySectionID string `json:"librarySectionID" xml:"librarySectionID,attr"`
		Type             string `json:"type" xml:"type,attr"`
		ParentRatingKey  string `json:"parentRatingKey" xml:"parentRatingKey,attr"`
		ParentYear       string `json:"parentYear" xml:"parentYear,attr"`
		ParentTitle      string `json:"parentTitle" xml:"parentTitle,attr"`
		Title            string `json:"title" xml:"title,attr"`
		Summary          string `json:"summary" xml:"summary,attr"`
		ViewCount        string `json:"viewCount" xml:"viewCount,attr"`
		LastViewedAt     string `json:"lastViewedAt" xml:"lastViewedAt,attr"`
		Year             string `json:"year" xml:"year,attr"`
		Thumb            string `json:"thumb" xml:"thumb,attr"`
		Art              string `json:"art" xml:"art,attr"`
		Duration         string `json:"duration" xml:"duration,attr"`
		AddedAt          string `json:"addedAt" xml:"addedAt,attr"`
		UpdatedAt        string `json:"updatedAt" xml:"updatedAt,attr"`
		ChapterSource    string `json:"chapterSource" xml:"chapterSource,attr"`
		Media            struct {
			VideoResolution string `json:"videoResolution" xml:"videoResolution,attr"`
			ID              string `json:"id" xml:"id,attr"`
			Duration        string `json:"duration" xml:"duration,attr"`
			Bitrate         string `json:"bitrate" xml:"bitrate,attr"`
			Width           string `json:"width" xml:"width,attr"`
			Height          string `json:"height" xml:"height,attr"`
			AspectRatio     string `json:"aspectRatio" xml:"aspectRatio,attr"`
			AudioChannels   string `json:"audioChannels" xml:"audioChannels,attr"`
			AudioCodec      string `json:"audioCodec" xml:"audioCodec,attr"`
			VideoCodec      string `json:"videoCodec" xml:"videoCodec,attr"`
			Container       string `json:"container" xml:"container,attr"`
			VideoFrameRate  string `json:"videoFrameRate" xml:"videoFrameRate,attr"`
			VideoProfile    string `json:"videoProfile" xml:"videoProfile,attr"`
			Part            struct {
				ID           string `json:"id" xml:"id,attr"`
				Key          string `json:"key" xml:"key,attr"`
				Duration     string `json:"duration" xml:"duration,attr"`
				File         string `json:"file" xml:"file,attr"`
				Size         string `json:"size" xml:"size,attr"`
				Container    string `json:"container" xml:"container,attr"`
				VideoProfile string `json:"videoProfile" xml:"videoProfile,attr"`
				Stream       []struct {
					ID                 string `json:"id" xml:"id,attr"`
					StreamType         string `json:"streamType" xml:"streamType,attr"`
					Default            string `json:"default" xml:"default,attr"`
					Codec              string `json:"codec" xml:"codec,attr"`
					Index              string `json:"index" xml:"index,attr"`
					Bitrate            string `json:"bitrate" xml:"bitrate,attr"`
					BitDepth           string `json:"bitDepth" xml:"bitDepth,attr"`
					Cabac              string `json:"cabac" xml:"cabac,attr"`
					ChromaSubsampling  string `json:"chromaSubsampling" xml:"chromaSubsampling,attr"`
					CodecID            string `json:"codecID" xml:"codecID,attr"`
					ColorRange         string `json:"colorRange" xml:"colorRange,attr"`
					ColorSpace         string `json:"colorSpace" xml:"colorSpace,attr"`
					Duration           string `json:"duration" xml:"duration,attr"`
					FrameRate          string `json:"frameRate" xml:"frameRate,attr"`
					FrameRateMode      string `json:"frameRateMode" xml:"frameRateMode,attr"`
					HasScalingMatrix   string `json:"hasScalingMatrix" xml:"hasScalingMatrix,attr"`
					HeaderStripping    string `json:"headerStripping" xml:"headerStripping,attr"`
					Height             string `json:"height" xml:"height,attr"`
					Level              string `json:"level" xml:"level,attr"`
					PixelFormat        string `json:"pixelFormat" xml:"pixelFormat,attr"`
					Profile            string `json:"profile" xml:"profile,attr"`
					RefFrames          string `json:"refFrames" xml:"refFrames,attr"`
					ScanType           string `json:"scanType" xml:"scanType,attr"`
					Width              string `json:"width" xml:"width,attr"`
					Selected           string `json:"selected" xml:"selected,attr"`
					Channels           string `json:"channels" xml:"channels,attr"`
					AudioChannelLayout string `json:"audioChannelLayout" xml:"audioChannelLayout,attr"`
					BitrateMode        string `json:"bitrateMode" xml:"bitrateMode,attr"`
					DialogNorm         string `json:"dialogNorm" xml:"dialogNorm,attr"`
					SamplingRate       string `json:"samplingRate" xml:"samplingRate,attr"`
				} `json:"Stream" xml:"Stream"`
			} `json:"Part" xml:"Part"`
		} `json:"Media" xml:"Media"`
		Field []struct {
			Name   string `json:"name" xml:"name,attr"`
			Locked string `json:"locked" xml:"locked,attr"`
		} `json:"Field" xml:"Field"`
	} `json:"Video" xml:"Video"`
	Directory []struct {
		RatingKey             string `json:"ratingKey" xml:"ratingKey,attr"`
		Key                   string `json:"key" xml:"key,attr"`
		GUID                  string `json:"guid" xml:"guid,attr"`
		LibrarySectionID      string `json:"librarySectionID" xml:"librarySectionID,attr"`
		ParentTitle           string `json:"parentTitle" xml:"parentTitle,attr"`
		ParentYear            string `json:"parentYear" xml:"parentYear,attr"`
		Studio                string `json:"studio" xml:"studio,attr"`
		Type                  string `json:"type" xml:"type,attr"`
		Title                 string `json:"title" xml:"title,attr"`
		TitleSort             string `json:"titleSort" xml:"titleSort,attr"`
		ContentRating         string `json:"contentRating" xml:"contentRating,attr"`
		Summary               string `json:"summary" xml:"summary,attr"`
		Index                 string `json:"index" xml:"index,attr"`
		Rating                string `json:"rating" xml:"rating,attr"`
		ViewCount             string `json:"viewCount" xml:"viewCount,attr"`
		LastViewedAt          string `json:"lastViewedAt" xml:"lastViewedAt,attr"`
		Year                  string `json:"year" xml:"year,attr"`
		Thumb                 string `json:"thumb" xml:"thumb,attr"`
		Art                   string `json:"art" xml:"art,attr"`
		Banner                string `json:"banner" xml:"banner,attr"`
		Theme                 string `json:"theme" xml:"theme,attr"`
		Duration              string `json:"duration" xml:"duration,attr"`
		OriginallyAvailableAt string `json:"originallyAvailableAt" xml:"originallyAvailableAt,attr"`
		LeafCount             string `json:"leafCount" xml:"leafCount,attr"`
		ViewedLeafCount       string `json:"viewedLeafCount" xml:"viewedLeafCount,attr"`
		ChildCount            string `json:"childCount" xml:"childCount,attr"`
		AddedAt               string `json:"addedAt" xml:"addedAt,attr"`
		UpdatedAt             string `json:"updatedAt" xml:"updatedAt,attr"`
		Genre                 []struct {
			ID  string `json:"id" xml:"id,attr"`
			Tag string `json:"tag" xml:"tag,attr"`
		} `json:"genre" xml:"Genre"`
		Role []struct {
			ID    string `json:"id" xml:"id,attr"`
			Tag   string `json:"tag" xml:"tag,attr"`
			Role  string `json:"role" xml:"role,attr"`
			Thumb string `json:"thumb" xml:"thumb,attr"`
		} `json:"role" xml:"Role"`
		Field struct {
			Name   string `json:"name" xml:"name,attr"`
			Locked string `json:"locked" xml:"locked,attr"`
		} `json:"field" xml:"Field"`
		Location string `json:"location" xml:"Location"`
	} `json:"directory" xml:"Directory"`
}

// SearchResultsEpisode contains metadata about an episode
type SearchResultsEpisode struct {
	Children []struct {
		Children []struct {
			Children []struct {
				ElementType           string `json:"_elementType"`
				Container             string `json:"container"`
				Duration              int    `json:"duration"`
				File                  string `json:"file"`
				Has64bitOffsets       bool   `json:"has64bitOffsets"`
				HasThumbnail          string `json:"hasThumbnail"`
				ID                    string `json:"id"`
				Key                   string `json:"key"`
				OptimizedForStreaming bool   `json:"optimizedForStreaming"`
				Size                  int    `json:"size"`
				VideoProfile          string `json:"videoProfile"`
			} `json:"_children"`
			ElementType           string `json:"_elementType"`
			AspectRatio           string `json:"aspectRatio"`
			AudioChannels         int    `json:"audioChannels"`
			AudioCodec            string `json:"audioCodec"`
			Bitrate               int    `json:"bitrate"`
			Container             string `json:"container"`
			Duration              int    `json:"duration"`
			Has64bitOffsets       bool   `json:"has64bitOffsets"`
			Height                int    `json:"height"`
			ID                    int    `json:"id"`
			OptimizedForStreaming int    `json:"optimizedForStreaming"`
			VideoCodec            string `json:"videoCodec"`
			VideoFrameRate        string `json:"videoFrameRate"`
			VideoProfile          string `json:"videoProfile"`
			VideoResolution       string `json:"videoResolution"`
			Width                 int    `json:"width"`
		} `json:"_children"`
		ElementType           string `json:"_elementType"`
		AddedAt               int    `json:"addedAt"`
		ChapterSource         string `json:"chapterSource"`
		Duration              int    `json:"duration"`
		Index                 int    `json:"index"`
		Key                   string `json:"key"`
		LastViewedAt          int    `json:"lastViewedAt"`
		OriginallyAvailableAt string `json:"originallyAvailableAt"`
		ParentKey             string `json:"parentKey"`
		ParentRatingKey       int    `json:"parentRatingKey"`
		Rating                string `json:"rating"`
		RatingKey             int    `json:"ratingKey"`
		Summary               string `json:"summary"`
		Thumb                 string `json:"thumb"`
		Title                 string `json:"title"`
		Type                  string `json:"type"`
		UpdatedAt             int    `json:"updatedAt"`
		ViewCount             int    `json:"viewCount"`
		Year                  int    `json:"year"`
	} `json:"_children"`
	ElementType              string `json:"_elementType"`
	AllowSync                string `json:"allowSync"`
	Art                      string `json:"art"`
	Banner                   string `json:"banner"`
	GrandparentContentRating string `json:"grandparentContentRating"`
	GrandparentRatingKey     string `json:"grandparentRatingKey"`
	GrandparentStudio        string `json:"grandparentStudio"`
	GrandparentTheme         string `json:"grandparentTheme"`
	GrandparentThumb         string `json:"grandparentThumb"`
	GrandparentTitle         string `json:"grandparentTitle"`
	Identifier               string `json:"identifier"`
	Key                      string `json:"key"`
	LibrarySectionID         string `json:"librarySectionID"`
	LibrarySectionTitle      string `json:"librarySectionTitle"`
	LibrarySectionUUID       string `json:"librarySectionUUID"`
	MediaTagPrefix           string `json:"mediaTagPrefix"`
	MediaTagVersion          string `json:"mediaTagVersion"`
	Nocache                  string `json:"nocache"`
	ParentIndex              string `json:"parentIndex"`
	ParentTitle              string `json:"parentTitle"`
	ParentYear               string `json:"parentYear"`
	SortAsc                  string `json:"sortAsc"`
	Theme                    string `json:"theme"`
	Thumb                    string `json:"thumb"`
	Title1                   string `json:"title1"`
	Title2                   string `json:"title2"`
	ViewGroup                string `json:"viewGroup"`
	ViewMode                 string `json:"viewMode"`
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
	Name                 string `json:"name" xml:"name,attr"`
	Product              string `json:"product" xml:"product,attr"`
	ProductVersion       string `json:"productVersion" xml:"productVersion,attr"`
	Platform             string `json:"platform" xml:"platform,attr"`
	PlatformVersion      string `json:"platformVersion" xml:"platformVersion,attr"`
	Device               int    `json:"device" xml:"device,attr"`
	ClientIdentifier     string `json:"clientIdentifier" xml:"clientIdentifier,attr"`
	CreatedAt            string `json:"createdAt" xml:"createdAt,attr"`
	LastSeenAt           string `json:"lastSeenAt" xml:"lastSeenAt,attr"`
	Provides             string `json:"provides" xml:"provides,attr"`
	Owned                string `json:"owned" xml:"owned,attr"`
	AccessToken          string `json:"accessToken" xml:"accessToken,attr"`
	HTTPSRequired        string `json:"httpsRequired" xml:"httpsRequired,attr"`
	Synced               string `json:"synced" xml:"synced,attr"`
	PublicAddressMatches string `json:"publicAddressMatches" xml:"publicAddressMatches,attr"`
	Presence             string `json:"presence" xml:"presence,attr"`
	Connection           []struct {
		Protocol string `json:"protocol" xml:"protocol,attr"`
		Address  string `json:"address" xml:"address,attr"`
		Port     string `json:"port" xml:"port,attr"`
		URI      string `json:"uri" xml:"uri,attr"`
		Local    int    `json:"local" xml:"local,attr"`
	} `json:"connection" xml:"Connection"`
}

// BaseAPIResponse info about the Plex Media Server
type BaseAPIResponse struct {
	Children []struct {
		ElementType string `json:"_elementType"`
		Count       string `json:"count"`
		Key         string `json:"key"`
		Title       string `json:"title"`
	} `json:"_children"`
	ElementType                   string `json:"_elementType"`
	AllowCameraUpload             string `json:"allowCameraUpload"`
	AllowChannelAccess            string `json:"allowChannelAccess"`
	AllowSharing                  string `json:"allowSharing"`
	AllowSync                     string `json:"allowSync"`
	BackgroundProcessing          string `json:"backgroundProcessing"`
	Certificate                   string `json:"certificate"`
	CompanionProxy                string `json:"companionProxy"`
	EventStream                   string `json:"eventStream "`
	FriendlyName                  string `json:"friendlyName"`
	HubSearch                     string `json:"hubSearch"`
	MachineIdentifier             string `json:"machineIdentifier"`
	Multiuser                     string `json:"multiuser"`
	MyPlex                        string `json:"myPlex"`
	MyPlexMappingState            string `json:"myPlexMappingState"`
	MyPlexSigninState             string `json:"myPlexSigninState"`
	MyPlexSubscription            string `json:"myPlexSubscription"`
	MyPlexUsername                string `json:"myPlexUsername"`
	OwnerFeatures                 string `json:"ownerFeatures"`
	Platform                      string `json:"platform"`
	PlatformVersion               string `json:"platformVersion"`
	PluginHost                    string `json:"pluginHost"`
	ReadOnlyLibraries             string `json:"readOnlyLibraries"`
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
	Updater                       string `json:"updater"`
	Version                       string `json:"version"`
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

// LibrarySections metadata of your library contents
type LibrarySections struct {
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
	ContentType      string
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

type CurrentSessionsVideo struct {
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
		ID    int    `xml:"id,attr"`
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
	Session              struct {
		ID        string `xml:"id,attr"`
		Bandwidth string `xml:"bandwidth,attr"`
		Location  string `xml:"location,attr"`
	}
	TranscodeSession struct {
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
}

// CurrentSessions is xml because plex returns a dynamic type (string or number) for the duration field
type CurrentSessions struct {
	XMLName xml.Name               `xml:"MediaContainer"`
	Size    string                 `xml:"size,attr"`
	Video   []CurrentSessionsVideo `xml:"Video"`
	Track   []struct {
		AddedAt              string `xml:"addedAt,attr"`
		Art                  string `xml:"art,attr"`
		ChapterSource        string `xml:"chapterSource,attr"`
		Duration             int    `xml:"duration,attr"`
		GrandparentArt       string `xml:"grandparentArt,attr"`
		GrandparentKey       string `xml:"grandparentKey,attr"`
		GrandparentRatingKey string `xml:"grandparentRatingKey,attr"`
		GrandparentThumb     string `xml:"grandparentThumb,attr"`
		GrandparentTitle     string `xml:"grandparentTitle,attr"`
		GUID                 string `xml:"guid,attr"`
		Index                string `xml:"index,attr"`
		Key                  string `xml:"key,attr"`
		LastViewedAt         string `xml:"lastViewedAt,attr"`
		LibrarySectionID     string `xml:"librarySectionID,attr"`
		ParentIndex          string `xml:"parentIndex,attr"`
		ParentKey            string `xml:"parentKey,attr"`
		ParentRatingKey      string `xml:"parentRatingKey,attr"`
		ParentTitle          string `xml:"parentTitle,attr"`
		RatingKey            string `xml:"ratingKey,attr"`
		SessionKey           string `xml:"sessionKey,attr"`
		Summary              string `xml:"summary,attr"`
		Tagline              string `xml:"tagline,attr"`
		Thumb                string `xml:"thumb,attr"`
		Title                string `xml:"title,attr"`
		Type                 string `xml:"type,attr"`
		UpdatedAt            string `xml:"updatedAt,attr"`
		ViewCount            int    `xml:"viewCount,attr"`
		ViewOffset           int    `xml:"viewOffset,attr"`
		Media                struct {
			AudioChannels string `xml:"audioChannels,attr"`
			AudioCodec    string `xml:"audioCodec,attr"`
			Bitrate       string `xml:"bitrate,attr"`
			Container     string `xml:"container,attr"`
			Duration      string `xml:"duration,attr"`
			ID            string `xml:"id,attr"`
			Part          struct {
				Container string `xml:"container,attr"`
				Duration  string `xml:"duration,attr"`
				File      string `xml:"file,attr"`
				ID        string `xml:"id,attr"`
				Key       string `xml:"key,attr"`
				Size      string `xml:"size,attr"`
				Stream    []struct {
					AudioChannelLayout string `xml:"audioChannelLayout,attr"`
					Bitrate            string `xml:"bitrate,attr"`
					BitrateMode        string `xml:"bitrateMode,attr"`
					Channels           string `xml:"channels,attr"`
					Codec              string `xml:"codec,attr"`
					Duration           string `xml:"duration,attr"`
					ID                 string `xml:"id,attr"`
					Index              string `xml:"index,attr"`
					SamplingRate       string `xml:"samplingRate,attr"`
					Selected           string `xml:"selected,attr"`
					StreamType         string `xml:"streamType,attr"`
				} `xml:"Stream"`
			} `xml:"Part"`
		} `xml:"Media"`
		User struct {
			ID    int    `xml:"id,attr"`
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
		TranscodeSession struct {
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
	} `xml:"Track"`
}
