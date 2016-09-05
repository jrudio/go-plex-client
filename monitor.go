package plex

import (
	"fmt"
	"time"
)

// Monitor holds the state for monitoring a plex server
type Monitor struct {
	PlexConn          *Plex
	Init              bool
	Interval          int
	Userlist          ListInterface
	doneChan          chan bool
	killTranscodeChan chan string
	videoSessionsChan chan CurrentSessions
	audioSessionsChan chan CurrentSessions
}

// ListInterface allows any datastore for users
type ListInterface interface {
	User(id int) (MonitoredUser, error)
	AddUser(id int, username, ratingKey string) error
	RemoveUser(id int) error
	SetField(id int, field string, value string) error
	Count() int
}

// MonitoredUser is a tracked user
type MonitoredUser struct {
	UserID         int    `json:"userID"`
	Duration       int    `json:"duration"`
	KillingSession bool   `json:"killingSession"`
	KilledSession  bool   `json:"killedSession"`
	RatingKey      string `json:"ratingKey"`
	IsTranscode    bool   `json:"isTranscode"`
	IsDirectPlay   bool   `json:"isDirectPlay"`
}

// Start grabs the session info on a specified interval
// and is considered the producer in the monitoring pipeline
func (m Monitor) Start() {
	if m.Interval == 0 {
		m.Interval = 1500
	}

	m.doneChan = make(chan bool)
	m.videoSessionsChan = make(chan CurrentSessions, 0)
	m.audioSessionsChan = make(chan CurrentSessions, 0)
	m.killTranscodeChan = make(chan string, 0)

	// listen for actions that need to talk to Plex
	go m.listenForPlexActions()

	// poll /status/sessions
	go m.pollSessions()

	go m.listenForNewVideoSessions()
	go m.listenForNewAudioSessions()

	m.Init = true
}

func (m Monitor) listenForPlexActions() {
	for transcodeKey := range m.killTranscodeChan {
		killedSession, err := m.PlexConn.KillTranscodeSession(transcodeKey)

		if err != nil {
			fmt.Printf("\t%s\n", err.Error())
		}

		fmt.Printf("\tSession killed: %t\n", killedSession)
	}
}

func (m Monitor) listenForNewVideoSessions() {
	for sessions := range m.videoSessionsChan {
		// fmt.Println("\tChecking video sessions")
		m.checkVideoSessions(sessions)
	}
}

func (m Monitor) listenForNewAudioSessions() {
	for sessions := range m.audioSessionsChan {
		// fmt.Println("\tChecking audio sessions")
		m.checkAudioSessions(sessions)
	}
}

func (m Monitor) pollSessions() {
	for {
		// Try not to bombard the server with too many requests
		time.Sleep(time.Duration(m.Interval) * time.Millisecond)

		if m.Userlist.Count() == 0 {
			fmt.Println("monitor list is empty")
			continue
		}

		fmt.Println("getting new sessions")

		sessions, err := m.PlexConn.GetSessions()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		// fmt.Println("Received new sessions")
		m.videoSessionsChan <- sessions
		m.audioSessionsChan <- sessions

	}
}

// checkVideoSessions and checkAudioSessions do the same thing but check different sections

func (m Monitor) checkVideoSessions(currentSessions CurrentSessions) {
	sessions := currentSessions.Video

	for _, s := range sessions {
		userID := s.User.ID
		username := s.User.Title

		key := s.RatingKey

		user, err := m.Userlist.User(userID)

		if err != nil {
			fmt.Printf("failed to retrieve user with id of %d: %v\n", userID, err)

			continue
		}

		if key == user.RatingKey {
			continue
		}

		// skip check if kill session has been initialized
		if user.KilledSession || user.KillingSession && user.IsTranscode {
			// fmt.Printf("stop monitoring user %d\n", user.UserID)

			continue
		}

		grandparentTitle := s.GrandparentTitle
		title := s.Title

		if grandparentTitle != "" {
			fmt.Printf("%s is watching a forbidden piece of media - %s: %s (%s)\n", username, grandparentTitle, title, key)
		} else {
			fmt.Printf("%s is watching a forbidden piece of media: %s (%s)\n", username, title, key)
		}

		transcodeKey := s.TranscodeSession.Key

		if transcodeKey != "" {
			fmt.Printf("\ttranscode session\n")

			m.Userlist.SetField(userID, "isTranscoding", "1")
			m.Userlist.SetField(userID, "killingSession", "1")

			m.killTranscodeChan <- transcodeKey

		} else {
			fmt.Println("\tdirect play session")

			// can't kill sessions like above
			m.Userlist.SetField(userID, "isDirectPlay", "1")

			// TODO we need to revoke plex user access to our server
		}

	}
}

func (m Monitor) checkAudioSessions(currentSessions CurrentSessions) {
	// sessions := currentSessions.Track

	// for _, s := range sessions {
	// 	userID := s.User.ID
	// 	username := s.User.Title

	// 	key := s.RatingKey

	// 	userInterface, hasUser := m.MonitoredList.Get(userID)

	// 	if !hasUser {
	// 		continue
	// 	}

	// 	user := userInterface.(MonitoredUser)

	// 	if key == user.RatingKey {
	// 		continue
	// 	}

	// 	// skip check if kill session has been initialized
	// 	if user.KilledSession || user.KillingSession && user.IsTranscode {
	// 		// fmt.Printf("stop monitoring user %d\n", user.UserID)

	// 		// Delete user id from list
	// 		m.RemoveUser(userID)
	// 		continue
	// 	}

	// 	grandparentTitle := s.GrandparentTitle
	// 	title := s.Title

	// 	if grandparentTitle != "" {
	// 		fmt.Printf("%s is watching a forbidden piece of media - %s: %s (%s)\n", username, grandparentTitle, title, key)
	// 	} else {
	// 		fmt.Printf("%s is watching a forbidden piece of media: %s (%s)\n", username, title, key)
	// 	}

	// 	transcodeKey := s.TranscodeSession.Key

	// 	if transcodeKey != "" {
	// 		fmt.Printf("\ttranscode session\n")

	// 		// user.KillingSession = true
	// 		user.IsTranscode = true

	// 		m.MonitoredList.Set(userID, user)

	// 		m.killTranscodeSessionChan <- transcodeKey

	// 	} else {
	// 		fmt.Println("\tdirect play session")

	// 		// can't kill sessions like above

	// 		user.IsDirectPlay = true

	// 		m.MonitoredList.Set(userID, user)

	// 		// TODO we need to revoke plex user access to our server
	// 	}

	// }
}
