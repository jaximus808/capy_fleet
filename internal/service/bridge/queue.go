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
	mut         sync.Mutex
	capacity    int
	actions     []Action
	cur_pointer uint
	end_pointer uint
	q_count     int
}

func CreateQueue(capacity int) *Queue {
	return &Queue{
		capacity:    capacity,
		actions:     make([]Action, capacity),
		cur_pointer: 0,
		end_pointer: 0,
	}
}

func CreateAction(author_id uint, msg *Packet) *Action {
	return &Action{
		author_id:      author_id,
		special_target: none,
		msg:            msg,
	}
}

func (q *Queue) inBounds() bool {

	return q.q_count < q.capacity
}

func (q *Queue) AddAction(action *Action) error {
	q.mut.Lock()
	defer q.mut.Unlock()
	if !q.inBounds() {

		return errors.New("queue is full")
	}
	q.actions[q.end_pointer] = *action
	q.incrEndPointer()
	return nil
}

func (q *Queue) ProcessAction() (Action, error) {
	q.mut.Lock()
	defer q.mut.Unlock()
	if q.Peak() {
		action := q.actions[q.cur_pointer]
		q.actions[q.cur_pointer] = Action{}
		q.incrCurPointer()
		return action, nil
	}
	return Action{}, errors.New("queue is empty")
}

func (q *Queue) incrCurPointer() {

	q.cur_pointer++
	if q.cur_pointer == uint(q.capacity) {
		q.cur_pointer = 0
	}
	q.q_count--
}
func (q *Queue) incrEndPointer() {

	q.end_pointer++
	if q.end_pointer == uint(q.capacity) {
		q.end_pointer = 0
	}
	q.q_count++
}

func (q *Queue) Peak() bool {
	return q.q_count > 0
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
