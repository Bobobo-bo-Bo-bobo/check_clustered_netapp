package shared

import (
	"io/ioutil"
)

// ReadFileContent - read content from a file
// Parameters:
//  f - file to read
// Result:
//  []bytes - data
//  error   - error
func ReadFileContent(f string) ([]byte, error) {
	content, err := ioutil.ReadFile(f)
	return content, err
}
