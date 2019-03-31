package controlers

/**
* @file sitesControler.go
* @brief 根据配置信息批量创建数据库和建立站点数据
* @author frankie@gmail.com
* @version v1.0
* @date 2019-03-11
 */

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os/exec"
	"path/filepath"
	"strings"

	"../common"
	"../database"
	"../definition"
	"../models"
	"github.com/jinzhu/gorm"
)

//GetDbByHost 根据域名获取数据库名字
func GetDbByHost(hostName string) string {
	dbname := strings.Replace(hostName, `.`, `_`, -1)
	return dbname
}

//ParseSistes 解析建站配置文件
func CreateDatabases() {
	fmt.Println("create database")
	siteConfig := getConf()
	keys := make([]definition.SiteConfigKey, 0)
	err := json.Unmarshal([]byte(siteConfig), &keys)
	if err == nil {
		fmt.Printf("%+v\n", keys)
	} else {
		fmt.Println(err)
		fmt.Printf("%+v\n", keys)
	}
	//批量创建数据库
	var dbCtreateStr string
	var dbName string
	for _, key := range keys {
		fmt.Println("create database for ", key.Domain)
		dbName = GetDbByHost(key.Domain)
		//创建数据库
		dbCtreateStr = "mysql -uroot -p79d75802a9d22ec2 -e " + `"create database IF NOT EXISTS ` + dbName + ` DEFAULT CHARSET utf8 COLLATE utf8_general_ci"`
		fmt.Println("createDb:", dbCtreateStr)
		exec.Command("bash", "-c", dbCtreateStr).CombinedOutput()
	}
	//倒入网站数据

	//倒入版块信息

	//倒入标签
	//倒入图片
	//倒入推荐数据
	//倒入友情链接
	fmt.Println("create database end")

}

//initForum 初始化版块信息
func initForum() bool {
	forumRes := `resource/forums.txt`
	var forums []string
	data, err := ioutil.ReadFile(forumRes)
	if err != nil {
		common.Log("get config fatal:" + err.Error())
		return false
	}
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		forums = append(forums, line)
	}
	newForums, err := Random(forums, 5)
	if err != nil {
		common.Log1("random forum err:", err)
		return false
	}
	fmt.Println("forms:", newForums)
	var forum definition.Forum
	//写入数据库
	for _, value := range newForums {

		_, find := models.GetForumByName(value)
		if !find {
			forum.ForumName = value
			models.CreateForumMode(forum)
		}
	}
	return true
}

//getConf 获取站点配置信息
func getConf() string {
	path := "/etc/sitesConfig.ini"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		common.Log("get config fatal:" + err.Error())
		return ""
	}
	return string(data)
}

//对slice进行洗牌，随机获取
func Random(strings []string, length int) ([]string, error) {
	var res []string //需要返回对结果
	if len(strings) <= 0 {
		return nil, errors.New("the length of the parameter strings should not be less than 0")
	}

	if length <= 0 || len(strings) <= length {
		return nil, errors.New("the size of the parameter length illegal")
	}

	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	for i := 0; i < length; i++ {
		res = append(res, strings[i])
	}
	return res, nil
}

func InitSiteConifg(key definition.SiteConfigKey) {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		return
	}
	var siteConfig definition.SiteConfigDb
	logo, banners := GetImageForSite()
	database := GetDbByHost(key.Domain)
	//检查是否存在配置表记录，如果不存在则需要创建
	//倒入网站kdhe
	db.Model(&siteConfig).Update(definition.SiteConfigDb{Domain: key.Domain, Database: database, Title: key.Title, Description: key.Description, Keywords: key.Keywords, Temple: key.Temple, Logo: logo, Banner: banners})

}

//ParseSistes 解析建站配置文件
func CreateSitesConfig() {
	siteConfig := getConf()
	keys := make([]definition.SiteConfigKey, 0)
	err := json.Unmarshal([]byte(siteConfig), &keys)
	if err == nil {
		fmt.Printf("%+v\n", keys)
	} else {
		fmt.Println(err)
		fmt.Printf("%+v\n", keys)
	}

	for _, key := range keys {

		initForum()
		InitSiteConifg(key)
		//倒入标签
		//倒入图片
		//倒入推荐数据
		//倒入友情链接

	}
	fmt.Println("parse site config ok")
}

//GetImageForSite 为网站获取图片
func GetImageForSite() (string, string) {
	var bannerRes string
	var logoRes string
	var count int
	count = 0
	bannerRes = ""
	banners, _ := filepath.Glob("static/images/banner/*")
	for _, banner := range banners {
		if count > 0 {
			bannerRes = bannerRes + "|" + banner
		} else {
			bannerRes = banner
		}
		count++

	}
	logos, _ := filepath.Glob("static/images/logo/*")
	for _, logo := range logos {
		logoRes = logo
	}
	return logoRes, bannerRes
}
