package entities

import "github.com/mochi-yu/webapp-auth-practice/ent"

type MessageRequest struct {
	Content    string `json:"content"`
	AuthorName string `json:"author_name"`
}

func (mr *MessageRequest) ToEntMessage() (dest *ent.Message) {
	dest = &ent.Message{
		Content:    mr.Content,
		AuthorName: mr.AuthorName,
	}

	return
}

type MessageResponse struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	AuthorName string `json:"author_name"`
	CreatedAt  string `json:"created_at"`
}
