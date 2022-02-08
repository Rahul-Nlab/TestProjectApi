package business

type Users struct {
	U_id        int    `json:"u_id"`
	First_name  string `json:"first_name"`
	Middle_name string `json:"middle_name"`
	Last_name   string `json:"last_name"`
	// Address Addresses
}

type Addresses struct {
	A_id    int    `json:"a_id"`
	Street  string `json:"street"`
	Area    string `json:"area"`
	Pincode int    `json:"pincode"`
	City    string `json:"city"`
}
