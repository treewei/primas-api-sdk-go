package group

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type GetGroupMemberWhitelistResponse struct {
	core.Response
	Data []dtcpv1.GroupMemberWhitelistGet `json:"data"`
}

func GetGroupMemberWhitelist(group_id string, page, pageSize int, application_status string) (*GetGroupMemberWhitelistResponse, error) {
	if group_id == "" {
		return nil, errors.New("group_id is empty")
	}

	queryParams := make(map[string]interface{}, 0)
	url := config.CONST_Server + `/groups/` + group_id + `/whitelist/members`

	if page > 0 {
		queryParams["page"] = page
	}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}
	if application_status != "" {
		if application_status != dtcpv1.CONST_DTCP_ApplicationStatus_Pending &&
			application_status != dtcpv1.CONST_DTCP_ApplicationStatus_Approved &&
			application_status != dtcpv1.CONST_DTCP_ApplicationStatus_Declined {
			return nil, errors.New("application_status is error")
		}
		queryParams["application_status"] = application_status
	}

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj GetGroupMemberWhitelistResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
