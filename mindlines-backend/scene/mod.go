package scene

import "github.com/open-spaced-repetition/go-fsrs/v3"

type Module struct {
	Service    Service
	Controller Controller
}

func CreateModule() Module {
	service := Service{
		fsrs: *fsrs.NewFSRS(fsrs.DefaultParam()),
	}
	controller := Controller{
		Service: service,
		FSRS:    *fsrs.NewFSRS(fsrs.DefaultParam()),
	}

	return Module{
		service,
		controller,
	}
}
