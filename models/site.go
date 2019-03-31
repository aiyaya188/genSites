/**
* @file site.go
* @brief 站群数据库操作
* @author frankie@gmail.com
* @version v1.0
* @date 2019-03-11
 */

package models

import (
	"../common"
	"../database"
	"../definition"
	"github.com/jinzhu/gorm"
)

/* --------------------------------------------------------------------------*/
/**
* @brief GetForums 获取版块
*
* @returns
 */
/* ----------------------------------------------------------------------------*/
//获取版块名字
func GetForums() []definition.Forum {
	var db *gorm.DB
	var forumSet []definition.Forum
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return nil
	}
	//db.Find(&forumSet)
	db.Order("created_at desc").Find(&forumSet)
	return forumSet
}

func GetSiteConfig() definition.SiteConfigDb {
	var db *gorm.DB
	var config definition.SiteConfigDb
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return config
	}
	//db.Find(&forumSet)
	db.Order("created_at desc").First(&config)
	return config
}

/* --------------------------------------------------------------------------*/
/**
* @brief CreateArticleMode
*
* @param definition.Article
*
* @returns
 */
/* ----------------------------------------------------------------------------*/
//GreatArticleMode 新增文章
func CreateArticleMode(article definition.Article) bool {
	var db *gorm.DB
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return false
	}
	if db.NewRecord(&article) {
		db.Create(&article)
		return true
	}
	return false
}

/* --------------------------------------------------------------------------*/
/**
* @brief GetArticlesModel
*
* @returns
 */
/* ----------------------------------------------------------------------------*/
func GetArticlesModel() []definition.Article {
	var db *gorm.DB
	var articleSet []definition.Article
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return nil
	}
	//db.Find(&forumSet)
	db.Order("created_at desc").Find(&articleSet)
	return articleSet
}

func GetArticlesByForum(forumId uint) []definition.Article {
	var db *gorm.DB
	var articleSet []definition.Article
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return nil
	}
	//db.Find(&forumSet)

	db.Where(&definition.Article{ForumId: forumId}).Find(&articleSet)
	//db.Order("created_at desc").Find(&articleSet)
	return articleSet
}

/* --------------------------------------------------------------------------*/
/**
* @brief CreateForumMode
*
* @param definition.Forum
*
* @returns
 */
/* ----------------------------------------------------------------------------*/
func CreateForumMode(forum definition.Forum) bool {
	var db *gorm.DB
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return false
	}
	if db.NewRecord(&forum) {
		db.Create(&forum)
		return true
	}
	return false
}
func GetForumByName(forumName string) (string, bool) {
	var db *gorm.DB
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return "", false
	}
	var forum definition.Forum
	//db.Where(&definition.Videos{Md5: videoSample.Md5}).First(&video)
	db.Where(&definition.Forum{ForumName: forumName}).First(&forum)
	if forum.ForumName == "" {
		return "", false
	}
	return forum.ForumName, true

}
func GetArticleById(articleId uint) (definition.Article, bool) {

	var db *gorm.DB
	var article definition.Article
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return article, false
	}
	//db.Where(&definition.Videos{Md5: videoSample.Md5}).First(&video)
	db.First(&article, articleId)
	if article.Title == "" {

		return article, false
	}
	return article, true

}
func GetForumById(forumId uint) (definition.Forum, bool) {
	var db *gorm.DB
	var forum definition.Forum
	//err int
	db = database.GetDB("master")
	if db == nil {
		common.Log("get db nill")
		return forum, false
	}
	//db.Where(&definition.Videos{Md5: videoSample.Md5}).First(&video)
	db.First(&forum, forumId)
	if forum.ForumName == "" {
		return forum, false
	}
	return forum, true

}
