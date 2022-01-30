package main

import (
	"book_management/infra/persistence"
	"book_management/usecase"
	"fmt"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/julienschmidt/httprouter"
	"learningGoCleanArchitecture/infrastructure/api/handler"
	"log"
	"net/http"
)

func main() {
	// 依存関係を注入（DI まではいきませんが一応注入っぽいことをしてる）
	// DI ライブラリを使えば、もっとスマートになるはず
	bookPersistence := persistence.NewBookPersistence()
	bookUseCase := usecase.NewBookUseCase(bookPersistence)
	bookHandker := handler.NewUserHandler(bookUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/v1/books", bookHandker.Index)

	// サーバ起動
	fmt.Println("==============================")
	fmt.Println("Server Start >> http://localhost:3000")
	fmt.Println("==============================")
	log.Fatal(http.ListenAndServe(":3000", router))
}
