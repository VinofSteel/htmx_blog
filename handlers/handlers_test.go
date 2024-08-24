package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sanitizeSlug(t *testing.T) {
	type sanitizeSlugTests struct {
		have string
		want string
		name string
	}

	testCases := []sanitizeSlugTests{
		{"asdj#!!@$@#%^^&&^&*1235%%%%kdkdkas", "asdj1235kdkdkas", "Many symbols case"},
		{"asd%%% 2", "asd-2", "Breaking symbols case"},
		{"Valid Slug With Spaces", "valid-slug-with-spaces", "Standard slug case"},
		{"invalid%20slug%20with%20spaces", "invalid-slug-with-spaces", "Slug with spaces"},
		{"!!!@@@###", "", "Only symbols case"},
		{"", "", "Empty string"},
	}

	for _, testCase := range testCases {
		t.Logf("Running sanitizeSlug %s\n", testCase.name)
		slug := sanitizeSlug(testCase.have)

		assert.Equal(t, testCase.want, slug, "Expected slug mismatch")
	}
}
