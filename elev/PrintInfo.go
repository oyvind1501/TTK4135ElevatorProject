package elev

import "fmt"




func PrintHallTable() {
	fmt.Println("-----------------Hall Order Table:----------------------")
	if len(HallOrderTable) == 0 {
		fmt.Println("	No hall")
	} else {
		for _, tableElement := range HallOrderTable {
			switch tableElement.Direction {
			case DIR_UP:
				if tableElement.Status == STATUS_AVAILABLE {
					fmt.Println(string(tableElement.Command) + " " + "UP" + " " + strconv.Itoa(tableElement.Floor) + " " + strconv.Itoa(tableElement.TimeReserved.Minute()) + ":" + strconv.Itoa(tableElement.TimeReserved.Second()) + " " + tableElement.ReserveID + " " + "AVAILABLE")
				}
				if tableElement.Status == STATUS_OCCUPIED {
					fmt.Println(string(tableElement.Command) + " " + "UP" + " " + strconv.Itoa(tableElement.Floor) + " " + strconv.Itoa(tableElement.TimeReserved.Minute()) + ":" + strconv.Itoa(tableElement.TimeReserved.Second()) + " " + tableElement.ReserveID + " " + "OCCUPIED")
				}

			case DIR_DOWN:
				if tableElement.Status == STATUS_AVAILABLE {
					fmt.Println(string(tableElement.Command) + " " + "DOWN" + " " + strconv.Itoa(tableElement.Floor) + " " + strconv.Itoa(tableElement.TimeReserved.Minute()) + ":" + strconv.Itoa(tableElement.TimeReserved.Second()) + " " + tableElement.ReserveID + " " + "AVAILABLE")
				}
				if tableElement.Status == STATUS_OCCUPIED {
					fmt.Println(string(tableElement.Command) + " " + "DOWN" + " " + strconv.Itoa(tableElement.Floor) + " " + strconv.Itoa(tableElement.TimeReserved.Minute()) + ":" + strconv.Itoa(tableElement.TimeReserved.Second()) + " " + tableElement.ReserveID + " " + "OCCUPIED")
				}
			}
		}
	}
	fmt.Println("--------------------------------------------------------")
}

func PrintCabTable() {
	fmt.Println("-----------------Cab Order Table------------------------")
	if len(CabOrderTable) == 0 {
		fmt.Println("	No cab")
	} else {
		for _, tableElement := range CabOrderTable {
			fmt.Println("	Order at Floor:\t\t" + strconv.Itoa(tableElement.Floor))
		}
	}
	fmt.Println("--------------------------------------------------------")
}

func PrintCommunicationTable() {
	var nodeType string
	var nodeTypeId string

	for _, clients := range ClientTable {
		if clients.ClientInfo.Id == masterId {
			nodeType = "Master"
			nodeTypeId = clients.ClientInfo.Id
		}
		if clients.ClientInfo.Id == backupId {
			nodeType = "Backup"
			nodeTypeId = clients.ClientInfo.Id
		}
		if clients.ClientInfo.Id != masterId && clients.ClientInfo.Id != backupId {
			nodeType = "Node"
			nodeTypeId = clients.ClientInfo.Id
		}
	}
	fmt.Println("----------------Node Network Information----------------")
	fmt.Println("	Type:		", nodeType)
	fmt.Println("	Id:		", nodeTypeId)
	fmt.Println("--------------------------------------------------------")
}

func PrintStateInfo() {
	fmt.Println("-----------------Elevator Info--------------------------")
	fmt.Println("	Target floor:\t\t" + strconv.Itoa(TargetFloor))
	fmt.Println("	Last floor: \t\t" + strconv.Itoa(LastFloor))
	switch ElevatorDirection {
	case DIR_STOP:
		fmt.Println("	Direction:\t\tDIR_STOP")
	case DIR_UP:
		fmt.Println("	Direction:\t\tDIR_UP")
	case DIR_DOWN:
		fmt.Println("	Direction:\t\tDIR_DOWN")
	}
	fmt.Println("--------------------------------------------------------")
}

func PrintElevatorInfo() {
	for {
		PrintCommunicationTable()
		fmt.Println()
		PrintStateInfo()
		fmt.Println()
		PrintCabTable()
		fmt.Println()
		PrintHallTable()
		fmt.Println()

		time.Sleep(500 * time.Millisecond)
	}
}