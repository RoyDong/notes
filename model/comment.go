package model

import (
    "time"
)

const (
    CommentStateDeleted = 1
)

type Comment struct {
    id int64

    Content string
    State int
    CreatedAt, UpdatedAt time.Time
}

func (c *Comment) Id() int64 {
    return c.id
}












var CommentModel = &CommentModel{"comment"}


type commentModel struct {
    table string
}


func (m *commentModel) Save(c *Comment) bool {

}