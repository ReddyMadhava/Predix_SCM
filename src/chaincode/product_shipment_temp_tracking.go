/*
Copyright TCS Ltd 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

         http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
	"strings"
	"time"
	//"github.com/hyperledger/fabric/core/crypto/primitives"
	//"github.com/op/go-logging"
)

//var myLogger = logging.MustGetLogger("track_shipment_temperature")

type trackShipmentTemp struct{}

func (t *trackShipmentTemp) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Errorf("Chaincode will write the tempature of package while its getting shipped and checks the temperature of the package")
	ct := time.Now()
	err := stub.CreateTable("detailedOrder", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "order_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "package_id", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "current_temperature", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "location", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "max_temperature", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "time", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "carrier", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "creation_date", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "min_temp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "shipping_address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "order_date", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "events", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "contract_status", Type: shim.ColumnDefinition_STRING, Key: false},
	})

	_, err = stub.InsertRow("detailedOrder", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: "123"}},
			&shim.Column{Value: &shim.Column_String_{String_: "p123"}},
			&shim.Column{Value: &shim.Column_String_{String_: "10"}},
			&shim.Column{Value: &shim.Column_String_{String_: "IN"}},
			&shim.Column{Value: &shim.Column_String_{String_: "20"}},
			&shim.Column{Value: &shim.Column_String_{String_: "11AM"}},
			&shim.Column{Value: &shim.Column_String_{String_: "BD"}},
			&shim.Column{Value: &shim.Column_String_{String_: ct.String()}},
			&shim.Column{Value: &shim.Column_String_{String_: "-2"}},
			&shim.Column{Value: &shim.Column_String_{String_: "A.O. Fox Memorial Hospital, Oneonta"}},
			&shim.Column{Value: &shim.Column_String_{String_: "01 Jan 2017"}},
			&shim.Column{Value: &shim.Column_String_{String_: "Departure"}},
			&shim.Column{Value: &shim.Column_String_{String_: "satisfied"}}},
	})

	err1 := stub.CreateTable("uniqueOrder", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "order_id", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "package_id", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "current_temperature", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "location", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "max_temperature", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "time", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "carrier", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "creation_date", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "min_temp", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "shipping_address", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "order_date", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "events", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "contract_status", Type: shim.ColumnDefinition_STRING, Key: false},
	})

	_, err1 = stub.InsertRow("uniqueOrder", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: "123"}},
			&shim.Column{Value: &shim.Column_String_{String_: "p123"}},
			&shim.Column{Value: &shim.Column_String_{String_: "10"}},
			&shim.Column{Value: &shim.Column_String_{String_: "IN"}},
			&shim.Column{Value: &shim.Column_String_{String_: "20"}},
			&shim.Column{Value: &shim.Column_String_{String_: "11AM"}},
			&shim.Column{Value: &shim.Column_String_{String_: "BD"}},
			&shim.Column{Value: &shim.Column_String_{String_: ct.String()}},
			&shim.Column{Value: &shim.Column_String_{String_: "-2"}},
			&shim.Column{Value: &shim.Column_String_{String_: "A.O. Fox Memorial Hospital, Oneonta"}},
			&shim.Column{Value: &shim.Column_String_{String_: "01 Jan 2017"}},
			&shim.Column{Value: &shim.Column_String_{String_: "Departure"}},
			&shim.Column{Value: &shim.Column_String_{String_: "satisfied"}}},
	})
	if err != nil && err1 != nil {
		return nil, err
	} else {
		return nil, nil
	}
}

func (t *trackShipmentTemp) insertRow(stub shim.ChaincodeStubInterface, args []string, tablename string, status string) ([]byte, error) {
	fmt.Errorf("writing record in the table - [%s], %s,%s,%s,%s,%s,%s,%s", len(args), args[0], args[1], args[2], args[3], args[4], args[5], args[6])
	ct := time.Now()
	_, err := stub.InsertRow(tablename, shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: args[0]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[1]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[2]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[3]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[4]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[5]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[6]}},
			&shim.Column{Value: &shim.Column_String_{String_: ct.String()}},
			&shim.Column{Value: &shim.Column_String_{String_: args[7]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[8]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[9]}},
			&shim.Column{Value: &shim.Column_String_{String_: args[10]}},
			&shim.Column{Value: &shim.Column_String_{String_: status}}},
	})
	fmt.Errorf("Inserted - [%s]", args)

	if err != nil {
		return nil, errors.New("Failed to insert record")
	}
	return nil, err
}

func (t *trackShipmentTemp) updateRow(stub shim.ChaincodeStubInterface, tablename string, status string, row shim.Row) ([]byte, error) {
	ct := time.Now()
	_, err := stub.ReplaceRow(tablename, shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: row.Columns[0].GetString_()}},
			&shim.Column{Value: &shim.Column_String_{String_: row.Columns[1].GetString_()}},
			&shim.Column{Value: &shim.Column_String_{String_: row.Columns[2].GetString_()}},
			&shim.Column{Value: &shim.Column_String_{String_: row.Columns[3].GetString_()}},
			&shim.Column{Value: &shim.Column_String_{String_: row.Columns[4].GetString_()}},
			&shim.Column{Value: &shim.Column_String_{String_: row.Columns[5].GetString_()}},
			&shim.Column{Value: &shim.Column_String_{String_: row.Columns[6].GetString_()}},
			&shim.Column{Value: &shim.Column_String_{String_: ct.String()}},
			&shim.Column{Value: &shim.Column_String_{String_: row.Columns[8].GetString_()}},
			&shim.Column{Value: &shim.Column_String_{String_: row.Columns[9].GetString_()}},
			&shim.Column{Value: &shim.Column_String_{String_: row.Columns[10].GetString_()}},
			&shim.Column{Value: &shim.Column_String_{String_: row.Columns[11].GetString_()}},
			&shim.Column{Value: &shim.Column_String_{String_: status}}},
	})
	if err != nil {
		return nil, errors.New("Failed to update record")
	}
	return nil, err
}

func (t *trackShipmentTemp) updateAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Errorf("Length of args - [%d]", len(args))

	current_temp, _ := strconv.Atoi(string(args[2]))
	max_temp, _ := strconv.Atoi(string(args[4]))
	min_temp, _ := strconv.Atoi(string(args[7]))

	if len(args) != 11 {
		return nil, errors.New("Incorrect number of arguments. Expecting 11")
	}

	fmt.Errorf("Function Invoked Current - [%d] max temp - [%d]", current_temp, max_temp)

	if max_temp >= current_temp && min_temp <= current_temp {
		fmt.Errorf("Condition Satisfied")
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: args[0]}}
		columns = append(columns, col1)
		row, err1 := stub.GetRow("uniqueOrder", columns)

		if len(row.Columns) == 0 {
			fmt.Errorf("First Time writing record")
			_, err1 = t.insertRow(stub, args, "detailedOrder", "satisfied")
			_, err1 = t.insertRow(stub, args, "uniqueOrder", "satisfied")

		} else {
			_, err1 = t.insertRow(stub, args, "detailedOrder", "satisfied")
		}
		if err1 != nil {
			return nil, err1
		}

	} else {
		fmt.Errorf("Condition Not Satisfied")
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: args[0]}}
		fmt.Errorf("Order ID", col1)
		columns = append(columns, col1)
		row, err1 := stub.GetRow("uniqueOrder", columns)
		fmt.Errorf("row", row)

		/*if len(row.Columns) == 0 {
		          fmt.Errorf("Writing the record in else for first time")
		          _, err1 = t.insertRow(stub, args, "detailedOrder", "Inspection is required")
		          _, err1 = t.insertRow(stub, args, "uniqueOrder", "Inspection is required")
		  } else if row.Columns[12].GetString_() == "satisfied" {
		          _, err1 = t.insertRow(stub, args, "detailedOrder", "Inspection is required")
		          _, err1 = t.updateRow(stub, "uniqueOrder", "Inspection is required", row)
		  } else {
		          _, err1 = t.insertRow(stub, args, "detailedOrder", "Inspection is required")
		  } */
		_, err1 = t.insertRow(stub, args, "detailedOrder", "Inspection is required")
		if len(row.Columns) == 0 {
			fmt.Errorf("Writing the record in else for first time")
			_, err1 = t.insertRow(stub, args, "uniqueOrder", "Inspection is required")
		} else if row.Columns[12].GetString_() == "satisfied" {
			_, err1 = t.updateRow(stub, "uniqueOrder", "Inspection is required", row)
		}
		if err1 != nil {
			return nil, err1
		}
	}
	return nil, nil

}

