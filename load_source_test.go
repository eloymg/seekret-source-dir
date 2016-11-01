package sourcedir

import (
	"github.com/apuigsech/seekret"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadObjectsFromEmptyDirectoryReturnsNoContent(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "empty_directory")

	sourceDir := SourceDir{}
	loadOptions := seekret.LoadOptions{}

	objectList, err := sourceDir.LoadObjects(tempDir, loadOptions)

	assert.Nil(t, err, "Error should be nil")
	assert.Empty(t, objectList, "ObjectList should be empty")

	os.Remove(tempDir)
}

func TestLoadObjectsFromDirectoryWithHiddenFilesReturnsNoContent(t *testing.T) {
	sourceDir := SourceDir{}
	loadOptions := seekret.LoadOptions{}

	objectList, err := sourceDir.LoadObjects("./fixtures/dir_with_hidden_files", loadOptions)

	assert.Nil(t, err, "Error should be nil")
	assert.Empty(t, objectList, "ObjectList should be empty")
}

func TestLoadObjectsFromDirectoryWithHiddenFilesReturnsContentIfHiddenOptionIsSet(t *testing.T) {
	sourceDir := SourceDir{}
	loadOptions := seekret.LoadOptions{"hidden": true}

	objectList, err := sourceDir.LoadObjects("./fixtures/dir_with_hidden_files", loadOptions)

	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, objectList, "ObjectList should not be empty")
	assert.Len(t, objectList, 1, "Only one object should be present")
}

func TestLoadObjectsFromDirectoryWithTwoFilesReturnsTwoObjects(t *testing.T) {
	sourceDir := SourceDir{}
	loadOptions := seekret.LoadOptions{}

	objectList, err := sourceDir.LoadObjects("./fixtures/dir_with_two_files", loadOptions)

	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, objectList, "ObjectList should not be empty")
	assert.Len(t, objectList, 2, "Two objects should be present")
}

func TestLoadObjectsFromTwoLevelDirectoryThatContainsTwoFilesReturnsTwoObjectsIfRecursive(t *testing.T) {
	sourceDir := SourceDir{}
	loadOptions := seekret.LoadOptions{"recursive": true}

	objectList, err := sourceDir.LoadObjects("./fixtures/recursive_dir_with_two_files", loadOptions)

	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, objectList, "ObjectList should not be empty")
	assert.Len(t, objectList, 2, "Two objects should be present")
}

func TestLoadObjectsFromInvalidDirectoryShouldReturnAnError(t *testing.T) {
	sourceDir := SourceDir{}
	loadOptions := seekret.LoadOptions{}

	objectList, err := sourceDir.LoadObjects("./fixtures/invalid", loadOptions)

	assert.NotNil(t, err, "Error should be nil")
	assert.Empty(t, objectList, "ObjectList should be empty")
}