// CREATE TABLE Definition (
// 	ID SERIAL PRIMARY KEY,
// 	WordID INT NOT NULL,
// 	Definition TEXT NOT NULL,
// 	FOREIGN KEY (WordID) REFERENCES Word(WordID)
// );
package models
// Definition represents the Definition table
type Definition struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	WordID string `gorm:"not null" json:"word_id"`
	Definition string `gorm:"not null" json:"definition"`
	Word Word `gorm:"foreignkey:WordID"`
}