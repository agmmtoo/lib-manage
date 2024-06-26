package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	bookService "github.com/agmmtoo/lib-manage/internal/core/book"
	libraryService "github.com/agmmtoo/lib-manage/internal/core/library"
	"github.com/agmmtoo/lib-manage/internal/core/libraryapp"
	loanService "github.com/agmmtoo/lib-manage/internal/core/loan"
	membershipService "github.com/agmmtoo/lib-manage/internal/core/membership"
	settingService "github.com/agmmtoo/lib-manage/internal/core/setting"
	staffService "github.com/agmmtoo/lib-manage/internal/core/staff"
	subscriptionService "github.com/agmmtoo/lib-manage/internal/core/subscription"
	userService "github.com/agmmtoo/lib-manage/internal/core/user"
	"github.com/agmmtoo/lib-manage/internal/infra/http"
	"github.com/agmmtoo/lib-manage/internal/infra/psql"
	"github.com/agmmtoo/lib-manage/pkg/libraryapp/config"
	"github.com/joho/godotenv"
	// "github.com/agmmtoo/lib-manage/internal/infra/slog"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db, err := psql.NewLibraryAppDB(os.Getenv(config.ENV_KEY_DB_URL))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// logger := slog.NewLogger()

	user := userService.New(db)
	book := bookService.New(db)
	library := libraryService.New(db)
	loan := loanService.New(db)
	staff := staffService.New(db)
	setting := settingService.New(db)
	membership := membershipService.New(db)
	subscription := subscriptionService.New(db)

	service := libraryapp.New(user, book, library, loan, staff, setting, membership, subscription)

	port := fmt.Sprintf(":%s", os.Getenv(config.ENV_KEY_PORT))
	server, err := http.NewServer(port, service)

	if err != nil {
		panic(err)
	}

	go server.Start()

	defer server.Stop()

	<-c
}
