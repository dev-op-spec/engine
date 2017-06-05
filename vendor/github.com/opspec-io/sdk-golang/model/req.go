package model

import "time"

type EventFilter struct {
	// filter to events from these root op id's
	RootOpIds []string
	// filter to events occurring after & including this time
	Since *time.Time
}

type GetEventStreamReq struct {
	Filter *EventFilter
}

type KillOpReq struct {
	OpId string
}

type StartOpReq struct {
	// map of args keyed by input name
	Args map[string]*Value
	Pkg  *DCGOpCallPkg
}
