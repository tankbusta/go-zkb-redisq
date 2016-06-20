//zkb_redisq.go handles POLLIN' for killmails
package zkbq

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// KillChan is the channel on which a kill will be sent
type KillChan chan *KillPackage

// ZKBPoller is responsible for constantly polling RedisQ
type ZKBPoller struct {
	Kill  KillChan
	Error chan error
	// Unexported fields below
	client *http.Client
	wg     *sync.WaitGroup
	stop   bool
}

// NewZKBPoller creates a Zkillboard Poller and sends killmails over ch
func NewZKBPoller() (zpoll *ZKBPoller, err error) {
	// Safety check to prevent NPE's
	if zpoll == nil {
		zpoll = &ZKBPoller{}
	}

	// ZKB RedisQ has a timeout of 10 seconds before it returns payload: null
	zpoll.client = &http.Client{Timeout: time.Second * 11}
	zpoll.wg = &sync.WaitGroup{}
	zpoll.Kill = make(KillChan)
	zpoll.Error = make(chan error)

	zpoll.wg.Add(1)
	go zpoll.work()
	return
}

func (s *ZKBPoller) getPayload() (*KillPackage, error) {
	var (
		kpkg *KillPackage
		dec  *json.Decoder
	)

	req, err := http.NewRequest("GET", ZkillboardServer, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", GoZKBStringVer)
	req.Header.Add("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&kpkg); err != nil {
		return nil, err
	}
	return kpkg, nil
}

func (s *ZKBPoller) work() {
	defer s.wg.Done()
	for s.stop == false {
		kp, err := s.getPayload()
		if err != nil {
			s.Error <- err
			continue
		}
		s.Kill <- kp
	}
}

// Wait for the ZKBPoller to stop
func (s *ZKBPoller) Wait() {
	s.wg.Wait()
}

// Stop the ZKBPoller from running
func (s *ZKBPoller) Stop() {
	s.stop = true
}
