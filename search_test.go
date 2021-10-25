package plex

import "testing"

func TestExtractKeyFromRatingKey(t *testing.T) {
	keys := [][]string{
		// Shows: test - expect
		{"/library/metadata/18/children", "18"},
		{"/library/metadata/1/children", "1"},
		// Movies: test - expect
		{"/library/metadata/797", "797"},
		{"/library/metadata/33", "33"},
		{"/library/metadata/700", "700"},
		{"/library/metadata/7", "7"},
	}

	p := Plex{}

	for _, k := range keys {
		key := p.ExtractKeyFromRatingKey(k[0])

		if key != k[1] {
			t.Errorf("Expected: %s \n Got: %s", k[1], key)
		}
	}
}

func TestExtractKeyFromRatingKeyRegex(t *testing.T) {
	keys := [][]string{
		// Shows: test - expect
		{"/library/metadata/18/children", "18"},
		{"/library/metadata/1/children", "1"},
		// Movies: test - expect
		{"/library/metadata/797", "797"},
		{"/library/metadata/33", "33"},
		{"/library/metadata/700", "700"},
		{"/library/metadata/7", "7"},
	}

	p := Plex{}

	for _, k := range keys {
		key := p.ExtractKeyFromRatingKeyRegex(k[0])

		if key != k[1] {
			t.Errorf("Expected: %s \n Got: %s", k[1], key)
		}
	}
}

func TestExtractKeyAndThumbFromURL(t *testing.T) {
	thumbs := [][]string{
		// test - expect
		{"/library/metadata/1/thumb/1459739349", "1", "1459739349"},
		{"/library/metadata/551/thumb/1455861333", "551", "1455861333"},
		{"/library/metadata/786/thumb/1463374779", "786", "1463374779"},
	}

	p := Plex{}

	for _, _thumb := range thumbs {
		key, thumb := p.ExtractKeyAndThumbFromURL(_thumb[0])

		if key != _thumb[1] {
			t.Errorf("Expected: %s \n Got: %s\n", _thumb[1], key)
		}

		if thumb != _thumb[2] {
			t.Errorf("Expected: %s\n Got: %s\n", _thumb[2], thumb)
		}
	}
}

func BenchmarkExtractKeyAndThumbFromURL(b *testing.B) {
	query := "/library/metadata/18/children"

	p := Plex{}

	for ii := 0; ii < b.N; ii++ {
		p.ExtractKeyAndThumbFromURL(query)
	}
}

func BenchmarkExtractKeyFromRatingKey(b *testing.B) {
	ratingKey := "/library/metadata/18/children"

	p := Plex{}

	for ii := 0; ii < b.N; ii++ {
		p.ExtractKeyFromRatingKey(ratingKey)
	}
}

func BenchmarkExtractKeyFromRatingKeyRegex(b *testing.B) {
	ratingKey := "/library/metadata/18/children"

	p := Plex{}

	for ii := 0; ii < b.N; ii++ {
		p.ExtractKeyFromRatingKeyRegex(ratingKey)
	}
}
