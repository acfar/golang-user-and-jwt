package controller

import(
	"log"
	"net/http"
	"refactorGolang/database"
	"refactorGolang/models"
	"refactorGolang/utils"
)

type Controller struct{}

var arrNews []models.News

var connection database.Connection


func (c Controller) GetData() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var news models.News
		var error models.Error

		db := connection.Connect()
		rows, err := db.Query("Select a.person_id, author, channel, content from identity_info a inner join news b on a.person_id=b.person_id")

		if err!=nil{
			log.Print(err)
		}

		for rows.Next(){
			if err := rows.Scan(&news.Id, &news.Author,&news.Channel,&news.Content,);
				err!=nil{
					log.Print(err)
					error.Message="Server Error"
					utils.RespondWithError(w,http.StatusInternalServerError,error)
					return
				} else {
					arrNews=append(arrNews, news)
				}
		}
		w.Header().Set("Content-Type","application/json")
		utils.RespondJson(w,arrNews)
	}
}
