package api

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/vikasd145/article_project/pkg/article_admin"
	article "github.com/vikasd145/article_project/pkg/gen"
	"log"
	"net/http"
	"strconv"
)

func GetArticle(w http.ResponseWriter, r *http.Request) {
	res := article.GetArticleResponse{
		Status:               proto.Int32(http.StatusOK),
		Message:              proto.String("Success"),
	}
	if r.URL.Path == "/articles/" {
		log.Println("Print all articles")
		info,err := article_admin.GlobalAdmin.ArticleDb.GetAllArticle(r.Context())
		if err!= nil{
			_ = fmt.Errorf("Error in getting all articles from db err:%v",err)
			res.Status = proto.Int32(http.StatusInternalServerError)
			res.Message = proto.String("Get db problem Please try again")
			FinalWork(w,res,http.StatusInternalServerError)
			return
		}
		res.Data = ConsturctGetAllArticleReponse(info)
	} else {
		path := r.URL.Path
		if len(path) <= 10 {
			fmt.Errorf("Wrong request")
			res.Status = proto.Int32(http.StatusBadRequest)
			res.Message = proto.String("Article id wrong format")
			FinalWork(w,res,http.StatusBadRequest)
			return
		}
		articleId, err := strconv.ParseInt(path[10:],10,64)
		if err != nil {
			fmt.Printf("Wrong format article id : %v err : %v", path[10:], err)
			res.Status = proto.Int32(http.StatusBadRequest)
			res.Message = proto.String("Article id not int")
			FinalWork(w,res,http.StatusBadRequest)
			return
		}
		fmt.Printf("Give article with id: %v", articleId)
		info,err := article_admin.GlobalAdmin.ArticleDb.GetArticleByID(r.Context(),articleId)
		if err!= nil{
			_ = fmt.Errorf("Error in getting all articles from db err:%v",err)
			res.Status = proto.Int32(http.StatusInternalServerError)
			res.Message = proto.String("Get db problem Please try again")
			FinalWork(w,res,http.StatusInternalServerError)
			return
		}
		res.Data = ConstructGetArticleByIdResponse(info)
	}
	FinalWork(w,res,http.StatusOK)
}

func ConsturctGetAllArticleReponse(data []article.Article) []*article.GetArticleData{
	res := make([]*article.GetArticleData,0,len(data))
	for key := range data{
		temp := &article.GetArticleData{
			Id:                  data[key].ArticleId,
			Title:               data[key].Title,
			Content:              data[key].Content,
			Author:               data[key].Author,
		}
		res = append(res,temp)
	}
	return res
}

func ConstructGetArticleByIdResponse(data *article.Article)[]*article.GetArticleData{
	res := make([]*article.GetArticleData,0,1)
	temp := &article.GetArticleData{
		Id:                  data.ArticleId,
		Title:               data.Title,
		Content:              data.Content,
		Author:               data.Author,
	}
	res = append(res,temp)
	return res
}
