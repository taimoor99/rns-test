package services

import (
	"bookexchange/app/models"
	"bookexchange/app/repository"
	"github.com/gomodule/redigo/redis"
)

func NewRedisConnection(uri string) redis.Conn {
	c, err := redis.DialURL(uri)
	if err != nil {
		panic(err)
	}
	return c
}

type DbSession struct {
	db redis.Conn
}

var WinAmountC = make(chan models.BookOrder)

func NewDbClient(conn redis.Conn) repository.BooksOrderDB {
	return &DbSession{db: conn}
}

func (d DbSession) AddBook(order models.BookOrder) error {
	panic("implement me")
}

func (DbSession) RemoveBook() error {
	panic("implement me")
}

func (d DbSession) Insert(t *models.Tree, v models.BookOrder) *models.Tree {
	if t == nil {
		return &models.Tree{nil, v, nil}
	}

	if t.Value.Side != v.Side  && t.Value.Price == v.Price{
		delete(t)
		return t
	}

	if v.Price < t.Value.Price {
		t.Left = d.Insert(t.Left, v)
		return t
	}
	t.Right = d.Insert(t.Right, v)
	return t
}

func delete(t *models.Tree) *models.Tree {
	t = nil
	return t
}