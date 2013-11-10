package model

import (
    "fmt"
    "time"
    "strings"
    "github.com/roydong/potato"
)


const (
    TopicStateEdit = 0
    TopicStatePublished = 1
)

type Topic struct {
    id int64

    Title, Content string
    State int
    CreatedAt, UpdatedAt time.Time
}

func (t *Topic) Id() int64 {
    return t.id
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

var TopicModel = &topicModel{"topic"}

type topicModel struct {
    table string
}

type Scanner interface{
    Scan(args ...interface{}) error
}

func (m *topicModel) Search(q map[string]string, page, limit int) []*Topic {
    sql := fmt.Sprintf("SELECT `id`,`title`,`content`,`state`,`created_at`,`updated_at` FROM `%s`", m.table)
    l := len(q)
    args := make([]interface{}, 0, l + 2)
    if l > 0 {
        c := make([]string, 0, l)
        for k, v := range q {
            c = append(c, fmt.Sprintf("`%s` REGEXP ? ", k))
            args = append(args, v)
        }

        sql = sql + " WHERE " + strings.Join(c, " AND ")
    }

    if page < 1 { page = 1}
    args = append(args, (page - 1) * limit, limit)

    rows, e := potato.D.Query(sql + " LIMIT ?,?", args...)
    if e != nil {
        potato.L.Println(e)
        return nil
    }

    topics := make([]*Topic, 0, limit)
    for rows.Next() {
        if t := m.loadTopic(rows); t != nil {
            topics = append(topics, t)
        }
    }

    return topics
}

func (m *topicModel) loadTopic(row Scanner) *Topic {
    t := new(Topic)
    var ct, ut int64
    if e := row.Scan(&t.id, &t.Title, &t.Content , &t.State, &ct, &ut); e != nil {
        potato.L.Println(e)
        return nil
    }

    t.CreatedAt = time.Unix(0, ct)
    t.UpdatedAt = time.Unix(0, ut)
    return t
}

func (m *topicModel) Find(id int) *Topic {
    sql := fmt.Sprintf("select `id`,`title`,`content`,`state`,`created_at`,`updated_at` from %s where `id`='%d'", m.table, id)

    return m.loadTopic(potato.D.QueryRow(sql))
}

func (m *topicModel) Save(t *Topic) bool {
    if t.Id() > 0 {
        return m.Update(t)
    }

    return m.Add(t)
}

func (m *topicModel) Update(t *Topic) bool {
    now := time.Now()
    t.UpdatedAt = now
    _,e := potato.D.Exec(fmt.Sprintf("UPDATE `%s` SET" +
            " `title`=?,`content`=?,`state`=?,`created_at`=?,`updated_at`=?" +
            " WHERE `id`=?", m.table),
            t.Title, t.Content, t.State, t.CreatedAt.UnixNano(), now.UnixNano(), t.id)
    if e != nil {
        potato.L.Println(e)
        return false
    }

    return true
}

func (m *topicModel) Add(t *Topic) bool {
    now := time.Now()
    t.CreatedAt = now
    t.UpdatedAt = now
    t.id = potato.D.Insert(fmt.Sprintf("INSERT INTO `%s`" +
            "(`title`,`content`,`state`,`created_at`,`updated_at`)" +
            "VALUES(?,?,?,?,?)", m.table),
            t.Title, t.Content, t.State, now.UnixNano(), now.UnixNano())

    return t.id > 0
}
