package payment

type Result struct {
	Status string
	Message string
}

type User struct {
	ID    int   
	Name     string 
	Email    string 
	Password string 
	Balance float64
}

type LoginRequest struct {
	Email    string
	Password string
}

type Value struct {
	Amount float64
}

type Send struct {
	ID int
	Amount float64
}

type BalanceRequest struct {
	ID int
}
                      
type BalanceResponse struct {
	StatusCode int
	Balance float64
}
                      
var users []User
var currentUserId int

func SignUp(req *User) Result {
	var result = Result{}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		result = Result{
			Status: "Failed",
			Message: "Invalid signup credentials!",
		}
	} else {
		users = append(users, *req)
		result = Result{
			Status: "Success",
			Message: "Signup Successful 👍!",
		}
	}

	return result
}

func Login(req *LoginRequest) Result {
	var result = Result{}

	if req.Email == "" || req.Password == "" {
		result = Result{
			Status: "Failed",
			Message: "Invalid login credentials!",
		}
	} else {
		// verify email and password
		for i, u := range users {
			if u.Email == req.Email && u.Password == req.Password {	
				// assign current user ID
				currentUserId = users[i].ID
				result = Result{
					Status: "Success",
					Message: "Login Successful 👍!",
				}
				break
			} else {
				result = Result{
					Status: "Failed",
					Message: "Password or E-mail do not match!!",
				}
			}
		}
		
	}

	return result
}

func Deposit(req *Value) Result {
	var result = Result{}

	if req.Amount == 0 {

		result = Result{
			Status: "Failed",
			Message: "Please input an amount!",
		}

	} else {
		for i, u := range users {
			if u.ID == currentUserId {
				users[i].Balance += req.Amount 
				result = Result{
					Status: "Success",
					Message: "Deposit Successful 👍!",
				}
				break
			}
		}
	}

	return result
}

func Transfer(req *Send) Result {
	var result = Result{}

	if req.Amount == 0 {

		result = Result{
			Status: "Failed",
			Message: "Please input an amount!",
		}

	} else {
		for i, u := range users {
			if u.ID == currentUserId {
				// check if balance is sufficient before debiting
				if users[i].Balance < req.Amount {
					result = Result{
						Status: "Failed",
						Message: "Insufficient balance!",
					}
				} else {
					users[i].Balance -= req.Amount
					// credit recipient
					for i, u := range users {
						if u.ID == req.ID {
							users[i].Balance += req.Amount 
							result = Result{
								Status: "Success",
								Message: "Transfer Successful 👍!",
							}
							break
						}
					}
				}
				break
			}
		}
	}

	return result
}

func TransferOut(req *Value) Result {
	var result = Result{}

	if req.Amount == 0 {

		result = Result{
			Status: "Failed",
			Message: "Please input an amount!",
		}

	} else {
		// check if balance is sufficient before debiting
		for i, u := range users {
			if u.ID == currentUserId {
				if users[i].Balance < req.Amount {
					result = Result{
						Status: "Failed",
						Message: "Insufficient balance!",
					}
				} else {
					users[i].Balance -= req.Amount
					result = Result{
						Status: "Success",
						Message: "Transfer Successful 👍!",
					}
				}
				break
			}
		}
	}

	return result
}

func GetBalance(req *BalanceRequest) BalanceResponse {
	var result = BalanceResponse{}

	if req.ID == 0 {

		result = BalanceResponse{
			StatusCode: 404,
		}

	} else {
		for i, u := range users {
			if u.ID == req.ID {
				result = BalanceResponse{
					StatusCode: 200,
					Balance: users[i].Balance,
				}
				break
			}
		}
	}
	return result
}
