package models

import "time"

//Employee contains makers registration form
type Employee struct {
	EmployeeID string    `pg:"employee_id,notnull,pk" json:"employee_id"`
	ScheduleID uint32    `json:"schedule_id" pg:"schedule_id" binding:"required"`
	FullName   string    `json:"full_name" pg:"full_name" binding:"required"`
	RfidID     string    `json:"rfid_id" pg:"rfid_id,unique" binding:"required"`
	NIK        uint32    `json:"nik" pg:"nik,unique" binding:"required"`
	Role       string    `json:"role" pg:"role" binding:"required"`
	CreatedAt  time.Time `pg:"created_at,default:now()" json:"created_at"`
	UpdatedAt  time.Time `pg:"updated_at" json:"updated_at"`
}
