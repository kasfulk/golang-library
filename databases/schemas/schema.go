package schemas

type Book struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Title      string `gorm:"type:varchar(255)" json:"title"`
	CategoryId string `gorm:"type:text" json:"category_id"`
	Author     string `gorm:"type:varchar(255)" json:"author"`
}

type BookCategory struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	CategoryName string `gorm:"type:varchar(255)" json:"category_name"`
}
