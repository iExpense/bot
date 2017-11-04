package fireclient

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"github.com/iexpense/bot/iparser"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

const (
	dbScope   = "https://www.googleapis.com/auth/firebase.database"
	userScope = "https://www.googleapis.com/auth/userinfo.email"
)

type Fireclient struct {
}

// Reference
// https://firebase.google.com/docs/database/rest/auth
func NewFireclient() (*Fireclient, error) {
	keyFilePath := viper.GetString("firebase.keyfile_path")
	if keyFilePath == "" {
		return nil, fmt.Errorf("key firebase.keyfile_path is empty")
	}

	opt := option.WithCredentialsFile(keyFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Printf("error initializing app: %v", err)
		return nil, err
	}

	_, err = app.Auth(context.Background())
	if err != nil {
		log.Printf("error getting Auth client: %v\n", err)
		return nil, err
	}

	return &Fireclient{}, nil
}

func (f *Fireclient) HandleExpenseCommand(cmd *iparser.Command) (string, error) {
	return "", fmt.Errorf("not implemented yet")
}
