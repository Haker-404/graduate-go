package service

import (
	"awesomeProject/userservice/pkg/model"
	"bufio"
	"flag"
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	// searcher是协程安全的
	searcher = engine.Engine{}
	users    = map[string]model.User{}
	userInfo = flag.String("user_info", "../third_party/user_info", "学生信息")
)

func SearchRun() {
	// 初始化
	searcher.Init(types.EngineInitOptions{})
	defer searcher.Close()
	file, err := os.Open(*userInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "||||")
		user := model.User{}
		user.Seq = data[0]
		user.Name = data[1]
		users[user.Seq] = user
	}
	for seq, userItem := range users {
		seq, _ := strconv.ParseUint(seq, 10, 64)
		searcher.IndexDocument(seq, types.DocumentIndexData{
			Content: userItem.Name,
		}, false)
	}
	searcher.FlushIndex()

}
