package simulator

import (
	"strconv"

	"github.com/google/uuid"
)

type uaGroup struct {
	platforms []string
	browsers  []string
}

var (
	// Categorize the data into logical device ecosystems
	desktopUA = uaGroup{
		platforms: []string{
			"Windows NT 10.0; Win64; x64",
			"Macintosh; Intel Mac OS X 10_15_7",
			"X11; Linux x86_64",
			"X11; CrOS x86_64 14541.0.0",
		},
		browsers: []string{
			// Google Chrome & Chromium-based (Updated to recent major version builds)
			"Chrome/146.0.0.0",
			"Chrome/145.0.6422.112",

			// Mozilla Firefox (Updated to recent rapid-release branches)
			"Firefox/135.0",
			"Firefox/136.0",

			// Microsoft Edge & Opera
			"Edge/146.0.2567.48",
			"OPR/112.0.0.0",
		},
	}

	mobileUA = uaGroup{
		platforms: []string{
			"iPhone; CPU iPhone OS 18_2 like Mac OS X",
			"iPad; CPU OS 18_2 like Mac OS X",
			"Linux; Android 14; K",
			"Linux; Android 15; K",
		},
		browsers: []string{
			"Version/18.2 Mobile/15E148 Safari/605.1.15",
			"Version/18.0 Mobile/15E148 Safari/604.1",
			"Chrome/146.0.0.0 Mobile Safari/537.36",
			"Firefox/135.0 Mobile",
		},
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
	var group uaGroup

	// 50/50 chance to pick a Desktop or Mobile ecosystem profile
	if Randomize.Intn(2) == 0 {
		group = desktopUA
	} else {
		group = mobileUA
	}

	platform := group.platforms[Randomize.Intn(len(group.platforms))]
	browser := group.browsers[Randomize.Intn(len(group.browsers))]

	// IMPROVEMENT: Concatenation lets the Go compiler calculate memory size allocations instantly
	return "Mozilla/5.0 (" + platform + ") AppleWebKit/537.36 (KHTML, like Gecko) " + browser + " Safari/537.36"
}

// GenerateUUID generates a unique string identifier
func GenerateUUID() string {
	// OPTIMIZATION: Bypasses the dual allocation overhead of uuid.New().String()
	// uuid.NewString() allocates memory directly into the final format layout
	return uuid.NewString()
}