func (t *trackShipmentTemp) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getUniqueOrderQuery" {
		return t.getUniqueOrderQuery(stub, args)
	} else if function == "getCompleteOrderDetails" {
		return t.getCompleteOrderDetails(stub, args)
	} else if function == "getOrderStatus" {
		return t.getOrderStatus(stub, args)
	} else if function == "getAllCompleteOrderDetails" {
		return t.getAllCompleteOrderDetails(stub, args)
	}

	err_msg := "Received unknown function [%s] for query."
	fmt.Errorf(err_msg, function)
	return nil, fmt.Errorf(err_msg, function)
}

func (t *trackShipmentTemp) getCompleteOrderDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	order_id := args[0]
	fmt.Errorf("Recieved Parameter [%s]", string(order_id))
	if order_id != "" {
		fmt.Errorf("Order is not null ", string(order_id))
		var columns []shim.Column
		var err_msg string
		var completed_trans []string
		col1 := shim.Column{Value: &shim.Column_String_{String_: order_id}}
		fmt.Errorf("col1 ", col1)
		columns = append(columns, col1)
		fmt.Errorf("col1 ", columns)
		rows, err := stub.GetRows("detailedOrder", columns)
		fmt.Errorf("col1 ", rows)
		if err != nil {
			err_msg = "Error: Not able query detailedOrder table. %s"
			fmt.Errorf(err_msg, err)
			return nil, fmt.Errorf(err_msg, err)
		}
		fmt.Errorf("Rows: %s", rows)

		for row := range rows {
			order_id := row.Columns[0].GetString_()
			package_id := row.Columns[1].GetString_()
			current_temp := row.Columns[2].GetString_()
			location := row.Columns[3].GetString_()
			max_temp := row.Columns[4].GetString_()
			time := row.Columns[5].GetString_()
			carrier := row.Columns[6].GetString_()
			creation_date := row.Columns[7].GetString_()
			min_temp := row.Columns[8].GetString_()
			shipping_address := row.Columns[9].GetString_()
			order_date := row.Columns[10].GetString_()
			events := row.Columns[11].GetString_()
			status := row.Columns[12].GetString_()
			strRes := "{\"order_id\": \"" + order_id + "\"," + "\"package_id\": \"" + package_id + "\"," + "\"current_temp\": \"" + current_temp + "\"," + "\"location\": \"" + location + "\"," + "\"max_temp\": \"" + max_temp + "\"," + "\"Time\": \"" + time + "\"," + "\"carrier\": \"" + carrier + "\"," + "\"creation_date\": \"" + creation_date + "\"," + "\"min_temp\": \"" + min_temp + "\"," + "\"shipping_address\": \"" + shipping_address + "\"," + "\"order_date\": \"" + order_date + "\"," + "\"events\": \"" + events + "\"," + "\"status\": \"" + status + "\"}"
			completed_trans = append(completed_trans, strRes)
		}
		return []byte(strings.Join(completed_trans, ",")), nil
	} else {
		return nil, errors.New("Excepting Order Id")
	}

	return nil, nil

}

