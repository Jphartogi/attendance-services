package main

import (
	"context"
	"fmt"

	"github.com/asim/go-micro/v3"
	attendance "github.com/jphartogi/user-attendance-services/srv/attendance-srv/proto/attendance"
)

func main() {
	// create a new service
	service := micro.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := attendance.NewAttendanceService("attendance-service", service.Client())
	count := 5
	for i := 0; i < count; i++ {
		// Make request
		rsp, err := cl.Register(context.Background(), &attendance.Employee{
			FullName:  "Joshua Phartogi",
			GroupId:   5,
			StartTime: "18:04:05.000Z",
			EndTime:   "12:04:05.000Z",
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(rsp.Msg)
	}

}
