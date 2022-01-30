package persistence

import (
	"book_management/domain/model"
	"book_management/domain/repository"
	"context"
	"time"
)

// repository という名前にしたいが domain 配下の repository とパッケージ名が被ってしまうため persistence で大替

type bookPersistence struct{}

// NewBookPersistence : Book データに関する Persistence を生成
func NewBookPersistence() repository.BookRepository {
	return &bookPersistence{}
}

// GetAll : DB から Book データを全件取得 (BookRepository インターフェースの GetAll() を実装したもの)
//-> 本来は DB からデータを取得するが、簡略化のために省略（モックデータを返却）
func (bp bookPersistence) GetAll(ctx context.Context) ([]*model.Book, error) {
	book1 := model.Book{}
	book1.Id = 1
	book1.Title = "test"
	book1.Author = "taro"
	book1.IssuedAt = time.Now().Add(-24 * time.Hour)

	book2 := model.Book{}
	book2.Id = 2
	book2.Title = "test2"
	book2.Author = "jiro"
	book2.IssuedAt = time.Now().Add(-24 * 8 * time.Hour)

	return []*model.Book{&book1, &book2}, nil
}
