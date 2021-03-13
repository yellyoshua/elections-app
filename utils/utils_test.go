package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUniqueID(t *testing.T) {
	testSimple := func(t *testing.T) {
		filename := "Some.fil.e.jgp"
		fileExtension := ExtractFilenameExtension(filename)
		token, _ := GenerateUniqueID(fileExtension)
		tokenExtension := ExtractFilenameExtension(token)

		assert.Equal(t, "jgp", tokenExtension)
		assert.Equal(t, "jgp", fileExtension)
	}

	testWithManyCharacters := func(t *testing.T) {
		filename := "^&_asd.E(e).a.jgp"
		fileExtension := ExtractFilenameExtension(filename)
		token, _ := GenerateUniqueID(fileExtension)
		tokenExtension := ExtractFilenameExtension(token)

		assert.Equal(t, "jgp", tokenExtension)
		assert.Equal(t, "jgp", fileExtension)
	}

	testWithUppercase := func(t *testing.T) {
		filename := "^&_asd.E(e).a.JGP"
		fileExtension := ExtractFilenameExtension(filename)
		token, _ := GenerateUniqueID(fileExtension)
		tokenExtension := ExtractFilenameExtension(token)

		assert.Equal(t, "jgp", tokenExtension)
		assert.Equal(t, "jgp", fileExtension)
	}

	testWithDemoTxt := func(t *testing.T) {
		filename := "demo.txt"
		fileExtension := ExtractFilenameExtension(filename)
		token, _ := GenerateUniqueID(fileExtension)
		tokenExtension := ExtractFilenameExtension(token)

		assert.Equal(t, "txt", tokenExtension)
		assert.Equal(t, "txt", fileExtension)
	}

	testWithNoExtension := func(t *testing.T) {
		filename := "demo"
		fileExtension := ExtractFilenameExtension(filename)
		token, _ := GenerateUniqueID(fileExtension)
		tokenExtension := ExtractFilenameExtension(token)

		assert.Equal(t, "", tokenExtension)
		assert.Equal(t, "", fileExtension)
	}

	testSimple(t)
	testWithManyCharacters(t)
	testWithUppercase(t)
	testWithDemoTxt(t)
	testWithNoExtension(t)
}
