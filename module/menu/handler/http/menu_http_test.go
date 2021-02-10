package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cpartogi/foodmenu/module/menu/mocks"
	"github.com/cpartogi/foodmenu/schema/response"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var errorMenu = errors.New("error menu")

func TestMenuHandlerNewMenuHandler(t *testing.T) {
	e := echo.New()
	mockMenu := new(mocks.Usecase)
	NewMenuHandler(e, mockMenu)
}

func TestMenuType(t *testing.T) {
	type input struct {
		transaction_id string
		date_from      string
		date_to        string
	}

	type output struct {
		err        error
		statusCode int
	}
	cases := []struct {
		name           string
		expectedInput  input
		expectedOutput output
		configureMock  func(
			payload input,
			mockMenu *mocks.Usecase,
		)
	}{
		{
			name: "#1 success menu type",
			expectedInput: input{
				transaction_id: "31123121",
				date_from:      "2021-02-01",
				date_to:        "2021-02-05",
			},
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mtResponse := []response.MenuType{}
				mockMenu.
					On("MenuType", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(mtResponse, nil)
			},
		},
		{
			name: "#2 internal server error menu type",
			expectedInput: input{
				transaction_id: "423423",
				date_from:      "2021-02-01",
				date_to:        "2021-02-05",
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mtResponse := []response.MenuType{}
				mockMenu.
					On("MenuType", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(mtResponse, errorMenu)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockMenu := new(mocks.Usecase)

			transaction_id := testCase.expectedInput.transaction_id

			e := echo.New()

			req, err := http.NewRequest(echo.GET, "/v1/menus/typelist",
				strings.NewReader(string(transaction_id)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/menus/typelist")

			testCase.configureMock(
				testCase.expectedInput,
				mockMenu,
			)

			handler := MenuHandler{
				menuUsecase: mockMenu,
			}

			err = handler.MenuType(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}
