package controllers

import (
	"googleauth/models"
	"log"

	"gorm.io/gorm"
)

func SetupWordRoutes() {

}
func CreateWord(db *gorm.DB) {
	word := models.Word{
		Word: "example",
		Definitions: []models.Definition{
			{Definition: "A thing characteristic of its kind or illustrating a general rule."},
		},
		PartOfSpeeches: []models.PartOfSpeech{
			{PartOfSpeech: "noun"},
		},
		ExampleSentences: []models.ExampleSentence{
			{Sentence: "This is an example of a simple Go application."},
		},
	}
	// Create a new word entry
	if err := db.Create(&word).Error; err != nil {
		log.Printf("Failed to create word: %v", err)
	} else {
		log.Println("Word created successfully")
	}
}

func ReadAllWords(db *gorm.DB) {
	var words []models.Word
	// fetch all words and preload related definitions, parts of speech, and example sentences
	if err := db.Preload("Definitions").Preload("PartOfSpeeches").Preload("ExampleSentences").Find(&words).Error; err != nil {
		log.Printf("Failed to fetch words: %v", err)
	} else {
		log.Printf("Words retrieved successfully: %v", words)
	}
}
