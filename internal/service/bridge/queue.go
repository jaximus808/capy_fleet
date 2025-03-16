package bridge

import (
	"errors"
	"sync"
)

const (
	none uint = iota
	all
	allexcept
)

type Action struct {
	client_id      []uint
	special_target uint
	author_id      uint
	msg            *Packet
}

type Queue struct {
	mut      sync.Mutex
	capacity int
	actions  []Action
}

func CreateQueue(capacity int) *Queue {
	return &Queue{
		capacity: capacity,
		actions:  make([]Action, capacity),
	}
}

func CreateAction(author_id uint, msg *Packet) *Action {
	return &Action{
		author_id:      author_id,
		special_target: none,
		msg:            msg,
	}
}

func (q *Queue) AddAction(action *Action) error {
	q.mut.Lock()
	defer q.mut.Unlock()
	if len(q.actions) == q.capacity {

		return errors.New("queue is full")
	}
	q.actions = append(q.actions, *action)
	return nil
}

func (q *Queue) ProcessAction() (Action, error) {
	q.mut.Lock()
	defer q.mut.Unlock()
	if len(q.actions) > 0 {
		action := q.actions[0]
		q.actions = q.actions[1:]
		return action, nil
	}
	return Action{}, errors.New("queue is empty")
}

func (q *Queue) Peak() bool {
	return len(q.actions) > 0
}

func (a *Action) GetPacket() *Packet {
	return a.msg
}
func (a *Action) GetSpecialTarget() uint {
	return a.special_target
}
func (a *Action) GetTargets() []uint {
	return a.client_id
}
func (a *Action) GetAuthor() uint {
	return a.author_id
}
func (a *Action) AddTarget(uid uint) {
	a.client_id = append(a.client_id, uid)
}

func (a *Action) AddTargets(uids []uint) {
	a.client_id = append(a.client_id, uids...)
}

func (a *Action) SetSpecial(special_target uint) {
	if special_target > 2 {
		return
	}
	a.special_target = special_target
}
