package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"sync"

	"github.com/jinzhu/gorm"

	//_ "github.com/go-sql-driver/mysql"
	"../common"
	"../definition"

	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbKey int
	//DBM 公共数据件
	DBM *gorm.DB
	//DBS 数据库组
	DBS        []*gorm.DB
	checkCount int
	//  OffsetMap  map[string]string
	OffsetMap sync.Map
	// 站群信息

)

//GetDbByHost 根据域名获取数据库名字
func GetDbByHost(hostName string) string {
	dbname := strings.Replace(hostName, `.`, `_`, -1)
	return dbname
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

//ParseSistes 解析建站配置文件
func createDatabases() {
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
		dbCtreateStr = "mysql -uroot -pAa20192019 -e " + `"create database IF NOT EXISTS ` + dbName + ` DEFAULT CHARSET utf8 COLLATE utf8_general_ci"`
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

func init() {
	fmt.Println("init db start")
	createDatabases()

	DBM = GetDB("master")
	if DBM == nil {
		return
	}
	//OffsetMap = make(map[string]string)
	DBM.AutoMigrate(&definition.Article{})
	DBM.AutoMigrate(&definition.Forum{})
	DBM.AutoMigrate(&definition.SiteConfigDb{})
	//创建配置表
	var config definition.SiteConfigDb
	exitConf := ifExitConfig()
	if exitConf == false {
		fmt.Println("create config table")
		if DBM.NewRecord(&config) {
			DBM.Create(&config)
		}
	}
	fmt.Println("init db end")
}

func ifExitConfig() bool {
	var db *gorm.DB
	//err int
	db = GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return false
	}
	var config definition.SiteConfigDb
	db.First(&config)
	if config == (definition.SiteConfigDb{}) {
		return false
	}
	return true

}

// GetDB : 获取数据库操作
func GetDB(dbType string) *gorm.DB {
	var db *gorm.DB
	if dbType == "master" {
		var err error
		//fmt.Println("get db master")
		if DBM == nil {
			dsn := common.GetConfig("mysql", "masterDsn").String()
			fmt.Println(fmt.Sprintf("dsn is %s", dsn))
			db, err = DbConn(dsn)
			if err == nil {
				return db
			}
			fmt.Println("connect db err:", err)
			fmt.Println("db connetct fail")
			return nil
		}
		//fmt.Println("get old db")
		return DBM
	}

	if dbType == "slave" {
		if len(DBS) < 1 {
			fmt.Println("slave new conn")
			slaveCount, _ := common.GetConfig("mysql", "slaveCount").Int()
			DBS = make([]*gorm.DB, slaveCount)
			for i := 0; i < slaveCount; i++ {
				dsn := "slaveDsn" + strconv.Itoa(i+1)
				dsn = common.GetConfig("mysql", dsn).String()
				fmt.Println("slave dns is : " + dsn)
				dbs, err := DbConn(dsn)
				if err == nil {
					DBS[i] = dbs
				} else {
					fmt.Println(err)
				}
			}
		}
		if len(DBS) > 0 {
			if dbKey > len(DBS)-1 || dbKey < 1 {
				dbKey = 0
			} else {
				dbKey++
			}
			return DBS[dbKey]
		}

	}
	return nil
}

//DbConn : 数据库连接
func DbConn(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		//fmt.Println(fmt.Sprint("database connect fail ,db:%s,err:%d",db,err))
		return db, err
	}
	fmt.Println("恭喜，数据库连接成功")
	//db.LogMode(true)
	// 全局禁用表名复数
	db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		//return "yf_" + defaultTableName
		return defaultTableName

	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(200)
	return db, err
}

/*
func CheckDb() {
	fmt.Println("checkdb 执行...")

	for i := 0; i < 1; i++ {
		if DBM != nil {

			// Raw SQL
			rows, dbmErr := DBM.Raw("select 1 from mysql.db limit 1").Rows()

			if dbmErr != nil {
				fmt.Println("master fail ! 报警处理=================================")
				fmt.Println(dbmErr)

				//panic("db fail")

				//尝试重连接
				//GetDB("master")

			} else {
				defer rows.Close()
				fmt.Println(strconv.Itoa(checkCount) + "--主数据库查询正常\n")
			}
		} else {
			fmt.Println(strconv.Itoa(checkCount) + "主数据库没连接")
		}

		checkCount++
	}

	for i := 0; i < len(DBS); i++ {
		if DBS[i] != nil {

			// Raw SQL
			rows, dbsErr := DBS[i].Raw("select 1 from mysql.db limit 1").Rows()
			if dbsErr != nil {
				fmt.Println("slave fail ! 报警处理")
				fmt.Println(dbsErr)

			} else {
				defer rows.Close()
				fmt.Printf("从数据库%d查询正常\n", i)
			}
		} else {
			fmt.Println(strconv.Itoa(len(DBS)) + "从数据库没连接")
		}
	}

}
*/
