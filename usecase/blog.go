package usecase

import (
	"context"

	"github.com/yyh-gl/go-api-server-by-ddd/domain/model"
	"github.com/yyh-gl/go-api-server-by-ddd/domain/repository"
)

// BookUseCase :BookにおけるUseCaseのインターフェース
type BookUseCase interface {
	GetAll(context.Context) ([]*model.Book, error)
}

type bookUseCase struct {
	bookRepository repository.BookRepository
}

// NewBookUseCase: Bookデータに関するUseCaseを生成
func NewBookUseCase(br repository.BookRepository) BookUseCase {
	return &bookUseCase{
		bookRepository: br,
	}
}

// GetAll : Book データを全件取得するためのユースケース
//  -> 本システムではあまりユースケース層の恩恵を受けれないが、もう少し大きなシステムになってくると、
//    「ドメインモデルの調節者」としての役割が見えてくる
func (bu BookUseCase) GetAll(ctx context.Context) (books []*model.Book, err error) {
	// Persistence(Repository)を呼び出し
	books, err = bu.bookRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
