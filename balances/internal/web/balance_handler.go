package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	getbalance "github.com.br/joaoguimb/fc-ms-wallet/balances/internal/usecase/get_balance"
)

type WebBalanceHandler struct {
	GetBalanceUseCase getbalance.GetBalanceUseCase
}

func NewWebBalanceHanlder(getBalanceUseCase getbalance.GetBalanceUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		GetBalanceUseCase: getBalanceUseCase,
	}
}

func (wh *WebBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountId := r.URL.Query().Get("account_id")
	fmt.Println(accountId)

	output, err := wh.GetBalanceUseCase.Execute(&getbalance.GetBalanceInputDTO{
		AccountID: accountId,
	})
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
