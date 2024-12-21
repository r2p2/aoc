package d22

import (
	"maps"
	"slices"
	"strconv"
	"strings"
)

func Part1(input string) string {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		n, _ := strconv.Atoi(line)
		n2000 := nextNth(n, 2000)
		sum += n2000
	}
	return strconv.Itoa(sum)
}

func Part2(input string) string {
	bds := []map[decission]int{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		n, _ := strconv.Atoi(line)
		bds = append(bds, byingDecissions(n))
	}

	// extract unique_decissions
	uniqueDecissions := map[decission]bool{}
	for i := range bds {
		for k := range maps.Keys(bds[i]) {
			if _, ok := uniqueDecissions[k]; ok {
				continue
			}

			uniqueDecissions[k] = true
		}
	}

	// apply each decision on each bds, sum bananas
	bananabilities := []int{}
	for decission := range maps.Keys(uniqueDecissions) {
		sumBananas := 0
		for j := range bds {
			if bananas, ok := bds[j][decission]; ok {
				sumBananas += bananas
			}
		}
		bananabilities = append(bananabilities, sumBananas)
	}

	// return max bananas
	return strconv.Itoa(slices.Max(bananabilities))
}

func nextNth(secret, n int) int {
	for i := 0; i < n; i++ {
		secret = nextSecret(secret)
	}
	return secret
}

func byingDecissions(secret int) (decissions map[decission]int) {
	decissions = make(map[decission]int)
	bananas := nextNthExtended(secret, 2000)
	for i := 1; i <= len(bananas)-4; i++ {
		curr := decission{
			bananas[i] - bananas[i-1],
			bananas[i+1] - bananas[i+1-1],
			bananas[i+2] - bananas[i+2-1],
			bananas[i+3] - bananas[i+3-1],
		}
		if _, ok := decissions[curr]; ok {
			// Don't overwrite existing decissions
			continue
		}

		decissions[curr] = bananas[i+3]
	}

	return
}

type decission struct {
	a, b, c, d int
}

func nextNthExtended(secret, n int) (bananas []int) {
	bananas = append(bananas, secret%10)
	for i := 0; i < n; i++ {
		secret = nextSecret(secret)
		bananas = append(bananas, secret%10)
	}
	return
}

func nextSecret(secret int) int {
	const (
		prune = 16777216
	)
	secret = (secret*64 ^ secret) % prune
	secret = (secret/32.0 ^ secret) % prune
	secret = (secret*2048 ^ secret) % prune
	return secret
}

func mod(a, b int) int {
	return (a%b + b) % b
}
