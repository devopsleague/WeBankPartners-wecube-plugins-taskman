package com.webank.taskman.dto;


import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

import java.util.Date;
import java.util.StringJoiner;

@ApiModel
public class TaskInfoDTO {


    @ApiModelProperty(value = "任务id",position = 1)
    private String id;
    @ApiModelProperty(value = "请求id",position = 2)
    private String requestId;
    @ApiModelProperty(value = "任务模板id",position = 3)
    private String taskTempId;
    @ApiModelProperty(value = "前置任务ID",position = 4)
    private String parentId;
    @ApiModelProperty(value = "流程节点ID",position = 5)
    private String nodeDefId;
    @ApiModelProperty(value = "流程节点名称",position = 6)
    private String nodeName;
    @ApiModelProperty(value = "任务名称",position = 7)
    private String name;
    @ApiModelProperty(value = "描述",position = 8)
    private String description;
    @ApiModelProperty(value = "处理人",position = 9)
    private String reporter;
    @ApiModelProperty(value = "处理时间",position = 10)
    private Date reportTime;
    @ApiModelProperty(value = "紧急程度",position = 11)
    private String emergency;
    @ApiModelProperty(value = "处理结果",position = 12)
    private String result;
    @ApiModelProperty(value = "任务状态",position = 13)
    private Integer status;
    @ApiModelProperty(value = "上报角色",position = 14)
    private String reportRole;
    @ApiModelProperty(value = "附件ID",position = 14)
    private String attachFileId;

    public String getId() {
        return id;
    }

    public TaskInfoDTO setId(String id) {
        this.id = id;
        return this;
    }

    public String getRequestId() {
        return requestId;
    }

    public TaskInfoDTO setRequestId(String requestId) {
        this.requestId = requestId;
        return this;
    }

    public String getTaskTempId() {
        return taskTempId;
    }

    public TaskInfoDTO setTaskTempId(String taskTempId) {
        this.taskTempId = taskTempId;
        return this;
    }

    public String getParentId() {
        return parentId;
    }

    public TaskInfoDTO setParentId(String parentId) {
        this.parentId = parentId;
        return this;
    }

    public String getNodeDefId() {
        return nodeDefId;
    }

    public TaskInfoDTO setNodeDefId(String nodeDefId) {
        this.nodeDefId = nodeDefId;
        return this;
    }

    public String getNodeName() {
        return nodeName;
    }

    public TaskInfoDTO setNodeName(String nodeName) {
        this.nodeName = nodeName;
        return this;
    }

    public String getName() {
        return name;
    }

    public TaskInfoDTO setName(String name) {
        this.name = name;
        return this;
    }

    public String getDescription() {
        return description;
    }

    public TaskInfoDTO setDescription(String description) {
        this.description = description;
        return this;
    }

    public String getReporter() {
        return reporter;
    }

    public TaskInfoDTO setReporter(String reporter) {
        this.reporter = reporter;
        return this;
    }

    public Date getReportTime() {
        return reportTime;
    }

    public TaskInfoDTO setReportTime(Date reportTime) {
        this.reportTime = reportTime;
        return this;
    }

    public String getEmergency() {
        return emergency;
    }

    public TaskInfoDTO setEmergency(String emergency) {
        this.emergency = emergency;
        return this;
    }

    public String getResult() {
        return result;
    }

    public TaskInfoDTO setResult(String result) {
        this.result = result;
        return this;
    }

    public Integer getStatus() {
        return status;
    }

    public TaskInfoDTO setStatus(Integer status) {
        this.status = status;
        return this;
    }

    public String getReportRole() {
        return reportRole;
    }

    public TaskInfoDTO setReportRole(String reportRole) {
        this.reportRole = reportRole;
        return this;
    }

    public String getAttachFileId() {
        return attachFileId;
    }

    public TaskInfoDTO setAttachFileId(String attachFileId) {
        this.attachFileId = attachFileId;
        return this;
    }

    @Override
    public String toString() {
        return new StringJoiner(", ", TaskInfoDTO.class.getSimpleName() + "[", "]")
                .add("id='" + id + "'")
                .add("requestId='" + requestId + "'")
                .add("taskTempId='" + taskTempId + "'")
                .add("parentId='" + parentId + "'")
                .add("nodeDefId='" + nodeDefId + "'")
                .add("nodeName='" + nodeName + "'")
                .add("name='" + name + "'")
                .add("description='" + description + "'")
                .add("reporter='" + reporter + "'")
                .add("reportTime=" + reportTime)
                .add("emergency='" + emergency + "'")
                .add("result='" + result + "'")
                .add("status=" + status)
                .add("reportRole='" + reportRole + "'")
                .add("attachFileId='" + attachFileId + "'")
                .toString();
    }
}