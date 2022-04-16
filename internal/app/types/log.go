package types

import "time"

type Log struct {
	ObjectCode 	string      		`ch:"objectCode"`
	ActionCode 	string      		`ch:"actionCode"`
	Data 		map[string]string 	`ch:"data"`
	CreatedAt	time.Time    		`ch:"createdAt"`
}