package schemas

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255)" json:"title"`
	Password string `gorm:"type:varchar(255)" json:"category_id"`
	Email    string `gorm:"type:varchar(255)" json:"author"`
}

type Book struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Title      string `gorm:"type:varchar(255)" json:"title"`
	CategoryId string `gorm:"type:int" json:"category_id"`
	BookedBy   string `gorm:"type:int" json:"booked_by"`
	Author     string `gorm:"type:varchar(255)" json:"author"`
}
