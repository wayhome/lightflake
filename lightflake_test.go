package lightflake

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestFlake(t *testing.T) {
	token, err := Generate(0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
	timestamp, workerid := ParseFlake(token)
	t.Log(EPOCH.Unix())
	t.Log(timestamp)
	t.Log(workerid)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, token)
	t.Logf("% x", buf.Bytes())
}
