package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/echo-lin/puffer/Common"
	"github.com/echo-lin/puffer/models"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

func AllAccounts() (accounts []models.Account, err error) {
	filepath, err := userPath()
	if err != nil {
		return nil, err
	}
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil
	}

	reader := bufio.NewReader(file)
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read file, error code: %v", Common.ErrCode["failedToReadAccountFile"])
	}
	json.Unmarshal(content, &accounts)
	return accounts, nil
}

func AddAccounts(newAccount models.Account) error {
	filepath, err := userPath()
	if err != nil {
		return err
	}
	// 如果文件已存在，则读取已存在的内容
	if _, err := os.Stat(filepath); err != nil {
		// 账号存储文件不存在，则创建配置文件
		// 创建文件
		file, err := os.Create(filepath)
		if err != nil {
			return fmt.Errorf("error creating file: %v", err)
		}

		// 设置文件权限，只允许root用户访问
		err = file.Chmod(0600) // 0600表示只有所有者可以读写
		if err != nil {
			return fmt.Errorf("error permission: %v", err)
		}
	}

	accountList, err := AllAccounts()
	if err != nil {
		return fmt.Errorf("failed to query account list, error code: %v", Common.ErrCode["failedToQueryAccountList"])
	}

	flag := false
	for _, val := range accountList {
		if val.Username == newAccount.Username && val.Domain == newAccount.Domain {
			flag = true
			break
		}
	}
	if flag {
		return fmt.Errorf("account already exists in domain %v, error code: %v", newAccount.Domain, Common.ErrCode["failedToQueryAccountList"])
	}
	accountList = append(accountList, newAccount)
	accountListJson, err := json.Marshal(accountList)
	if err != nil {
		return fmt.Errorf("in the process of adding accounts, the JSON serialization of the account list failed, error code: %v", Common.ErrCode["failedToJsonAccountList"])
	}

	file, err := os.OpenFile(filepath, os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open file, error code: %v", Common.ErrCode["failedToOpenAccountFile"])
	}
	_, err = file.WriteString(string(accountListJson))
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func DeleteAccount(account models.Account) error {
	accountList, err := AllAccounts()
	if err != nil {
		return fmt.Errorf("failed to query account list, error code: %v", Common.ErrCode["failedToQueryAccountList"])
	}

	exist := false
	var accountIndex int
	for index, val := range accountList {
		if val.Username == account.Username && val.Domain == account.Domain {
			accountIndex = index
			exist = true
			break
		}
	}
	if !exist {
		return fmt.Errorf("the account [%v] is not exist in domain %v", account.Username, account.Domain)
	}

	newAccountList := append(accountList[:accountIndex], accountList[accountIndex+1:]...)
	newAccountListJson, err := json.Marshal(newAccountList)
	if err != nil {
		return fmt.Errorf("in the process of adding accounts, the JSON serialization of the account list failed, error code: %v", Common.ErrCode["failedToJsonAccountList"])
	}

	filepath, err := userPath()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to open file, error code: %v", Common.ErrCode["failedToOpenAccountFile"])
	}
	_, err = file.WriteString(string(newAccountListJson))
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func userPath() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve user path: %v")
	}
	filepath := filepath.Join(currentUser.HomeDir, ".puffer.json")
	return filepath, nil
}
