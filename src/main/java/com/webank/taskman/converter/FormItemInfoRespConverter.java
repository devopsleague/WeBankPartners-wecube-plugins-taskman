package com.webank.taskman.converter;
import com.webank.taskman.base.BaseConverter;
import com.webank.taskman.domain.FormItemInfo;
import com.webank.taskman.dto.req.SaveFormItemInfoReq;
import com.webank.taskman.dto.resp.FormItemInfoResp;
import org.mapstruct.Mapper;
import org.mapstruct.ReportingPolicy;

@Mapper(componentModel = "spring",uses = {},unmappedTargetPolicy = ReportingPolicy.IGNORE)
public interface FormItemInfoRespConverter extends BaseConverter<FormItemInfoResp, FormItemInfo> {


}