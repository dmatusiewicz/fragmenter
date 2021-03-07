package fragmenter

import (
	"bytes"
	"errors"
	"log"
)

var (
	errNoMatchFound = errors.New("Cound not find a pattern in the data")
)

// Fragment fragments an array of bytes into array of bytes arrays. The expected result is to have data divided into sections. It is crawling throughout the entire data slice.
func Fragment(data, pattern []byte) [][]byte {
	var fragmentedData = [][]byte{}
	for i := 0; i < len(data); {
		x, y, err := locateFirstFragment(data[i:], pattern)
		if err != nil {
			log.Fatal(err)
		}
		if x != 0 {
			y += x
		}
		fragmentedData = append(fragmentedData, data[x+i:y+i])
		i += y
	}
	return fragmentedData
}

// locateFirstFragment helper function for locating the fragment that is matching the pattern in the data slice. Returns an error when cannot find a single pattern. The crowler logic is not a concern of this fucntion.
func locateFirstFragment(data []byte, pattern []byte) (int, int, error) {
	if bytes.Contains(pattern, []byte("\n")) {
		pattern = bytes.Replace(pattern, []byte("\n"), nil, 1)
	}
	start := bytes.Index(data, pattern)
	end := bytes.Index(data[start+len(pattern):], pattern)
	if start == -1 {
		return -1, -1, errNoMatchFound
	}
	if end == -1 {
		end = len(data)
		return start, end, nil
	}
	return start, end + len(pattern), nil
}
