package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/WeBankPartners/go-common-lib/guid"
	"github.com/WeBankPartners/wecube-plugins-taskman/taskman-server/common/log"
	"github.com/WeBankPartners/wecube-plugins-taskman/taskman-server/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetTaskFormStruct(procInstId, nodeDefId string) (result models.TaskMetaResult, err error) {
	result = models.TaskMetaResult{Status: "OK", Message: "Success"}
	var items []*models.FormItemTemplateTable
	err = x.SQL("select * from form_item_template where form_template in (select form_template from task_template where node_def_id=? and request_template in (select request_template from request where proc_instance_id=?))", nodeDefId, procInstId).Find(&items)
	if err != nil {
		return
	}
	if len(items) == 0 {
		err = fmt.Errorf("Can not find task item template with procInstId:%s nodeDefId:%s ", procInstId, nodeDefId)
		return
	}
	resultData := models.TaskMetaResultData{FormMetaId: items[0].FormTemplate}
	for _, item := range items {
		if item.Entity == "" {
			continue
		}
		resultData.FormItemMetas = append(resultData.FormItemMetas, &models.TaskMetaResultItem{FormItemMetaId: item.Id, PackageName: item.PackageName, EntityName: item.Entity, AttrName: item.Name})
	}
	result.Data = resultData
	return
}

