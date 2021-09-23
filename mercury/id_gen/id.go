package id_gen

import (
	"fmt"
	"github.com/sony/sonyflake"
)

var sonyFlake *sonyflake.Sonyflake
var sonyMachineID uint16

func getMachineID() (uint16,error) {
	return sonyMachineID, nil
}

func Init(machineId uint16) (err error) {
	sonyMachineID = machineId

	st := sonyflake.Settings{}
	st.MachineID = getMachineID

	sonyFlake = sonyflake.NewSonyflake(st)

	return
}

func GetId() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not inited")
		return
	}
	id, err = sonyFlake.NextID()

	return
}
