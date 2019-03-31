//getIndexPage 获取
/**
* @file indexControler.go
* @brief 根据访问域名生成对应模版的index页面
* @author frankie@gmail.com
* @version 1.0
* @date 2019-03-10
 */

package controlers

import (
	"net/http"
	"strings"

	"../definition"
	"../models"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	//host := c.Request.Host
	//fmt.Println("host:", host)
	var pageData definition.IndexPageData
	pageData.SiteConfig = models.GetSiteConfig()
	pageData.Forums = models.GetForums()          //获取版块
	pageData.Articles = models.GetArticlesModel() //获取版块
	pageData.Banner = strings.Split(pageData.SiteConfig.Banner, `|`)
	//templatePath := "views/www_abc_com/index.html" //模版路径
	templatePath := "/tpl1/index.html" //模版路径
	c.HTML(http.StatusOK, templatePath, gin.H{
		"Config":   pageData.SiteConfig,
		"Forums":   pageData.Forums,
		"Articles": pageData.Articles,
		"Banners":  pageData.Banner,
	})
}
