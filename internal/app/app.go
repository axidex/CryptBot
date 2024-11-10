package app

type App interface {
	Run()
	Stop()
	//MustRun()
}
