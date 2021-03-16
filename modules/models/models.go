package models

type Mother struct {
	UserId int `json:"user_id"`
	NicNo string `json:"nic_no"`
	Surname string `json:"surname"`
	FirstName string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName string `json:"last_name"`
	PreferredName string `json:"preferred_name"`
	DateOfBirth string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
	CityOfBirth string `json:"city_of_birth"`
	MaritalState string `json:"marital_state"`
	Religion string `json:"religion"`
	NoOfPregnancy string `json:"no_of_pregnancy"`
	DeathBirths string `json:"death_births"`
	LastPeriodDate string `json:"last_period_date"`
	BloodGroup string `json:"blood_group"`
	Alcoholic string `json:"alcoholic"`
	Diseases string `json:"diseases"`
	Address string `json:"address"`
	Email string `json:"email_address"`
	MobilePhoneNo string `json:"mobile_phone_no"`
	EmergencyMobilePhoneNo string `json:"emergency_mobile_phone_no"`
	EmergencyAddress string `json:"emergency_address"`
	LandPhoneNo string `json:"land_phone_no"`
	Education string `json:"education"`
	Occupation string `json:"occupation"`
	Employer string `json:"employer"`
	Remarks string `json:"remarks"`

	Password string `json:"login_password"`
}

type Father struct {
	UserId int `json:"user_id"`
	NicNo string `json:"nic_no"`
	MotherNicNo string `json:"mother_nic_no"`
	Surname string `json:"surname"`
	FirstName string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName string `json:"last_name"`
	PreferredName string `json:"preferred_name"`
	DateOfBirth string `json:"date_of_birth"`
	CountryOfBirth string `json:"country_of_birth"`
	CityOfBirth string `json:"city_of_birth"`
	MaritalState string `json:"marital_state"`
	Religion string `json:"religion"`
	BloodGroup string `json:"blood_group"`
	Alcoholic string `json:"alcoholic"`
	Diseases string `json:"diseases"`
	Address string `json:"address"`
	Email string `json:"email_address"`
	MobilePhoneNo string `json:"mobile_phone_no"`
	EmergencyMobilePhoneNo string `json:"emergency_mobile_phone_no"`
	EmergencyAddress string `json:"emergency_address"`
	LandPhoneNo string `json:"land_phone_no"`
	Education string `json:"education"`
	Occupation string `json:"occupation"`
	Employer string `json:"employer"`
	Remarks string `json:"remarks"`

	Password string `json:"password"`
}

type Child struct {
	MotherNicNo string `json:"mother_nic_no"`
	FatherNicNo string `json:"father_nic_no"`
	UserId int `json:"user_id"`
	Surname string `json:"surname"`
	FirstName string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName string `json:"last_name"`
	PreferredName string `json:"preferred_name"`
	DateOfBirth string `json:"date_of_birth"`
	DateOfRegistration string `json:"date_of_registration"`
	CountryOfBirth string `json:"country_of_birth"`
	CityOfBirth string `json:"city_of_birth"`
	Religion string `json:"religion"`
	BloodGroup string `json:"blood_group"`
	DeliveryMethod string `json:"delivery_method"`
	NoOfApgar string `json:"no_of_apgar"`
	Weight string `json:"weight"`
	HeadRound string `json:"head_round"`
	Height string `json:"height"`
	VitaminK string `json:"vitamin_k"`
	FeedFirst string `json:"feed_first"`
	FeedCorrect string `json:"feed_correct"`
	FeedPosition string `json:"feed_position"`
	Remarks string `json:"remarks"`

	MotherId int `json:"mother_id"`
	MotherName string `json:"mother_name"`
	MotherMobilePhoneNo string `json:"mobile_phone_no"`
	MotherAddress string `json:"address"`
}

type NIC struct {
	NicNo string `json:"nic_no"`
	UserType int `json:"user_type"`
}

type User struct {
	UserId string `json:"user_id"`
	UserType int `json:"user_type"`
}

type Weight struct {
	UserId string `json:"user_id"`
	Month string `json:"month"`
	Weight string `json:"weight"`
}

type VaccineNotification struct {
	UserId int `json:"user_id"`
	ChildFirstName string `json:"child_first_name"`
	ChildLastName string `json:"child_last_name"`
	DateOfBirth string `json:"date_of_birth"`
	VaccineCode string `json:"vaccine_code"`
	VaccineName string `json:"vaccine_name"`
	MotherFirstName string `json:"mother_first_name"`
	MotherLastName string `json:"mother_last_name"`
	Email string `json:"email"`
	PhoneNo string `json:"phone_no"`
}

type GeoLocation struct {
	UserId string `json:"user_id"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
