package model



type Scanner interface{
    Scan(args ...interface{}) error
}