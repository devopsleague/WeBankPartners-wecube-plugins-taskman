package com.webank.taskman.service;


import com.baomidou.mybatisplus.extension.service.IService;
import com.webank.taskman.base.QueryResponse;
import com.webank.taskman.domain.FormTemplate;
import com.webank.taskman.dto.req.FormTemplateSaveReqDto;
import com.webank.taskman.dto.resp.FormTemplateRespDto;

public interface FormTemplateService extends IService<FormTemplate> {
    QueryResponse<FormTemplateRespDto> selectFormTemplate(Integer current, Integer limit, FormTemplateSaveReqDto req);

    void deleteFormTemplate(String id);

    FormTemplateRespDto detailFormTemplate(FormTemplateSaveReqDto saveFormTemplateReq);

    FormTemplateRespDto saveFormTemplateByReq(FormTemplateSaveReqDto formTemplateReq);

    FormTemplateRespDto queryDetailByTemp(Integer tempType,String tempId);

}