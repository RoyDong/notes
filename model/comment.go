package model

import (
    "fmt"
    "time"
    "github.com/roydong/potato"
)

const (
    CommentStateDeleted = 1
)

type Comment struct {
    id int64

    tid int64
    topic *Topic

    uid int64
    user *User

    Content string
    State int
    CreatedAt, UpdatedAt time.Time
}

func (c *Comment) Id() int64 {
    return c.id
}

func (c *Comment) Tid() int64 {
    return c.tid
}

func (c *Comment) Topic() *Topic {
    if c.topic == nil {
        c.topic = TopicModel.FindById(c.tid)
    }

    return c.topic
}

func (c *Comment) SetTopic(t *Topic) {
    c.tid = t.id
    c.topic = t
}

func (c *Comment) Uid() int64 {
    return c.uid
}

func (c *Comment) User() *User {
    if c.user == nil {
        c.user = UserModel.Find(c.uid)
    }

    return c.user
}

func (c *Comment) SetUser(u *User) {
    c.uid = u.id
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


var CommentModel = &commentModel{"comment"}


type commentModel struct {
    table string
}

func (m *commentModel) FindBy(key string, v interface{}, order string, limit ...int) []*Comment {
    sql := fmt.Sprintf(
            "SELECT `id`,`tid`,`uid`,`content`,`state`,`created_at`,`updated_at` " +
            "FROM `%s` WHERE `%s`=? ORDER BY %s", m.table, key, order)

    if len(limit) == 1 {
        sql = fmt.Sprintf("%s LIMIT %d", sql, limit[0])
    } else if len(limit) == 2 {
        sql = fmt.Sprintf("%s LIMIT %d, %d", sql, limit[0], limit[1])
    }

    rows, e := potato.D.Query(sql, v)
    if e != nil {
        potato.L.Println(e)
        return nil
    }

    comments := make([]*Comment, 0)
    for rows.Next() {
        if c := m.loadComment(rows); c != nil {
            comments = append(comments, c)
        }
    }

    return comments
}

func (m *commentModel) loadComment(row Scanner) *Comment {
    c := new(Comment)
    var ct, ut int64
    if e := row.Scan(&c.id, &c.tid, &c.uid, &c.Content , &c.State, &ct, &ut); e != nil {
        potato.L.Println(e)
        return nil
    }

    c.CreatedAt = time.Unix(0, ct)
    c.UpdatedAt = time.Unix(0, ut)
    return c
}

func (m *commentModel) Save(c *Comment) bool {
    if c.Id() > 0 {
        return m.Update(c)
    }

    return m.Add(c)
}

func (m *commentModel) Update(c *Comment) bool {
    if c.Topic() == nil {
        return false
    }

    now := time.Now()
    c.UpdatedAt = now
    _,e := potato.D.Exec(fmt.Sprintf("UPDATE `%s` SET" +
            " `tid`=?, `uid`=?,`content`=?,`state`=?,`updated_at`=?" +
            " WHERE `id`=?", m.table),
            c.tid, c.uid , c.Content, c.State, now.UnixNano(), c.id)

    if e != nil {
        potato.L.Println(e)
        return false
    }

    return true
}

func (m *commentModel) Add(c *Comment) bool {
    if c.Topic() == nil {
        return false
    }

    now := time.Now()
    c.CreatedAt = now
    c.UpdatedAt = now
    c.id = potato.D.Insert(fmt.Sprintf("INSERT INTO `%s`" +
            "(`tid`,`uid`,`content`,`state`,`created_at`,`updated_at`)" +
            "VALUES(?,?,?,?,?,?)", m.table),
            c.tid, c.uid, c.Content, c.State, now.UnixNano(), now.UnixNano())

    return c.id > 0
}
