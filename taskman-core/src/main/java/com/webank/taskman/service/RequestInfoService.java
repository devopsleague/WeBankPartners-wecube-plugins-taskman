package com.webank.taskman.service;


import com.baomidou.mybatisplus.extension.service.IService;
import com.webank.taskman.base.QueryResponse;
import com.webank.taskman.commons.TaskmanRuntimeException;
import com.webank.taskman.domain.RequestInfo;
import com.webank.taskman.dto.CreateTaskDto;
import com.webank.taskman.dto.req.RequestInfoQueryReqDto;
import com.webank.taskman.dto.req.RequestInfoSaveReqDto;
import com.webank.taskman.dto.resp.RequestInfoResqDto;
import com.webank.taskman.support.core.dto.DynamicWorkflowInstCreationInfoDto;
import com.webank.taskman.support.core.dto.DynamicWorkflowInstInfoDto;

public interface RequestInfoService extends IService<RequestInfo> {

    QueryResponse<RequestInfoResqDto> selectRequestInfoPage
            (Integer current, Integer limit, RequestInfoQueryReqDto req) throws TaskmanRuntimeException;

    RequestInfoResqDto saveRequestInfoByDto(CreateTaskDto req);

    DynamicWorkflowInstInfoDto createNewWorkflowInstance(CreateTaskDto req);


    RequestInfoResqDto selectDetail(String id);
}