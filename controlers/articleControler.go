/**
* @file articleControler.go
* @brief 文章的增删查改
* @author frankie@gmail.com
* @version v1.0
* @date 2019-03-11
 */

package controlers

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"../common"
	"../definition"
	"../models"
	"github.com/gin-gonic/gin"
)

/* --------------------------------------------------------------------------*/
/**
* @brief CreateArtile 从文件中读取文章内容插入数据库
*
* @returns
 */
/* ----------------------------------------------------------------------------*/
func CreateArticle() string {
	var article definition.Article
	article.Title = "abcdefefefe"
	article.ArticleContent = "ssdfdfdfdfdfddfabcdefefefe"
	article.Status = 1
	article.ReadTimes = 123
	article.ForumId = 1
	models.CreateArticleMode(article)
	return ""
}

func CrawArticle() {
	articlePath, _ := filepath.Glob("resource/article/*")
	var article definition.Article
	var res bool
	for _, path := range articlePath {
		article, res = readArticle(path)
		if res {
			models.CreateArticleMode(article)
		}
	}

}
func readArticle(path string) (definition.Article, bool) {
	var article definition.Article
	data, err := ioutil.ReadFile(path)
	if err != nil {
		common.Log1("get config fatal:" + err.Error())
		return article, false
	}
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)
	count := 0
	summary := ""
	content := ""
	for scanner.Scan() {
		line := scanner.Text()
		if count == 0 {
			article.Title = line
		}
		if count > 0 && count < 4 {
			summary = summary + line
		}

		if count > 0 {
			content = content + line
		}
		count++
	}
	article.Summary = summary
	article.ArticleContent = content
	article.Status = 1
	article.ReadTimes = 123
	article.ForumId = 1
	return article, true
}

func GetArticle(c *gin.Context) {
	//host := c.Request.Host
	articleIdStr := c.Param("id")
	fmt.Println("articleId:", articleIdStr)
	articleId, err := strconv.Atoi(articleIdStr)
	if err != nil {
		return
	}
	var res bool
	var pageData definition.ArticlePageData
	pageData.SiteConfig = models.GetSiteConfig()
	pageData.Forums = models.GetForums()
	pageData.ArticleData, res = models.GetArticleById(uint(articleId))

	if res == false {
		return
	}
	forumName, err1 := models.GetForumById(uint(pageData.ArticleData.ForumId))
	if err1 == false {
		return
	}
	pageData.Articles = models.GetArticlesByForum(uint(pageData.ArticleData.ForumId)) //获取版块
	//pageData.Banner = strings.Split(pageData.SiteConfig.Banner, `|`)
	BannerPath := strings.Split(pageData.SiteConfig.Banner, `|`)
	for _, v := range BannerPath {
		pageData.Banner = append(pageData.Banner, "../"+v)
	}
	pageData.SiteConfig.Logo = "../" + pageData.SiteConfig.Logo
	templatePath := "/tpl1/article.html" //模版路径
	c.HTML(http.StatusOK, templatePath, gin.H{
		"Config":       pageData.SiteConfig,
		"Forums":       pageData.Forums,
		"Article":      pageData.ArticleData,
		"Articles":     pageData.Articles,
		"Banners":      pageData.Banner,
		"CurrentForum": forumName,
	})

}
