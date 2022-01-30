package repository

import (
	"book_management/domain/model"
	"context"
)

// BookRepository : Book における Repository のインターフェース
// -> 依存性逆転の法則により infra 層は domain 層（本インターフェース）に依存
type BookRepository interface {
	GetAll(ctx context.Context) ([]*model.Book, error)
}
