package goauth2

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"net/url"

	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
)

func BuildAutorizationRequestURL(oauth2Conf oauth2.Config) (string, error) {
	stateBytes := make([]byte, 16)
	if _, err := rand.Read(stateBytes); err != nil {
		return "", err
	}
	state := fmt.Sprintf("%x", stateBytes)
	return oauth2Conf.AuthCodeURL(state, oauth2.AccessTypeOffline), nil
}

func GetAuthorizationCode(ctx context.Context, authorizationURL string, redirectURI string) (string, error) {
	// redirectURIからポート番号を取得するためにパース
	u, err := url.Parse(redirectURI)
	if err != nil {
		return "", err
	}
	mux := http.NewServeMux()
	quit := make(chan string)
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		quit <- r.URL.Query().Get("code")
	})
	mux.HandleFunc("GET /favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", u.Port()),
		Handler: mux,
	}
	// 認可レスポンスの受け取り先としてhttpサーバーを起動
	go srv.ListenAndServe()
	// ブラウザを開いて同意確認
	if err := open.Start(authorizationURL); err != nil {
		return "", err
	}
	// 認可コードを待つ
	authzCode := <-quit
	// httpサーバーを停止
	if err := srv.Shutdown(ctx); err != nil {
		return "", err
	}
	return authzCode, nil
}
