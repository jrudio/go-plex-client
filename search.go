package plex

import "regexp"

// SearchPlex searches just like Search, but omits the last 4 results which are not relevant
// SearchPlex searches just like Search, but omits the last 4 results which are not relevant
func (p *Plex) SearchPlex(searchTerm string) (SearchHubContainer, error) {
	results, err := p.HubSearch(searchTerm)

	if err != nil {
		return SearchHubContainer{}, err
	}

	// Logic to limit results to top 4 (safely)
	// Since results are now in Hubs, we might want to limit the Metadata in the first Hub?
	// For now, returning full results to avoid logic assumptions on Hub structure.

	return results, nil
}

// ExtractKeyAndThumbFromURL extracts the rating key and thumbnail id from the url
func (p *Plex) ExtractKeyAndThumbFromURL(_url string) (string, string) {
	count := len(_url)

	var key string
	var thumbID string

	if count > 18 {
		k := _url[18:]
		keyCount := len(k)

		if keyCount > 7 {
			// Get key
			for ii := 0; ii < keyCount; ii++ {
				if k[ii:ii+1] == "/" {
					key = k[:ii]
					thumbID = k[ii+7:]
					break
				}

			}
		}
	}

	return key, thumbID
}

// ExtractKeyFromRatingKey extracts the key from the rating key url
func (*Plex) ExtractKeyFromRatingKey(key string) string {
	keyCount := len(key)

	if keyCount > 18 {
		key = key[18:]
		keyCount = len(key)

		if keyCount > 9 {
			for ii := 0; ii < keyCount; ii++ {
				if key[ii:ii+1] == "/" {
					key = key[:ii]
					break
				}
			}
		}

		return key
	}

	return key
}

// ExtractKeyFromRatingKeyRegex extracts the key from a rating key url via regex
func (p *Plex) ExtractKeyFromRatingKeyRegex(key string) string {
	r := regexp.MustCompile(`\d+`)
	result := r.FindStringSubmatch(key)

	if len(key) == 0 {
		return ""
	}

	return result[0]
}
