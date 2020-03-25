package main

import (
	"github.com/tazzaoui/risearch/lib"
	"testing"
)

// Sanity check to make sure exact match is top result
func TestExactMatch(t *testing.T) {
	to_query := "../data/img/COCO_val2014_000000000073.jpg"

	top_match := lib.GetMatches(to_query)[0]

	if top_match.Img != to_query {
		t.Errorf("Top match not exact! Wanted: %s Got: %s", to_query, top_match.Img)
	}
}
