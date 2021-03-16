package database

import (
	"back-end/modules/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"time"
)

type ProductDetails struct {
	gorm.Model
	Code string `json:"Code"`
	Price uint
	Prices uint
}

var DB *sql.DB

func OpenPostgresDatabaseConnection() {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	DB, err = sql.Open(os.Getenv("DB_DIALECT"), connectionString)
	if err != nil {
		panic(err)

	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}
}

func GetAllUsers() interface{} {
	return nil
}

func GetUser(user Login) (interface{}, error) {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("SELECT u.id, u.first_name, u.last_name, u.user_type_id " +
		"FROM Login l, User u WHERE l.user_id = u.id AND l.user_name = '%s' AND l.password = '%s' AND u.status = 1 " +
		"AND l.status = 1", user.NIC, user.Password)

	rows, err := DB.Query(sqlQuery)
	authUser := Login{}

	if err == nil {
		var id, typeId int
		var firstName, lastName string

		for rows.Next() {
			_ = rows.Scan(&id, &firstName, &lastName, &typeId)
			authUser.Id = id
			authUser.FirstName = firstName
			authUser.LastName = lastName
			authUser.UserTypeId = typeId
		}
	}
	return authUser, err
}

func GetEmail(user Login) (interface{}, error) {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("SELECT u.id, u.nic_no, c.email, l.user_name, l.password FROM Contact c, User u, Login l " +
		"WHERE c.user_id = u.id AND l.user_id = u.id AND u.nic_no = '%s' AND u.status = 1", user.NIC)

	rows, err := DB.Query(sqlQuery)
	authUser := Login{}

	if err == nil {
		var id int
		var nicNo, email, userName, password string

		for rows.Next() {
			_ = rows.Scan(&id, &nicNo, &email, &userName, &password)
			authUser.Id = id
			authUser.NIC = nicNo
			authUser.UserName = userName
			authUser.Password = password
			authUser.Email = email
		}
	}
	return authUser, err
}

func SaveMother(mother models.Mother) error {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("INSERT INTO User (surname, first_name, middle_name, last_name, preferred_name, nic_no, " +
		"date_of_birth, country_of_birth, city_of_birth, marital_state, religion, blood_group, alcoholic, diseases, education, occupation, employer, remarks, user_type_id, status) " +
		"VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d, %d)",
		mother.Surname, mother.FirstName, mother.MiddleName, mother.LastName, mother.PreferredName, mother.NicNo, mother.DateOfBirth, mother.CountryOfBirth,
		mother.CityOfBirth, mother.MaritalState, mother.Religion, mother.BloodGroup, mother.Alcoholic, mother.Diseases, mother.Education, mother.Occupation,
		mother.Employer, mother.Remarks, 2, 1)

	res, err := DB.Exec(sqlQuery)

	if err == nil {
		var userId int64
		userId, err = res.LastInsertId()

		if err != nil {
			return err
		}

		sqlQuery = fmt.Sprintf("INSERT INTO Mother (no_of_pregnancy, death_births, last_period_date, user_id) VALUES ('%s', '%s', '%s', %d)",
			mother.NoOfPregnancy, mother.DeathBirths, mother.LastPeriodDate, userId)

		_, err := DB.Exec(sqlQuery)

		if err != nil {
			return err
		}

		sqlQuery = fmt.Sprintf("INSERT INTO Contact (address, email, mobile_phone_no, land_phone_no, emergency_address, emergency_mobile_phone_no, user_id) " +
			"VALUES ('%s', '%s', '%s', '%s', '%s', '%s', %d)",
			mother.Address, mother.Email, mother.MobilePhoneNo, mother.LandPhoneNo, mother.EmergencyAddress, mother.EmergencyMobilePhoneNo, userId)

		_, err = DB.Exec(sqlQuery)

		if err != nil {
			return err
		}

		sqlQuery = fmt.Sprintf("INSERT INTO Login (user_name, password, user_id) VALUES ('%s', '%s', %d)", mother.NicNo, mother.Password, userId)

		_, err = DB.Exec(sqlQuery)

		if err != nil {
			return err
		}
	}
	return err
}