func PluginTaskCreate(input *models.PluginTaskCreateRequestObj, callRequestId, dueDate string, nextOptions []string) (result *models.PluginTaskCreateOutputObj, err error) {
	result = &models.PluginTaskCreateOutputObj{CallbackParameter: input.CallbackParameter, ErrorCode: "0", ErrorMessage: "", Comment: ""}
	var requestTable []*models.RequestTable
	err = x.SQL("select id,form,request_template,emergency from request where proc_instance_id=?", input.ProcInstId).Find(&requestTable)
	if err != nil {
		return result, fmt.Errorf("Try to check proc_instance_id:%s is in request fail,%s ", input.ProcInstId, err.Error())
	}
	var actions []*execAction
	nowTime := time.Now().Format(models.DateTimeFormat)
	newTaskFormObj := models.FormTable{Id: guid.CreateGuid(), Name: "form_" + input.TaskName}
	input.RoleName = remakeTaskReportRole(input.RoleName)
	newTaskObj := models.TaskTable{Id: guid.CreateGuid(), Name: input.TaskName, Status: "created", Form: newTaskFormObj.Id, Reporter: input.Reporter, ReportRole: input.RoleName, Description: input.TaskDescription, CallbackUrl: input.CallbackUrl, CallbackParameter: input.CallbackParameter, NextOption: strings.Join(nextOptions, ","), Handler: input.Handler}
	var taskFormInput models.PluginTaskFormDto
	if input.TaskFormInput != "" {
		err = json.Unmarshal([]byte(input.TaskFormInput), &taskFormInput)
		if err != nil {
			return result, fmt.Errorf("Try to json unmarshal taskFormInput to json data fail,%s ", err.Error())
		}
	} else {
		// Custom task create
		customExpireTime := ""
		dueMin, _ := strconv.Atoi(dueDate)
		if dueMin > 0 {
			customExpireTime = time.Now().Add(time.Duration(dueMin) * time.Minute).Format(models.DateTimeFormat)
		}
		taskInsertAction := execAction{Sql: "insert into task(id,name,description,status,proc_def_id,proc_def_key,node_def_id,node_name,callback_url,callback_parameter,reporter,report_role,report_time,expire_time,emergency,callback_request_id,next_option,handler,created_by,created_time,updated_by,updated_time) value (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"}
		taskInsertAction.Param = []interface{}{newTaskObj.Id, newTaskObj.Name, newTaskObj.Description, newTaskObj.Status, newTaskObj.ProcDefId, newTaskObj.ProcDefKey, newTaskObj.NodeDefId, newTaskObj.NodeName, newTaskObj.CallbackUrl, newTaskObj.CallbackParameter, newTaskObj.Reporter, newTaskObj.ReportRole, nowTime, customExpireTime, newTaskObj.Emergency, callRequestId, newTaskObj.NextOption, newTaskObj.Handler, "system", nowTime, "system", nowTime}
		actions = append(actions, &taskInsertAction)
		err = transaction(actions)
		return
	}
	itemTemplateMap := getFormItemTemplateMap(taskFormInput.FormMetaId)
	newTaskObj.ProcDefId = taskFormInput.ProcDefId
	newTaskObj.ProcDefKey = taskFormInput.ProcDefKey
	if len(requestTable) > 0 {
		newTaskObj.Request = requestTable[0].Id
		newTaskObj.NodeDefId = taskFormInput.TaskNodeDefId
		newTaskObj.Emergency = requestTable[0].Emergency
		var taskTemplateTable []*models.TaskTemplateTable
		x.SQL("select * from task_template where request_template=? and node_def_id=?", requestTable[0].RequestTemplate, taskFormInput.TaskNodeDefId).Find(&taskTemplateTable)
		if len(taskTemplateTable) > 0 {
			newTaskObj.TaskTemplate = taskTemplateTable[0].Id
			newTaskObj.NodeName = taskTemplateTable[0].NodeName
			newTaskObj.ExpireTime = calcExpireTime(nowTime, taskTemplateTable[0].ExpireDay)
			newTaskObj.Name = taskTemplateTable[0].Name
			newTaskObj.Description = taskTemplateTable[0].Description
			newTaskObj.Reporter = requestTable[0].Reporter
			newTaskObj.ReportTime = nowTime
			newTaskObj.Handler = taskTemplateTable[0].Handler
			newTaskFormObj.FormTemplate = taskTemplateTable[0].FormTemplate
		} else {
			log.Logger.Warn("Can not find any taskTemplate", log.String("requestTemplate", requestTable[0].RequestTemplate), log.String("nodeDefId", taskFormInput.TaskNodeDefId))
			err = fmt.Errorf("Can not find any taskTemplate in request:%s with nodeDefId:%s ", newTaskObj.Request, taskFormInput.TaskNodeDefId)
			return
		}
	}
	taskInsertAction := execAction{Sql: "insert into task(id,name,description,form,status,request,task_template,proc_def_id,proc_def_key,node_def_id,node_name,callback_url,callback_parameter,reporter,report_role,report_time,emergency,cache,callback_request_id,next_option,expire_time,handler,created_by,created_time,updated_by,updated_time) value (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"}
	taskInsertAction.Param = []interface{}{newTaskObj.Id, newTaskObj.Name, newTaskObj.Description, newTaskObj.Form, newTaskObj.Status, newTaskObj.Request, newTaskObj.TaskTemplate, newTaskObj.ProcDefId, newTaskObj.ProcDefKey, newTaskObj.NodeDefId, newTaskObj.NodeName, newTaskObj.CallbackUrl, newTaskObj.CallbackParameter, newTaskObj.Reporter, newTaskObj.ReportRole, nowTime, newTaskObj.Emergency, input.TaskFormInput, callRequestId, newTaskObj.NextOption, newTaskObj.ExpireTime, newTaskObj.Handler, "system", nowTime, "system", nowTime}
	actions = append(actions, &taskInsertAction)
	actions = append(actions, &execAction{Sql: "insert into form(id,name,form_template) value (?,?,?)", Param: []interface{}{newTaskFormObj.Id, newTaskFormObj.Name, newTaskFormObj.FormTemplate}})
	for _, formDataEntity := range taskFormInput.FormDataEntities {
		tmpGuidList := guid.CreateGuidList(len(formDataEntity.FormItemValues))
		for i, formDataItem := range formDataEntity.FormItemValues {
			tmpItemGroup := itemTemplateMap[formDataItem.FormItemMetaId].ItemGroup
			actions = append(actions, &execAction{Sql: "insert into form_item(id,form,form_item_template,name,value,item_group,row_data_id) value (?,?,?,?,?,?,?)", Param: []interface{}{tmpGuidList[i], newTaskFormObj.Id, formDataItem.FormItemMetaId, formDataItem.AttrName, formDataItem.AttrValue, tmpItemGroup, formDataItem.Oid}})
		}
	}
	customItems, tmpErr := getLastCustomFormItem(newTaskObj.Request, newTaskFormObj.FormTemplate, newTaskFormObj.Id)
	if tmpErr != nil {
		err = fmt.Errorf("Try to get custom item fail,%s ", tmpErr.Error())
		return
	}
	if len(customItems) > 0 {
		tmpGuidList := guid.CreateGuidList(len(customItems))
		for i, customItem := range customItems {
			actions = append(actions, &execAction{Sql: "insert into form_item(id,form,form_item_template,name,value,item_group,row_data_id) value (?,?,?,?,?,?,?)", Param: []interface{}{tmpGuidList[i], customItem.Form, customItem.FormItemTemplate, customItem.Name, customItem.Value, customItem.ItemGroup, customItem.RowDataId}})
		}
	}
	err = transactionWithoutForeignCheck(actions)
	return
}

