package controller

import "testing"

type DBMock struct {
}

func (db *DBMock) Insert(key, value string) error {
	return nil
}

func (db *DBMock) Find(key string) (string, error) {
	return key, nil
}

func TestURLShortener_Get(t *testing.T) {

}

func TestURLShortener_Save(t *testing.T) {

}
