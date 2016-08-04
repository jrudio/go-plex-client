package plex

// GetMediaTypeID returns plex's media type id
func GetMediaTypeID(mediaType string) string {
	switch mediaType {
	case "movie":
		return "1"
	case "show":
		return "2"
	case "season":
		return "3"
	case "episode":
		return "4"
	case "trailer":
		return "5"
	case "comic":
		return "6"
	case "person":
		return "7"
	case "artist":
		return "8"
	case "album":
		return "9"
	case "track":
		return "10"
	case "photoAlbum":
		return "11"
	case "picture":
		return "12"
	case "photo":
		return "13"
	case "clip":
		return "14"
	case "playlistItem":
		return "15"
	default:
		return mediaType
	}
}

// GetMediaType is a helper function that returns the media type. Usually, used after GetMetadata().
func GetMediaType(info MediaMetadata) string {
	if dType := info.Directory.Type; dType != "" {
		return dType
	}

	if vType := info.Video.Type; vType != "" {
		return vType
	}

	return ""
}
