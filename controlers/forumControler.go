package controlers

import (
	"net/http"
	"strconv"
	"strings"

	"../definition"
	"../models"
	"github.com/gin-gonic/gin"
)

func GetForum(c *gin.Context) {
	//host := c.Request.Host
	forumIdStr := c.Param("id")
	forumId, err := strconv.Atoi(forumIdStr)
	if err != nil {
		return
	}
	forumName, err1 := models.GetForumById(uint(forumId))
	if err1 == false {
		return
	}
	var pageData definition.IndexPageData
	pageData.SiteConfig = models.GetSiteConfig()
	pageData.Forums = models.GetForums()                         //获取版块
	pageData.Articles = models.GetArticlesByForum(uint(forumId)) //获取版块
	//pageData.Banner = strings.Split(pageData.SiteConfig.Banner, `|`)
	BannerPath := strings.Split(pageData.SiteConfig.Banner, `|`)
	for _, v := range BannerPath {
		pageData.Banner = append(pageData.Banner, "../"+v)
	}
	pageData.SiteConfig.Logo = "../" + pageData.SiteConfig.Logo
	templatePath := "/tpl1/forum.html" //模版路径
	c.HTML(http.StatusOK, templatePath, gin.H{
		"Config":       pageData.SiteConfig,
		"Forums":       pageData.Forums,
		"Articles":     pageData.Articles,
		"Banners":      pageData.Banner,
		"CurrentForum": forumName,
	})

}
