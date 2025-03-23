package models

type LoginData struct {
  Email string;
  Password string;
}

type LoginSucess struct{
  Jwt string `json:"jwt"`
}