func getLastCustomFormItem(requestId, taskFormTemplateId, newTaskFormId string) (result []*models.FormItemTable, err error) {
	result = []*models.FormItemTable{}
	if requestId == "" || taskFormTemplateId == "" {
		return
	}
	var formItemTemplates []*models.FormItemTemplateTable
	err = x.SQL("select * from form_item_template where entity='' and form_template=?", taskFormTemplateId).Find(&formItemTemplates)
	if len(formItemTemplates) == 0 || err != nil {
		return
	}
	groupNameTemplateIdMap := make(map[string]string)
	filterList := []string{}
	for _, v := range formItemTemplates {
		groupNameTemplateIdMap[fmt.Sprintf("%s_%s", v.ItemGroup, v.Name)] = v.Id
		filterList = append(filterList, fmt.Sprintf("(item_group='%s' and name='%s')", v.ItemGroup, v.Name))
	}
	var formItems []*models.FormItemTable
	err = x.SQL("select * from form_item where ("+strings.Join(filterList, " or ")+") and form in (select form from request where id=? union select form from task where request=?) order by item_group,name,id desc", requestId, requestId).Find(&formItems)
	groupNameExistMap := make(map[string]int)
	for _, v := range formItems {
		tmpKey := fmt.Sprintf("%s_%s", v.ItemGroup, v.Name)
		if _, b := groupNameExistMap[tmpKey]; b {
			continue
		}
		result = append(result, &models.FormItemTable{Form: newTaskFormId, FormItemTemplate: groupNameTemplateIdMap[tmpKey], Name: v.Name, ItemGroup: v.ItemGroup, Value: v.Value, RowDataId: v.RowDataId})
		groupNameExistMap[tmpKey] = 1
	}
	return
}

func getFormItemTemplateMap(formTemplateId string) map[string]*models.FormItemTemplateTable {
	resultMap := make(map[string]*models.FormItemTemplateTable)
	var itemTemplateTable []*models.FormItemTemplateTable
	x.SQL("select * from form_item_template where form_template=?", formTemplateId).Find(&itemTemplateTable)
	for _, v := range itemTemplateTable {
		resultMap[v.Id] = v
	}
	return resultMap
}

func remakeTaskReportRole(reportRoles string) string {
	resultList := []string{}
	for _, v := range strings.Split(reportRoles, ",") {
		if v != "" {
			resultList = append(resultList, v)
		}
	}
	if len(resultList) > 0 {
		return fmt.Sprintf(",%s,", strings.Join(resultList, ","))
	}
	return ""
}

