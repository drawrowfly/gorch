package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"regexp"
	"strconv"
)

func createArchwayWallet(home string) (string, string, error) {

	randomWalletName := strconv.Itoa(rand.Intn(10000000-1+1) + 1)
	cmd, err := exec.Command("bash", "-c", "echo -e 'password\npassword' | archwayd keys add "+randomWalletName+"  --home ~/"+home).Output()

	if err != nil {
		return "", "", err
	}
	output := string(cmd)
	walletAddressRegx := regexp.MustCompile(`address: (archway[\w].*)`)
	walletAddressResult := walletAddressRegx.FindAllStringSubmatch(output, -1)

	return walletAddressResult[0][1], randomWalletName, nil
}

func getArchwayWalletList(home string) ([]WalletList, error) {
	cmd, err := exec.Command("bash", "-c", "echo 'password' | archwayd keys list --home ~/"+home).Output()

	if err != nil {
		return []WalletList{}, err
	}
	output := string(cmd)

	walletAddressRegx := regexp.MustCompile(`- name: \"([\w].*)"\n.*\n.*address: (archway[\w].*)`)
	walletAddressResult := walletAddressRegx.FindAllStringSubmatch(output, -1)

	var walletList = []WalletList{}

	for _, k := range walletAddressResult {
		walletList = append(walletList, WalletList{
			WalletName: k[1],
			Address:    k[2],
		})
	}

	return walletList, nil
}

func deleteArchwayWallet(name string, home string) (string, error) {
	cmd, err := exec.Command("bash", "-c", "echo 'password' | archwayd keys delete "+name+" -y --home ~/"+home).Output()

	if err != nil {
		return "", err
	}
	output := string(cmd)

	return output, nil
}

func farmArchwayWallet(home string) (string, error) {

	walletList, err := getArchwayWalletList(home)
	if err != nil {
		return "", err
	}

	for _, k := range walletList {

		balance, err := getWalletBalance(k.Address)
		if err != nil {
			continue
		}

		_, err = exec.Command("bash", "-c", "echo 'password' | archwayd tx bank send "+k.Address+" archway1tqr8wagu7zxy0sc5lk8js04qpydm0tzslvr7dg "+balance+"utorii --chain-id torii-1 -y --home ~/"+home).Output()

		if err != nil {
			continue
		}

		deleteArchwayWallet(k.WalletName, home)
		fmt.Println("DONE:", k.Address, balance)
	}

	return "", nil
}

func getWalletBalance(wallet string) (string, error) {

	cmd, err := exec.Command("bash", "-c", "archwayd q bank balances "+wallet).Output()

	if err != nil {
		return "", err
	}

	output := string(cmd)

	walletBalanceRegx := regexp.MustCompile(`- amount.*"([0-9]+)"`)
	walletBalanceResult := walletBalanceRegx.FindAllStringSubmatch(output, -1)

	return walletBalanceResult[0][1], nil
}
