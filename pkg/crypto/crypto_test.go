package crypto_test

import (
	"api-boilerplate/pkg/crypto"
	"testing"
)

func Test_Crypto(t *testing.T) {
	t.Run("Can hash and compare", should_be_able_to_hash_and_compare_strings)
	t.Run("Can detect unequal hashes", should_return_error_when_comparing_unequal_hashes)
	t.Run("Generates a different salt every time", should_generate_a_different_salt_each_time)
}

func should_be_able_to_hash_and_compare_strings(t *testing.T) {
	// Arrange
	c := crypto.Crypto{}
	testInput := "testInput"

	// Act
	generatedHash, generateError := c.Generate(testInput)
	compareError := c.Compare(generatedHash, testInput)

	// Assert
	if generateError != nil {
		t.Error("error generating hash")
	}

	if testInput == generatedHash {
		t.Error("generated hash is the same as input")
	}

	if compareError != nil {
		t.Error("error comparing hash to input")
	}
}

func should_return_error_when_comparing_unequal_hashes(t *testing.T) {
	// Arrange
	c := crypto.Crypto{}
	testInput := "testInput"
	testCompare := "testComapre"

	// Act
	generatedHash, generateError := c.Generate(testInput)
	compareError := c.Compare(generatedHash, testCompare)

	// Assert
	if generateError != nil {
		t.Error("error generating hash")
	}

	if testInput == generatedHash {
		t.Error("generated hash is the same as input")
	}

	if compareError == nil {
		t.Error("compare should not have been successful")
	}
}

func should_generate_a_different_salt_each_time(t *testing.T) {
	// Arrange
	c := crypto.Crypto{}
	testInput := "testInput"

	hash1, _ := c.Generate(testInput)
	hash2, _ := c.Generate(testInput)

	if hash1 == hash2 {
		t.Error("subsequent hashes should not be equal")
	}
}
