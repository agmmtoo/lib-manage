package main

import (
	"os"
	"os/signal"
	"syscall"

	bookService "github.com/agmmtoo/lib-manage/internal/core/book"
	libraryService "github.com/agmmtoo/lib-manage/internal/core/library"
	"github.com/agmmtoo/lib-manage/internal/core/libraryapp"
	loanService "github.com/agmmtoo/lib-manage/internal/core/loan"
	userService "github.com/agmmtoo/lib-manage/internal/core/user"
	"github.com/agmmtoo/lib-manage/internal/infra/http"
	"github.com/agmmtoo/lib-manage/internal/infra/psql"
	// "github.com/agmmtoo/lib-manage/internal/infra/slog"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	db, err := psql.NewLibraryAppDB("postgres://liber:liber@localhost:5432/libraryapp?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// logger := slog.NewLogger()

	user := userService.New(db)
	book := bookService.New(db)
	library := libraryService.New(db)
	loan := loanService.New(db)

	service := libraryapp.New(user, book, library, loan)

	server, err := http.NewServer(":8080", service)

	if err != nil {
		panic(err)
	}

	go server.Start()

	defer server.Stop()

	<-c
}
