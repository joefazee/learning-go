package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckUsername(t *testing.T) {
	assert.True(t, CheckUsername("test123"))
}

// TestCheckUsernameTable table-driven test
func TestCheckUsernameTable(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"too short", "test1", false},
		{"empty", "", false},
		{"contains admin", "adminuser", false},
		{"valid", "greatusername", true},
	}

	for _, tc := range testCases {
		got := CheckUsername(tc.input)
		assert.Equal(t, tc.want, got, "got %v\nwant %v", got, tc.want)

	}
}

// TestCheckUsernameTable table-driven test with sub-test
func TestCheckUsernameTableWithSubTest(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{"too short", "test1", false},
		{"empty", "", false},
		{"contains admin", "adminuser", false},
		{"valid", "greatusername", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := CheckUsername(tc.input)
			assert.Equal(t, tc.want, got, "got %v\nwant %v", got, tc.want)
		})
	}
}

func TestLoginSuccess(t *testing.T) {
	err, ok := Login("testusername")
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestLoginFailure(t *testing.T) {
	err, ok := Login("name")
	assert.Error(t, err)
	assert.False(t, ok)
}
