// CREATE TABLE Word (
// 	ID SERIAL PRIMARY KEY,
// 	Word VARCHAR(255) UNIQUE NOT NULL,
// 	CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );
package models

import "time"
// Word represents the Word table
type Word struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Word string `gorm:"type:varchar(255);unique;not null" json:"word"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	// Relationships
	Definitions []Definition `gorm:"foreignKey:WordID"`
	PartOfSpeeches []PartOfSpeech `gorm:"foreignKey:WordID"`
	ExampleSentences []ExampleSentence `gorm:"foreignKey:WordID"`
}
