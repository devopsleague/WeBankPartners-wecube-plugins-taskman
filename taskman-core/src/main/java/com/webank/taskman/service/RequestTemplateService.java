package com.webank.taskman.service;


import com.baomidou.mybatisplus.extension.service.IService;
import com.webank.taskman.base.QueryResponse;
import com.webank.taskman.commons.TaskmanRuntimeException;
import com.webank.taskman.domain.RequestTemplate;
import com.webank.taskman.dto.RequestTemplateDto;
import com.webank.taskman.dto.req.RequestTemplateQueryReqDto;
import com.webank.taskman.dto.req.RequestTemplateSaveReqDto;
import com.webank.taskman.dto.resp.RequestTemplateRespDto;

public interface RequestTemplateService extends IService<RequestTemplate> {

    RequestTemplateDto saveRequestTemplate(RequestTemplateSaveReqDto saveRequestTemplateReq);

    void deleteRequestTemplateService(String id) throws TaskmanRuntimeException;

    QueryResponse<RequestTemplateDto> selectRequestTemplatePage
            (Integer current, Integer limit, RequestTemplateQueryReqDto req) throws TaskmanRuntimeException;

    RequestTemplateRespDto detailRequestTemplate(String id) throws TaskmanRuntimeException;

   }