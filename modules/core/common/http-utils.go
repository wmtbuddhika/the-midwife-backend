package common

import (
	"back-end/modules/core/authentication/login"
	"back-end/modules/database"
	"back-end/modules/models"
	"back-end/modules/notification"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"time"
)


func Login(w http.ResponseWriter, r *http.Request) {

	var user = database.Login{}
	var response = database.Response{}
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
		RespondWithObject(response, w)
	} else {
		err := json.Unmarshal(data, &user)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
			RespondWithObject(response, w)
		} else {
			u, err := database.GetUser(user)
			user = u.(database.Login)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
				RespondWithObject(response, w)
			} else {
				if user.Id > 0 {
					token := login.GenerateToken(user)

					w.Header().Set("Token", token)

					response.Status = true
					response.StatusCode = http.StatusOK
					response.Data = user
					response.Token = token

					RespondWithObject(response, w)
				} else {
					response.Status = false
					response.StatusCode = http.StatusUnauthorized
					RespondWithObject(response, w)
				}
			}
		}
	}
}

func Forgot(w http.ResponseWriter, r *http.Request) {

	var user = database.Login{}
	var response = database.Response{}
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
		RespondWithObject(response, w)
	} else {
		err := json.Unmarshal(data, &user)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
			RespondWithObject(response, w)
		} else {
			u, err := database.GetEmail(user)
			user = u.(database.Login)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
				RespondWithObject(response, w)
			} else {
				if user.Id > 0 {
					message := fmt.Sprintf("<h1>Retreive Your Password</h1>" +
						"<p>We got forgot password request from NIC No: %s. <br>" +
						"Please find your login credentils for the Midwife System. <br><br> " +
						"Username : %s <br><br>" +
						"Password : %s </p><br><br>" +
						"Regards, <br>" +
						"<strong>Midwife System</strong>", user.NIC, user.UserName, user.Password)
					notification.SendEmail(models.SiteEmail, []string{user.Email}, models.ForgotPasswordSubject, message)
					response.Status = true
					response.StatusCode = http.StatusOK
					response.Data = "success"

					RespondWithObject(response, w)
				} else {
					response.Status = false
					response.StatusCode = http.StatusUnauthorized
					RespondWithObject(response, w)
				}
			}
		}
	}
}

func Authenticate(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := login.DecodeToken(r)

			if err != nil {
				claims := token.Claims.(jwt.MapClaims)
				if time.Now().Sub(time.Unix(int64(claims["exp"].(float64)), 0)) > 30*time.Second {
					w.WriteHeader(http.StatusBadRequest)
					return
				} else {
					w.Header().Set("Token", login.RefreshToken(claims))
					w.WriteHeader(http.StatusOK)
					endpoint(w,r)
				}
			}
			if token.Valid {
				w.WriteHeader(http.StatusOK)
				endpoint(w,r)
			}
		} else {
			response := database.Response{}
			response.Status = false
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(response)
		}
	})
}

func SaveMother(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}
	mother := models.Mother{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		err := json.Unmarshal(data, &mother)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
		} else {
			err := database.SaveMother(mother)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
			} else {
				if mother.Email != "" {
					notification.SendVerificationEmail(mother.Email)
				}

				response.Data = mother
				response.Status = true
				response.StatusCode = http.StatusOK
			}
		}
	}
	RespondWithObject(response, w)
}

func SaveFather(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}
	father := models.Father{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		err := json.Unmarshal(data, &father)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
		} else {
			err := database.SaveFather(father)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
			} else {
				response.Data = father
				response.Status = true
				response.StatusCode = http.StatusOK
			}
		}
	}
	RespondWithObject(response, w)
}

func SaveChild(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}
	child := models.Child{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		err := json.Unmarshal(data, &child)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
		} else {
			err := database.SaveChild(child)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
			} else {
				response.Data = child
				response.Status = true
				response.StatusCode = http.StatusOK
			}
		}
	}
	RespondWithObject(response, w)
}

func CheckNic(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}
	nic := models.NIC{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		err := json.Unmarshal(data, &nic)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
		} else {
			userId, err := database.GetUserFromNic(nic)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
			} else {
				response.Data = userId
				response.Status = true
				response.StatusCode = http.StatusOK
			}
		}
	}

	RespondWithObject(response, w)
}

func GetWeight(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}
	user := models.User{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		err := json.Unmarshal(data, &user)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
		} else {
			weightList, err := database.GetWeightByUserId(user)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
			} else {
				response.Data = weightList
				response.Status = true
				response.StatusCode = http.StatusOK
			}
		}
	}

	RespondWithObject(response, w)
}

func SaveWeight(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}
	weight := models.Weight{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		err := json.Unmarshal(data, &weight)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
		} else {
			err := database.SaveWeight(weight)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
			} else {
				response.Data = "success"
				response.Status = true
				response.StatusCode = http.StatusOK
			}
		}
	}

	RespondWithObject(response, w)
}

func SaveLocation(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}
	geo := models.GeoLocation{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		err := json.Unmarshal(data, &geo)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
		} else {
			err := database.SaveLocation(geo)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
			} else {
				response.Data = "success"
				response.Status = true
				response.StatusCode = http.StatusOK
			}
		}
	}

	RespondWithObject(response, w)
}

func GetChildren(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}
	nic := models.NIC{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		err := json.Unmarshal(data, &nic)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
		} else {
			children, err := database.GetChildrenByNic(nic)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
			} else {
				response.Data = children
				response.Status = true
				response.StatusCode = http.StatusOK
			}
		}
	}

	RespondWithObject(response, w)
}

func GetMother(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}
	user := models.User{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		err := json.Unmarshal(data, &user)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
		} else {
			mother, err := database.GetMother(user.UserId)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
			} else {
				response.Data = mother
				response.Status = true
				response.StatusCode = http.StatusOK
			}
		}
	}

	RespondWithObject(response, w)
}

func GetChild(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}
	user := models.User{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		err := json.Unmarshal(data, &user)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
		} else {
			child, err := database.GetChild(user.UserId)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
			} else {
				response.Data = child
				response.Status = true
				response.StatusCode = http.StatusOK
			}
		}
	}

	RespondWithObject(response, w)
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}
	geo := models.GeoLocation{}

	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		err := json.Unmarshal(data, &geo)
		if err != nil {
			response.Status = false
			response.StatusCode = http.StatusInternalServerError
		} else {
			location, err := database.GetLocation(geo.UserId)
			if err != nil {
				response.Status = false
				response.StatusCode = http.StatusInternalServerError
			} else {
				response.Data = location
				response.Status = true
				response.StatusCode = http.StatusOK
			}
		}
	}

	RespondWithObject(response, w)
}

func GetAllMothers(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}

	mothers, err := database.GetAllMothers()
	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		response.Data = mothers
		response.Status = true
		response.StatusCode = http.StatusOK
	}

	RespondWithObject(response, w)
}

func GetAllChildren(w http.ResponseWriter, r *http.Request) {
	response := database.Response{}

	children, err := database.GetAllChildren()
	if err != nil {
		response.Status = false
		response.StatusCode = http.StatusInternalServerError
	} else {
		response.Data = children
		response.Status = true
		response.StatusCode = http.StatusOK
	}

	RespondWithObject(response, w)
}

func RespondWithObject(response database.Response, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(response.StatusCode)
	_ = json.NewEncoder(w).Encode(response)
}