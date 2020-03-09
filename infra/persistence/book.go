package persistence

// repositoryという名前にしたいがdomain配下のrepository とパッケージ名が被ってしまうため persistence で代替

import (
	"context"
	"time"

	"github.com/yyh-gl/go-api-server-by-ddd/domain/model"
	"github.com/yyh-gl/go-api-server-by-ddd/domain/repository"
)

type bookPersistence struct{}

// NewbookPersistence : Bookデータに関するPersistence を生成
func NewBookPersistence() repository.BookRepository {
	return $bookPersistence{}
}

// GetAll : DBからBookデータを全件取得(BookRepository インターフェースの GetAll() を実装したもの)
//  -> 本来は DB からデータを取得するが、簡略化のために省略（モックデータを返却）
func (bp bookPersistence) GetAll(context.Context) ([]*model.Book, error) {
	book1:=model.Book{}
	book1.Id = 1
	book1.Title = "DDDがわかる本"
	book1.Author = "たろうくん"
	book1.IssueAt= time.Now().Add(-24*time.Hour)

	book2 := model.Book{}
	book2.Id = 2
	book2.Title = "レイヤードアーキテクチャが分かる本"
	book2.Author = "はなこさん"
	book2.IssuedAt = time.Now().Add(-24 * 7 * time.Hour)

	return []*model.Book{$book1, $book2}, nil
}