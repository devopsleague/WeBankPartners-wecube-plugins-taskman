package com.webank.taskman.dto.req;

import com.webank.taskman.domain.FormInfo;
import com.webank.taskman.domain.FormItemInfo;

import java.util.List;

public class ProcessingTasksReq {
    private String recordId;

    private List<FormItemInfoReq> formItemInfoList;

    public String getRecordId() {
        return recordId;
    }

    public void setRecordId(String recordId) {
        this.recordId = recordId;
    }

    public List<FormItemInfoReq> getFormItemInfoList() {
        return formItemInfoList;
    }

    public void setFormItemInfoList(List<FormItemInfoReq> formItemInfoList) {
        this.formItemInfoList = formItemInfoList;
    }

    @Override
    public String toString() {
        return "ProcessingTasksReq{" +
                "recordId='" + recordId + '\'' +
                ", formItemInfoList=" + formItemInfoList +
                '}';
    }
}