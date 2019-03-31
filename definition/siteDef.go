package definition

import (
	"github.com/jinzhu/gorm"
)

type PublicKey struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

//网站信息保存在缓存
type SiteConfigKey struct {
	SiteId string `json:"siteId"`
	Domain string `json:"domain"`
	//Database     string `json:"database"`
	//DatabaseType string `json:"databaseType"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Temple      string `json:"temple"`
	Logo        string `json:"logo"`
	Banner      string `json:"banner"`
}

//网站信息保存数据库
type SiteConfigDb struct {
	gorm.Model
	Domain      string `gorm:"size:255"`
	Database    string `gorm:"size:255"`
	Title       string `gorm:"size:255"`
	Description string `gorm:"size:255"`
	Keywords    string `gorm:"size:255"`
	Temple      string `gorm:"size:255"`
	Logo        string `gorm:"size:255"`
	Banner      string `gorm:"size:255"`
}

//推荐信息
type RecommendDb struct {
	gorm.Model
	Data  string `gorm:"size:3000"`
	Image string `gorm:"size:255"`
}

//文章字段
type Article struct {
	gorm.Model
	Link           string `gorm:"size:255"`        //标题
	Title          string `gorm:"size:255"`        //标题
	ArticleContent string `gorm:"type:mediumtext"` //正文
	Summary        string `gorm:"size:600"`        //摘要
	Status         int    //文章状态
	ReadTimes      int    //阅读次数
	ForumId        uint   //版块ID
}

//tag
type Tag struct {
	gorm.Model
	TagName string `gorm:"size:255"`
}

//tag video 关系表
type TagVideo struct {
	gorm.Model
	TagId   int64
	VideoId int64
}

//广告推荐版块
type AdvanceModel struct {
	gorm.Model
	Content string `gorm:"type:mediumtext"`
	Picture string
}

//版块字段
type Forum struct {
	gorm.Model
	ForumName string `gorm:"size:255"`
}

//文章页面渲染元素
type ArticlePageData struct {
	Forums      []Forum
	ArticleData Article
	Articles    []Article
	SiteConfig  SiteConfigDb
	Banner      []string
}

//inndex页面渲染元素
type IndexPageData struct {
	Forums     []Forum
	Articles   []Article
	SiteConfig SiteConfigDb
	Banner     []string
}
