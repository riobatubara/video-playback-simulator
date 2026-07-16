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
	{"vid001", "Edge of Tomorrow", []string{"action", "sci-fi", "thriller", "time-travel", "aliens"}, 113},
	{"vid002", "The Grand Budapest Hotel", []string{"comedy", "drama", "adventure", "indie", "quirky"}, 99},
	{"vid003", "The Matrix", []string{"sci-fi", "action", "cyberpunk", "dystopian", "classic"}, 136},
	{"vid004", "La La Land", []string{"romance", "drama", "music", "musical", "hollywood", "jazz"}, 128},
	{"vid005", "Mad Max: Fury Road", []string{"action", "adventure", "post-apocalyptic", "car-chase", "survival"}, 120},
	{"vid006", "Inception", []string{"sci-fi", "action", "heist", "mind-bending", "psychological", "thriller"}, 148},
	{"vid007", "The Dark Knight", []string{"action", "drama", "superhero", "crime", "neo-noir", "thriller"}, 142},
	{"vid008", "Interstellar", []string{"sci-fi", "drama", "adventure", "space-travel", "cosmic", "science"}, 154},
	{"vid009", "Pulp Fiction", []string{"crime", "thriller", "drama", "pop-culture", "retro", "dark-comedy"}, 142},
	{"vid010", "Inglourious Basterds", []string{"action", "drama", "historical", "suspense", "war", "satire"}, 130},
}

// Internal map index for instant O(1) lookups instead of looping
var libraryIndex map[string]Video

func init() {
	// IMPROVEMENT: Build a lookup map on startup to eliminate loop bottlenecks
	libraryIndex = make(map[string]Video, len(videoLibrary))
	for _, v := range videoLibrary {
		libraryIndex[v.ID] = v
	}
}

// RandomVideo returns a randomly selected video from the library.
func RandomVideo() Video {
	// Keeps your exact shared global variable from the other file
	return videoLibrary[Randomize.Intn(len(videoLibrary))]
}

// VideoNameByID returns the name of a video given its ID.
func VideoNameByID(id string) string {
	if v, exists := libraryIndex[id]; exists {
		return v.Name
	}
	return "Unknown Title"
}

// TagsByID returns a comma-separated string of tags for a given video ID.
func TagsByID(id string) string {
	if v, exists := libraryIndex[id]; exists {
		return strings.Join(v.Tags, ",")
	}
	return ""
}

// DurationByID returns the duration of a video in seconds given its ID.
func DurationByID(id string) int {
	if v, exists := libraryIndex[id]; exists {
		return v.Duration
	}
	return 90
}
