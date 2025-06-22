package simulator

const (
	EventSDKVer    = "sdk_ver"
	EventGeoIP     = "geoip"
	EventUADev     = "uadev"
	EventDeviceID  = "duuid"
	EventUserID    = "user_id"
	EventVideoID   = "video_id"
	EventVideoName = "video_name"
	EventTags      = "tags"

	EventLoad     = "load"
	EventPlay     = "play"
	EventPlaying  = "playing"
	EventBuffer   = "buffer"
	EventPause    = "pause"
	EventSeek     = "seek"
	EventBitrate  = "bitrate"
	EventComplete = "complete"
	EventUnload   = "unload"
	EventError    = "error"
)

// You could expand this later with type-safe structs or helper functions
// for building and validating event payloads if needed.
