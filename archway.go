package main

import (
	"math/rand"
	"os/exec"
	"regexp"
	"strconv"
)

func createArchwayWallet() (string, string, error) {

	randomWalletName := strconv.Itoa(rand.Intn(10000000-1+1) + 1)
	cmd, err := exec.Command("bash", "-c", "echo -e 'password\npassword' | archwayd keys add "+randomWalletName+"  --home ~/w1").Output()

	if err != nil {
		return "", "", err
	}
	output := string(cmd)
	walletAddressRegx := regexp.MustCompile(`address: (archway[\w].*)`)
	walletAddressResult := walletAddressRegx.FindAllStringSubmatch(output, -1)

	return walletAddressResult[0][1], randomWalletName, nil
}

func getArchwayWalletList() (string, error) {
	cmd, err := exec.Command("bash", "-c", "echo 'password' | archwayd keys list --home ~/w1").Output()

	if err != nil {
		return "", err
	}
	output := string(cmd)
	// walletAddressRegx := regexp.MustCompile(`address: (archway[\w].*)`)
	// walletAddressResult := walletAddressRegx.FindAllStringSubmatch(output, -1)

	return output, nil
}
