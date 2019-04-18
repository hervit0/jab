package consecutive

import (
	"strings"
)

func LongestConsec(chain []string, k int) string {
	chainLength := len(chain)
	if chainLength == 0 || k > chainLength || k <= 0 {
		return ""
	}

	longestConsecutiveChain := ""
	for n := 0; n <= chainLength-k; n++ {
		subChain := strings.Join(chain[n:n+k], "")

		if len(subChain) > len(longestConsecutiveChain) {
			longestConsecutiveChain = subChain
		}
	}

	return longestConsecutiveChain
}
