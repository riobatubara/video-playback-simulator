package simulator

import (
	"strings"
)

type Video struct {
	ID       string
	Name     string
	Tags     []string
	Duration int // in seconds
}

var videoLibrary = []Video{
	{"vid001", "Edge of Tomorrow", []string{"action", "sci-fi", "thriller"}, 95},
	{"vid002", "The Grand Budapest Hotel", []string{"comedy", "drama"}, 100},
	{"vid003", "The Matrix", []string{"sci-fi", "action"}, 110},
	{"vid004", "La La Land", []string{"romance", "drama", "music"}, 128},
	{"vid005", "Mad Max: Fury Road", []string{"action", "adventure"}, 105},
}

// RandomVideo returns a randomly selected video from the library.
func RandomVideo() Video {
	return videoLibrary[Randomize.Intn(len(videoLibrary))]
}

// VideoNameByID returns the name of a video given its ID.
func VideoNameByID(id string) string {
	for _, v := range videoLibrary {
		if v.ID == id {
			return v.Name
		}
	}
	return "Unknown Title"
}

// TagsByID returns a comma-separated string of tags for a given video ID.
func TagsByID(id string) string {
	for _, v := range videoLibrary {
		if v.ID == id {
			return strings.Join(v.Tags, ",")
		}
	}
	return ""
}

// DurationByID returns the duration of a video in seconds given its ID.
func DurationByID(id string) int {
	for _, v := range videoLibrary {
		if v.ID == id {
			return v.Duration
		}
	}
	return 90
}
