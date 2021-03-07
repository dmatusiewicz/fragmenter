package fragmenter

import (
	"bytes"
	_ "embed"
	"testing"
)

//go:embed test.yaml
var haveData []byte
var pattern = []byte("---\n")
var wantData = [][]byte{
	[]byte("---\nsection1\n"),
	[]byte("---\nsection2\nbigger\n"),
	[]byte("---\nsection3\neven bigger\neven bigger\n"),
	[]byte("---\nsection4\nsmaller\n"),
}

func TestFragment(t *testing.T) {

	gotData := Fragment(haveData, pattern)

	gotDataLen := len(gotData)
	wantDataLen := len(wantData)

	if wantDataLen == gotDataLen {
		for i, data := range wantData {
			if bytes.Compare(data, gotData[i]) != 0 {
				t.Errorf("Data in index: %d dosen't match.\nGot:\n%s\nShould have:\n%s", i, gotData[i], data)
			}
		}
	} else if wantDataLen > gotDataLen {
		t.Errorf("wantDataLen(%d) is longer than gotDataLen(%d). Possible data loss!", wantDataLen, gotDataLen)
	} else {
		t.Errorf("wantDataLen(%d) is shorter than gotDataLen(%d). Possible data addition!", wantDataLen, gotDataLen)
	}
}
