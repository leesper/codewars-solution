package kata

import "fmt"

func PrinterError(s string) string {
	denom := len(s)
	numer := 0
	for _, r := range s {
		if r > 'm' {
			numer++
		}
	}
	return fmt.Sprintf("%d/%d", numer, denom)
}
