package services

import (
	"bbs-go/model"
	"bbs-go/pkg/common"
	"bbs-go/repositories"
	"net/http"

	"github.com/mlogclub/simple/common/dates"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"github.com/sirupsen/logrus"
)

var OperateLogService = newOperateLogService()

func newOperateLogService() *operateLogService {
	return &operateLogService{}
}

type operateLogService struct {
}

func (s *operateLogService) Get(id int64) *model.OperateLog {
	return repositories.OperateLogRepository.Get(sqls.DB(), id)
}

func (s *operateLogService) Take(where ...interface{}) *model.OperateLog {
	return repositories.OperateLogRepository.Take(sqls.DB(), where...)
}

func (s *operateLogService) Find(cnd *sqls.Cnd) []model.OperateLog {
	return repositories.OperateLogRepository.Find(sqls.DB(), cnd)
}

func (s *operateLogService) FindOne(cnd *sqls.Cnd) *model.OperateLog {
	return repositories.OperateLogRepository.FindOne(sqls.DB(), cnd)
}

func (s *operateLogService) FindPageByParams(params *params.QueryParams) (list []model.OperateLog, paging *sqls.Paging) {
	return repositories.OperateLogRepository.FindPageByParams(sqls.DB(), params)
}

func (s *operateLogService) FindPageByCnd(cnd *sqls.Cnd) (list []model.OperateLog, paging *sqls.Paging) {
	return repositories.OperateLogRepository.FindPageByCnd(sqls.DB(), cnd)
}

func (s *operateLogService) Count(cnd *sqls.Cnd) int64 {
	return repositories.OperateLogRepository.Count(sqls.DB(), cnd)
}

func (s *operateLogService) Create(t *model.OperateLog) error {
	return repositories.OperateLogRepository.Create(sqls.DB(), t)
}

func (s *operateLogService) Update(t *model.OperateLog) error {
	return repositories.OperateLogRepository.Update(sqls.DB(), t)
}

func (s *operateLogService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.OperateLogRepository.Updates(sqls.DB(), id, columns)
}

func (s *operateLogService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.OperateLogRepository.UpdateColumn(sqls.DB(), id, name, value)
}

func (s *operateLogService) Delete(id int64) {
	repositories.OperateLogRepository.Delete(sqls.DB(), id)
}

func (s *operateLogService) AddOperateLog(userId int64, opType, dataType string, dataId int64,
	description string, r *http.Request) {

	operateLog := &model.OperateLog{
		UserId:      userId,
		OpType:      opType,
		DataType:    dataType,
		DataId:      dataId,
		Description: description,
		CreateTime:  dates.NowTimestamp(),
	}
	if r != nil {
		operateLog.Ip = common.GetRequestIP(r)
		operateLog.UserAgent = common.GetUserAgent(r)
		operateLog.Referer = r.Header.Get("Referer")
	}
	if err := repositories.OperateLogRepository.Create(sqls.DB(), operateLog); err != nil {
		logrus.Error(err)
	}
}
