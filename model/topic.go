package model

import (
    "time"
    "github.com/roydong/potato"
)



type Topic struct {
    id int

    Title, Content string
    CreatedAt, UpdatedAt time.Time
}

func (t *Topic) Id() int {
    return t.id
}




var TopicModel = &topicModel{"topic"}

type topicModel struct {
    table string
}

func (m *topicModel) Find(id int) *Topic {
    stmt := fmt.Sprintf("select `id`,`title`,`content`,`created_at`,`updated_at` from %s where `id`='%d'", m.tabel, id)

    row := potato.D.QueryRow(stmt)
    t := new(Topic)
    var ct, ut int64
    if e := row.Scan(&t.id, &t.Title, t.Content , &ct, &ut); e != nil {
        return nil
    }

    t.CreatedAt = time.Unix(0, ct)
    t.UpdatedAt = time.Unix(0, ut)
    return t
}

func (m *topicModel) Save(t *Topic) {
    if t.Id() > 0 {

        return false
    }

    return m.Add(t)
}

func (m *topicModel) Add(t *Topic) bool {
    now := time.Now()
    t.CreatedAt = now
    t.UpdatedAt = now
    u.id = potato.D.Insert(fmt.Sprintf("INSERT INTO `%s`" +
            "(`title`,`content`,`created_at`,`updated_at`)" +
            "VALUES(?,?,?,?)", m.tabel),
            t.Title, t.Content, now.UnixNano(), now.UnixNano())

    return u.id > 0
}
