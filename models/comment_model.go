package models

// CommentModel 评论表
type CommentModel struct {
	MODEL
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"` //父级评论
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`  //子级评论列表
	ParentCommentID    *uint           `gorm:"size:10" json:"parent_comment_id"`                //父级评论id
	Content            string          `gorm:"size:256" json:"content"`                         //评论内容
	DiggCount          int             `gorm:"size:8;default:0;" json:"digg_count"`             //点赞数
	CommentCount       int             `gorm:"size:8;default:0;" json:"comment_count"`          //子评论数
	Article            ArticleModel    `json:"article"`                                         //关联的文章
	ArticleID          uint            `json:"article_id"`                                      //文章id
	User               UserModel       `json:"user"`                                            //关联的用户
	//	NickName           string          `gorm:"size:42" json:"nick_name"`                        //用户的昵称
	//	Avatar             string          `json:"avatar"`
	UserID uint `json:"user_id"` //评论的用户
}
