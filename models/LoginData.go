package models

type LoginData struct {
  Email string;
  Password string;
}
type SignUpData struct{
  Name string;
  Age int;
  Email string;
  Password string;
}
type LoginSucess struct{
  Jwt string `json:"jwt"`
}

type SignUpSucess struct{
  Status string  `json:"status"`
}
