package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type GorchCreateWalletResponse struct {
	Status     bool   `json:"status"`
	WalletName string `json:"wallet_name"`
	Address    string `json:"address"`
}

type WalletList struct {
	WalletName string `json:"wallet_name"`
	Address    string `json:"address"`
}

func createWalletHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	// Get --home option
	vars := mux.Vars(r)
	home := vars["home"]

	status := true

	address, walletName, err := createArchwayWallet(home)
	if err != nil {
		status = false
	}
	json.NewEncoder(w).Encode(GorchCreateWalletResponse{Status: status, WalletName: walletName, Address: address})
}

func getWalletListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	// Get --home option
	vars := mux.Vars(r)
	home := vars["home"]

	walletList, err := getArchwayWalletList(home)
	if err != nil {
		json.NewEncoder(w).Encode(walletList)
		return
	}

	json.NewEncoder(w).Encode(walletList)
}
