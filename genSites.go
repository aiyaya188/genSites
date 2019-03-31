package main

import (
	//"net/http"

	"os"
	"path/filepath"
	"strings"

	"./common"
	"./controlers"
	"github.com/gin-gonic/gin"
	//"./models"
	//	"fmt"
	//"os/exec"
	// "bytes"
	//"io/ioutil"
	// "net/url"
)

func setTemplate(engine *gin.Engine) {
	engine.LoadHTMLGlob(filepath.Join(getCurrentDirectory(), "./views/**/*"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		//seelog.Critical(err)
		common.Log1("error:", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {
	/*
		go controlers.TranCodeRun()
		http.HandleFunc("/upload", controlers.HandleUpload)
		http.HandleFunc("/notify", controlers.CloudNotifyProcess)
		http.HandleFunc("/delVideo", controlers.DelVideoProcess)
		http.ListenAndServe(":8989", nil)
	*/

	//controlers.ParseSites()
	//return
	//controlers.ParseTemple()
	//controlers.CreateArticle()
	//return
	/*
		http.HandleFunc("/", controlers.ShowIndex)
		//http.Handle("/", http.FileServer(http.Dir("/Users/frankie/project/goSites/genSites/tpl")))
		http.ListenAndServe(":8989", nil)
	*/
	if len(os.Args) > 1 {
		if os.Args[1] == "buildConfig" {
			controlers.CreateSitesConfig()
		}
		if os.Args[1] == "buildArticles" {
			controlers.CrawArticle()
		}

		return

	}

	router := gin.Default()
	setTemplate(router)
	//测试随机
	//router.GET("/testRan", controlers.initForum)
	//首页
	router.GET("/index", controlers.GetIndex)
	router.GET("/", controlers.GetIndex)
	//文章浏览
	router.GET("/article/:id", controlers.GetArticle)
	//版块
	router.GET("/forum/:id", controlers.GetForum)
	//静态资源
	router.Static("/static", filepath.Join(getCurrentDirectory(), "./static"))
	router.Run(":8989")

}
