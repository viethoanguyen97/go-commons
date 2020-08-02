package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(arr []byte) (int, error) {
	for i := range arr {
		arr[i] = 65
	}

	return len(arr), nil
}

func main() {
	reader.Validate(MyReader{})

}
