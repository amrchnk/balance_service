package models

import (

)

type Balance struct{
    Id int `json:"-"`
    UserId int `json:"user_id" db:"user_id"`
    Balance float64 `json:"balance"`
}

type TransferQuery struct{
    SenderId int `json:"sender_id"`
    ReceiverId int `json:"receiver_id"`
    Sum float64 `json:sum`
}

type TransferResponse struct{
    SenderId int `json:"sender_id"`
    SenderSum float64 `json:"sender_sum"`
    ReceiverId int `json:"receiver_id"`
    ReceiverSum float64 `json:"receiver_sum"`
}
