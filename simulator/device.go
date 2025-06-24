// package simulator

// import (
// 	"fmt"

// 	"github.com/google/uuid"
// )

// var (
// 	publicIPBlocks = []int{
// 		23, 31, 45, 52, 66, 72, 96, 104, 107, 129, 136, 142,
// 		151, 158, 162, 172, 173, 184, 192, 198, 199, 208, 216,
// 	}

// 	userAgents = []string{
// 		// Desktop
// 		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36",
// 		"Mozilla/5.0 (Macintosh; Intel Mac OS X 13_5_1) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3 Safari/605.1.15",
// 		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.6367.91 Safari/537.36",
// 		"Mozilla/5.0 (Windows NT 10.0; rv:124.0) Gecko/20100101 Firefox/124.0",

// 		// Mobile
// 		"Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
// 		"Mozilla/5.0 (Linux; Android 14; Pixel 7 Pro) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Mobile Safari/537.36",
// 		"Mozilla/5.0 (Linux; Android 13; SM-G991B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Mobile Safari/537.36",
// 		"Mozilla/5.0 (iPad; CPU OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Mobile/15E148 Safari/604.1",
// 	}
// )

// func RandomIP() string {
// 	o1 := publicIPBlocks[Randomizer.Intn(len(publicIPBlocks))]
// 	o2 := Randomizer.Intn(256)
// 	o3 := Randomizer.Intn(256)
// 	o4 := Randomizer.Intn(256)

// 	return fmt.Sprintf("%d.%d.%d.%d", o1, o2, o3, o4)
// }

// func RandomUA() string {
// 	return userAgents[Randomizer.Intn(len(userAgents))]
// }

// func GenerateUUID() string {
// 	return uuid.New().String()
// }

package simulator

import (
	"fmt"

	"github.com/google/uuid"
)

var (
	IPBlocks = []int{
		23, 31, 45, 52, 66, 72, 96, 104, 107, 129, 136, 142,
		151, 158, 162, 172, 173, 184, 192, 198, 199, 208, 216,
	}

	platforms = []string{"Windows NT 10.0", "Macintosh; Intel Mac OS X 10_15_7", "X11; Linux x86_64", "iPhone; CPU iPhone OS 15_2 like Mac OS X", "Android 11"}
	browsers  = []string{"Chrome/91.0.4472.124", "Firefox/89.0", "Safari/604.1", "Edge/91.0.864.48", "Opera/76.0.4017.123"}
)

// RandomIP generates a pseudo-random IPv4 address.
func RandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		Randomize.Intn(223)+1,
		Randomize.Intn(256),
		Randomize.Intn(256),
		Randomize.Intn(256),
	)
}

// RandomUA returns a fake User-Agent string.
func RandomUA() string {
	platform := platforms[Randomize.Intn(len(platforms))]
	browser := browsers[Randomize.Intn(len(browsers))]
	return fmt.Sprintf("Mozilla/5.0 (%s) AppleWebKit/537.36 (KHTML, like Gecko) %s Safari/537.36", platform, browser)
}

func GenerateUUID() string {
	return uuid.New().String()
}
