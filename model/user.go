package model

import (
    "io"
    "crypto/rand"
    "crypto/sha512"
    "encoding/hex"
)

var UserModel = &userModel{"user"}

type User struct {
    id int64
    passwd, salt string

    Name, Email string
    CreatedAt, UpdatedAt int64
}

type UserForm struct {
    Id int64
    Name, Email, Passwd string
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

    u.salt = string(rnd)
    u.passwd = UserModel.HashPasswd(passwd, u.salt)
}

func (u *User) CheckPasswd(passwd string) bool {
    return UserModel.HashPasswd(passwd, u.salt) == u.passwd
}


type userModel struct {
    tabel string
}

func (m *userModel) User(id int64) *User {
    return nil
}

func (m *userModel) FindByEmail(email string) *User {

    return nil
}

func (m *userModel) Save(u *User) {

}

func (m *userModel) HashPasswd(passwd string, salt string) string {
    hash := sha512.New()

    if _, e := hash.Write([]byte(passwd + salt)); e != nil {
        panic("could not hash password")
    }

    return hex.EncodeToString(hash.Sum(nil))
}
