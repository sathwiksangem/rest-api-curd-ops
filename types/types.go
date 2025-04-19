package types

type Item struct {
	ID    string `json:"id" dynamodbav:"id"`
	Name  string `json:"name" dynamodbav:"name"`
	Email string `json:"email" dynamodbav:"email"`
}
