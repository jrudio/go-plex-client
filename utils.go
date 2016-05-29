package plex

import (
	"bytes"
	"net/http"
)

func (p *Plex) options(query string) (*http.Response, error) {
	client := p.HTTPClient

	req, reqErr := http.NewRequest("OPTIONS", query, nil)

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}

func (p *Plex) get(query string) (*http.Response, error) {
	client := p.HTTPClient

	req, reqErr := http.NewRequest("GET", query, nil)

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	req.Header.Add("Accept", p.headers.Accept)
	req.Header.Add("X-Plex-Platform", p.headers.Platform)
	req.Header.Add("X-Plex-Platform-Version", p.headers.PlatformVersion)
	req.Header.Add("X-Plex-Provides", p.headers.Provides)
	req.Header.Add("X-Plex-Client-Identifier", p.headers.ClientIdentifier)
	req.Header.Add("X-Plex-Product", p.headers.Product)
	req.Header.Add("X-Plex-Version", p.headers.Version)
	req.Header.Add("X-Plex-Device", p.headers.Device)
	// req.Header.Add("X-Plex-Container-Size", p.headers.ContainerSize)
	// req.Header.Add("X-Plex-Container-Start", p.headers.ContainerStart)
	req.Header.Add("X-Plex-Token", p.Token)

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}

func (p *Plex) delete(query string) (*http.Response, error) {
	client := p.HTTPClient

	req, reqErr := http.NewRequest("DELETE", query, nil)

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	req.Header.Add("Accept", p.headers.Accept)
	req.Header.Add("X-Plex-Platform", p.headers.Platform)
	req.Header.Add("X-Plex-Platform-Version", p.headers.PlatformVersion)
	req.Header.Add("X-Plex-Provides", p.headers.Provides)
	req.Header.Add("X-Plex-Client-Identifier", p.headers.ClientIdentifier)
	req.Header.Add("X-Plex-Product", p.headers.Product)
	req.Header.Add("X-Plex-Version", p.headers.Version)
	req.Header.Add("X-Plex-Device", p.headers.Device)
	// req.Header.Add("X-Plex-Container-Size", p.headers.ContainerSize)
	// req.Header.Add("X-Plex-Container-Start", p.headers.ContainerStart)
	req.Header.Add("X-Plex-Token", p.Token)

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}

func (p *Plex) post(query string, body []byte) (*http.Response, error) {
	client := p.HTTPClient

	req, reqErr := http.NewRequest("POST", query, bytes.NewBuffer(body))

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	// req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", p.headers.Accept)
	req.Header.Add("X-Plex-Platform", p.headers.Platform)
	req.Header.Add("X-Plex-Platform-Version", p.headers.PlatformVersion)
	req.Header.Add("X-Plex-Provides", p.headers.Provides)
	req.Header.Add("X-Plex-Client-Identifier", p.headers.ClientIdentifier)
	req.Header.Add("X-Plex-Product", p.headers.Product)
	req.Header.Add("X-Plex-Version", p.headers.Version)
	req.Header.Add("X-Plex-Device", p.headers.Device)
	// req.Header.Add("X-Plex-Container-Size", p.headers.ContainerSize)
	// req.Header.Add("X-Plex-Container-Start", p.headers.ContainerStart)
	req.Header.Add("X-Plex-Token", p.Token)

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}

func (p *Plex) put(query string) (*http.Response, error) {
	client := p.HTTPClient

	req, reqErr := http.NewRequest("PUT", query, nil)

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	// req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", p.headers.Accept)
	req.Header.Add("X-Plex-Platform", p.headers.Platform)
	req.Header.Add("X-Plex-Platform-Version", p.headers.PlatformVersion)
	req.Header.Add("X-Plex-Provides", p.headers.Provides)
	req.Header.Add("X-Plex-Client-Identifier", p.headers.ClientIdentifier)
	req.Header.Add("X-Plex-Product", p.headers.Product)
	req.Header.Add("X-Plex-Version", p.headers.Version)
	req.Header.Add("X-Plex-Device", p.headers.Device)
	// req.Header.Add("X-Plex-Container-Size", p.headers.ContainerSize)
	// req.Header.Add("X-Plex-Container-Start", p.headers.ContainerStart)
	req.Header.Add("X-Plex-Token", p.Token)

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}
