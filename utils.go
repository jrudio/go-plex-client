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

func (p *Plex) get(query string, h headers) (*http.Response, error) {
	client := p.HTTPClient

	req, reqErr := http.NewRequest("GET", query, nil)

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	req.Header.Add("Accept", h.Accept)
	req.Header.Add("X-Plex-Platform", h.Platform)
	req.Header.Add("X-Plex-Platform-Version", h.PlatformVersion)
	req.Header.Add("X-Plex-Provides", h.Provides)
	req.Header.Add("X-Plex-Client-Identifier", h.ClientIdentifier)
	req.Header.Add("X-Plex-Product", h.Product)
	req.Header.Add("X-Plex-Version", h.Version)
	req.Header.Add("X-Plex-Device", h.Device)
	// req.Header.Add("X-Plex-Container-Size", h.ContainerSize)
	// req.Header.Add("X-Plex-Container-Start", h.ContainerStart)
	req.Header.Add("X-Plex-Token", p.Token)

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}

func (p *Plex) delete(query string, h headers) (*http.Response, error) {
	client := p.HTTPClient

	req, reqErr := http.NewRequest("DELETE", query, nil)

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	req.Header.Add("Accept", h.Accept)
	req.Header.Add("X-Plex-Platform", h.Platform)
	req.Header.Add("X-Plex-Platform-Version", h.PlatformVersion)
	req.Header.Add("X-Plex-Provides", h.Provides)
	req.Header.Add("X-Plex-Client-Identifier", h.ClientIdentifier)
	req.Header.Add("X-Plex-Product", h.Product)
	req.Header.Add("X-Plex-Version", h.Version)
	req.Header.Add("X-Plex-Device", h.Device)
	// req.Header.Add("X-Plex-Container-Size", h.ContainerSize)
	// req.Header.Add("X-Plex-Container-Start", h.ContainerStart)
	req.Header.Add("X-Plex-Token", p.Token)

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}

func (p *Plex) post(query string, body []byte, h headers) (*http.Response, error) {
	client := p.HTTPClient

	req, reqErr := http.NewRequest("POST", query, bytes.NewBuffer(body))

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	// req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", h.Accept)
	req.Header.Add("X-Plex-Platform", h.Platform)
	req.Header.Add("X-Plex-Platform-Version", h.PlatformVersion)
	req.Header.Add("X-Plex-Provides", h.Provides)
	req.Header.Add("X-Plex-Client-Identifier", h.ClientIdentifier)
	req.Header.Add("X-Plex-Product", h.Product)
	req.Header.Add("X-Plex-Version", h.Version)
	req.Header.Add("X-Plex-Device", h.Device)
	// req.Header.Add("X-Plex-Container-Size", h.ContainerSize)
	// req.Header.Add("X-Plex-Container-Start", h.ContainerStart)
	req.Header.Add("X-Plex-Token", p.Token)

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}

func (p *Plex) put(query string, h headers) (*http.Response, error) {
	client := p.HTTPClient

	req, reqErr := http.NewRequest("PUT", query, nil)

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	// req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", h.Accept)
	req.Header.Add("X-Plex-Platform", h.Platform)
	req.Header.Add("X-Plex-Platform-Version", h.PlatformVersion)
	req.Header.Add("X-Plex-Provides", h.Provides)
	req.Header.Add("X-Plex-Client-Identifier", h.ClientIdentifier)
	req.Header.Add("X-Plex-Product", h.Product)
	req.Header.Add("X-Plex-Version", h.Version)
	req.Header.Add("X-Plex-Device", h.Device)
	// req.Header.Add("X-Plex-Container-Size", h.ContainerSize)
	// req.Header.Add("X-Plex-Container-Start", h.ContainerStart)
	req.Header.Add("X-Plex-Token", p.Token)

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}
