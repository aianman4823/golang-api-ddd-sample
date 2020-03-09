package repository

import (
	"context"

	"github.com/yyh-gl/go-api-server-by-ddd/domain/model"
)

// BookRepoository: BookにおけるRepositoryのインターフェース
//  -> 依存性逆転の法則により infra 層は domain 層（本インターフェース）に依存
type BookRepoository interface {
	GetAll(context.Context) ([]*model.Book, error)
}