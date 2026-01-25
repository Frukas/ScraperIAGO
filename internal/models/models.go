package models

var Categories = []string{"自動車", "金融・投資", "資源", "農業", "経済指標", "アメリカ",
	"アルゼンチン", "ペルー", "チリ", "コロンビア", "ボリビア", "パラグアイ"}

type Article struct {
	Id        uint16 `gorm:"primarykey" json:"Id"`
	Address   string `json:"address" gorm:"unique"`
	Title     string `json:"Title"`
	Category  string `json:"Category"`
	IsCorrect bool   `json:"iscorrect"`
}
