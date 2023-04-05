package log

import (
	"errors"
	"fmt"
	configPkg "github.com/Adesubomi/lightning-node-manager/pkg/config"
	"github.com/getsentry/sentry-go"

	"time"
)

func ConnectToSentry(sentryConfig configPkg.SentryConfig, env configPkg.AppEnv) error {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:     fmt.Sprintf(sentryConfig.DNS),
		Debug:   !env.Is(configPkg.AppEnvProduction),
		Release: fmt.Sprintf(sentryConfig.Release),
	}); err != nil {
		sentry.CaptureException(err)
		fmt.Printf("?? Sentry could not be initialized: %v\n", err)
	}

	return nil
}

func ReportMessage(m string) {
	go func() {
		fmt.Println(">> message log :: ", m)
		sentry.CaptureMessage(m)
	}()
}

func ReportError(err error) {
	go func() {
		if err != nil {
			msg := fmt.Sprintf(
				">> [%v] ERROR :: %v",
				time.Now().Format("Mon, 1 Jan 2022 13:59"),
				err.Error(),
			)

			fmt.Println(msg)
			sentry.CaptureException(err)
		} else {
			logicError := errors.New("trying to report a non-error")
			fmt.Println(">> error log :: ", logicError)
			sentry.CaptureException(err)
		}
	}()
}
