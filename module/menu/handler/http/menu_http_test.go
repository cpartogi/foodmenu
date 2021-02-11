package http

import (
	"encoding/json"
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

func TestMenuAdd(t *testing.T) {
	type input struct {
		req map[string]interface{}
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
			name: "#1 success insert data",
			expectedInput: input{
				req: map[string]interface{}{
					"menu_detail":  "a",
					"menu_name":    "b",
					"menu_picture": "c",
					"menu_price":   1,
					"menu_type_id": 1,
					"warteg_id":    "d",
				},
			},
			expectedOutput: output{nil, http.StatusCreated},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mnResponse := response.MenuAdd{}
				mnResponse.MenuDetail = "a"
				mnResponse.MenuName = "b"
				mnResponse.MenuPicture = "c"
				mnResponse.MenuPrice = 1
				mnResponse.MenuTypeId = 1
				mnResponse.WartegId = "d"
				mockMenu.
					On("MenuAdd", mock.Anything, mock.Anything).
					Return(mnResponse, nil)
			},
		},
		{
			name: "#2 unprocessable add menu",
			expectedInput: input{
				req: map[string]interface{}{
					"menu_detail":  "a",
					"menu_name":    "b",
					"menu_picture": "c",
					"menu_price":   "1",
					"menu_type_id": 1,
					"warteg_id":    "d",
				},
			},
			expectedOutput: output{nil, http.StatusUnprocessableEntity},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mnResponse := response.MenuAdd{}

				mockMenu.
					On("MenuAdd", mock.Anything, mock.Anything).
					Return(mnResponse, nil)
			},
		},
		{
			name: "#3 bad request add menu",
			expectedInput: input{
				req: map[string]interface{}{
					"menu_detail": "a",
					"menu_name":   "b",
				},
			},
			expectedOutput: output{nil, http.StatusBadRequest},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mnResponse := response.MenuAdd{}

				mockMenu.
					On("MenuAdd", mock.Anything, mock.Anything).
					Return(mnResponse, nil)
			},
		},
		{
			name: "#4 internal server error add menu",
			expectedInput: input{
				req: map[string]interface{}{
					"menu_detail":  "a",
					"menu_name":    "b",
					"menu_picture": "c",
					"menu_price":   1,
					"menu_type_id": 1,
					"warteg_id":    "d",
				},
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mnResponse := response.MenuAdd{}

				mockMenu.
					On("MenuAdd", mock.Anything, mock.Anything).
					Return(mnResponse, errorMenu)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockMenu := new(mocks.Usecase)

			payload, err := json.Marshal(testCase.expectedInput.req)

			assert.NoError(t, err)

			e := echo.New()

			req, err := http.NewRequest(echo.POST, "/v1/menu",
				strings.NewReader(string(payload)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/menu")

			testCase.configureMock(
				testCase.expectedInput,
				mockMenu,
			)

			handler := MenuHandler{
				menuUsecase: mockMenu,
			}

			err = handler.MenuAdd(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}

func TestMenuDelete(t *testing.T) {
	type input struct {
		menu_id string
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
			name: "#1 success delete menu",
			expectedInput: input{
				menu_id: "abc",
			},
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mnResponse := response.MenuDelete{}
				mnResponse.MenuId = "a"

				mockMenu.
					On("MenuDelete", mock.Anything, mock.Anything).
					Return(mnResponse, nil)
			},
		},
		{
			name: "#2 internal server error delete menu",
			expectedInput: input{
				menu_id: "abc",
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mnResponse := response.MenuDelete{}

				mockMenu.
					On("MenuDelete", mock.Anything, mock.Anything).
					Return(mnResponse, errorMenu)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockMenu := new(mocks.Usecase)

			menu_id := testCase.expectedInput.menu_id

			e := echo.New()

			req, err := http.NewRequest(echo.DELETE, "/v1/menu/:menu_id",
				strings.NewReader(string(menu_id)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/menu/")

			testCase.configureMock(
				testCase.expectedInput,
				mockMenu,
			)

			handler := MenuHandler{
				menuUsecase: mockMenu,
			}

			err = handler.MenuDelete(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}

func TestMenuUpdate(t *testing.T) {
	type input struct {
		menu_id string
		req     map[string]interface{}
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
			name: "#1 success update",
			expectedInput: input{
				menu_id: "abc",
				req: map[string]interface{}{
					"menu_detail":  "a",
					"menu_name":    "b",
					"menu_picture": "c",
					"menu_price":   1,
					"menu_type_id": 1,
					"warteg_id":    "d",
				},
			},
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mnResponse := response.MenuUpdate{}

				mockMenu.
					On("MenuUpdate", mock.Anything, mock.Anything).
					Return(mnResponse, nil)
			},
		},
		{
			name: "#2 bad request update",
			expectedInput: input{
				menu_id: "abc",
				req: map[string]interface{}{
					"menu_detail":  "a",
					"menu_name":    "b",
					"menu_picture": "c",
					"menu_price":   1,
					"warteg_id":    "d",
				},
			},
			expectedOutput: output{nil, http.StatusBadRequest},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mnResponse := response.MenuUpdate{}

				mockMenu.
					On("MenuUpdate", mock.Anything, mock.Anything).
					Return(mnResponse, nil)
			},
		},
		{
			name: "#3 unprocessable update",
			expectedInput: input{
				menu_id: "abc",
				req: map[string]interface{}{
					"menu_detail":  "a",
					"menu_name":    "b",
					"menu_picture": "c",
					"menu_price":   "1",
					"menu_type_id": 1,
					"warteg_id":    "d",
				},
			},
			expectedOutput: output{nil, http.StatusUnprocessableEntity},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mnResponse := response.MenuUpdate{}

				mockMenu.
					On("MenuUpdate", mock.Anything, mock.Anything).
					Return(mnResponse, nil)
			},
		},
		{
			name: "#4 internal server error update",
			expectedInput: input{
				menu_id: "abc",
				req: map[string]interface{}{
					"menu_detail":  "a",
					"menu_name":    "b",
					"menu_picture": "c",
					"menu_price":   1,
					"menu_type_id": 1,
					"warteg_id":    "d",
				},
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mnResponse := response.MenuUpdate{}

				mockMenu.
					On("MenuUpdate", mock.Anything, mock.Anything).
					Return(mnResponse, errorMenu)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockMenu := new(mocks.Usecase)

			//	menu_id := testCase.expectedInput.menu_id

			j, err := json.Marshal(testCase.expectedInput.req)

			e := echo.New()

			req, err := http.NewRequest(echo.PUT, "/v1/menu/:menu_id",
				strings.NewReader(string(j)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/menu/")

			testCase.configureMock(
				testCase.expectedInput,
				mockMenu,
			)

			handler := MenuHandler{
				menuUsecase: mockMenu,
			}

			err = handler.MenuUpdate(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}

func TestMenuList(t *testing.T) {
	type input struct {
		warteg_id    string
		menu_type_id int
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
			name: "#1 error menu list",
			expectedInput: input{
				warteg_id:    "asdfsdfsd",
				menu_type_id: 1,
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockMenu *mocks.Usecase,
			) {
				mnResponse := response.MenuList{}

				mockMenu.
					On("MenuList", mock.Anything, mock.Anything).
					Return(mnResponse, nil)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockMenu := new(mocks.Usecase)

			warteg_id := testCase.expectedInput.warteg_id

			e := echo.New()

			req, err := http.NewRequest(echo.GET, "/v1/menus/list",
				strings.NewReader(string(warteg_id)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/menus/list")

			testCase.configureMock(
				testCase.expectedInput,
				mockMenu,
			)

			handler := MenuHandler{
				menuUsecase: mockMenu,
			}

			err = handler.MenuList(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}
