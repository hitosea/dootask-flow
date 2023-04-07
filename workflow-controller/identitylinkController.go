package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"workflow/workflow-engine/service"

	"workflow/util"
)

// FindParticipantByProcInstID 根据流程id查询流程参与者
func FindParticipantByProcInstID(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		util.ResponseErr(writer, "只支持get方法！！")
		return
	}
	request.ParseForm()
	if len(request.Form["procInstId"]) == 0 {
		util.ResponseErr(writer, "流程 procInstId 不能为空")
		return
	}
	procInstId, err := strconv.Atoi(request.Form["procInstId"][0])
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}
	result, err := service.FindParticipantByProcInstID(procInstId)
	if err != nil {
		util.ResponseErr(writer, err)
		return
	}
	fmt.Fprintf(writer, result)

}
