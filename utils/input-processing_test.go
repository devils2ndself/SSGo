package utils_test

import (
	"testing"

	utils "github.com/devils2ndself/SSGo/utils"

	. "github.com/stretchr/testify/assert"
)

// GetNameAndExt

func Test_GetNameAndExt_NameExtension(t *testing.T) {
	name, ext := utils.GetNameAndExt("test.txt")
	Equal(t, "test", name)
	Equal(t, ".txt", ext)
}

func Test_GetNameAndExt_MultipleDots(t *testing.T) {
	name, ext := utils.GetNameAndExt(".test.txt.zip")
	Equal(t, ".test.txt", name)
	Equal(t, ".zip", ext)
}

func Test_GetNameAndExt_NoExt(t *testing.T) {
	name, ext := utils.GetNameAndExt("test")
	Equal(t, "test", name)
	Equal(t, "", ext)
}

func Test_GetNameAndExt_NoName(t *testing.T) {
	name, ext := utils.GetNameAndExt(".test")
	Equal(t, "", name)
	Equal(t, ".test", ext)
}
