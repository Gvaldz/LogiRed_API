package core

import (
	"context"
	"sync"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var (
	fbApp  *firebase.App
	fbOnce sync.Once
	fbErr  error
)

// GetFirebaseApp devuelve la instancia compartida del Firebase App (singleton).
// Solo se inicializa una vez; llamadas posteriores ignoran credentialsFile.
func GetFirebaseApp(credentialsFile string) (*firebase.App, error) {
	fbOnce.Do(func() {
		opt := option.WithCredentialsFile(credentialsFile)
		fbApp, fbErr = firebase.NewApp(context.Background(), nil, opt)
	})
	return fbApp, fbErr
}
