// Package main
package main

import (
	"context"
	"time"

	db "github.com/Jphartogi/attendance-services/srv/attendance-srv/database"
	"github.com/Jphartogi/attendance-services/srv/attendance-srv/employee"
	attendance "github.com/Jphartogi/attendance-services/srv/attendance-srv/proto/attendance"
	"github.com/Jphartogi/attendance-services/srv/attendance-srv/utils"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/util/log"
	"google.golang.org/grpc"
)

func main() {
	go func() {
		for {
			grpc.DialContext(context.TODO(), "127.0.0.1:9091")
			time.Sleep(time.Second)
		}
	}()

	service := micro.NewService(
		micro.Name("attendance-service"),
	)

	// optionally setup command line usage
	service.Init()

	// initialize Database
	err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	db.InitDB()

	// Register Handlers
	attendance.RegisterAttendanceHandler(service.Server(), new(employee.EmployeeAttendance))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