func SaveFather(father models.Father) error {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("INSERT INTO User (surname, first_name, middle_name, last_name, preferred_name, nic_no, " +
		"date_of_birth, country_of_birth, city_of_birth, marital_state, religion, " +
		"blood_group, alcoholic, diseases, education, occupation, employer, remarks, user_type_id, status) " +
		"VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d, %d)",
		father.Surname, father.FirstName, father.MiddleName, father.LastName, father.PreferredName, father.NicNo, father.DateOfBirth, father.CountryOfBirth,
		father.CityOfBirth, father.MaritalState, father.Religion, father.BloodGroup, father.Alcoholic, father.Diseases, father.Education, father.Occupation, father.Employer, father.Remarks, 3, 1)

	res, err := DB.Exec(sqlQuery)

	if err == nil {
		var userId int64
		userId, err = res.LastInsertId()

		if err != nil {
			return err
		}

		sqlQuery = fmt.Sprintf("INSERT INTO Contact (address, email, mobile_phone_no, land_phone_no, emergency_address, emergency_mobile_phone_no, user_id) " +
			"VALUES ('%s', '%s', '%s', '%s', '%s', '%s', %d)",
			father.Address, father.Email, father.MobilePhoneNo, father.LandPhoneNo, father.EmergencyAddress, father.EmergencyMobilePhoneNo, userId)

		_, err = DB.Exec(sqlQuery)

		if err != nil {
			return err
		}

		sqlQuery = fmt.Sprintf("INSERT INTO Family (mother_nic_no, father_nic_no) VALUES ('%s', '%s')", father.MotherNicNo, father.NicNo)

		_, err = DB.Exec(sqlQuery)
	}

	return err
}


func SaveChild(child models.Child) error {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("INSERT INTO User (surname, first_name, middle_name, last_name, preferred_name, date_of_birth, country_of_birth, " +
		"city_of_birth, religion, blood_group, remarks, user_type_id, status) " +
		"VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d, %d)",
		child.Surname, child.FirstName, child.MiddleName, child.LastName, child.PreferredName, child.DateOfBirth, child.CountryOfBirth,
		child.CityOfBirth, child.Religion, child.BloodGroup, child.Remarks, 4, 1)

	res, err := DB.Exec(sqlQuery)

	if err == nil && child.MotherNicNo != "" {
		var childId int64
		childId, err = res.LastInsertId()

		sqlQuery = fmt.Sprintf("INSERT INTO Family (mother_nic_no, father_nic_no) VALUES ('%s', '%s')", child.MotherNicNo, child.FatherNicNo)

		res, err = DB.Exec(sqlQuery)

		if err == nil {
			var familyId int64
			familyId, err = res.LastInsertId()

			sqlQuery := fmt.Sprintf("INSERT INTO FamilyChildren (family_id, user_id) VALUES (%d, %d)", familyId, childId)
			res, err = DB.Exec(sqlQuery)
		}

		sqlQuery = fmt.Sprintf("INSERT INTO Children (date_of_registration, delivery_method, no_of_apgar, weight, head_round, height, " +
			"vitamin_k, feed_first, feed_correct, feed_position, user_id) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %d)",
			child.DateOfRegistration, child.DeliveryMethod, child.NoOfApgar, child.Weight, child.HeadRound, child.Height, child.VitaminK, child.FeedFirst,
			child.FeedCorrect, child.FeedPosition, childId)

		_, err := DB.Exec(sqlQuery)

		if err != nil {
			return err
		}

		err = SaveVaccine(child, childId)
	}

	return err
}

func SaveVaccine(child models.Child, childId int64) error {
	layout := "2006-01-02"
	dob, err := time.Parse(layout, child.DateOfBirth)

	if err != nil {
		return err
	}

	sqlQuery := fmt.Sprintf("SELECT id, month FROM Vaccine WHERE status = 1 ")

	rows, err := DB.Query(sqlQuery)

	if err != nil {
		return err
	}

	var id, month int

	if rows != nil {
		for rows.Next() {
			rows.Scan(&id, &month)

			effectiveDate := dob.AddDate(0, month, 0)

			sqlQuery := fmt.Sprintf("INSERT INTO ChildrenVaccine (user_id, vaccine_id, effective_date, status) VALUES (%d, %d, '%s', %d)",
				childId, id, effectiveDate, 0)

			_, err := DB.Exec(sqlQuery)

			if err != nil {
				return err
			}
		}
	}
	return err
}

