package models

type Transaction struct{
    UserId int `json:"user_id" db:"user_id"`
    Type string `json:"type" db:"type_t"`
    Amount float64 `json:"amount" db:"amount"`
    Description string `json:"description" db:"description"`
    Created string `json:"created" db:"created"`
}

type AddressReq struct{
    Direction string
    Page int
    Sort string
    Records int
}