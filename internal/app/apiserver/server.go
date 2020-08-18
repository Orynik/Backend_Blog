package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Kentabr5427/http-rest-api/internal/app/model"
	"github.com/Kentabr5427/http-rest-api/internal/app/store"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName        = "Kentabr5427"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
)

type ctxKey int8

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}
	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowCredentials()))
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"http://192.168.0.102:8080", "http://orynik.gitlab.io"})))
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")

	s.router.HandleFunc("/articles/post", s.handleArticlesPost()).Methods("POST")
	s.router.HandleFunc("/articles/get", s.handleArticlesGet()).Methods("GET")

	s.router.HandleFunc("/categories/post", s.handleCategoriesPost()).Methods("POST")
	s.router.HandleFunc("/categories/get", s.handleCategoriesGet()).Methods("GET")

	s.router.HandleFunc("/comments/post", s.handleCommentsPost()).Methods("POST")
	s.router.HandleFunc("/comments/get", s.handleCommentsGet()).Methods("GET")

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.autheticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")
}

func (s *server) handleCategoriesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		arr, err := s.store.Category().GetCategories()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, arr)
	}
}

func (s *server) handleCategoriesPost() http.HandlerFunc {
	type request struct {
		ID       int    `json:"id"`
		Category string `json:"category"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return

		}
		a := &model.Category{
			Category: req.Category,
		}
		if err := s.store.Category().Create(a); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, a)
	}
}

//................
func (s *server) handleArticlesPost() http.HandlerFunc {
	type request struct {
		Title    string `json:"title"`
		Date     string `json:"date"`
		Category string `json:"category"`
		Author   string `json:"author"`
		Content  string `json:"content"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return

		}
		a := &model.Article{
			Title:    req.Title,
			Date:     req.Date,
			Category: req.Category,
			Author:   req.Author,
			Content:  req.Content,
		}
		if err := s.store.Article().Create(a); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, a)
	}
}

//.....................
func (s *server) handleArticlesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		arr, err := s.store.Article().GetArticles()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, arr)
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
} //...

func (s *server) handleCommentsGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		arr, err := s.store.Comment().GetComments()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		s.respond(w, r, http.StatusOK, arr)
	}
}

func (s *server) handleCommentsPost() http.HandlerFunc {
	type request struct {
		ID      int    `json:"id"`
		Article int    `json:"idArticle"`
		Author  string `json:"author"`
		Text    string `json:"text"`
		Date    string `json:"date"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		a := &model.Comment{
			ID:      req.ID,
			Article: req.Article,
			Author:  req.Author,
			Text:    req.Text,
			Date:    req.Date,
		}
		if err := s.store.Comment().Create(a); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, a)
	}
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start))

	})
}

func (s *server) autheticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))

	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		if err := s.sessionStore.Save(r, w, session); err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}

}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
