package model

import (
    "io"
    "time"
    "strings"
    "crypto/rand"
    "crypto/md5"
    "crypto/sha512"
    "encoding/hex"
    "github.com/roydong/potato"
    "github.com/roydong/potato/orm"
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
    Id int64 `column:"id"`
    Passwd string `column:"passwd"`
    Salt string `column:"salt"`
    Name string `column:"name"`
    Email string `column:"email"`
    CreatedAt time.Time `column:"created_at"`
    UpdatedAt time.Time `column:"updated_at"`
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

    u.Salt = hex.EncodeToString(hash.Sum(nil))
    u.Passwd = UserModel.HashPasswd(passwd, u.Salt)
}

func (u *User) CheckPasswd(passwd string) bool {
    return UserModel.HashPasswd(passwd, u.Salt) == u.Passwd
}


type userModel struct {
    *orm.Model
}

var UserModel = &userModel{orm.NewModel("user", new(User))}

func (m *userModel) Find(id int64) *User {
    var u *User
    rows, e := orm.NewStmt().Select("u.*").From("User", "u").
            Where("u.id = ?").Query(id)

    if e == nil {
        rows.ScanRow(&u)
    }

    return u
}

func (m *userModel) FindByEmail(email string) *User {
    var u *User
    rows, e := orm.NewStmt().Select("u.*").
            From("User", "u").Where("u.email = ?").Query(email)

    if e == nil {
        rows.ScanRow(&u)
    }

    return u
}


func (m *userModel) Exists(email string) bool {
    n,_ := orm.NewStmt().Count("User", "u").
            Where("u.email = ?").Exec(email)

    return n > 0
}


func (m *userModel) HashPasswd(passwd string, salt string) string {
    hash := sha512.New()

    if _, e := hash.Write([]byte(passwd + salt)); e != nil {
        panic("could not hash password")
    }

    return hex.EncodeToString(hash.Sum(nil))
}