func ListTask(param *models.QueryRequestParam, userRoles []string, operator string) (pageInfo models.PageInfo, rowData []*models.TaskListObj, err error) {
	rowData = []*models.TaskListObj{}
	roleFilterSql := "1=1"
	if len(userRoles) > 0 {
		roleFilterList := []string{}
		for _, v := range userRoles {
			roleFilterList = append(roleFilterList, "report_role like '%,"+v+",%'")
		}
		roleFilterSql = strings.Join(roleFilterList, " or ")
	}
	filterSql, _, queryParam := transFiltersToSQL(param, &models.TransFiltersParam{IsStruct: true, StructObj: models.TaskTable{}, PrimaryKey: "id", Prefix: "t1"})
	baseSql := fmt.Sprintf("select t1.id,t1.name,t1.description,t1.form,t1.status,t1.`version`,t1.request,t1.task_template,t1.node_name,t1.reporter,t1.report_time,t1.emergency,t1.handler,t1.created_by,t1.created_time,t1.updated_by,t1.updated_time,t1.expire_time from (select * from task where task_template in (select task_template from task_template_role where role_type='USE' and `role` in ('"+strings.Join(userRoles, "','")+"')) union select * from task where task_template is null and (%s)) t1 where t1.del_flag=0 %s ", roleFilterSql, filterSql)
	if param.Paging {
		pageInfo.StartIndex = param.Pageable.StartIndex
		pageInfo.PageSize = param.Pageable.PageSize
		pageInfo.TotalRows = queryCount(baseSql, queryParam...)
		pageSql, pageParam := transPageInfoToSQL(*param.Pageable)
		baseSql += pageSql
		queryParam = append(queryParam, pageParam...)
	}
	err = x.SQL(baseSql, queryParam...).Find(&rowData)
	if err != nil {
		return models.PageInfo{}, nil, err
	}
	var requestIdList []string
	for _, v := range rowData {
		buildTaskOperation(v, operator)
		requestIdList = append(requestIdList, v.Request)
	}
	var requestTables []*models.RequestTable
	x.SQL("select t1.id,t1.name,t1.reporter,t1.report_time,t2.name as request_template from request t1 left join request_template t2 on t1.request_template=t2.id where t1.id in ('" + strings.Join(requestIdList, "','") + "')").Find(&requestTables)
	requestMap := make(map[string]*models.RequestTable)
	for _, v := range requestTables {
		requestMap[v.Id] = v
	}
	nowTime := time.Now().Format(models.DateTimeFormat)
	for _, v := range rowData {
		if _, b := requestMap[v.Request]; b {
			v.RequestObj = *requestMap[v.Request]
			v.Reporter = "taskman"
		}
		if v.ExpireTime != "" {
			timeObj := models.ExpireObj{ReportTime: v.ReportTime, ExpireTime: v.ExpireTime, NowTime: nowTime}
			calcExpireObj(&timeObj)
			v.ExpirePercentObj = timeObj
		}
	}
	return
}

func GetTask(taskId string) (result models.TaskQueryResult, err error) {
	taskObj, tmpErr := getSimpleTask(taskId)
	if tmpErr != nil {
		return result, tmpErr
	}
	if taskObj.Request == "" {
		taskForm, tmpErr := queryTaskForm(&taskObj)
		if tmpErr != nil {
			return result, tmpErr
		}
		result.Data = []*models.TaskQueryObj{&taskForm}
		result.TimeStep = []*models.TaskQueryTimeStep{}
		return
	}
	// get request
	var requests []*models.RequestTable
	x.SQL("select * from request where id=?", taskObj.Request).Find(&requests)
	if len(requests) == 0 {
		return result, fmt.Errorf("Can not find request with id:%s ", taskObj.Request)
	}
	var requestCache models.RequestPreDataDto
	err = json.Unmarshal([]byte(requests[0].Cache), &requestCache)
	if err != nil {
		return result, fmt.Errorf("Try to json unmarshal request cache fail,%s ", err.Error())
	}
	requestQuery := models.TaskQueryObj{RequestId: taskObj.Request, RequestName: requests[0].Name, Reporter: requests[0].Reporter, ReportTime: requests[0].ReportTime, Comment: requests[0].Result, AttachFiles: []string{requests[0].AttachFile}, Editable: false}
	requestQuery.ExpireTime = requests[0].ExpireTime
	requestQuery.ExpectTime = requests[0].ExpectTime
	requestQuery.FormData = requestCache.Data
	result.Data = []*models.TaskQueryObj{&requestQuery}
	result.TimeStep, err = getRequestTimeStep(requests[0].RequestTemplate)
	if err != nil {
		return result, err
	}
	// get task list
	var taskList []*models.TaskTable
	x.SQL("select * from task where request=? and report_time<='"+taskObj.ReportTime+"' order by created_time", taskObj.Request).Find(&taskList)
	for _, v := range taskList {
		tmpTaskForm, tmpErr := queryTaskForm(v)
		if tmpErr != nil {
			err = tmpErr
			break
		}
		result.Data = append(result.Data, &tmpTaskForm)
	}
	if err != nil {
		return
	}
	for _, v := range result.TimeStep {
		if v.TaskTemplateId == taskObj.TaskTemplate {
			v.Active = true
			break
		}
	}
	return
}

