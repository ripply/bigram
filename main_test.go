package main

import "testing"

func TestExample(t *testing.T) {
	example := "the quick brown fox and the quick blue hare."
	expected := make(map[string]int)
	expected["the quick"] = 2
	expected["quick brown"] = 1
	expected["brown fox"] = 1
	expected["fox and"] = 1
	expected["and the"] = 1
	expected["quick blue"] = 1
	expected["blue hare"] = 1

	test(t, example, expected)
}

func test(t *testing.T, input string, expected map[string]int) {
	t.Logf("Testing\n\t%s", input)
	bigrams := computeBigram(input)
	if len(expected) != len(bigrams) {
		t.Errorf(
			"Got different bigram map size: %d, expected %d",
			len(bigrams),
			len(expected),
		)
	}

	for key, value := range bigrams {
		if expected[key] != value {
			t.Errorf(
				"Got different bigram histogram count for '%s': %d, expected %d",
				key,
				value,
				expected[key],
			)
		}
	}
}
