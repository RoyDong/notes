package model

import (
    "io"
    "fmt"
    "time"
    "strings"
    "crypto/rand"
    "crypto/md5"
    "crypto/sha512"
    "encoding/hex"
    "github.com/roydong/potato"
)

type UserForm struct {
    Id int64
    Name, Email, Passwd, Message string
}

func (f *UserForm) LoadData(r *potato.Request) {
    f.Email,_  = r.String("email")
    f.Passwd,_ = r.String("passwd")
    f.Name,_   = r.String("name")
}

func (f *UserForm) Valid() bool {
    f.Email = strings.Trim(f.Email, " ")
    if !strings.HasSuffix(f.Email, "@roydong.com") {
        f.Message = "email is not allowd"
        return false
    }

    f.Passwd = strings.Trim(f.Passwd, " ")
    if len(f.Passwd) < 6 {
        f.Message = "password must more than 6"
        return false
    }

    f.Name = strings.Trim(f.Name, " ")
    return true
}

type User struct {
    id int64
    passwd, salt string

    Name, Email string
    CreatedAt, UpdatedAt time.Time
}

func (u *User) Id() int64 {
    return u.id
}

/**
 * set a hash password
 */
func (u *User) SetPasswd(passwd string) {
    rnd := make([]byte, 32)
    if _,e := io.ReadFull(rand.Reader, rnd); e != nil {
        panic("could not generate random salt")
    }

    hash := md5.New()
    if _, e := hash.Write(rnd); e != nil {
        panic("could not hash salt")
    }

    u.salt = hex.EncodeToString(hash.Sum(nil))
    u.passwd = UserModel.HashPasswd(passwd, u.salt)
}

func (u *User) CheckPasswd(passwd string) bool {
    return UserModel.HashPasswd(passwd, u.salt) == u.passwd
}


type userModel struct {
    tabel string
}

var UserModel = &userModel{"user"}

func (m *userModel) User(id int64) *User {
    return nil
}

func (m *userModel) Find(id int64) *User {
    stmt := fmt.Sprintf("select `id`,`email`,`name`,`passwd`,`salt`,`created_at`,`updated_at` from %s where `id`='%d'", m.tabel, id)

    row := potato.D.QueryRow(stmt)
    return m.loadUser(row)
}

func (m *userModel) FindByEmail(email string) *User {
    stmt := fmt.Sprintf("select `id`,`email`,`name`,`passwd`,`salt`,`created_at`,`updated_at` from %s where `email`='%s'", m.tabel, email)

    row := potato.D.QueryRow(stmt)
    return m.loadUser(row)
}

func (m *userModel) loadUser(row Scanner) *User {
    u := new(User)
    var ct, ut int64
    if e := row.Scan(&u.id, &u.Email, &u.Name, &u.passwd, &u.salt, &ct, &ut); e != nil {
        return nil
    }

    u.CreatedAt = time.Unix(0, ct)
    u.UpdatedAt = time.Unix(0, ut)
    return u
}


func (m *userModel) Exists(email string) bool {
    stmt := fmt.Sprintf("select count(id) c from %s where `email`='%s'", m.tabel, email)
    row := potato.D.QueryRow(stmt)
    var count int
    if e := row.Scan(&count); e != nil {
        potato.L.Println(e)
    }

    return count > 0
}

func (m *userModel) Save(u *User) bool {
    if u.Id() > 0 {
        return false
    }

    return m.Add(u)
}

func (m *userModel) Add(u *User) bool {
    t := time.Now()
    u.CreatedAt = t
    u.UpdatedAt = t
    u.id = potato.D.Insert(fmt.Sprintf("INSERT INTO `%s`" +
            "(`email`,`name`,`passwd`,`salt`,`created_at`,`updated_at`)" +
            "VALUES(?,?,?,?,?,?)", m.tabel),
            u.Email, u.Name, u.passwd, u.salt, t.UnixNano(), t.UnixNano())

    return u.id > 0
}

func (m *userModel) HashPasswd(passwd string, salt string) string {
    hash := sha512.New()

    if _, e := hash.Write([]byte(passwd + salt)); e != nil {
        panic("could not hash password")
    }

    return hex.EncodeToString(hash.Sum(nil))
}
