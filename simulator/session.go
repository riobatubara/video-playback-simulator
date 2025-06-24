package simulator

import (
	"fmt"
	"time"

	"video-playback-simulator/utils"
)

type Session struct {
	Index               int
	SessionID           string
	UserID              string
	Video               Video
	Logger              func(int, int64, string, string, string)
	Emit                func(string, string, string)
	enablePause         bool
	enableSeek          bool
	enableBitrateChange bool
}

func NewSession(index int) *Session {
	video := RandomVideo()

	s := &Session{
		Index:     index,
		SessionID: GenerateSessionID(),
		UserID:    RandomUserID(),
		Video:     video,
		Logger: func(idx int, ts int64, sessid, event, value string) {
			utils.LogEvent(idx, ts, sessid, event, value)
		},
		Emit: utils.EmitPayload,
	}

	// Random session capabilities
	s.enablePause = Randomize.Intn(2) == 0
	s.enableSeek = Randomize.Intn(2) == 0
	s.enableBitrateChange = Randomize.Intn(2) == 0

	return s
}

func (s *Session) Run() {
	s.sendInitialMetadata()

	duration := s.Video.Duration
	playPosition := 0
	tsStart := time.Now().UnixMilli()

	s.EmitEvent("load", fmt.Sprintf("%d", tsStart))
	time.Sleep(300 * time.Millisecond)

	vb, ab := RandomBitrate()
	s.EmitEvent("bitrate", fmt.Sprintf("%d,%d", vb, ab))
	s.EmitEvent("play", fmt.Sprintf("%d", time.Now().UnixMilli()))

	buffered := false

	for playPosition < duration {
		time.Sleep(500 * time.Millisecond)

		// Buffer before playing
		bufferCount := Randomize.Intn(3) + 1
		for i := 0; i < bufferCount; i++ {
			s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
			time.Sleep(300 * time.Millisecond)
			buffered = true
		}

		playPosition += 5
		s.EmitEvent("playing", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(playPosition)))
		buffered = false

		switch Randomize.Intn(20) {
		case 2:
			if s.enablePause {
				s.EmitEvent("pause", fmt.Sprintf("%d", time.Now().UnixMilli()))
				time.Sleep(time.Duration(Randomize.Intn(5)+1) * time.Second)
				s.EmitEvent("play", fmt.Sprintf("%d", time.Now().UnixMilli()))
				s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
				time.Sleep(300 * time.Millisecond)
				s.EmitEvent("playing", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(playPosition)))
			}
		case 3:
			if s.enableSeek {
				seekTo := Randomize.Intn(duration)
				s.EmitEvent("seek", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(seekTo)))
				playPosition = seekTo

				if Randomize.Intn(2) == 0 {
					vb, ab := RandomBitrate()
					s.EmitEvent("bitrate", fmt.Sprintf("%d,%d", vb, ab))
				}
				s.EmitEvent("buffer", fmt.Sprintf("%d", time.Now().UnixMilli()))
				time.Sleep(300 * time.Millisecond)
				s.EmitEvent("playing", fmt.Sprintf("%d,%.2f", time.Now().UnixMilli(), float64(playPosition)))
			}
		case 5:
			if s.enableBitrateChange {
				vb, ab := RandomBitrate()
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
		"geoip":      RandomIP(),
		"uadev":      RandomUA(),
		"duuid":      GenerateUUID(),
		"user_id":    s.UserID,
		"video_id":   s.Video.ID,
		"video_name": s.Video.Name,
		"tags":       TagsByID(s.Video.ID),
	}
	for k, v := range meta {
		s.EmitEvent(k, v)
	}
}
