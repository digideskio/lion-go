package gcloud_test

import (
	"os"

	"go.pedge.io/protolog"
	"go.pedge.io/protolog/gcloud"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/logging/v1beta3"
)

func Example() error {
	projectID, _ := os.LookupEnv("GCLOUD_PROJECT_ID")
	logName := "protolog"
	client, err := google.DefaultClient(
		context.Background(),
		logging.LoggingWriteScope,
	)
	if err != nil {
		return err
	}
	service, err := logging.New(client)
	if err != nil {
		return err
	}
	logger := protolog.NewStandardLogger(
		gcloud.NewPusher(
			service.Projects.Logs.Entries,
			projectID,
			logName,
		),
	)
	logger.Infoln("Hello from protolog!")
	return nil
}
