package simulator

import (
	"strconv"
)

var (
	firstNames = []string{
		"alex", "jordan", "maria", "chris", "taylor", "sam", "jamie", "morgan", "casey", "pat",
		"liam", "olivia", "noah", "emma", "oliver", "ava", "elijah", "charlotte", "william", "sophia",
		"james", "amelia", "benjamin", "isabella", "lucas", "mia", "henry", "evelyn", "alexander", "harper",
	}
	lastNames = []string{
		"smith", "johnson", "lee", "martinez", "davis", "brown", "clark", "walker", "hill", "turner",
		"jones", "garcia", "rodriguez", "wilson", "thomas", "anderson", "taylor", "moore", "jackson", "martin",
		"white", "lopez", "harris", "sanchez", "clark", "ramirez", "lewis", "robinson", "walker", "young",
	}
	domains = []string{"gmail.com", "yahoo.com", "outlook.com", "hotmail.com", "protonmail.com", "mail.com"}
)

// RandomUserID generates a fake email-like user ID.
func RandomUserID() string {
	fn := firstNames[Randomize.Intn(len(firstNames))]
	ln := lastNames[Randomize.Intn(len(lastNames))]
	num := Randomize.Intn(9000) + 1000
	domain := domains[Randomize.Intn(len(domains))]

	// IMPROVEMENT: Avoid fmt.Sprintf parsing reflection.
	// Explicit string concatenation is optimized directly by the Go compiler.
	return fn + "." + ln + strconv.Itoa(num) + "@" + domain
}
