package usecase

import (
	"book_management/domain/model"
	"book_management/domain/repository"
	"context"
)

// BookUseCase : Book における UseCase のインターフェース
type BookUseCase interface {
	GetAll(ctx context.Context) ([]*model.Book, error)
}

type bookUseCase struct {
	bookRepository repository.BookRepository
}

// NewBookUseCase : Book データに関する UseCase を生成
func NewBookUseCase(br repository.BookRepository) BookUseCase {
	return &bookUseCase{
		bookRepository: br,
	}
}

// GetAll : Book データを全件取得するためのユースケース
//  -> 本システムではあまりユースケース層の恩恵を受けれないが、もう少し大きなシステムになってくると、
//    「ドメインモデルの調節者」としての役割が見えてくる
func (bu bookUseCase) GetAll(ctx context.Context) (books []*model.Book, err error) {
	// Persistence (Repository)を呼び出し
	books, err = bu.bookRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return books, nil
}
