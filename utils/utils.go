package utils

import (
	"NetworkList/models"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"os/exec"
	"time"
)

func RunCommand(AccountList []models.Account) bool {
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

func Login(account *models.Account) bool {
	cmdPrefix := `curl -s --connect-timeout 10 --header \
	"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9" --compressed \
	--header "Accept-Language: en-US,en;q=0.9" --header "Cache-Control: max-age=0" --header "Connection: keep-alive" --header "Origin: http://172.20.20.1:801" \
	--header "Referer: http://172.20.20.1:801/srun_portal_pc.php?ac_id=3&" --header "Upgrade-Insecure-Requests: 1" \
	--user-agent "Mozilla/5.0 (Windows; U; Windows NT 4.0) AppleWebKit/533.43.4 (KHTML, like Gecko) Version/4.0.5 Safari/533.43.4" \
	--data-binary "action=login&ac_id=3&user_ip=&nas_ip=&user_mac=&url=&username=`
	cmdMiddle := `&password=`
	cmdPostfix := `" "http://172.20.20.1:801/srun_portal_pc.php?ac_id=3&" > /dev/null`
	runCmd := cmdPrefix + account.Account + cmdMiddle + account.Password + cmdPostfix
	logs.Info("running command: ", runCmd)
	_, err := exec.Command("/bin/sh", "-c", runCmd).Output()
	if err != nil {
		logs.Info(err)
		return false
	}
	// test connection
	logs.Info("Checking Internet Connection...")
	if NetWorkStatus() {
		logs.Info("Internet Connection OK")
		// fmt.Println("You are using DLUT-EDA provided by", account.Owner)
		return true
	} else {
		logs.Info("Internet Connection Failed")
		return false
	}
}

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
