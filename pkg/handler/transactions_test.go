package handler

import (
//     "bytes"
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

func TestHandler_getAllTransactions(t *testing.T){
    type mockBehavior func(s *mock_service.MockTransactions,input models.AddressReq)

    testTable:=[]struct{
        name string
        Request models.AddressReq
        IncRec string
        mockBehavior mockBehavior
        expectedStatusCode int
        expectedRequestBody string
    }{
        {
            name:"OK",
            Request: models.AddressReq{
                Direction:"up",
                Page: 1,
                Sort:"amount",
                Records:10,
            },
            IncRec:"10",
            mockBehavior: func(s *mock_service.MockTransactions,input models.AddressReq){
                s.EXPECT().GetAllTransactions(input).Return([]models.Transaction{
                    {UserId:2,Type:"type",Amount:1200,Created:"12.12.12"},
                    {UserId:2,Type:"type",Amount:1200,Created:"12.12.12"},
                },nil)
            },
            expectedStatusCode:200,
            expectedRequestBody:`[{"user_id":2,"type":"type","amount":1200,"description":"","created":"12.12.12"},{"user_id":2,"type":"type","amount":1200,"description":"","created":"12.12.12"}]`,
        },
        {
            name:"Response error",
            Request: models.AddressReq{
                Direction:"up",
                Page: 1,
                Sort:"amount",
                Records:10,
            },
            IncRec:"10",
            mockBehavior: func(s *mock_service.MockTransactions,input models.AddressReq){
                s.EXPECT().GetAllTransactions(input).Return([]models.Transaction{},errors.New("some server error"))
            },
            expectedStatusCode:500,
            expectedRequestBody:`{"message":"some server error","status":500}`,
        },
        {
            name:"Page error",
            Request: models.AddressReq{
                Direction:"up",
                Page: 0,
                Sort:"amount",
                Records:10,
            },
            IncRec:"10",
            mockBehavior: func(s *mock_service.MockTransactions,input models.AddressReq){
                s.EXPECT().GetAllTransactions(input).Return([]models.Transaction{
                    {UserId:2,Type:"type",Amount:1200,Created:"12.12.12"},
                    {UserId:2,Type:"type",Amount:1200,Created:"12.12.12"},
                },nil)
            },
            expectedStatusCode:400,
            expectedRequestBody:`{"message":"Page number out of range","status":400}`,
        },
        {
            name:"Input error",
            Request: models.AddressReq{
                Direction:"up",
                Page: 0,
                Sort:"amount",
                Records:10,
            },
            IncRec:"sd",
            mockBehavior: func(s *mock_service.MockTransactions,input models.AddressReq){},
            expectedStatusCode:400,
            expectedRequestBody:`{"message":"Invalid type of data","status":400}`,
        },
    }

    for _,testCase:=range testTable{
        t.Run(testCase.name,func(t *testing.T){
            c:=gomock.NewController(t)
            defer c.Finish()

            trn:=mock_service.NewMockTransactions(c)
            testCase.mockBehavior(trn,testCase.Request)

            services:=&service.Service{Transactions:trn}
            handler:=NewHandler(services)

            //Test server
            r:=gin.New()
            r.GET("/transactions/",handler.getAllTransactions)

            //Test request
            w:=httptest.NewRecorder()
            req:=httptest.NewRequest("GET",fmt.Sprintf("/transactions/?sort=%v&page=%v&records=%v&direction=%v",testCase.Request.Sort,testCase.Request.Page,testCase.IncRec,testCase.Request.Direction),nil)
            r.ServeHTTP(w,req)

            assert.Equal(t,testCase.expectedStatusCode,w.Code)
            assert.Equal(t,testCase.expectedRequestBody,w.Body.String())
        })
    }
}

func TestHandler_getTransactionByUserId(t *testing.T){
    type mockBehavior func(s *mock_service.MockTransactions,id int, input models.AddressReq)
    testTable:=[]struct{
        name string
        Id int
        IncId string
        IncRec string
        Request models.AddressReq
        mockBehavior mockBehavior
        expectedStatusCode int
        expectedRequestBody string
    }{
        {
            name:"OK",
            Id:1,
            IncId:"1",
            Request: models.AddressReq{
                Direction:"up",
                Page: 1,
                Sort:"amount",
                Records:10,
            },
            IncRec:"10",
            mockBehavior: func(s *mock_service.MockTransactions,id int,input models.AddressReq){
                s.EXPECT().GetTransactionByUserId(id,input).Return([]models.Transaction{
                    {UserId:1,Type:"type",Amount:1200,Created:"12.12.12"},
                    {UserId:1,Type:"type",Amount:1200,Created:"12.12.12"},
                },nil)
            },
            expectedStatusCode:200,
            expectedRequestBody:`[{"user_id":1,"type":"type","amount":1200,"description":"","created":"12.12.12"},{"user_id":1,"type":"type","amount":1200,"description":"","created":"12.12.12"}]`,
        },
        {
            name:"Id error",
            Id:1,
            Request: models.AddressReq{
                Direction:"up",
                Page: 1,
                Sort:"amount",
                Records:10,
            },
            IncRec:"10",
            mockBehavior: func(s *mock_service.MockTransactions,id int,input models.AddressReq){},
            expectedStatusCode:400,
            expectedRequestBody:`{"message":"Invalid id","status":400}`,
        },
        {
            name:"Input error",
            Id:1,
            IncId:"1",
            Request: models.AddressReq{
                Direction:"",
                Page: 1,
                Sort:"",
                Records:10,
            },
            IncRec:"d",
            mockBehavior: func(s *mock_service.MockTransactions,id int,input models.AddressReq){},
            expectedStatusCode:400,
            expectedRequestBody:`{"message":"Invalid type of data","status":400}`,
        },
        {
            name:"Page error",
            Id:1,
            IncId:"1",
            Request: models.AddressReq{
                Direction:"up",
                Page: 0,
                Sort:"amount",
                Records:10,
            },
            IncRec:"10",
            mockBehavior: func(s *mock_service.MockTransactions,id int,input models.AddressReq){
                s.EXPECT().GetTransactionByUserId(id,input).Return([]models.Transaction{
                    {UserId:1,Type:"type",Amount:1200,Created:"12.12.12"},
                    {UserId:1,Type:"type",Amount:1200,Created:"12.12.12"},
                },nil)
            },
//             mockBehavior: func(s *mock_service.MockTransactions,id int,input models.AddressReq){
//                 s.EXPECT().GetTransactionByUserId(id,input).Return([]models.Transaction{
//                     {UserId:2,Type:"type",Amount:1200,Created:"12.12.12"},
//                     {UserId:2,Type:"type",Amount:1200,Created:"12.12.12"},
//                 },nil)
//             },
            expectedStatusCode:400,
            expectedRequestBody:`{"message":"Page number out of range","status":400}`,
        },
        {
            name:"Response error",
            Id:1,
            IncId:"1",
            Request: models.AddressReq{
                Direction:"up",
                Page: 1,
                Sort:"amount",
                Records:10,
            },
            IncRec:"10",
            mockBehavior: func(s *mock_service.MockTransactions,id int,input models.AddressReq){
                s.EXPECT().GetTransactionByUserId(id,input).Return([]models.Transaction{},errors.New("some server error"))
            },
            expectedStatusCode:500,
            expectedRequestBody:`{"message":"some server error","status":500}`,
        },
    }

    for _,testCase:=range testTable{
        t.Run(testCase.name,func(t *testing.T){
            c:=gomock.NewController(t)
            defer c.Finish()
            trn:=mock_service.NewMockTransactions(c)
            testCase.mockBehavior(trn,testCase.Id,testCase.Request)

            services:=&service.Service{Transactions:trn}
            handler:=NewHandler(services)

            //Test server
            r:=gin.New()
            r.GET("/transactions/:id",handler.getTransactionByUserId)

            //Test request
            w:=httptest.NewRecorder()
            req:=httptest.NewRequest("GET",fmt.Sprintf("/transactions/%v?sort=%v&page=%v&records=%v&direction=%v",testCase.IncId,testCase.Request.Sort,testCase.Request.Page,testCase.IncRec,testCase.Request.Direction),nil)
            r.ServeHTTP(w,req)

            assert.Equal(t,testCase.expectedStatusCode,w.Code)
            assert.Equal(t,testCase.expectedRequestBody,w.Body.String())
        })
    }
}
