package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yunify/qsctl/v2/constants"
)

func TestParseFlow(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected constants.FlowType
	}{
		{"xxxx", "qs://xxxx", constants.FlowToRemote},
		{"qs://xxxx", "xxxx", constants.FlowToLocal},
		{"xxxx", "xxxx", constants.FlowInvalid},
		{"qs://xxxx", "qs://xxxx", constants.FlowInvalid},
		{"xxxx", "", constants.FlowAtLocal},
		{"qs://xxxx", "", constants.FlowAtRemote},
	}

	for _, v := range cases {
		x := ParseFlow(v.input1, v.input2)
		assert.Equal(t, v.expected, x)
	}
}

func TestParseKey(t *testing.T) {
	cases := []struct {
		input              string
		expectedKeyType    constants.KeyType
		expectedBucketName string
		expectedKey        string
	}{
		{"qs://xxxx-bucket/abc", constants.KeyTypeObject, "xxxx-bucket", "abc"},
		{"qs://abcdef", constants.KeyTypeBucket, "abcdef", ""},
		{"qs://abcdef/", constants.KeyTypeBucket, "abcdef", ""},
		{"qs://abcdef/def/ghi", constants.KeyTypeObject, "abcdef", "def/ghi"},
		{"qs://abcdef/def/ghi/", constants.KeyTypePseudoDir, "abcdef", "def/ghi/"},
		{"abcdef", constants.KeyTypeBucket, "abcdef", ""},
		{"abcdef/", constants.KeyTypeBucket, "abcdef", ""},
		{"abcdef/def/ghi", constants.KeyTypeObject, "abcdef", "def/ghi"},
		{"abcdef/👾 🙇 💁 🙅 🙆 🙋 🙎 🙍", constants.KeyTypeObject, "abcdef", "👾 🙇 💁 🙅 🙆 🙋 🙎 🙍"},
	}

	for k, v := range cases {
		actualKeyType, actualBucketName, actualKey, err := ParseKey(v.input)
		assert.Equal(t, v.expectedKeyType, actualKeyType, k)
		assert.Equal(t, v.expectedBucketName, actualBucketName, k)
		assert.Equal(t, v.expectedKey, actualKey, k)
		assert.NoError(t, err, k)
	}
}

func TestIsValidBucketName(t *testing.T) {
	cases := []struct {
		name          string
		input         string
		expectedValid bool
	}{
		{"start with letter", "a-bucket-test", true},
		{"start with digit", "0-bucket-test", true},
		{"start with strike", "-bucket-test", false},
		{"end with strike", "bucket-test-", false},
		{"too short", "abcd", false},
		{"too long (64)", "abcdefghijklmnopqrstuvwxyz123456abcdefghijklmnopqrstuvwxyz123456", false},
		{"contains illegal char", "abcdefg_1234", false},
	}

	for _, v := range cases {
		valid := IsValidBucketName(v.input)
		assert.Equal(t, valid, v.expectedValid)
	}
}