func queryTaskForm(taskObj *models.TaskTable) (taskForm models.TaskQueryObj, err error) {
	taskForm = models.TaskQueryObj{TaskId: taskObj.Id, TaskName: taskObj.Name, Description: taskObj.Description, RequestId: taskObj.Request, Reporter: taskObj.Reporter, ReportTime: taskObj.ReportTime, Comment: taskObj.Result, Status: taskObj.Status, AttachFiles: []string{}, NextOption: []string{}, ExpireTime: taskObj.ExpireTime, FormData: []*models.RequestPreDataTableObj{}}
	if taskObj.Status != "done" {
		taskForm.Editable = true
	}
	if taskObj.NextOption != "" {
		taskForm.NextOption = strings.Split(taskObj.NextOption, ",")
		taskForm.ChoseOption = taskObj.ChoseOption
	}
	if taskObj.Request == "" {
		return
	}
	if taskObj.AttachFile != "" {
		taskForm.AttachFiles = []string{taskObj.AttachFile}
	}
	var itemTemplates []*models.FormItemTemplateTable
	err = x.SQL("select * from form_item_template where form_template in (select form_template from task_template where id=?) order by item_group,sort", taskObj.TaskTemplate).Find(&itemTemplates)
	if err != nil {
		return
	}
	if len(itemTemplates) == 0 {
		return taskForm, fmt.Errorf("Can not find any form item template with task:%s ", taskObj.Id)
	}
	formResult := getItemTemplateTitle(itemTemplates)
	taskForm.FormData = formResult
	var items []*models.FormItemTable
	x.SQL("select * from form_item where form=? order by item_group,row_data_id", taskObj.Form).Find(&items)
	if len(items) == 0 {
		return
	}
	itemRowMap := make(map[string][]string)
	rowItemMap := make(map[string][]*models.FormItemTable)
	for _, item := range items {
		if tmpRows, b := itemRowMap[item.ItemGroup]; b {
			existFlag := false
			for _, v := range tmpRows {
				if item.RowDataId == v {
					existFlag = true
					break
				}
			}
			if !existFlag {
				itemRowMap[item.ItemGroup] = append(itemRowMap[item.ItemGroup], item.RowDataId)
			}
		} else {
			itemRowMap[item.ItemGroup] = []string{item.RowDataId}
		}
		if _, b := rowItemMap[item.RowDataId]; b {
			rowItemMap[item.RowDataId] = append(rowItemMap[item.RowDataId], item)
		} else {
			rowItemMap[item.RowDataId] = []*models.FormItemTable{item}
		}
	}
	for _, formTable := range formResult {
		if rows, b := itemRowMap[formTable.ItemGroup]; b {
			for _, row := range rows {
				tmpRowObj := models.EntityTreeObj{Id: row, DataId: row, PackageName: formTable.PackageName, EntityName: formTable.Entity}
				tmpRowObj.EntityData = make(map[string]interface{})
				for _, rowItem := range rowItemMap[row] {
					tmpRowObj.EntityData[rowItem.Name] = rowItem.Value
				}
				formTable.Value = append(formTable.Value, &tmpRowObj)
			}
		}
	}
	taskForm.FormData = formResult
	return
}

func getRequestTimeStep(requestTemplateId string) (result []*models.TaskQueryTimeStep, err error) {
	var requestTemplateTable []*models.RequestTemplateTable
	err = x.SQL("select id,name from request_template where id=?", requestTemplateId).Find(&requestTemplateTable)
	if err != nil {
		return
	}
	if len(requestTemplateTable) == 0 {
		return result, fmt.Errorf("Can not find requestTemplate with id:%s ", requestTemplateId)
	}
	result = append(result, &models.TaskQueryTimeStep{RequestTemplateId: requestTemplateTable[0].Id, Name: "Start", Active: false})
	var taskTemplateTable []*models.TaskTemplateTable
	x.SQL("select id,name from task_template where request_template=?", requestTemplateId).Find(&taskTemplateTable)
	for _, v := range taskTemplateTable {
		result = append(result, &models.TaskQueryTimeStep{RequestTemplateId: requestTemplateId, TaskTemplateId: v.Id, Name: v.Name, Active: false})
	}
	return
}

func getSimpleTask(taskId string) (result models.TaskTable, err error) {
	var taskTable []*models.TaskTable
	err = x.SQL("select * from task where id=?", taskId).Find(&taskTable)
	if err != nil {
		return
	}
	if len(taskTable) == 0 {
		return result, fmt.Errorf("Can not find any task with id:%s ", taskId)
	}
	result = *taskTable[0]
	return
}

