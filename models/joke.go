package models

import (
    "gorm.io/gorm"
)

type Joke struct {
    ID   uint   `json:"id"`
    Text string `json:"text"`
}

func GetRandomJoke(db *gorm.DB) (Joke, error) {
    var joke Joke
    if err := db.Order("RAND()").First(&joke).Error; err != nil {
        return joke, err
    }
    return joke, nil
}
