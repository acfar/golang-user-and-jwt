package models

type News struct{
	Id int `json:"id"`
	Author string `json:"author"`
	Channel string `json:"channel"`
	Content string `json:"content"`
	Person_id int `json:"person_id"`
}