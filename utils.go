package plex

import (
	"net/http"
	"time"
)

func (r *request) options(query string) (*http.Response, error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

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

func (r *request) get(query string) (*http.Response, error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	req, reqErr := http.NewRequest("GET", query, nil)

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	req.Header.Add("Accept", r.Accept)
	req.Header.Add("X-Plex-Platform", r.Platform)
	req.Header.Add("X-Plex-Platform-Version", r.PlatformVersion)
	req.Header.Add("X-Plex-Provides", r.Provides)
	req.Header.Add("X-Plex-Client-Identifier", r.ClientIdentifier)
	req.Header.Add("X-Plex-Product", r.Product)
	req.Header.Add("X-Plex-Version", r.Version)
	req.Header.Add("X-Plex-Device", r.Device)
	req.Header.Add("X-Plex-Container-Size", r.ContainerSize)
	req.Header.Add("X-Plex-Container-Start", r.ContainerStart)
	req.Header.Add("X-Plex-Token", r.Token)

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}

func (r *request) delete(query string) (*http.Response, error) {
	client := http.Client{
		Timeout: 3 * time.Second,
	}

	req, reqErr := http.NewRequest("DELETE", query, nil)

	if reqErr != nil {
		return &http.Response{}, reqErr
	}

	req.Header.Add("Accept", r.Accept)
	req.Header.Add("X-Plex-Platform", r.Platform)
	req.Header.Add("X-Plex-Platform-Version", r.PlatformVersion)
	req.Header.Add("X-Plex-Provides", r.Provides)
	req.Header.Add("X-Plex-Client-Identifier", r.ClientIdentifier)
	req.Header.Add("X-Plex-Product", r.Product)
	req.Header.Add("X-Plex-Version", r.Version)
	req.Header.Add("X-Plex-Device", r.Device)
	req.Header.Add("X-Plex-Container-Size", r.ContainerSize)
	req.Header.Add("X-Plex-Container-Start", r.ContainerStart)
	req.Header.Add("X-Plex-Token", r.Token)

	resp, err := client.Do(req)

	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}

func (r *request) setRequestHeaders(newRequestInfo request) {
	r.Platform = newRequestInfo.Platform
	r.PlatformVersion = newRequestInfo.PlatformVersion
	r.Provides = newRequestInfo.Provides
	r.ClientIdentifier = newRequestInfo.ClientIdentifier
	r.Product = newRequestInfo.Product
	r.Version = newRequestInfo.Version
	r.Device = newRequestInfo.Device
	r.ContainerSize = newRequestInfo.ContainerSize
	r.ContainerStart = newRequestInfo.ContainerStart
	r.Token = newRequestInfo.Token
	r.Accept = newRequestInfo.Accept
}
