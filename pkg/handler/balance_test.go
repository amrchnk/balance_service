package handler

import (
    "bytes"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/amrchnk/balance_service/pkg/models"
    "github.com/amrchnk/balance_service/pkg/service"
    mock_service "github.com/amrchnk/balance_service/pkg/service/mocks"
    "net/http/httptest"
    "github.com/golang/mock/gomock"
    "errors"
    "fmt"
)

func TestHandler_changeUserBalance(t *testing.T){
    type mockBehavior func(s *mock_service.MockBalance,balance models.Balance, tr_type string)

    testTable:=[]struct{
        name string
        inputBody string
        tr_type string
        inputBalance models.Balance
        mockBehavior mockBehavior
        expectedStatusCode int
        expectedRequestBody string
    }{
        {
            name:"OK",
            inputBody:`{"user_id":2,"balance":50}`,
            tr_type:"increase",
            inputBalance: models.Balance{
                UserId:2,
                Balance:50,
            },
            mockBehavior: func(s *mock_service.MockBalance,balance models.Balance, tr_type string){
                s.EXPECT().ChangeUserBalance(balance,tr_type).Return("50",nil)
            },
            expectedStatusCode:200,
            expectedRequestBody:`{"id":2,"message":"Current balance in rubles: 50","status":200}`,
        },
        {
            name:"Empty fields",
            inputBody:`{"user_id":2,"balance":50}`,
            mockBehavior: func(s *mock_service.MockBalance,balance models.Balance, tr_type string){},
            expectedStatusCode:400,
            expectedRequestBody:`{"error":"Unknown type of operation","status":400}`,
        },
        {
            name:"Binding error",
            inputBody:`{"user_id":"d","balance":50}`,
            tr_type:"increase",
            mockBehavior: func(s *mock_service.MockBalance,balance models.Balance, tr_type string){},
            expectedStatusCode:400,
            expectedRequestBody:`{"error":"invalid data in body","status":400}`,
        },
        {
            name:"Response error",
            tr_type:"increase",
            inputBody:`{"user_id":2,"balance":50}`,
            inputBalance: models.Balance{
                UserId:2,
                Balance:50,
            },
            mockBehavior: func(s *mock_service.MockBalance,balance models.Balance, tr_type string){
                s.EXPECT().ChangeUserBalance(balance,tr_type).Return("", errors.New("some server error"))
            },
            expectedStatusCode:500,
            expectedRequestBody:`{"message":"some server error","status":500}`,
        },
    }

    for _,testCase:=range testTable{
        t.Run(testCase.name,func(t *testing.T){
            c:=gomock.NewController(t)
            defer c.Finish()

            bal:=mock_service.NewMockBalance(c)
            testCase.mockBehavior(bal,testCase.inputBalance,testCase.tr_type)

            services:=&service.Service{Balance:bal}
            handler:=NewHandler(services)

            //Test server
            r:=gin.New()
            r.GET("/balance/:type",handler.changeUserBalance)

            //Test request
            w:=httptest.NewRecorder()
            req:=httptest.NewRequest("GET",fmt.Sprintf("/balance/%v",testCase.tr_type),bytes.NewBufferString(testCase.inputBody))
            r.ServeHTTP(w,req)

            assert.Equal(t,testCase.expectedStatusCode,w.Code)
            assert.Equal(t,testCase.expectedRequestBody,w.Body.String())
        })
    }
}

