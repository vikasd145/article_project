package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/vikasd145/article_project/pkg/article_admin"
	article "github.com/vikasd145/article_project/pkg/gen"
	"net/http"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w,r.Body,1024*1024*5) //do not read request more than 5 MB just a random size to save server from Attack
	dec := json.NewDecoder(r.Body)
	createReq := &article.CreateRequest{}
	err := dec.Decode(createReq)
	res := article.CreateResponse{
		Status:               proto.Int32(http.StatusOK),
		Message:              proto.String("Success"),
	}
	if err != nil{
		_ = fmt.Errorf("CreateArticle: Error in parsing request err:%v",err)
		res.Status = proto.Int32(http.StatusRequestEntityTooLarge)
		res.Message = proto.String("Not Processed")
		FinalWork(w,res,http.StatusRequestEntityTooLarge)
		return
	}
	 err = CheckCreateRest(createReq)
	if err != nil{
		_ = fmt.Errorf("Check request returned error:%v",err)
		res.Status = proto.Int32(http.StatusBadRequest)
		res.Message = proto.String(err.Error())
		FinalWork(w,res,http.StatusBadRequest)
		return
	}
	articleData := &article.Article{
		Title:                createReq.Title,
		Author:               createReq.Author,
		Content:              createReq.Content,
	}
	articleid,err := article_admin.GlobalAdmin.ArticleDb.SetArticle(r.Context(),articleData)
	res.Data = &article.CreateArticleData{Id:proto.Int64(articleid)}
	FinalWork(w,res,http.StatusOK)
}


func CheckCreateRest(data *article.CreateRequest)error{
	if data.GetTitle() == ""{
		return errors.New("Title cannot be empty")
	}
	if data.GetContent() == ""{
		return errors.New("Content cannot be empty")
	}
	if data.GetAuthor() == ""{
		return errors.New("Author cannot be empty")
	}
	return nil
}