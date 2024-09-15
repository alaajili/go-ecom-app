package dto

type UserLoginDto struct {
	Email		string	`json:"email"`
	Password	string	`json:"password"`
}

type UserRegisterDto struct {
	UserLoginDto
	Phone	string `json:"phone"`
}

type VerificationCodeDto struct {
	Code	int	`json:"code"`
}

type BecomeSellerDto struct {
	FirstName		string	`json:"first_name"`
	LastName		string	`json:"last_name"`
	Phone			string	`json:"phone"`
	AccountNumber	uint	`json:"account_number"`
	SwiftCode		string	`json:"swift_code"`
	PaymentType		string	`json:"payment_type"`
}