func GetUserFromNic(nic models.NIC) (int, error) {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("SELECT u.id FROM User u WHERE u.nic_no = '%s' AND u.user_type_id = %d AND u.status = 1",
		nic.NicNo, nic.UserType)

	rows, err := DB.Query(sqlQuery)
	var id int

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&id)
		}
	}

	return id, err
}

func GetWeightByUserId(user models.User) ([]models.Weight, error) {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("SELECT w.month, w.weight FROM Weight w WHERE w.user_id = %s ", user.UserId)

	rows, err := DB.Query(sqlQuery)

	var weightList []models.Weight
	var month string
	var amount string

	if rows != nil {
		for rows.Next() {
			var weight = models.Weight{}
			rows.Scan(&month, &amount)
			weight.Month = month
			weight.Weight = amount
			weightList = append(weightList, weight)
		}
	}

	return weightList, err
}

func GetChildrenByNic(nic models.NIC) ([]models.Child, error) {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("SELECT u.id, u.first_name, u.last_name, u.date_of_birth FROM User u, Family f, FamilyChildren fc " +
		"WHERE u.id = fc.user_id AND f.id = fc.family_id AND f.mother_nic_no = '%s' AND u.status = 1", nic.NicNo)

	rows, err := DB.Query(sqlQuery)

	var children []models.Child
	var id int
	var firstName, lastName, dob string

	if rows != nil {
		for rows.Next() {
			var child = models.Child{}
			rows.Scan(&id, &firstName, &lastName, &dob)
			child.UserId = id
			child.FirstName = firstName
			child.LastName = lastName
			child.DateOfBirth = dob
			children = append(children, child)
		}
	}

	return children, err
}

func GetAllMothers() ([]models.Mother, error) {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("SELECT id, first_name, last_name, nic_no FROM User WHERE user_type_id = 2 AND status = 1")

	rows, err := DB.Query(sqlQuery)

	var mothers []models.Mother
	var id int
	var firstName, lastName, nicNo string

	if rows != nil {
		for rows.Next() {
			var mother = models.Mother{}
			rows.Scan(&id, &firstName, &lastName, &nicNo)
			mother.UserId = id
			mother.FirstName = firstName
			mother.LastName = lastName
			mother.NicNo = nicNo
			mothers = append(mothers, mother)
		}
	}

	return mothers, err
}

func GetMother(userId string) (models.Mother, error) {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("SELECT u.id, u.first_name, u.last_name, u.nic_no, c.mobile_phone_no, c.address " +
		"FROM User u, Contact c WHERE u.id = c.user_id AND u.id = %s AND u.status = 1", userId)

	rows, err := DB.Query(sqlQuery)

	var mother = models.Mother{}
	var id int
	var firstName, lastName, nicNo, mobilePhoneNo, address string

	if rows != nil {
		for rows.Next() {
			rows.Scan(&id, &firstName, &lastName, &nicNo, &mobilePhoneNo, &address)
			mother.UserId = id
			mother.FirstName = firstName
			mother.LastName = lastName
			mother.NicNo = nicNo
			mother.MobilePhoneNo = mobilePhoneNo
			mother.Address = address
		}
	}

	return mother, err
}

func GetChild(userId string) (models.Child, error) {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("SELECT u.id, u.first_name child_name, u.last_name, u.date_of_birth, m.id mother_id, m.first_name mother_name, c.mobile_phone_no, c.address " +
		"FROM User m, Contact c, Family f, FamilyChildren fc LEFT JOIN User u ON fc.user_id = u.id WHERE m.nic_no = f.mother_nic_no " +
		"AND f.id = fc.family_id AND fc.user_id = u.id AND m.id = c.user_id AND u.id = %s AND u.status = 1", userId)

	rows, err := DB.Query(sqlQuery)

	var child = models.Child{}
	var id, motherId int
	var firstName, lastName, dob, motherName, mobilePhoneNo, address string

	if rows != nil {
		for rows.Next() {
			rows.Scan(&id, &firstName, &lastName, &dob, &motherId, &motherName, &mobilePhoneNo, &address)
			child.UserId = id
			child.FirstName = firstName
			child.LastName = lastName
			child.DateOfBirth = dob
			child.MotherId = motherId
			child.MotherName = motherName
			child.MotherMobilePhoneNo = mobilePhoneNo
			child.MotherAddress = address
		}
	}

	return child, err
}

