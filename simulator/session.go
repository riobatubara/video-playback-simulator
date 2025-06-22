package simulator

import (
	"fmt"
	"math/rand"
	"time"

	"video-playback-simulator/utils"
)

type Session struct {
	Index     int
	SessionID string
	UserID    string
	VideoID   string
	Logger    func(int, int64, string, string, string)
	Emit      func(string, string, string)
}

func NewSession(index int) *Session {
	return &Session{
		Index:     index,
		SessionID: utils.GenerateSessionID(),
		UserID:    utils.RandomUserID(),
		VideoID:   utils.RandomVideoID(),
		Logger: func(idx int, ts int64, sessid, event, value string) {
			utils.LogEvent(idx, ts, sessid, event, value)
		},
		Emit: utils.EmitPayload,
	}
}

func (s *Session) Run() {
	s.sendInitialMetadata()

	duration := rand.Intn(60) + 30
	playPosition := 0

	s.EmitEvent("load", fmt.Sprintf("%d", time.Now().UnixMilli()))
	time.Sleep(300 * time.Millisecond)

	vb, ab := utils.RandomBitrate()
	s.EmitEvent("bitrate", fmt.Sprintf("%d,%d", vb, ab))
	s.EmitEvent("play", fmt.Sprintf("%d", time.Now().UnixMilli()))
	s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
	time.Sleep(300 * time.Millisecond)
	s.EmitEvent("playing", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(playPosition)))

	lastWasPlaying := true
	bufferedSinceLastPlay := true

	for playPosition < duration {
		time.Sleep(500 * time.Millisecond)

		// Random event chance
		eventRoll := rand.Intn(20)
		switch eventRoll {
		case 2: // pause
			s.EmitEvent("pause", fmt.Sprintf("%d", time.Now().UnixMilli()))
			time.Sleep(800 * time.Millisecond) // simulate wait before resume

			s.EmitEvent("play", fmt.Sprintf("%d", time.Now().UnixMilli()))
			s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
			time.Sleep(300 * time.Millisecond)
			s.EmitEvent("playing", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(playPosition)))
			lastWasPlaying = true
			bufferedSinceLastPlay = true
			continue

		case 3: // seek
			seekTo := rand.Intn(duration)
			s.EmitEvent("seek", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(seekTo)))
			s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
			time.Sleep(300 * time.Millisecond)
			playPosition = seekTo
			s.EmitEvent("playing", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(playPosition)))
			lastWasPlaying = true
			bufferedSinceLastPlay = true
			continue

		case 5: // network issue - bitrate + buffer
			vb, ab := utils.RandomBitrate()
			s.EmitEvent("bitrate", fmt.Sprintf("%d,%d", vb, ab))
			s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
			time.Sleep(300 * time.Millisecond)

			lastWasPlaying = false
			bufferedSinceLastPlay = true
		}

		// If last was playing, insert buffer to break repeat
		// If last was playing, emit 1–2 buffers before next playing
		if lastWasPlaying {
			bufferCount := rand.Intn(2) + 1 // 1 to 2
			for i := 0; i < bufferCount; i++ {
				s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
				time.Sleep(200 * time.Millisecond)
			}
			bufferedSinceLastPlay = true
		}

		playPosition += 5
		s.EmitEvent("playing", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(playPosition)))
		lastWasPlaying = true
		bufferedSinceLastPlay = false
	}

	// Ensure buffer before complete
	if !bufferedSinceLastPlay {
		s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
	}
	s.EmitEvent("complete", fmt.Sprintf("%d", time.Now().UnixMilli()))
	s.EmitEvent("unload", fmt.Sprintf("%d", time.Now().UnixMilli()))
}

func (s *Session) EmitEvent(name, value string) {
	ts := time.Now().UnixMilli()
	s.Emit(name, s.SessionID, value)
	s.Logger(s.Index, ts, s.SessionID, name, value)
}

func (s *Session) sendInitialMetadata() {
	meta := map[string]string{
		"sdk_ver":    "1.0.2",
		"geoip":      utils.RandomIP(),
		"uadev":      utils.RandomUA(),
		"duuid":      utils.GenerateUUID(),
		"user_id":    s.UserID,
		"video_id":   s.VideoID,
		"video_name": utils.VideoNameByID(s.VideoID),
		"tags":       utils.TagsByID(s.VideoID),
	}
	for k, v := range meta {
		s.EmitEvent(k, v)
	}
}
