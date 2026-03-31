package core

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2/google"
)

type FCMService struct {
	ProjectID       string
	CredentialsFile string
}

func NewFCMService(projectID, credentialsFile string) *FCMService {
	return &FCMService{
		ProjectID:       projectID,
		CredentialsFile: credentialsFile,
	}
}

type FCMMessage struct {
	Token string
	Title string
	Body  string
	Data  map[string]string
}

func (f *FCMService) getAccessToken(ctx context.Context) (string, error) {
	data, err := os.ReadFile(f.CredentialsFile)
	if err != nil {
		return "", fmt.Errorf("error al leer credenciales FCM: %w", err)
	}

	creds, err := google.CredentialsFromJSON(ctx, data,
		"https://www.googleapis.com/auth/firebase.messaging",
	)
	if err != nil {
		return "", fmt.Errorf("error al parsear credenciales FCM: %w", err)
	}

	token, err := creds.TokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("error al obtener access token FCM: %w", err)
	}

	return token.AccessToken, nil
}

func (f *FCMService) SendToToken(ctx context.Context, msg FCMMessage) error {
	accessToken, err := f.getAccessToken(ctx)
	if err != nil {
		return err
	}

	payload := map[string]interface{}{
		"message": map[string]interface{}{
			"token": msg.Token,
			"notification": map[string]string{
				"title": msg.Title,
				"body":  msg.Body,
			},
			"data": msg.Data,
			"android": map[string]interface{}{
				"priority": "high",
			},
			"apns": map[string]interface{}{
				"payload": map[string]interface{}{
					"aps": map[string]interface{}{
						"sound": "default",
					},
				},
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error al serializar payload FCM: %w", err)
	}

	url := fmt.Sprintf(
		"https://fcm.googleapis.com/v1/projects/%s/messages:send",
		f.ProjectID,
	)

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error al crear request FCM: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar notificación FCM: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errBody map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&errBody)
		return fmt.Errorf("FCM error %d: %v", resp.StatusCode, errBody)
	}

	return nil
}

func (f *FCMService) SendToMany(ctx context.Context, tokens []string, msg FCMMessage) []error {
	var errs []error
	for _, token := range tokens {
		msg.Token = token
		if err := f.SendToToken(ctx, msg); err != nil {
			errs = append(errs, fmt.Errorf("token %s...: %w", token[:10], err))
		}
	}
	return errs
}

func IsInvalidToken(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "UNREGISTERED") ||
		strings.Contains(msg, "INVALID_ARGUMENT") ||
		strings.Contains(msg, "NOT_FOUND")
}