func (t *trackShipmentTemp) getUniqueOrderQuery(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	order_id := args[0]
	fmt.Errorf("Recieved Parameter [%s]", string(order_id))
	if order_id != "" {
		fmt.Errorf("Order is not null ", string(order_id))
		var columns []shim.Column
		var err_msg string
		// var completed_trans []string
		col1 := shim.Column{Value: &shim.Column_String_{String_: order_id}}
		fmt.Errorf("col1 ", col1)
		columns = append(columns, col1)
		fmt.Errorf("col1 ", columns)
		row, err := stub.GetRow("uniqueOrder", columns)
		fmt.Errorf("col1 ", row)
		if err != nil {
			err_msg = "Error: Not able query uniqueOrder table. %s"
			fmt.Errorf(err_msg, err)
			return nil, fmt.Errorf(err_msg, err)
		}
		if len(row.Columns) == 0 || err != nil {
			err_msg = "Error: Order_ID not found in table uniqueOrder table. %s"
			fmt.Errorf("Order_ID not found in table uniqueOrder table [%s]", string(order_id))
			return nil, fmt.Errorf(err_msg)
		}

		fmt.Errorf("Row: %s", row)

		order_id := row.Columns[0].GetString_()
		package_id := row.Columns[1].GetString_()
		current_temp := row.Columns[2].GetString_()
		location := row.Columns[3].GetString_()
		max_temp := row.Columns[4].GetString_()
		time := row.Columns[5].GetString_()
		carrier := row.Columns[6].GetString_()
		creation_date := row.Columns[7].GetString_()
		min_temp := row.Columns[8].GetString_()
		shipping_address := row.Columns[9].GetString_()
		order_date := row.Columns[10].GetString_()
		events := row.Columns[11].GetString_()
		status := row.Columns[12].GetString_()
		order_details := "{\"order_id\": \"" + order_id + "\"," + "\"package_id\": \"" + package_id + "\"," + "\"current_temp\": \"" + current_temp + "\"," + "\"location\": \"" + location + "\"," + "\"max_temp\": \"" + max_temp + "\"," + "\"Time\": \"" + time + "\"," + "\"carrier\": \"" + carrier + "\"," + "\"creation_date\": \"" + creation_date + "\"," + "\"min_temp\": \"" + min_temp + "\"," + "\"shipping_address\": \"" + shipping_address + "\"," + "\"order_date\": \"" + order_date + "\"," + "\"events\": \"" + events + "\"," + "\"status\": \"" + status + "\"}"
		return []byte(order_details), nil

	} else {
		return nil, errors.New("Excepting Order Id")
	}
	return nil, nil
}

