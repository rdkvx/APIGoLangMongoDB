package model

//Model cryptocurrency to be used on the application
type CryptoCurrency struct {
	Id        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Symbol    string `json:"symbol" bson:"symbol"`
	Votes     int    `json:"votes" bson:"votes"`
	CreatedAT string `json:"createdat" bson:"createdat"`
	UpdatedAT string `json:"updatedat" bson:"updatedat"`
}
