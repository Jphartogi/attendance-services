package employee

import (
	"context"
	"time"

	"github.com/Jphartogi/attendance-services/srv/attendance-srv/proto/attendance"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/google/uuid"
)

type EmployeeAttendance struct {
}

func (s *EmployeeAttendance) Register(ctx context.Context, req *attendance.Employee, rsp *attendance.RegisterEmployeeResponse) error {
	log.Log("Received Say.Hello request")
	employeeID := uuid.New()
	layout := "15:04:05.000Z"
	startTime, _ := time.Parse(layout, req.GetStartTime())
	endTime, _ := time.Parse(layout, req.GetEndTime())
	rsp.Msg = "Hello " + req.GetFullName() + " \nYour attendance start time is " + startTime.String() + "End time is " + endTime.String() + "\nAnd your id is " + employeeID.String()
	return nil
}
