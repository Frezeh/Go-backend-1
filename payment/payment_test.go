package payment

import "testing"

func TestSignUp(t *testing.T) {
	tests := []User{
		{
			ID: 1,
			Name: "User A",
			Email: "UserA@gmail.com",
			Password: "12345",
		},
		{
			ID: 2,
			Name: "User B",
			Email: "UserB@gmail.com",
			Password: "56789",
		},
	  }

	for _, test := range tests {
		result := SignUp(&test)

		if result.Status != "Success" {
			t.Errorf("FAILED. expected %s got %s\n", "Success", result)
		} else {
			t.Logf("PASSED. expected %s got %s\n", "Success", result)
		}
	}
}

func TestLogin(t *testing.T) {	
	var argument = LoginRequest{
		Email: "UserA@gmail.com",
		Password: "12345",
	}

	result := Login(&argument)

	if result.Status != "Success" {
		t.Errorf("FAILED. expected %s got %s\n", "Success", result)
	} else {
		t.Logf("PASSED. expected %s got %s\n", "Success", result)
	}
}

func TestDeposit(t *testing.T) {	
	var argument = Value{
		Amount: 100,
	}

	result := Deposit(&argument)

	if result.Status != "Success" {
		t.Errorf("FAILED. expected %s got %s\n", "Success", result)
	} else {
		t.Logf("PASSED. expected %s got %s\n", "Success", result)
	}
}

func TestTransfer(t *testing.T) {
	var argument = Send{
		ID: 2,
		Amount: 50,
	}

	result := Transfer(&argument)

	if result.Status != "Success" {
		t.Errorf("FAILED. expected %s got %s\n", "Success", result)
	} else {
		t.Logf("PASSED. expected %s got %s\n", "Success", result)
	}
}

func TestTransferOut(t *testing.T) {
	var argument = Value{
		Amount: 50,
	}

	result := TransferOut(&argument)

	if result.Status != "Success" {
		t.Errorf("FAILED. expected %s got %s\n", "Success", result)
	} else {
		t.Logf("PASSED. expected %s got %s\n", "Success", result)
	}
}

func TestBalance(t *testing.T) {
	tests := []BalanceRequest{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	}

	for _, test := range tests {
		result := GetBalance(&test)

		if result.StatusCode != 200 {
			t.Errorf("FAILED. expected %d got %d \n", 200, result.StatusCode)
		} else {
			t.Logf("PASSED. expected %d got %d\n", 200, result.StatusCode)
		}
	}
}