func ApproveTask(taskId, operator, userToken string, param models.TaskApproveParam) error {
	err := SaveTaskForm(taskId, operator, param)
	if err != nil {
		return err
	}
	requestParam, callbackUrl, getDataErr := getApproveCallbackParam(taskId)
	if getDataErr != nil {
		return getDataErr
	}
	if param.ChoseOption != "" {
		requestParam.ResultCode = param.ChoseOption
	}
	for _, v := range requestParam.Results.Outputs {
		v.Comment = param.Comment
	}
	requestBytes, _ := json.Marshal(requestParam)
	req, newReqErr := http.NewRequest(http.MethodPost, models.Config.Wecube.BaseUrl+callbackUrl, bytes.NewReader(requestBytes))
	if newReqErr != nil {
		return fmt.Errorf("Try to new http request fail,%s ", newReqErr.Error())
	}
	req.Header.Set("Authorization", userToken)
	req.Header.Set("Content-Type", "application/json")
	resp, respErr := http.DefaultClient.Do(req)
	if respErr != nil {
		return fmt.Errorf("Try to do http request fail,%s ", respErr.Error())
	}
	var respResult models.CallbackResult
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	log.Logger.Info("Callback response", log.String("body", string(b)))
	err = json.Unmarshal(b, &respResult)
	if err != nil {
		return fmt.Errorf("Try to json unmarshal response body fail,%s ", err.Error())
	}
	nowTime := time.Now().Format(models.DateTimeFormat)
	if respResult.Status != "OK" {
		if strings.Contains(respResult.Message, "None process instance found") {
			x.Exec("update task set status='done',updated_by=?,updated_time=? where id=?", operator, nowTime, taskId)
		}
		return fmt.Errorf("Callback fail,%s ", respResult.Message)
	}
	_, err = x.Exec("update task set report_time=?,callback_data=?,result=?,chose_option=?,status=?,updated_by=?,updated_time=? where id=?", nowTime, string(requestBytes), param.Comment, param.ChoseOption, "done", operator, nowTime, taskId)
	if err != nil {
		return fmt.Errorf("Callback succeed,but update database task row fail,%s ", err.Error())
	}
	return nil
}

func getApproveCallbackParam(taskId string) (result models.PluginTaskCreateResp, callbackUrl string, err error) {
	result = models.PluginTaskCreateResp{ResultCode: "0"}
	taskObj, tmpErr := getSimpleTask(taskId)
	if tmpErr != nil {
		return result, callbackUrl, tmpErr
	}
	callbackUrl = taskObj.CallbackUrl
	resultObj := models.PluginTaskCreateOutput{RequestId: taskObj.CallbackRequestId}
	if taskObj.Cache == "" {
		resultObj.Outputs = []*models.PluginTaskCreateOutputObj{{CallbackParameter: taskObj.CallbackParameter, Comment: taskObj.Result, ErrorCode: "0"}}
		result.Results = resultObj
		return
	}
	var taskFormOutput models.PluginTaskFormDto
	err = json.Unmarshal([]byte(taskObj.Cache), &taskFormOutput)
	if err != nil {
		return result, callbackUrl, fmt.Errorf("Try to json unmarshal cache data fail:%s ", err.Error())
	}
	var items []*models.TaskFormItemQueryObj
	err = x.SQL("select t1.*,t2.attr_def_data_type,t2.element_type from form_item t1 left join form_item_template t2 on t1.form_item_template=t2.id where form=?", taskObj.Form).Find(&items)
	if err != nil {
		return result, callbackUrl, fmt.Errorf("Try to query form item fail:%s ", err.Error())
	}
	itemValueMap := make(map[string]interface{})
	for _, v := range items {
		itemValueMap[fmt.Sprintf("%s_%s", v.FormItemTemplate, v.RowDataId)] = v.Value
	}
	for _, formEntity := range taskFormOutput.FormDataEntities {
		for _, itemValueObj := range formEntity.FormItemValues {
			tmpKey := fmt.Sprintf("%s_%s", itemValueObj.FormItemMetaId, itemValueObj.Oid)
			itemValueObj.AttrValue = itemValueMap[tmpKey]
		}
	}
	formBytes, _ := json.Marshal(taskFormOutput)
	resultObj.Outputs = []*models.PluginTaskCreateOutputObj{{CallbackParameter: taskObj.CallbackParameter, Comment: taskObj.Result, ErrorCode: "0", TaskFormOutput: string(formBytes)}}
	result.Results = resultObj
	return
}

