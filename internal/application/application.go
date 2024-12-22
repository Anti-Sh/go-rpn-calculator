package application

import (
	"encoding/json"
	"errors"
	"github.com/Anti-Sh/go-rpn-calculator/internal/config"
	"github.com/Anti-Sh/go-rpn-calculator/pkg/calculator"
	"net/http"
)

type Application struct {
	config *config.Config
}

func NewApplication() *Application {
	return &Application{
		config: config.NewConfigFromEnv(),
	}
}

func (a *Application) RunServer() error {
	mux := http.NewServeMux()

	calculate := http.HandlerFunc(CalcHandler)
	mux.Handle("/api/v1/calculate", UnhandledErrorMiddleware(calculate))

	return http.ListenAndServe(":"+a.config.Port, mux)
}

type (
	CalcRequest struct {
		Expression string `json:"expression"`
	}
	CalcSuccessResponse struct {
		Result float64 `json:"result"`
	}
	CalcErrorResponse struct {
		Error string `json:"error"`
		code  int
	}
)

func NewCalcSuccessResponse(result float64) *CalcSuccessResponse {
	return &CalcSuccessResponse{
		Result: result,
	}
}

func NewCalcErrorResponse(err error) *CalcErrorResponse {
	var (
		msg  string
		code int
	)

	if errors.Is(err, calculator.ErrInvalidExpression) ||
		errors.Is(err, calculator.ErrDivisionByZero) ||
		errors.Is(err, calculator.ErrUnknownToken) ||
		errors.Is(err, calculator.ErrUnknownOperator) {

		msg = "Expression is not valid"
		code = http.StatusUnprocessableEntity
	} else if errors.Is(err, ErrInvalidRequestMethod) {
		msg = "Invalid request method"
		code = http.StatusMethodNotAllowed
	} else {
		msg = "Internal server error"
		code = http.StatusInternalServerError
	}

	return &CalcErrorResponse{
		Error: msg,
		code:  code,
	}
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		HandleCalcError(w, ErrInvalidRequestMethod)
		return
	}

	request := new(CalcRequest)
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		HandleCalcError(w, err)
		return
	}

	calc := calculator.NewCalculator(request.Expression)
	executionResult, err := calc.Execute()
	if err != nil {
		HandleCalcError(w, err)
		return
	}

	res := NewCalcSuccessResponse(executionResult)
	SendJSON(w, http.StatusOK, res)
}

func HandleCalcError(w http.ResponseWriter, err error) {
	res := NewCalcErrorResponse(err)
	SendJSON(w, res.code, res)
}

func UnhandledErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				HandleCalcError(w, ErrInternalServerError)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func SendJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		HandleCalcError(w, err)
	}
}
