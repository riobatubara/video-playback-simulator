package simulator

import (
	"fmt"
	"math/rand"
	"time"

	"video-playback-simulator/utils"
)

type Session struct {
	Index               int
	SessionID           string
	UserID              string
	VideoID             string
	Logger              func(int, int64, string, string, string)
	Emit                func(string, string, string)
	enablePause         bool
	enableSeek          bool
	enableBitrateChange bool
}

func NewSession(index int) *Session {
	s := &Session{
		Index:     index,
		SessionID: utils.GenerateSessionID(),
		UserID:    utils.RandomUserID(),
		VideoID:   utils.RandomVideoID(),
		Logger: func(idx int, ts int64, sessid, event, value string) {
			utils.LogEvent(idx, ts, sessid, event, value)
		},
		Emit: utils.EmitPayload,
	}

	// Randomly assign capabilities to each session
	s.enablePause = rand.Intn(2) == 0
	s.enableSeek = rand.Intn(2) == 0
	s.enableBitrateChange = rand.Intn(2) == 0

	return s
}

func (s *Session) Run() {
	s.sendInitialMetadata()
	duration := rand.Intn(60) + 30 // Simulate 30–90 seconds video
	playPosition := 0
	tsStart := time.Now().UnixMilli()
	s.EmitEvent("load", fmt.Sprintf("%d", tsStart))
	time.Sleep(300 * time.Millisecond)

	vb, ab := utils.RandomBitrate()
	s.EmitEvent("bitrate", fmt.Sprintf("%d,%d", vb, ab))
	s.EmitEvent("play", fmt.Sprintf("%d", time.Now().UnixMilli()))

	buffered := false

	for playPosition < duration {
		time.Sleep(500 * time.Millisecond)

		bufferCount := rand.Intn(3) + 1 // 1–3 buffers between playing
		for i := 0; i < bufferCount; i++ {
			s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
			time.Sleep(300 * time.Millisecond)
			buffered = true
		}

		playPosition += 5
		s.EmitEvent("playing", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(playPosition)))
		buffered = false

		// Optional events by session config
		switch rand.Intn(20) {
		case 2:
			if s.enablePause {
				s.EmitEvent("pause", fmt.Sprintf("%d", time.Now().UnixMilli()))
				delay := time.Duration(rand.Intn(5)+1) * time.Second
				time.Sleep(delay)
				s.EmitEvent("play", fmt.Sprintf("%d", time.Now().UnixMilli()))
				s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
				time.Sleep(300 * time.Millisecond)
				s.EmitEvent("playing", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(playPosition)))
			}
		case 3:
			if s.enableSeek {
				seekTo := rand.Intn(duration)
				s.EmitEvent("seek", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(seekTo)))
				playPosition = seekTo
				if rand.Intn(2) == 0 {
					vb, ab := utils.RandomBitrate()
					s.EmitEvent("bitrate", fmt.Sprintf("%d,%d", vb, ab))
				}
				s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
				time.Sleep(300 * time.Millisecond)
				s.EmitEvent("playing", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(playPosition)))
			}
		case 5:
			if s.enableBitrateChange {
				vb, ab := utils.RandomBitrate()
				s.EmitEvent("bitrate", fmt.Sprintf("%d,%d", vb, ab))
				s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
				buffered = true
			}
		}
	}

	if !buffered {
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
