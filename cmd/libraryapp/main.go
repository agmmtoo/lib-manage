package main

import (
	"os"
	"os/signal"

	"github.com/agmmtoo/lib-manage/internal/core/users"
	"github.com/agmmtoo/lib-manage/internal/infra/http"
	"github.com/agmmtoo/lib-manage/internal/infra/psql"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)

	db, err := psql.NewLibraryAppDB("postgres://liber:liber@localhost:5432/libraryapp?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userService := users.NewService(db)

	server, err := http.NewServer(":8080", userService)

	if err != nil {
		panic(err)
	}

	go func() {
		err := server.Start()
		if err != nil {
			panic(err)
		}
	}()

	defer server.Stop()

	<-c
}