func (t *trackShipmentTemp) getOrderStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	order_id := args[0]
	var err_msg string
	fmt.Errorf("Recieved Parameter [%s]", string(order_id))
	if order_id != "" {
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: order_id}}
		columns = append(columns, col1)

		row, err := stub.GetRow("uniqueOrder", columns)
		if err != nil {
			fmt.Errorf("Failed retriving status [%s]: [%s]", string(order_id), err)
			return nil, errors.New(fmt.Sprintf("Failed retriving status [%s]: [%s]", string(order_id), err))
		}

		if len(row.Columns) == 0 || err != nil {
			err_msg = "Error: Order_ID not found in table uniqueOrder table. %s"
			fmt.Errorf("Order_ID not found in table uniqueOrder table [%s]", string(order_id))
			return nil, fmt.Errorf(err_msg)
		}

		order_id := row.Columns[0].GetString_()
		status := row.Columns[12].GetString_()
		order_details := "{\"order_id\": \"" + order_id + "\"status\": \"" + status + "\"," + "\"}"
		return []byte(order_details), nil

	} else {
		return nil, errors.New("Excepting Order Id")
	}

}

func (t *trackShipmentTemp) getAllCompleteOrderDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var tablename = args[0]

	if tablename == "" {
		return nil, errors.New("Expecting Table Name")
	}

	var columns []shim.Column
	rows, err := stub.GetRows(tablename, columns)
	if err != nil {
		fmt.Errorf("Failed query table [%s]: [%s]", string(tablename), err)
		return nil, errors.New(fmt.Sprintf("Failed query table [%s]: [%s]", string(tablename), err))
	}
	fmt.Errorf("Rows: %s", len(rows))

	var completed_trans []string
	for row := range rows {
		order_id := row.Columns[0].GetString_()
		package_id := row.Columns[1].GetString_()
		current_temp := row.Columns[2].GetString_()
		location := row.Columns[3].GetString_()
		max_temp := row.Columns[4].GetString_()
		time := row.Columns[5].GetString_()
		carrier := row.Columns[6].GetString_()
		creation_date := row.Columns[7].GetString_()
		min_temp := row.Columns[8].GetString_()
		shipping_address := row.Columns[9].GetString_()
		order_date := row.Columns[10].GetString_()
		events := row.Columns[11].GetString_()
		status := row.Columns[12].GetString_()
		strRes := "{\"order_id\": \"" + order_id + "\"," + "\"package_id\": \"" + package_id + "\"," + "\"current_temp\": \"" + current_temp + "\"," + "\"location\": \"" + location + "\"," + "\"max_temp\": \"" + max_temp + "\"," + "\"Time\": \"" + time + "\"," + "\"carrier\": \"" + carrier + "\"," + "\"creation_date\": \"" + creation_date + "\"," + "\"min_temp\": \"" + min_temp + "\"," + "\"shipping_address\": \"" + shipping_address + "\"," + "\"order_date\": \"" + order_date + "\"," + "\"events\": \"" + events + "\"," + "\"status\": \"" + status + "\"}"
		completed_trans = append(completed_trans, strRes)
	}
	return []byte(strings.Join(completed_trans, ",")), nil
}

func (t *trackShipmentTemp) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "updateAsset" {
		return t.updateAsset(stub, args)
	}
	return nil, errors.New("Received unknown function invocation")
}

func main() {
	err := shim.Start(new(trackShipmentTemp))
	if err != nil {
		fmt.Printf("Error starting Track Shipment chaincode: %s", err)

	}
}

