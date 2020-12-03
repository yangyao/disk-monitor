package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func main() {

	// 1. 运行 shell 命令 blkid UUID  uuid
	// 2. 解析 shell 运行结果 得到 UUID
	// 3. 看看自己的 uuid 是否在所有的 uuid 中，
	// 4. 如果UUID不在的话发送 server 酱的推送

	cmd := exec.Command("blkid")
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatalln("run blkid command failed")
	}
	blkidSlice := strings.Split(string(stdout), "\t")
	err = godotenv.Load()
	if err != nil {
		log.Fatalln("load env field failed")
	}
	var match = `UUID="` + os.Getenv("UUID") + `"`

	found := false

	for _, line := range blkidSlice {
		blkid := strings.Split(line, " ")
		for _, each := range blkid {
			if each == match {
				found = true
				continue
			}
		}
	}
	if !found {
		fmt.Println("disk not found")
		_, err := http.PostForm(fmt.Sprintf("https://sc.ftqq.com/%s.send", os.Getenv("SC_KEY")), url.Values{
			"text": {"硬盘挂了"},
			"desp": {"重启下服务器？不行插拔下硬盘试试。"},
		})
		if err != nil {
			log.Fatalln("send wechat message failed")
		}
	}

}
