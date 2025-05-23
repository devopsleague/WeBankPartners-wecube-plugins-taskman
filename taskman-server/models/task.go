package models

import "time"

type TaskTable struct {
	Id                 string   `json:"id" xorm:"'id' pk" primary-key:"id"`
	Name               string   `json:"name" xorm:"name"`
	Description        string   `json:"description" xorm:"description"`
	Form               string   `json:"form" xorm:"form"`
	AttachFile         string   `json:"attachFile" xorm:"attach_file"`
	Status             string   `json:"status" xorm:"status"`
	Version            string   `json:"version" xorm:"version"`
	Request            string   `json:"request" xorm:"request"`
	Parent             string   `json:"parent" xorm:"parent"`
	TaskTemplate       string   `json:"taskTemplate" xorm:"task_template"`
	PackageName        string   `json:"packageName" xorm:"package_name"`
	EntityName         string   `json:"entityName" xorm:"entity_name"`
	ProcDefId          string   `json:"procDefId" xorm:"proc_def_id"`
	ProcDefKey         string   `json:"procDefKey" xorm:"proc_def_key"`
	ProcDefName        string   `json:"procDefName" xorm:"proc_def_name"`
	NodeDefId          string   `json:"nodeDefId" xorm:"node_def_id"`
	NodeName           string   `json:"nodeName" xorm:"node_name"`
	CallbackUrl        string   `json:"callbackUrl" xorm:"callback_url"`
	CallbackParameter  string   `json:"callbackParameter" xorm:"callback_parameter"`
	CallbackData       string   `json:"callbackData" xorm:"callback_data"`
	Emergency          int      `json:"emergency" xorm:"emergency"`
	Result             string   `json:"result" xorm:"result"`
	Cache              string   `json:"cache" xorm:"cache"`
	CallbackRequestId  string   `json:"callbackRequestId" xorm:"callback_request_id"`
	Reporter           string   `json:"reporter" xorm:"reporter"`
	ReportTime         string   `json:"reportTime" xorm:"report_time"`
	ReportRole         string   `json:"reportRole" xorm:"report_role"`
	Handler            string   `json:"handler" xorm:"handler"`
	NextOption         string   `json:"nextOption" xorm:"next_option"`
	ChoseOption        string   `json:"choseOption" xorm:"chose_option"`
	CreatedBy          string   `json:"createdBy" xorm:"created_by"`
	CreatedTime        string   `json:"createdTime" xorm:"created_time"`
	UpdatedBy          string   `json:"updatedBy" xorm:"updated_by"`
	UpdatedTime        string   `json:"updatedTime" xorm:"updated_time"`
	DelFlag            string   `json:"delFlag" xorm:"del_flag"`
	OperationOptions   []string `json:"operationOptions" xorm:"-"`
	ExpireTime         string   `json:"expireTime" xorm:"expire_time"`
	NotifyCount        int      `json:"notifyCount" xorm:"notify_count"`
	TemplateType       int      `json:"templateType" xorm:"template_type"` // 请求模板类型 0表示请求,1表示发布
	Type               string   `json:"type" xorm:"type"`
	Sort               int      `json:"sort" xorm:"sort"`
	TaskResult         string   `json:"taskResult" xorm:"task_result"`
	ConfirmResult      string   `json:"confirmResult" xorm:"confirm_result"`
	RequestCreatedTime string   `json:"requestCreatedTime" xorm:"request_created_time"` // 请求创建时间
}

func (TaskTable) TableName() string {
	return "task"
}

type TaskNotifyTable struct {
	Id                 string `json:"id" xorm:"'id' pk" primary-key:"id"`
	Task               string `json:"task" xorm:"task"`
	DoingNotifyCount   int    `json:"doingNotifyCount" xorm:"'doing_notify_count'"`
	TimeoutNotifyCount int    `json:"timeoutNotifyCount" xorm:"'timeout_notify_count'"`
	ErrMsg             string `json:"errMsg" xorm:"'err_msg'"`
	UpdatedTime        string `json:"updatedTime" xorm:"updated_time"`
}

type TaskDto struct {
	Id           string           `json:"id"`
	Name         string           `json:"name"`
	Type         string           `json:"type"`
	Sort         int              `json:"sort"`
	TaskTemplate string           `json:"taskTemplate"`
	Request      string           `json:"request" xorm:"request"`
	Handles      []*TaskHandleDto `json:"handles"`
}