func SaveTaskForm(taskId, operator string, param models.TaskApproveParam) error {
	var actions []*execAction
	taskObj, err := getSimpleTask(taskId)
	if err != nil {
		return err
	}
	nowTime := time.Now().Format(models.DateTimeFormat)
	actions = append(actions, &execAction{Sql: "update task set `result`=?,chose_option=?,updated_by=?,updated_time=? where id=?", Param: []interface{}{param.Comment, param.ChoseOption, operator, nowTime, taskId}})
	if taskObj.Request != "" {
		for _, tableForm := range param.Form {
			tmpColumnMap := make(map[string]int)
			for _, title := range tableForm.Title {
				tmpColumnMap[title.Name] = 1
			}
			for _, valueObj := range tableForm.Value {
				for k, v := range valueObj.EntityData {
					if _, b := tmpColumnMap[k]; b {
						actions = append(actions, &execAction{Sql: "update form_item set value=? where form=? and row_data_id=? and name=?", Param: []interface{}{v, taskObj.Form, valueObj.Id, k}})
					}
				}
			}
		}
	}
	return transaction(actions)
}

func ChangeTaskStatus(taskId, operator, operation string) (taskObj models.TaskTable, err error) {
	taskObj, err = getSimpleTask(taskId)
	if err != nil {
		return
	}
	if taskObj.Status == "done" {
		return taskObj, fmt.Errorf("Task aleary done with %s %s ", taskObj.UpdatedBy, taskObj.UpdatedTime)
	}
	var actions []*execAction
	nowTime := time.Now().Format(models.DateTimeFormat)
	if operation == "mark" {
		//if taskObj.Status == "doing" {
		//	return taskObj, fmt.Errorf("Task doing with %s %s ", taskObj.UpdatedBy, taskObj.UpdatedTime)
		//}
		actions = append(actions, &execAction{Sql: "update task set status=?,handler=?,updated_by=?,updated_time=? where id=?", Param: []interface{}{"marked", operator, operator, nowTime, taskId}})
	} else if operation == "start" {
		if operator != taskObj.Handler {
			return taskObj, fmt.Errorf("Task handler is %s ", taskObj.Handler)
		}
		actions = append(actions, &execAction{Sql: "update task set status=?,updated_by=?,updated_time=? where id=?", Param: []interface{}{"doing", operator, nowTime, taskId}})
	} else if operation == "quit" {
		if operator != taskObj.Handler {
			return taskObj, fmt.Errorf("Task handler is %s ", taskObj.Handler)
		}
		actions = append(actions, &execAction{Sql: "update task set status=?,updated_by=?,updated_time=? where id=?", Param: []interface{}{"marked", operator, nowTime, taskId}})
	}
	actions = append(actions, &execAction{Sql: "insert into operation_log(id,task,operation,operator,op_time) value (?,?,?,?,?)", Param: []interface{}{guid.CreateGuid(), taskId, operation, operator, nowTime}})
	err = transaction(actions)
	if err != nil {
		return taskObj, err
	}
	taskObj, _ = getSimpleTask(taskId)
	if taskObj.Status == "created" {
		taskObj.OperationOptions = []string{"mark"}
	} else if taskObj.Status == "marked" || taskObj.Status == "doing" {
		if taskObj.Handler == operator {
			taskObj.OperationOptions = []string{"start"}
		} else {
			taskObj.OperationOptions = []string{"mark"}
		}
	} else {
		taskObj.OperationOptions = []string{}
	}
	return taskObj, nil
}

func buildTaskOperation(taskObj *models.TaskListObj, operator string) {
	if taskObj.Status == "created" {
		taskObj.OperationOptions = []string{"mark"}
	} else if taskObj.Status == "marked" || taskObj.Status == "doing" {
		if taskObj.Handler == operator {
			taskObj.OperationOptions = []string{"start"}
		} else {
			taskObj.OperationOptions = []string{"mark"}
		}
	} else {
		taskObj.OperationOptions = []string{}
	}
}