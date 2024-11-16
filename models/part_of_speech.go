// CREATE TABLE PartOfSpeech (
// 	PartOfSpeechID SERIAL PRIMARY KEY,
// 	WordID INT NOT NULL,
// 	PartOfSpeech VARCHAR(50) NOT NULL,
// 	FOREIGN KEY (WordID) REFERENCES Word(WordID)
// );
package models

// PartOfSpeech represents the PartOfSpeech table
type PartOfSpeech struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	WordID string `gorm:"not null" json:"word_id"`
	PartOfSpeech string `json:"part_of_speech" gorm:"not null"`
	Word Word `gorm:"foreignkey:WordID"`
}