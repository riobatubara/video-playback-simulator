package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateSessionID() string {
	return uuid.New().String()
}

func GenerateUUID() string {
	return uuid.New().String()
}

func RandomUserID() string {
	return fmt.Sprintf("user%d@mail.com", rand.Intn(1000))
}

func RandomVideoID() string {
	return fmt.Sprintf("%d", rand.Intn(100))
}

func VideoNameByID(id string) string {
	// Dummy logic — could be replaced with actual mapping
	return "Pacific Rim"
}

func TagsByID(id string) string {
	return "Action,Adventure,Sci-Fi,International"
}

func RandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func RandomUA() string {
	oses := []string{"iOS", "Android", "Windows", "macOS"}
	browsers := []string{"Chrome", "Safari", "Firefox", "Edge"}
	devices := []string{"Phone", "Tablet", "Browser", "SmartTV"}
	return fmt.Sprintf("%s,16.%d,%s,%s,%s",
		oses[rand.Intn(len(oses))],
		rand.Intn(5)+1,
		browsers[rand.Intn(len(browsers))],
		"Browser",
		devices[rand.Intn(len(devices))],
	)
}

func CurrentTimestamp() int64 {
	return time.Now().UnixMilli()
}

func FormatTimestamp(t int64) string {
	return fmt.Sprintf("%d", t)
}

func RandomBitrate() (int, int) {
	videoBitrates := []int{1500, 2400, 3600, 4500, 6000} // example video bitrates in kbps
	audioBitrates := []int{96, 128, 160, 192, 256}       // example audio bitrates in kbps

	v := videoBitrates[rand.Intn(len(videoBitrates))]
	a := audioBitrates[rand.Intn(len(audioBitrates))]

	return v, a
}
