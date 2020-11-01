package Entity

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/vikasd145/article_project/internal/config"

	article "github.com/vikasd145/article_project/pkg/gen"
)

type ArticleTable interface {
	GetArticleByID(cont context.Context, id string) (*article.Article, error)
	SetArticle(cont context.Context, articleData *article.Article) error
	GetAllArticle(cont context.Context) ([]article.Article, error)
}

type DB struct {
	*sql.DB
}

func InitDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Printf("Unable to make cconnection to sql:%v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Printf("Connection not responding ping:%v", err)
		return nil, err
	}
	db.SetMaxIdleConns(1000)

	return &DB{db}, nil
}

func (dbcli *DB) GetArticleByID(cont context.Context, id string) (*article.Article, error) {
	queryTemp := "select * from article_tab where id=?"
	cont, cancel := context.WithTimeout(cont, config.GlobalDynamicConfig.DBContextQueryTimeout)
	defer cancel()
	row := dbcli.QueryRowContext(cont, queryTemp, id)
	tempArticle := article.Article{}
	err := row.Scan(tempArticle.ArticleId, tempArticle.Title, tempArticle.Content, tempArticle.Author)
	if err != nil {
		log.Printf("ErrorValue scanning getting tempArticle by id:%v :%v", id, err)
		return nil, err
	}
	return &tempArticle, nil
}

func (dbcli *DB) GetAllArticle(cont context.Context) ([]article.Article, error) {
	queryTemp := "select * from article_tab"
	cont, cancel := context.WithTimeout(cont, config.GlobalDynamicConfig.DBContextQueryTimeout)
	defer cancel()
	stmt, err := dbcli.PrepareContext(cont, queryTemp)
	if err != nil {
		_ := fmt.Errorf("GetAllArticle: Error in prepraing db query err:%v", err)
		return nil, err
	}
	defer stmt.Close()
	results, err := stmt.QueryContext(cont)
	if err != nil {
		_ := fmt.Errorf("GetAllArticle: QueryContext error:%v", err)
		return nil, err
	}
	articleData := make([]article.Article, 1024)
	for results.Next() {
		tempArticle := article.Article{}
		err := results.Scan(tempArticle.ArticleId, tempArticle.Title, tempArticle.Content, tempArticle.Author)
		if err != nil {
			log.Printf("ErrorValue scanning getting article err:%v", err)
			continue
		}
		articleData = append(articleData, tempArticle)
	}
	return articleData, nil
}

func (dbcli *DB) SetArticle(cont context.Context, articleData *article.Article) error {
	ins, err := dbcli.Prepare("UPDATE article_tab SET title_id=?, author=?, content=? where id = ?")
	if err != nil {
		log.Printf("Error preparing SetArticle:%v, article:%v", err, articleData)
		return err
	}
	_, err = ins.Exec(articleData.GetTitle(), articleData.GetAuthor(), articleData.GetContent(), articleData.GetArticleId())
	if err != nil {
		log.Printf("Error Executing SetArticle:%v, article:%v", err, articleData)
		return err
	}
	return nil
}
