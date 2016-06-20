//messages_test.go tests unmarshalling a response from the ZKB RedisQ service
package zkbq

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalKillmail(t *testing.T) {
	var evekill KillPackage

	kmb, err := ioutil.ReadFile("./test_data/sample_resp.json")
	assert.Nil(t, err)

	err = json.Unmarshal(kmb, &evekill)
	assert.Nil(t, err)

	assert.Equal(t, 54717661, evekill.Payload.KillID)
	assert.Len(t, evekill.Payload.Killmail.Attackers, 1)
	// XXX(tankbusta): More checks here
}
