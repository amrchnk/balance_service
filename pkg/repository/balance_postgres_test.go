package repository

import (
    "testing"
    "github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
    "github.com/amrchnk/balance_service/pkg/models"
//     "errors"
)

func TestBalance_ChangeUserBalance(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

    r := NewBalancePostgres(db)

    type args struct {
        balance models.Balance
        tr_type string
    }

	tests := []struct {
		name    string
		mock    func()
		input   args
		want    string
		wantErr bool
	}{
        {
			name: "Ok",
			mock: func() {
			    mock.ExpectBegin()
				rows := sqlmock.NewRows([]string{"balance"}).AddRow(230.45)
				mock.ExpectQuery("INSERT INTO balance").
					WithArgs(1,230.45).WillReturnRows(rows)
			    mock.ExpectCommit()
			},
			input: args{
				balance: models.Balance{
				    UserId:1,
				    Balance:230.45,
				},
				tr_type: "increase",
			},
			want: "230.45",
		},
//         {
// 			name: "Empty fields",
// 			input: args{
//                 balance: models.Balance{
//                     UserId:1,
//                 },
//                 tr_type: "increase",
//             },
// 			mock: func() {
//
// 			    mock.ExpectBegin()
// 				rows := sqlmock.NewRows([]string{"balance"}).AddRow().RowError(0, errors.New("insert error"))
// 				mock.ExpectQuery("INSERT INTO balance").
// 					WithArgs(1,0).WillReturnRows(rows)
// 			    mock.ExpectRollback()
// 			},
//
// 			wantErr: true,
// 		},
	}

    for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.ChangeUserBalance(tt.input.balance,tt.input.tr_type)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}