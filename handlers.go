package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

func createWalletHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	status := true

	address, walletName, err := createArchwayWallet()
	if err != nil {
		status = false
	}
	json.NewEncoder(w).Encode(GorchWallet{Status: status, WalletName: walletName, Address: address})
}

func getWalletListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	status := true

	list, err := getArchwayWalletList()
	if err != nil {
		status = false
	}

	walletAddressRegx := regexp.MustCompile(`name: \"([\w].*)"  type: local  address: (archway[\w].*)`)
	walletAddressResult := walletAddressRegx.FindAllStringSubmatch(list, -1)

	fmt.Println(list)

	fmt.Println(walletAddressResult)

	json.NewEncoder(w).Encode(GorchWallet{Status: status})
}
