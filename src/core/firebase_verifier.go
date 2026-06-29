package core

import (
	"context"
	"fmt"

	firebaseAuth "firebase.google.com/go/v4/auth"
)

type FirebaseVerifier struct {
	credentialsFile string
}

func NewFirebaseVerifier(credentialsFile string) *FirebaseVerifier {
	return &FirebaseVerifier{credentialsFile: credentialsFile}
}

func (v *FirebaseVerifier) VerifyIDToken(ctx context.Context, idToken string) (*firebaseAuth.Token, error) {
	app, err := GetFirebaseApp(v.credentialsFile)
	if err != nil {
		return nil, fmt.Errorf("error al inicializar Firebase: %w", err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cliente Firebase Auth: %w", err)
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("token de Google inválido: %w", err)
	}
	return token, nil
}
