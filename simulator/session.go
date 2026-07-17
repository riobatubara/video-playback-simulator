package simulator

import (
	"strconv"
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

	s.EmitEvent("load", strconv.FormatInt(tsStart, 10))
	time.Sleep(300 * time.Millisecond)

	vb, ab := RandomBitrate()
	s.EmitEvent("bitrate", strconv.Itoa(vb)+","+strconv.Itoa(ab))
	s.EmitEvent("play", strconv.FormatInt(time.Now().UnixMilli(), 10))

	buffered := false

	// FIX: Limit total random backward jumps to prevent infinite execution traps
	seekAttempts := 0

	for playPosition < duration {
		time.Sleep(500 * time.Millisecond)

		bufferCount := Randomize.Intn(3) + 1
		for i := 0; i < bufferCount; i++ {
			s.EmitEvent("buffer", strconv.FormatInt(time.Now().UnixMilli(), 10))
			time.Sleep(300 * time.Millisecond)
			buffered = true
		}

		playPosition += 5
		// FIX: String conversion avoids heavy runtime allocation chains
		s.EmitEvent("playing", strconv.FormatInt(time.Now().UnixMilli(), 10)+","+strconv.FormatFloat(float64(playPosition), 'f', 2, 64))
		buffered = false

		switch Randomize.Intn(20) {
		case 2:
			if s.enablePause {
				s.EmitEvent("pause", strconv.FormatInt(time.Now().UnixMilli(), 10))
				time.Sleep(time.Duration(Randomize.Intn(5)+1) * time.Second)
				s.EmitEvent("play", strconv.FormatInt(time.Now().UnixMilli(), 10))
				s.EmitEvent("buffer", strconv.FormatInt(time.Now().UnixMilli(), 10))
				time.Sleep(300 * time.Millisecond)
				s.EmitEvent("playing", strconv.FormatInt(time.Now().UnixMilli(), 10)+","+strconv.FormatFloat(float64(playPosition), 'f', 2, 64))
			}
		case 3:
			// FIX: Safety constraint allows up to 3 backward seek actions per session to prevent infinite execution traps
			if s.enableSeek && seekAttempts < 3 {
				seekTo := Randomize.Intn(duration)
				if seekTo < playPosition {
					seekAttempts++
				}
				s.EmitEvent("seek", strconv.FormatInt(time.Now().UnixMilli(), 10)+","+strconv.FormatFloat(float64(seekTo), 'f', 2, 64))
				playPosition = seekTo

				if Randomize.Intn(2) == 0 {
					vb, ab := RandomBitrate()
					s.EmitEvent("bitrate", strconv.Itoa(vb)+","+strconv.Itoa(ab))
				}
				s.EmitEvent("buffer", strconv.FormatInt(time.Now().UnixMilli(), 10))
				time.Sleep(300 * time.Millisecond)
				s.EmitEvent("playing", strconv.FormatInt(time.Now().UnixMilli(), 10)+","+strconv.FormatFloat(float64(playPosition), 'f', 2, 64))
			}
		case 5:
			if s.enableBitrateChange {
				vb, ab := RandomBitrate()
				s.EmitEvent("bitrate", strconv.Itoa(vb)+","+strconv.Itoa(ab))
				s.EmitEvent("buffer", strconv.FormatInt(time.Now().UnixMilli(), 10))
				buffered = true
			}
		}
	}

	if !buffered {
		s.EmitEvent("buffer", strconv.FormatInt(time.Now().UnixMilli(), 10))
	}
	s.EmitEvent("complete", strconv.FormatInt(time.Now().UnixMilli(), 10))
	s.EmitEvent("unload", strconv.FormatInt(time.Now().UnixMilli(), 10))
}

func (s *Session) EmitEvent(name, value string) {
	ts := time.Now().UnixMilli()
	s.Emit(name, s.SessionID, value)
	s.Logger(s.Index, ts, s.SessionID, name, value)
}

func (s *Session) sendInitialMetadata() {
	// FIX: Use an ordered structure array instead of maps to guarantee deterministic event firing timelines
	type metaField struct {
		key, val string
	}
	meta := []metaField{
		{"sdk_ver", "1.0.2"},
		{"geoip", RandomIP()},
		{"uadev", RandomUA()},
		{"duuid", GenerateUUID()},
		{"user_id", s.UserID},
		{"video_id", s.Video.ID},
		{"video_name", s.Video.Name},
		{"tags", TagsByID(s.Video.ID)},
	}

	for _, item := range meta {
		s.EmitEvent(item.key, item.val)
	}
}
