package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/Orynik/Backend_Blog/internal/app/store/sqlstore"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

//Start ..
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(DatabaseURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