type TaskListObj struct {
	Id               string       `json:"id" xorm:"'id' pk" primary-key:"id"`
	Name             string       `json:"name" xorm:"name"`
	Description      string       `json:"description" xorm:"description"`
	Status           string       `json:"status" xorm:"status"`
	Request          string       `json:"request" xorm:"request"`
	TaskTemplate     string       `json:"taskTemplate" xorm:"task_template"`
	NodeDefId        string       `json:"nodeDefId" xorm:"node_def_id"`
	NodeName         string       `json:"nodeName" xorm:"node_name"`
	Emergency        int          `json:"emergency" xorm:"emergency"`
	Result           string       `json:"result" xorm:"result"`
	Reporter         string       `json:"reporter" xorm:"reporter"`
	ReportTime       string       `json:"reportTime" xorm:"report_time"`
	ReportRole       string       `json:"reportRole" xorm:"report_role"`
	Handler          string       `json:"handler" xorm:"handler"`
	NextOption       string       `json:"nextOption" xorm:"next_option"`
	ChoseOption      string       `json:"choseOption" xorm:"chose_option"`
	CreatedBy        string       `json:"createdBy" xorm:"created_by"`
	CreatedTime      string       `json:"createdTime" xorm:"created_time"`
	UpdatedBy        string       `json:"updatedBy" xorm:"updated_by"`
	UpdatedTime      string       `json:"updatedTime" xorm:"updated_time"`
	OperationOptions []string     `json:"operationOptions" xorm:"-"`
	ExpireTime       string       `json:"expireTime" xorm:"expire_time"`
	HandleRoles      []string     `json:"handleRoles" xorm:"-"`
	RequestObj       RequestTable `json:"requestObj" xorm:"-"`
	ExpirePercentObj ExpireObj    `json:"expireObj"`
}

type TaskMetaResult struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Data    TaskMetaResultData `json:"data"`
}

type TaskMetaResultData struct {
	FormMetaId    string                `json:"formMetaId"`
	FormItemMetas []*TaskMetaResultItem `json:"formItemMetas"`
}

type TaskMetaResultItem struct {
	FormItemMetaId string `json:"formItemMetaId"`
	PackageName    string `json:"packageName"`
	EntityName     string `json:"entityName"`
	AttrName       string `json:"attrName"`
}

type PluginTaskCreateRequest struct {
	RequestId      string                        `json:"requestId"`
	DueDate        string                        `json:"dueDate"`
	AllowedOptions []string                      `json:"allowedOptions"`
	Inputs         []*PluginTaskCreateRequestObj `json:"inputs"`
}

type PluginTaskCreateRequestObj struct {
	CallbackParameter string `json:"callbackParameter"`
	ProcInstId        string `json:"procInstId"`
	CallbackUrl       string `json:"callbackUrl"`
	Reporter          string `json:"reporter"`
	Handler           string `json:"handler"`
	RoleName          string `json:"roleName"`
	TaskName          string `json:"taskName"`
	TaskDescription   string `json:"taskDescription"`
	TaskFormInput     string `json:"taskFormInput"`
}

type PluginTaskCreateResp struct {
	ResultCode    string                 `json:"resultCode"`
	ResultMessage string                 `json:"resultMessage"`
	Results       PluginTaskCreateOutput `json:"results"`
}

type PluginTaskCreateOutput struct {
	RequestId      string                       `json:"requestId"`
	AllowedOptions []string                     `json:"allowedOptions,omitempty"`
	Outputs        []*PluginTaskCreateOutputObj `json:"outputs"`
}

type PluginTaskCreateOutputObj struct {
	CallbackParameter string `json:"callbackParameter"`
	Comment           string `json:"comment"`
	TaskFormOutput    string `json:"taskFormOutput"`
	ErrorCode         string `json:"errorCode"`
	ErrorMessage      string `json:"errorMessage"`
	ErrorDetail       string `json:"errorDetail,omitempty"`
}

type PluginTaskFormDto struct {
	FormMetaId       string                  `json:"formMetaId"`
	ProcDefId        string                  `json:"procDefId"`
	ProcDefKey       string                  `json:"procDefKey"`
	ProcInstId       interface{}             `json:"procInstId"`
	ProcInstKey      string                  `json:"procInstKey"`
	TaskNodeDefId    string                  `json:"taskNodeDefId"`
	TaskNodeInstId   interface{}             `json:"taskNodeInstId"`
	FormDataEntities []*PluginTaskFormEntity `json:"formDataEntities"`
}

type PluginTaskFormEntity struct {
	FormMetaId       string                 `json:"formMetaId"`
	PackageName      string                 `json:"packageName"`
	EntityName       string                 `json:"entityName"`
	Oid              string                 `json:"oid"`
	EntityDataId     string                 `json:"entityDataId"`
	FullEntityDataId string                 `json:"fullEntityDataId"`
	EntityDataState  string                 `json:"entityDataState"`
	EntityDataOp     string                 `json:"entityDataOp"`
	BindFlag         string                 `json:"bindFlag"`
	FormItemValues   []*PluginTaskFormValue `json:"formItemValues"`
}

type PluginTaskFormValue struct {
	FormItemMetaId   string      `json:"formItemMetaId"`
	PackageName      string      `json:"packageName"`
	EntityName       string      `json:"entityName"`
	AttrName         string      `json:"attrName"`
	Oid              string      `json:"oid"`
	EntityDataId     string      `json:"entityDataId"`
	FullEntityDataId string      `json:"fullEntityDataId"`
	AttrValue        interface{} `json:"attrValue"`
}

type CallbackResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type TaskQueryResult struct {
	TimeStep []*TaskQueryTimeStep `json:"timeStep"`
	Data     []*TaskQueryObj      `json:"data"`
}

type TaskQueryTimeStep struct {
	RequestTemplateId string `json:"requestTemplateId"`
	TaskTemplateId    string `json:"taskTemplateId"`
	Name              string `json:"name"`
	Active            bool   `json:"active"`
	Done              bool   `json:"done"`
}

type TaskQueryObj struct {
	RequestId       string                    `json:"requestId"`
	RequestName     string                    `json:"requestName"`
	RequestTemplate string                    `json:"requestTemplate"`
	ProcInstanceId  string                    `json:"procInstanceId"`
	Description     string                    `json:"description"`
	TaskId          string                    `json:"taskId"`
	TaskName        string                    `json:"taskName"`
	Reporter        string                    `json:"reporter"`
	ReportTime      string                    `json:"reportTime"`
	Comment         string                    `json:"comment"`
	AttachFiles     []*AttachFileTable        `json:"attachFiles"`
	Editable        bool                      `json:"editable"`
	Status          string                    `json:"status"`
	NextOption      []string                  `json:"nextOption"`
	ChoseOption     string                    `json:"choseOption"`
	ExpireTime      string                    `json:"expireTime"`
	ExpectTime      string                    `json:"expectTime"`
	Handler         string                    `json:"handler"`
	HandleTime      string                    `json:"handleTime"`
	FormData        []*RequestPreDataTableObj `json:"formData"`
	IsHistory       bool                      `json:"isHistory"`
	HandleRoleName  string                    `json:"handleRoleName"`
	CreatedTime     string                    `json:"createTime"`
	RollbackDesc    string                    `json:"rollbackDesc"`
}

type TaskApproveParam struct {
	FormData      []*RequestPreDataTableObj `json:"formData"`
	Comment       string                    `json:"comment"`
	ChoseOption   string                    `json:"choseOption"`
	TaskHandleId  string                    `json:"taskHandleId"`  // 任务处理ID
	ProcDefResult string                    `json:"procDefResult"` // 编排选项
}

type ProcessTaskFormParam struct {
	Task              *TaskTable
	Operator          string
	TaskApproveParam  *TaskApproveParam
	RequestPreDataDto *RequestPreDataDto
	RequestId         string
	FormData          []*RequestPreDataTableObj
}

type TaskHandlerQueryData struct {
	Id          string `xorm:"id"`
	Handler     string `xorm:"handler"`
	DisplayName string `xorm:"display_name"`
}

type TaskHandleUpdateParam struct {
	TaskId           string `json:"taskId"`
	TaskHandleId     string `json:"taskHandleId"` // 任务处理Id
	LatestUpdateTime string `json:"latestUpdateTime"`
	ChangeReason     string `json:"changeReason"` //变更原因: assign 系统分配、claim 主动领取、give 转给我
}

type RequestPoolDataQueryRow struct {
	FormId           string    `xorm:"form_id"`
	Task             string    `xorm:"task"`
	FormTemplate     string    `xorm:"form_template"`
	ItemGroup        string    `xorm:"item_group"`
	ItemGroupType    string    `xorm:"item_group_type"`
	DataId           string    `xorm:"data_id"`
	FormItemId       string    `xorm:"form_item_id"`
	FormItemTemplate string    `xorm:"form_item_template"`
	Name             string    `xorm:"name"`
	Value            string    `xorm:"value"`
	UpdatedTime      time.Time `xorm:"updated_time"`
}

type RequestPoolDataQueryRows []*RequestPoolDataQueryRow

func (r RequestPoolDataQueryRows) DataParse() []*RequestPoolForm {
	var forms []*RequestPoolForm
	formItemMap := make(map[string][]*RequestPoolDataQueryRow)
	for _, v := range r {
		if itemList, ok := formItemMap[v.FormId]; ok {
			formItemMap[v.FormId] = append(itemList, v)
		} else {
			forms = append(forms, &RequestPoolForm{
				FormId:        v.FormId,
				Task:          v.Task,
				FormTemplate:  v.FormTemplate,
				ItemGroup:     v.ItemGroup,
				ItemGroupType: v.ItemGroupType,
				DataId:        v.DataId,
			})
			formItemMap[v.FormId] = []*RequestPoolDataQueryRow{v}
		}
	}
	for _, v := range forms {
		v.Items = formItemMap[v.FormId]
	}
	return forms
}

type RequestPoolForm struct {
	FormId        string                     `json:"formId"`
	Task          string                     `json:"task"`
	FormTemplate  string                     `json:"formTemplate"`
	ItemGroup     string                     `json:"itemGroup"`
	ItemGroupType string                     `json:"itemGroupType"`
	DataId        string                     `json:"dataId"`
	Items         []*RequestPoolDataQueryRow `json:"items"`
}
