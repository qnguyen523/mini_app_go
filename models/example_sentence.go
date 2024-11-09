// CREATE TABLE ExampleSentences (
// 	SentenceID SERIAL PRIMARY KEY,
// 	WordID INT NOT NULL,
// 	Sentence TEXT NOT NULL,
// 	FOREIGN KEY (WordID) REFERENCES Word(WordID)
// );
package models
// ExampleSentences represents the ExampleSentences table
type ExampleSentence struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	WordID int `json:"word_id" gorm:"not null"`
	Sentence string `json:"sentence" gorm:"not null"`
	Word Word `gorm:"foreignkey:WordID"`
}

