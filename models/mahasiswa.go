package models

type Mahasiswa struct {
	Id            int64  `gorm:"primaryKey" json:"id"`
	NamaMahasiswa string `gorm:"type:varchar(300)" json:"nama_mahasiswa"`
	Alamat        string `gorm:"type:varchar(300)" json:"alamat"`
}

func (Mahasiswa) TableName() string {
    return "mahasiswa"
}

func Migrate() {
    DB.AutoMigrate(&Mahasiswa{})
}
