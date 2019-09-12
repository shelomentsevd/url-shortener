package controller

import (
	"fmt"
	
        "github.com/shelomentsevd/url-shortener/db"
)

const urlLength = 5

type Controller interface {
	Save(url string) (string, error)
	Get(key string) error
}

type URLShortener struct {
	DB db.Database
}

func NewURLShortener(database db.Database) Controller {
	return &URLShortener{
		DB: database,
	}
}

func (us *URLShortener) Save(url string) (string, error) {
	key, err := us.DB.Find(url)
	if err != nil {
		return "", fmt.Errorf("URLShortener: database search error: %v", err)
	}

	if len(key) > 0 {
		return key, nil
	}

	key = generateKey(urlLength)

	us.DB.Insert(key, url)

	return key, nil
}

func (us *URLShortener) Get(key string) error {

	return nil
}
