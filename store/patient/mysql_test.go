package patient

import (
	"Hospital/model"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
)

func Test_Get(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	store := New(db)
	if err != nil {
		t.Errorf("Error While mocking")
	}
	rows := sqlmock.NewRows([]string{"id", "name", "phone", "discharge", "bloodgroup", "desc"}).AddRow(1, "zopsmart", "965462", true, "A+", "fit")

	TestCases :=
		[]struct {
			desc          string
			id            int
			output        []model.Patient
			mockQuery     interface{}
			expectedError error
		}{
			{
				desc:          "success",
				id:            1,
				output:        []model.Patient{{1, "zopsmart", "965462", true, "A+", "fit"}},
				mockQuery:     mock.ExpectQuery("SELECT * FROM patient WHERE id=?").WithArgs(1).WillReturnRows(rows),
				expectedError: nil,
			},
			{
				desc:          "fail",
				id:            3,
				output:        nil,
				mockQuery:     mock.ExpectQuery("SELECT * FROM patient WHERE id=?").WithArgs(3),
				expectedError: nil,
			},
		}
	for _, testCase := range TestCases {
		resp, err := store.Get(testCase.id)
		if err != nil && err.Error() != testCase.expectedError.Error() {
			t.Errorf("expected error: %v, got: %v", testCase.expectedError, err)
		}
		if !reflect.DeepEqual(resp, testCase.output) {
			t.Errorf("expected user: %v, got: %v", testCase.output, resp)
		}
	}
}
