package simulator

import (
	"strconv"

	"github.com/google/uuid"
)

var (
	platforms = []string{
		// Desktop Platforms (Modernized with exact frozen format tokens)
		"Windows NT 10.0; Win64; x64",
		"Macintosh; Intel Mac OS X 10_15_7",
		"X11; Linux x86_64",
		"X11; CrOS x86_64 14541.0.0", // ChromeOS

		// Mobile Platforms (Updated to modern OS standards)
		"iPhone; CPU iPhone OS 18_2 like Mac OS X",
		"iPad; CPU OS 18_2 like Mac OS X",
		"Linux; Android 14; K", // Android 13+ standardizes on "K" token for privacy reduction
		"Linux; Android 15; K",
	}

	browsers = []string{
		// Google Chrome & Chromium-based (Updated to recent major version builds)
		"Chrome/146.0.0.0",
		"Chrome/145.0.6422.112",

		// Apple Safari (Updated to match iOS 18 ecosystem)
		"Version/18.2 Safari/605.1.15",
		"Version/18.0 Safari/604.1",

		// Mozilla Firefox (Updated to recent rapid-release branches)
		"Firefox/135.0",
		"Firefox/136.0",

		// Microsoft Edge & Opera
		"Edge/146.0.2567.48",
		"OPR/112.0.0.0", // Opera uses OPR token
	}
)

// RandomIP generates a pseudo-random IPv4 address.
func RandomIP() string {
	// IMPROVEMENT: Direct string concatenation completely bypasses fmt runtime reflection
	return strconv.Itoa(Randomize.Intn(223)+1) + "." +
		strconv.Itoa(Randomize.Intn(256)) + "." +
		strconv.Itoa(Randomize.Intn(256)) + "." +
		strconv.Itoa(Randomize.Intn(256))
}

// RandomUA returns a fake User-Agent string.
func RandomUA() string {
	platform := platforms[Randomize.Intn(len(platforms))]
	browser := browsers[Randomize.Intn(len(browsers))]

	// IMPROVEMENT: Concatenation lets the Go compiler calculate memory size allocations instantly
	return "Mozilla/5.0 (" + platform + ") AppleWebKit/537.36 (KHTML, like Gecko) " + browser + " Safari/537.36"
}

// GenerateUUID generates a unique string identifier
func GenerateUUID() string {
	// OPTIMIZATION: Bypasses the dual allocation overhead of uuid.New().String()
	// uuid.NewString() allocates memory directly into the final format layout
	return uuid.NewString()
}
