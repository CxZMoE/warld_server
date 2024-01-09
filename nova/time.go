package nova

import (
	"fmt"
	"time"
)

var (
	timeRatio = 1 // 時間の流れスピード
)

type Time struct {
	timestamp uint64 // in Nano
}

func CreateTime() *Time {
	return &Time{
		timestamp: 0,
	}
}

func (t *Time) Start() {
	start := time.Now()
	go func() {
		for {
			t.timestamp = uint64(time.Since(start).Nanoseconds())
		}
	}()
}

func (t *Time) Tick() {
	t.timestamp += 1
}

func (t *Time) String() string {
	return fmt.Sprintf("[Time] %d", t.timestamp)

}

func (t *Time) ToUnixNano() {
	t.timestamp = uint64(time.Now().UnixNano())
}

func (t *Time) GetRealWorldTime() time.Time {
	return time.Now()
}
