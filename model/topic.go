package model

import (
    "fmt"
    "time"
    "github.com/roydong/potato"
)



type Topic struct {
    id int64

    Title, Content string
    CreatedAt, UpdatedAt time.Time
}

func (t *Topic) Id() int64 {
    return t.id
}


type TopicForm struct {
    Title, Content, Message string
}

func (f *TopicForm) LoadData(r *potato.Request) {
    f.Title,_ = r.String("title")
    f.Content,_ = r.String("content")
}

var TopicModel = &topicModel{"topic"}

type topicModel struct {
    table string
}

func (m *topicModel) Page(k, v string, page, size int) []*Topic {

}

func (m *topicModel) Find(id int) *Topic {
    stmt := fmt.Sprintf("select `id`,`title`,`content`,`created_at`,`updated_at` from %s where `id`='%d'", m.table, id)

    row := potato.D.QueryRow(stmt)
    t := new(Topic)
    var ct, ut int64
    if e := row.Scan(&t.id, &t.Title, &t.Content , &ct, &ut); e != nil {
        potato.L.Println(e)
        return nil
    }

    t.CreatedAt = time.Unix(0, ct)
    t.UpdatedAt = time.Unix(0, ut)
    return t
}

func (m *topicModel) Save(t *Topic) bool {
    if t.Id() > 0 {

        return false
    }

    return m.Add(t)
}

func (m *topicModel) Add(t *Topic) bool {
    now := time.Now()
    t.CreatedAt = now
    t.UpdatedAt = now
    t.id = potato.D.Insert(fmt.Sprintf("INSERT INTO `%s`" +
            "(`title`,`content`,`created_at`,`updated_at`)" +
            "VALUES(?,?,?,?)", m.table),
            t.Title, t.Content, now.UnixNano(), now.UnixNano())

    return t.id > 0
}
