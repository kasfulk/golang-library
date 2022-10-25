package schemas

type Post struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Title string `gorm:"type:varchar(255)" json:"description"`
	Body  string `gorm:"type:text" json:"body"`
	Image string `gorm:"type:varchar(255)" json:"image"`
}
