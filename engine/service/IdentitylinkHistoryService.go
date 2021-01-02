package service

import (
	"github.com/mumushuiding/util"
	"github.com/newjunwei/go-workflow/go-workflow/engine/model"
)

// FindParticipantHistoryByProcInstID 历史纪录查询
func FindParticipantHistoryByProcInstID(procInstID int) (string, error) {
	datas, err := model.FindParticipantHistoryByProcInstID(procInstID)
	if err != nil {
		return "", err
	}
	str, err := util.ToJSONStr(datas)
	if err != nil {
		return "", err
	}
	return str, nil
}
