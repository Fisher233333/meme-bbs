package repositories

import (
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"gorm.io/gorm"

	"bbs-go/model"
)

var TopicNodeRepository = newTopicNodeRepository()

func newTopicNodeRepository() *topicNodeRepository {
	return &topicNodeRepository{}
}

type topicNodeRepository struct {
}

func (r *topicNodeRepository) Get(db *gorm.DB, id int64) *model.TopicNode {
	ret := &model.TopicNode{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *topicNodeRepository) Take(db *gorm.DB, where ...interface{}) *model.TopicNode {
	ret := &model.TopicNode{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *topicNodeRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.TopicNode) {
	cnd.Find(db, &list)
	return
}

func (r *topicNodeRepository) FindOne(db *gorm.DB, cnd *sqls.Cnd) *model.TopicNode {
	ret := &model.TopicNode{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *topicNodeRepository) FindPageByParams(db *gorm.DB, params *params.QueryParams) (list []model.TopicNode, paging *sqls.Paging) {
	return r.FindPageByCnd(db, &params.Cnd)
}

func (r *topicNodeRepository) FindPageByCnd(db *gorm.DB, cnd *sqls.Cnd) (list []model.TopicNode, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.TopicNode{})

	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *topicNodeRepository) Create(db *gorm.DB, t *model.TopicNode) (err error) {
	err = db.Create(t).Error
	return
}

func (r *topicNodeRepository) Update(db *gorm.DB, t *model.TopicNode) (err error) {
	err = db.Save(t).Error
	return
}

func (r *topicNodeRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.TopicNode{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *topicNodeRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.TopicNode{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *topicNodeRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.TopicNode{}, "id = ?", id)
}
