package main

import (
	"net/http"
)

func TimesheetHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(GetTimesheets()))
}

func GetTimesheets() string {
	return `[
    {
        "project": "test",
        "hours": 99,
        "id": 100,
        "approved": false
    },
    {
        "project": "klefdms",
        "hours": 12323,
        "id": 101,
        "approved": false
    },
    {
        "project": "adas",
        "hours": 313,
        "id": 102,
        "approved": false
    }
]`
}
