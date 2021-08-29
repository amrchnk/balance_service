package handler

import (

)

func ValidateType (str string) bool{
    if str=="increase"||str=="decrease"{
        return true
    } else{
        return false
    }
}