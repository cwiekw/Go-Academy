package movie

import (
	"MovieManager/utils"
	"fmt"
	"strings"
)

type MovieDataBase struct {
	db map[uint64]Movie
}

func NewMovieDataBase() MovieDataBase {
	return MovieDataBase{
		db: make(map[uint64]Movie),
	}
}

func (mdb MovieDataBase) String() string {
	var sb strings.Builder

	for _, m := range mdb.db {
		sb.WriteString(fmt.Sprintf("%s\n", m))
	}

	return sb.String()
}

func (mdb MovieDataBase) GetAll() []Movie {
	result := make([]Movie, len(mdb.db))

	idx := 0
	for _, m := range mdb.db {
		result[idx] = m
		idx++
	}

	return result
}

func (mdb MovieDataBase) GetById(id uint64) (Movie, error) {
	m, exists := mdb.db[id]

	if !exists {
		return m, utils.NewEntityDoesNotExistError("Movie", id)
	}

	return m, nil
}

func (mdb MovieDataBase) Add(m Movie) uint64 {
	id := utils.GenerateId()
	m.Id = id
	mdb.db[id] = m
	return id
}

func (mdb MovieDataBase) Update(id uint64, u Movie) (bool, error) {
	m, err := mdb.GetById(id)

	if err != nil {
		return false, err
	}

	m.Name = u.Name
	m.Year = u.Year

	mdb.db[id] = m

	return true, nil
}

func (mdb MovieDataBase) Delete(id uint64) (bool, error) {
	_, err := mdb.GetById(id)

	if err != nil {
		return false, err
	}

	delete(mdb.db, id)

	return true, nil
}
