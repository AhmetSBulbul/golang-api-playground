package user

import (
	"strings"

	"github.com/AhmetSBulbul/golang-api-playground/entity"
)

//inmemory in memory repo
type inmemory struct {
	m map[entity.ID]*entity.User
}

//newInmem create new in memory repository
func newInmem() *inmemory {
	var m = map[entity.ID]*entity.User{}
	return &inmemory{m: m}
}

//Create an user
func (r *inmemory) Create(e *entity.User) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get an user
func (r *inmemory) Get(id entity.ID) (*entity.User, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Update an user
func (r *inmemory) Update(e *entity.User) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Search users
func (r *inmemory) Search(query string) ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		if strings.Contains(strings.ToLower(j.FirstName), query) {
			d = append(d, j)
		}
	}
	if len(d) == 0 {
		return nil, entity.ErrNotFound
	}
	return d, nil
}

//List users
func (r *inmemory) List() ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Delete an user
func (r *inmemory) Delete(id entity.ID) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