func TestHandler_getBalanceById(t *testing.T){
    type mockBehavior func(s *mock_service.MockBalance,input models.UserBalanceQuery)

    testTable:=[]struct{
        name string
        Id string
        Currency string
        inputQuery models.UserBalanceQuery
        Response models.UserBalanceResponse
        mockBehavior mockBehavior
        expectedStatusCode int
        expectedRequestBody string
    }{
        {
            name:"OK",
            Id:"2",
            Currency:"USD",
            inputQuery: models.UserBalanceQuery{
                UserId:2,
                Currency:"USD",
            },
            mockBehavior: func(s *mock_service.MockBalance,input models.UserBalanceQuery){
                s.EXPECT().GetBalanceById(input).Return(models.UserBalanceResponse{
                    UserId:2,
                    Currency:"USD",
                    Balance:30,
                },nil)
            },
            expectedStatusCode:200,
            expectedRequestBody:`{"user_id":2,"currency":"USD","balance":30}`,
        },
        {
            name:"Id error",
            Id:"sd",
            Currency:"USD",
            mockBehavior: func(s *mock_service.MockBalance,input models.UserBalanceQuery){},
            expectedStatusCode:400,
            expectedRequestBody:`{"message":"Invalid id","status":400}`,
        },
        {
            name:"Response error",
            Id:"2",
            Currency:"USD",
            inputQuery: models.UserBalanceQuery{
                UserId:2,
                Currency:"USD",
            },
            mockBehavior: func(s *mock_service.MockBalance,input models.UserBalanceQuery){
                s.EXPECT().GetBalanceById(input).Return(models.UserBalanceResponse{}, errors.New("some server error"))
            },
            expectedStatusCode:500,
            expectedRequestBody:`{"message":"some server error","status":500}`,
        },
    }

    for _,testCase:=range testTable{
        t.Run(testCase.name,func(t *testing.T){
            c:=gomock.NewController(t)
            defer c.Finish()

            bal:=mock_service.NewMockBalance(c)
            testCase.mockBehavior(bal,testCase.inputQuery)

            services:=&service.Service{Balance:bal}
            handler:=NewHandler(services)

            //Test server
            r:=gin.New()
            r.GET("/balance/:id",handler.getBalanceById)

            //Test request
            w:=httptest.NewRecorder()
            req:=httptest.NewRequest("GET",fmt.Sprintf("/balance/%v?currency=%v",testCase.Id,testCase.Currency),nil)
            r.ServeHTTP(w,req)

            assert.Equal(t,testCase.expectedStatusCode,w.Code)
            assert.Equal(t,testCase.expectedRequestBody,w.Body.String())
        })
    }
}

func TestHandler_transferMoney(t *testing.T){
    type mockBehavior func(s *mock_service.MockBalance,input models.TransferQuery)

    testTable:=[]struct{
        name string
        inputBody string
        Query models.TransferQuery
        mockBehavior mockBehavior
        expectedStatusCode int
        expectedRequestBody string
    }{
        {
            name:"OK",
            inputBody:`{"sender_id":2,"receiver_id":1,"sum":230}`,
            Query: models.TransferQuery{
                SenderId:2,
                ReceiverId:1,
                Sum:230,
            },
            mockBehavior: func(s *mock_service.MockBalance,input models.TransferQuery){
                s.EXPECT().TransferMoney(input).Return([]float64{23,420},nil)
            },
            expectedStatusCode:200,
            expectedRequestBody:`{"sender_id":2,"sender_balance":23,"receiver_id":1,"receiver_balance":420}`,
        },
        {
            name:"Binding error",
            inputBody:`{"sender_id":"2","receiver_id":1,"sum":230}`,
            mockBehavior: func(s *mock_service.MockBalance,input models.TransferQuery){},
            expectedStatusCode:400,
            expectedRequestBody:`{"message":"invalid data in body","status":400}`,
        },
        {
            name:"Response error",
            inputBody:`{"sender_id":2,"receiver_id":1,"sum":230}`,
            Query: models.TransferQuery{
                SenderId:2,
                ReceiverId:1,
                Sum:230,
            },
            mockBehavior: func(s *mock_service.MockBalance,input models.TransferQuery){
                s.EXPECT().TransferMoney(input).Return([]float64{},errors.New("some server error"))
            },
            expectedStatusCode:500,
            expectedRequestBody:`{"message":"some server error","status":500}`,
        },
    }

    for _,testCase:=range testTable{
        t.Run(testCase.name,func(t *testing.T){
            c:=gomock.NewController(t)
            defer c.Finish()

            bal:=mock_service.NewMockBalance(c)
            testCase.mockBehavior(bal,testCase.Query)

            services:=&service.Service{Balance:bal}
            handler:=NewHandler(services)

            //Test server
            r:=gin.New()
            r.POST("/balance/transfer",handler.transferMoney)

            //Test request
            w:=httptest.NewRecorder()
            req:=httptest.NewRequest("POST","/balance/transfer",
            bytes.NewBufferString(testCase.inputBody))
            r.ServeHTTP(w,req)

            assert.Equal(t,testCase.expectedStatusCode,w.Code)
            assert.Equal(t,testCase.expectedRequestBody,w.Body.String())
        })
    }
}