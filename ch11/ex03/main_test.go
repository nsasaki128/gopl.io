package main

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "empty", input: "", expected: true},
		{name: "1 word", input: "a", expected: true},
		{name: "2 word not palindrome", input: "ab", expected: false},
		{name: "2 word palindrome", input: "aa", expected: true},
		{name: "3 word not palindrome", input: "abb", expected: false},
		{name: "3 word palindrome", input: "aba", expected: true},
		{name: "palindrome", input: "nomelonnolemon", expected: true},
		{name: "palindrome2", input: "kayak", expected: true},
		{name: "palindrome3", input: "detartrated", expected: true},
		{name: "sentence palindrome1", input: "A man, a plan, a canal: Panama", expected: true},
		{name: "sentence palindrome2", input: "Evil I did dwell; lewd did I live.", expected: true},
		{name: "sentence palindrome3", input: "Able was I ere I saw Elba", expected: true},
		{name: "French palindrome", input: "été", expected: true},
		{name: "Panama palindrome", input: "Et se resservir, ivresse reste.", expected: true},
		{name: "non-palindrome1", input: "nomelonnoolemon", expected: false},
		{name: "non-palindrome2", input: "palindrome", expected: false}, // non-palindrome
		{name: "semi-palindrome", input: "desserts", expected: false},   // semi-palindrome
		{name: "Japanese palindrome", input: "たけやぶやけた", expected: true},
		{name: "Japanese no palindrome", input: "たけやぶがやけた", expected: false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsPalindrome(testCase.input)
			if actual != testCase.expected {
				t.Errorf("%q expects as palindrome: %v but actual is %v", testCase.input, testCase.expected, actual)
			}
		})
	}
}

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(23) + 2 // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < n; {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		if !unicode.IsLetter(r) {
			continue
		}
		runes[i] = r
		i++
	}
	for i := 0; i < (n+1)/2; i++ {
		if runes[i] != runes[n-1-i] {
			return string(runes)
		}
	}
	return randomNonPalindrome(rng)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
func TestNonRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		q := randomNonPalindrome(rng)
		if IsPalindrome(q) {
			t.Errorf("IsNonPalindrome(%q) = %t", q, IsPalindrome(q))
		}
	}
}
