package model

import (
    "fmt"
    "time"
    "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
)

const (
    CommentStateDeleted = 1
)

type Comment struct {
    Id int64 `column:"id"`

    Tid int64 `column:"tid"`
    topic *Topic

    Uid int64 `column:"uid"`
    user *User

    Content string `column:"content"`
    State int `column:"state"`
    CreatedAt time.Time `column:"created_at"`
    UpdatedAt time.Time `column:"updated_at"`
}

func (c *Comment) Topic() *Topic {
    if c.topic == nil {
        c.topic = TopicModel.FindById(c.Tid)
    }

    return c.topic
}

func (c *Comment) SetTopic(t *Topic) {
    c.Tid = t.Id
    c.topic = t
}

func (c *Comment) User() *User {
    if c.user == nil {
        c.user = UserModel.Find(c.Uid)
    }

    return c.user
}

func (c *Comment) SetUser(u *User) {
    c.Uid = u.Id
    c.user = u
}



type CommentForm struct {
    Content string
    State int
}

func (f *CommentForm) LoadData(r *potato.Request) {
    f.Content,_ = r.String("content")
    f.State,_ = r.Int("state")
}


var CommentModel = &commentModel{orm.NewModel("comment", new(Comment))}


type commentModel struct {
    *orm.Model
}

func (m *commentModel) FindBy(k string, v interface{}) []*Comment {
    stmt := orm.NewStmt().Select("c.*").From("Comment", "c").
            Where(fmt.Sprintf("`c`.`%s` = ?", k)).Asc("id")

    rows, e := stmt.Query(v)
    if e != nil {
        return nil
    }

    comments := make([]*Comment, 0)
    for rows.Next() {
        var c *Comment
        rows.ScanEntity(&c)
        comments = append(comments, c)
    }

    return comments
}
