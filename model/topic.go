package model

import (
    "log"
    "fmt"
    "time"
    "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
)


const (
    TopicStateDraft = 0
    TopicStatePublished = 1
    TopicStateDeleted = 2
)

type Topic struct {
    Id int64 `column:"id"`
    Title string `column:"title"`
    Content string `column:"content"`
    State int `column:"state"`
    CreatedAt time.Time `column:"created_at"`
    UpdatedAt time.Time `column:"updated_at"`
}

func (t *Topic) Comments() []*Comment {
    return CommentModel.FindBy("tid", t.Id)
}

type TopicForm struct {
    Title, Content, Message string
    State int
}

func (f *TopicForm) LoadData(r *potato.Request) {
    f.Title,_ = r.String("title")
    f.Content,_ = r.String("content")
    f.State,_ = r.Int("state")
}

type topicModel struct {
    *orm.Model
}

var TopicModel = &topicModel{orm.NewModel("topic", new(Topic))}

func (m *topicModel) Search(k, v string) []*Topic {
    stmt := orm.NewStmt().Select("t.*").From("Topic", "t").
        Desc("id").Where(fmt.Sprintf("`t`.`%s` LIKE ?", k))

    rows, e := stmt.Query("%" + v + "%")
    if e != nil {
        log.Println(e)
        return nil
    }

    topics := make([]*Topic, 0)
    for rows.Next() {
        var t *Topic
        rows.ScanEntity(&t)
        topics = append(topics, t)
    }

    return topics
}

func (m *topicModel) FindById(id int64) *Topic {
    var t *Topic
    rows, e := orm.NewStmt().Select("t.*").From("Topic", "t").
            Where("t.id = ?").Query(id)

    if e == nil && rows.Next() {
        rows.ScanEntity(&t)
    }

    return t
}
