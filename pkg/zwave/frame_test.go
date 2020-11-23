package zwave

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrame(t *testing.T) {
	sof := SOF(REQ, DiscoveryNodes)
	assert.Equal(t, sof.ToBytes(), []byte{0x01, 0x03, 0x00, 0x02, 0xfe})
	l, msg, err := NewFrame(sof.ToBytes())
	log.Printf("l: %d (len: %d) msg: %+v, err:%s", l, len(sof.ToBytes()), msg, err)
}
