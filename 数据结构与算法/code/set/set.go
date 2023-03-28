package set

import "sync"

type Set struct {
	set  map[string]struct{}
	len  int
	lock sync.Mutex
}

func NewSet() *Set {
	return &Set{
		set:  make(map[string]struct{}),
		len:  0,
		lock: sync.Mutex{},
	}
}
func (s *Set) Add(data string) {
	s.lock.Lock()
	s.set[data] = struct{}{}
	s.len = len(s.set)
	s.lock.Unlock()
}
func (s *Set) Remove(data string) {
	s.lock.Lock()
	if s.len == 0 {
		return
	}
	delete(s.set, data)
	s.len = len(s.set)
	s.lock.Unlock()
}
func (s *Set) Has(data string) bool {
	s.lock.Lock()
	if s.len == 0 {
		return false
	}
	_, ok := s.set[data]
	s.lock.Unlock()
	return ok
}
func (s *Set) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len
}
func (s *Set) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len == 0
}
func (s *Set) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.set = make(map[string]struct{})
	s.len = 0
}
func (s *Set) List() []string {
	s.lock.Lock()
	defer s.lock.Unlock()
	ret := make([]string, 0, s.len)
	for data := range s.set {
		ret = append(ret, data)
	}
	return ret
}
