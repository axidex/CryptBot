package app

import "context"

type App interface {
	Run() error
	Stop(error)
	//MustRun()
}

func HandleApp(ctx context.Context, app App) (execute func() error, interrupt func(error)) {
	currentCtx, cancel := context.WithCancel(ctx)
	return func() error {
			errorCh := make(chan error)
			go func() {
				err := app.Run()
				if err != nil {
					errorCh <- err
				}
			}()

			select {
			case <-currentCtx.Done():
				return currentCtx.Err()
			case err := <-errorCh:
				return err
			}
		},
		func(error) {
			cancel()
		}
}
