package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var FILE_DIR = "./DLUT-EDA.json"

type Account struct {
	Account  string
	Password string
	Owner    string
}

// type Account struct {
// 	Account  string `json:"account"`
// 	Password string `json:"password"`
// 	Owner    string `json:"owner"`
// }

func main() {
	// s := []Account{{"201992222", "06133017", "Augists"}, {"201992223", "06133017", "Augists"}}
	// writeJson(s)
	// fmt.Println(readJson())

	// cmd := exec.Command("whoami")
	// out, err := cmd.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(string(out)[:len(string(out))-1])

	var AccountList = []Account{}
	AccountList = readJson()
	runCommand(AccountList)

	// http.HandleFunc("/", handler)
	// http.HandleFunc("/add", addAccount)
	// fmt.Println("Server running at http://localhost:8080")
	// log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

/*
 * handler:
 * 		menu page
 * function:
 * 		1. add account -> /add
 * 		2. delete account -> /delete
 * 		3. modify account -> /modify
 * 		4. log in and redirect to selected page -> /login
 * 		5. exit webpage and send kill signal
 */
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// func runCommand(cmd string) (bool, error) {
// 	fmt.Println("run command: " + cmd)
// 	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
// 	if err != nil {
// 		fmt.Println(err)
// 		return false, err
// 	}
// 	fmt.Println(string(out))
// 	return true, nil
// }

func runCommand(AccountList []Account) bool {
	cmdPrefix := `curl -s --connect-timeout 10 --header \
	"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9" --compressed \
	--header "Accept-Language: en-US,en;q=0.9" --header "Cache-Control: max-age=0" --header "Connection: keep-alive" --header "Origin: http://172.20.20.1:801" \
	--header "Referer: http://172.20.20.1:801/srun_portal_pc.php?ac_id=3&" --header "Upgrade-Insecure-Requests: 1" \
	--user-agent "Mozilla/5.0 (Windows; U; Windows NT 4.0) AppleWebKit/533.43.4 (KHTML, like Gecko) Version/4.0.5 Safari/533.43.4" \
	--data-binary "action=login&ac_id=3&user_ip=&nas_ip=&user_mac=&url=&username=`
	cmdMiddle := `&password=`
	cmdPostfix := `" "http://172.20.20.1:801/srun_portal_pc.php?ac_id=3&" > /dev/null`
	for _, account := range AccountList {
		runCmd := cmdPrefix + account.Account + cmdMiddle + account.Password + cmdPostfix
		fmt.Println("running command: ", runCmd)
		_, err := exec.Command("/bin/sh", "-c", runCmd).Output()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// test connection
		fmt.Println("Checking Internet Connection...")
		if NetWorkStatus() {
			fmt.Println("Internet Connection OK")
			fmt.Println("You are using DLUT-EDA provided by", account.Owner)
			return true
		} else {
			fmt.Println("Internet Connection Failed")
		}
	}
	fmt.Println("Every account doesn't work fine\nPlease update your json file and Try again")
	return false
}

/*
 * NetWorkStatus:
 * 		test network connection
 */
func NetWorkStatus() bool {
	cmd := exec.Command("ping", "baidu.com", "-c", "1", "-W", "5")
	fmt.Println("NetWorkStatus Start:", time.Now().Unix())
	err := cmd.Run()
	fmt.Println("NetWorkStatus End  :", time.Now().Unix())
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		fmt.Println("Net Status , OK")
	}
	return true
}

func addAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
	accounts := readJson()
	accounts = append(accounts, account)
	writeJson(accounts)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func writeJson(s interface{}) {
	jsonFile, err := os.Create(FILE_DIR)

	if err != nil {
		fmt.Println("Error creating file")
		fmt.Println(err)
		os.Exit(1)
	}

	defer jsonFile.Close()
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(s)
	if err != nil {
		fmt.Println("Error encoding JSON")
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("Successfully encoded JSON")
	}
}

func readJson() []Account {
	jsonFile, err := os.Open(FILE_DIR)

	if err != nil {
		fmt.Println("Error opening file")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var accounts []Account
	json.Unmarshal(byteValue, &accounts)
	return accounts
}

// func (account *Account) String() string {
// 	return fmt.Sprintf("%s,%s,%s", account.Account, account.Password, account.Owner)
// }