func GetLocation(userId string) (models.GeoLocation, error) {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("SELECT c.latitude, c.longitude FROM Contact c WHERE c.user_id = %s", userId)

	rows, err := DB.Query(sqlQuery)

	var geo = models.GeoLocation{}
	var latitude, longitude float32

	if rows != nil {
		for rows.Next() {
			rows.Scan(&latitude, &longitude)
			geo.Latitude = latitude
			geo.Longitude = longitude
		}
	}

	return geo, err
}

func GetAllChildren() ([]models.Child, error) {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("SELECT id, first_name, last_name, (SELECT IFNULL(f.mother_nic_no,'') FROM Family f, FamilyChildren fc " +
		"WHERE f.id = fc.family_id AND fc.user_id = u.id) FROM User u WHERE user_type_id = 4 AND status = 1")

	rows, err := DB.Query(sqlQuery)

	var children []models.Child
	var id int
	var firstName, lastName, nicNo string

	if rows != nil {
		for rows.Next() {
			var child = models.Child{}
			rows.Scan(&id, &firstName, &lastName, &nicNo)
			child.UserId = id
			child.FirstName = firstName
			child.LastName = lastName
			child.MotherNicNo = nicNo
			children = append(children, child)
		}
	}

	return children, err
}

func SaveWeight(weight models.Weight) error {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("INSERT INTO Weight (user_id, month, weight) VALUES ('%s', '%s', '%s')",
		weight.UserId, weight.Month, weight.Weight)
	_, err := DB.Exec(sqlQuery)
	return err
}

func SaveLocation(geo models.GeoLocation) error {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("UPDATE Contact SET latitude = %f, longitude = %f WHERE user_id = %s", geo.Latitude, geo.Longitude, geo.UserId)
	_, err := DB.Exec(sqlQuery)
	return err
}

func GetChildrenVaccineByDate(date time.Time) ([]models.VaccineNotification, error) {
	OpenPostgresDatabaseConnection()

	sqlQuery := fmt.Sprintf("SELECT u.id, u.first_name, u.last_name, IFNULL(c.email,''), c.mobile_phone_no, IFNULL(v.name,''), IFNULL(v.code,''), ch.first_name, ch.last_name, ch.date_of_birth " +
		"FROM User u, Contact c, Family f, FamilyChildren fc, Vaccine v, ChildrenVaccine cv LEFT JOIN User ch ON cv.user_id = ch.id " +
		"WHERE c.user_id = u.id AND f.mother_nic_no = u.nic_no AND f.id = fc.family_id AND cv.vaccine_id = v.id AND cv.user_id = fc.user_id " +
		"AND cv.status = 0 AND cv.effective_date = '%s' AND u.status = 1", date.Format("2006-01-02"))

	rows, err := DB.Query(sqlQuery)

	var vaccines []models.VaccineNotification

	if rows != nil {
		var id int
		var motherFirstName, motherLastName, dob, vaccineName, vaccineCode, childFirstName, childLastName, email, phoneNo string

		for rows.Next() {
			var vaccine = models.VaccineNotification{}
			err = rows.Scan(&id, &motherFirstName, &motherLastName, &email, &phoneNo, &vaccineName, &vaccineCode, &childFirstName, &childLastName, &dob)

			if err != nil {
				return vaccines, err
			}

			vaccine.UserId = id
			vaccine.MotherFirstName = motherFirstName
			vaccine.MotherLastName = motherLastName
			vaccine.Email = email
			vaccine.PhoneNo = phoneNo

			vaccine.ChildFirstName = childFirstName
			vaccine.ChildLastName = childLastName
			vaccine.DateOfBirth = dob
			vaccine.VaccineCode = vaccineCode
			vaccine.VaccineName = vaccineName

			vaccines = append(vaccines, vaccine)
		}
	}

	return vaccines, err
}
