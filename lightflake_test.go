package lightflake

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestFlake(t *testing.T) {
	token := Generate(25)
	t.Log(token)
	timestamp, workerid := ParseFlake(token)
	t.Log(EPOCH.Unix())
	t.Log(timestamp)
	t.Log(workerid)
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, token)
	t.Logf("% x", buf.Bytes())
}
