package sender

import (
	"time"
)

type Locker struct {
	ch chan struct{}
}

func NewLocker() *Locker {
	return &Locker{ch: make(chan struct{}, 1)}
}

func (l *Locker) Lock() {
	l.ch <- struct{}{}
}

func (l *Locker) Unlock() {
	<-l.ch
}

func (l *Locker) SyncDelay(d time.Duration) {
	l.Lock()
	select {
	case <-time.After(d):
		l.Unlock()
	}
}
func (l *Locker) AsyncDelay(d time.Duration) {
	l.Lock()
	go func() {
		select {
		case <-time.After(d):
			l.Unlock()
		}
	}()
}

func (l *Locker) Close() {
	close(l.ch)
}
