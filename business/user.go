package business

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
)

type HttpHandlers struct {
	db *sql.DB
}

func New(db *sql.DB) HttpHandlers {
	return HttpHandlers{
		db: db,
	}
}

func (h HttpHandlers) GetUsers(id string) ([]Users, string) {
	var intId int
	var str string

	if id != "" {
		var err error
		intId, err = strconv.Atoi(id)
		if err != nil {
			return nil, "Index id must be a valid integer!"
		}
		str = "SELECT * FROM Users WHERE u_id = $1"
	} else {
		str = "SELECT * FROM Users WHERE $1 = $1"
	}

	rows, err := h.db.Query(str, intId)
	if err != nil {
		return nil, "There was a problem while executing the query"
	}

	var UserStruct []Users
	for rows.Next() {
		var userId int
		var fName, mName, lName string

		err := rows.Scan(&userId, &fName, &mName, &lName)
		if err != nil {
			return nil, "Error while scanning database."
		}
		UserStruct = append(UserStruct, Users{U_id: userId, First_name: fName, Middle_name: mName, Last_name: lName})
	}

	if UserStruct == nil {
		err := fmt.Sprintf("User id %v does not exit!", intId)
		return nil, err
	}

	return UserStruct, ""
}

func (h HttpHandlers) CreateUsers(id string, reqBody []byte) string {
	var intId int
	var str string
	var newUser Users
	json.Unmarshal(reqBody, &newUser)

	if id != "" {
		var err error
		intId, err = strconv.Atoi(id)
		if err != nil {
			return "Index id must be a valid integer!"
		}
		str = "INSERT INTO Users(first_name, middle_name, last_name, u_id) VALUES ($1, $2, $3, $4)"
		_, e := h.db.Exec(str, newUser.First_name, newUser.Middle_name, newUser.Last_name, intId)
		if e != nil {
			return e.Error()
		}

	} else {
		str = "INSERT INTO Users(first_name, middle_name, last_name) VALUES ($1, $2, $3)"
		_, e := h.db.Exec(str, newUser.First_name, newUser.Middle_name, newUser.Last_name)
		if e != nil {
			return e.Error()
		}

	}
	return "Successfully added!"
}

func (h HttpHandlers) ChangeUser(id string, reqBody []byte) string {
	var intId int
	var str string
	var updatedUser Users
	json.Unmarshal(reqBody, &updatedUser)

	if id == "" {
		return "Please enter a User Id for updating an entry."
	}

	var err error
	intId, err = strconv.Atoi(id)

	if err != nil {
		return "Index id must be a valid integer!"
	}

	str = "UPDATE Users SET first_name = $1, middle_name = $2, last_name = $3 WHERE u_id = $4"
	_, e := h.db.Exec(str, updatedUser.First_name, updatedUser.Middle_name, updatedUser.Last_name, intId)

	if e != nil {
		return e.Error()
	}

	return "Update Successful!"

}

func (h HttpHandlers) DeleteUser(id string) string {

	if id == "" {
		return "Please enter a User Id for deleting an entry."
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		return "Index id must be a valid integer!"
	}

	rows, err := h.db.Query("SELECT * FROM Users_Addresses WHERE u_id = $1", intId)
	if err != nil {
		// return "Here's the problem..........."
		return err.Error()
	}

	if !rows.Next() {

		rows1, err1 := h.db.Query("SELECT * FROM Users WHERE u_id = $1", intId)
		if err1 != nil {
			// return "Here's the problem..........."
			return err.Error()
		}

		if !rows1.Next() {
			return "No such user exists."
		}

		h.db.Exec("DELETE FROM Users WHERE u_id = $1;", intId)
		return "Deleted Successfully!"

	} else {
		fmt.Println(rows.Next())
		return "User has addresses, so cannot be deleted."
	}

}
