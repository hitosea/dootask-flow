package model

import "github.com/jinzhu/gorm"

// Procdef 流程定义表
type Procdef struct {
	Model
	Name       string `gorm:"comment:'名字'" json:"name,omitempty"`
	Version    int    `gorm:"comment:'版本'" json:"version,omitempty"`
	Resource   string `gorm:"size:10000;comment:'流程定义json字符串'"  json:"resource,omitempty"`
	Userid     string `gorm:"comment:'用户id'" json:"userid,omitempty"`
	Username   string `gorm:"comment:'用户名称'" json:"username,omitempty"`
	Company    string `gorm:"comment:'用户所在公司'" json:"company,omitempty"`
	DeployTime string `gorm:"comment:'部署时间'" json:"deployTime,omitempty"`
}

// Save save and return id
// 保存并返回ID
func (p *Procdef) Save() (ID int, err error) {
	err = db.Create(p).Error
	if err != nil {
		return 0, err
	}
	return p.ID, nil
}

// SaveTx SaveTx
func (p *Procdef) SaveTx(tx *gorm.DB) error {
	err := tx.Create(p).Error
	if err != nil {
		return err
	}
	return nil
}

// GetProcdefLatestByNameAndCompany :get latest procdef by name and company
// 根据名字和公司查询最新的流程定义
func GetProcdefLatestByNameAndCompany(name, company string) (*Procdef, error) {
	var p []*Procdef
	err := db.Where("name=? and company=?", name, company).Order("version desc").Find(&p).Error
	if err != nil || len(p) == 0 {
		return nil, err
	}
	return p[0], err
}

// GetProcdefByID 根据流程定义
func GetProcdefByID(id int) (*Procdef, error) {
	var p = &Procdef{}
	err := db.Where("id=?", id).Find(p).Error
	return p, err
}

// DelProcdefByID del by id
// 根据id删除
func DelProcdefByID(id int) error {
	err := db.Where("id = ?", id).Delete(&Procdef{}).Error
	return err
}

// DelProcdefByIDTx DelProcdefByIDTx
func DelProcdefByIDTx(id int, tx *gorm.DB) error {
	return tx.Where("id = ?", id).Delete(&Procdef{}).Error
}

// FindProcdefsWithCountAndPaged return result with total count and error
// 返回查询结果和总条数
func FindProcdefsWithCountAndPaged(pageIndex, pageSize int, maps map[string]interface{}) (datas []*Procdef, count int, err error) {
	err = db.Select("id,name,version,userid,deploy_time").Where(maps).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&datas).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	err = db.Model(&Procdef{}).Where(maps).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return datas, count, nil
}

// MoveProcdefToHistoryByIDTx 将流程定义移至历史纪录表
func MoveProcdefToHistoryByIDTx(ID int, tx *gorm.DB) error {
	err := tx.Exec("insert into procdef_history select * from procdef where id=?", ID).Error
	if err != nil {
		return err
	}
	return DelProcdefByIDTx(ID, tx)
}
