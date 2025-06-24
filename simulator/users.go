package simulator

import (
	"fmt"
)

var (
	firstNames = []string{"alex", "jordan", "maria", "chris", "taylor", "sam", "jamie", "morgan", "casey", "pat"}
	lastNames  = []string{"smith", "johnson", "lee", "martinez", "davis", "brown", "clark", "walker", "hill", "turner"}
	domains    = []string{"gmail.com", "yahoo.com", "outlook.com", "hotmail.com", "protonmail.com", "mail.com"}
)

// RandomUserID generates a fake email-like user ID.
func RandomUserID() string {
	fn := firstNames[Randomize.Intn(len(firstNames))]
	ln := lastNames[Randomize.Intn(len(lastNames))]
	num := Randomize.Intn(9000) + 1000
	domain := domains[Randomize.Intn(len(domains))]
	return fmt.Sprintf("%s.%s%d@%s", fn, ln, num, domain)
}
