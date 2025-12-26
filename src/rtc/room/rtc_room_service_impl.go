package room

import (
	"encoding/json"
	"strconv"

	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
	"github.com/netease-im/yunxin-server-sdk-golang/src/core/http"
	. "github.com/netease-im/yunxin-server-sdk-golang/src/rtc/result"
)

// rtcRoomService RTC房间服务实现
type rtcRoomService struct {
	httpClient core.YunxinApiHttpClient
}

// NewRtcRoomService 创建RTC房间服务实例
func NewRtcRoomService(httpClient core.YunxinApiHttpClient) RtcRoomService {
	return &rtcRoomService{
		httpClient: httpClient,
	}
}

// CreateRoom 创建房间
func (s *rtcRoomService) CreateRoom(request *RtcCreateRoomRequest) (*RtcResult[*RtcCreateRoomResponse], error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, CreateRoom, nil, nil, string(body))
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtcCreateRoomResponse](apiResponse)
}

// GetRoomByCid 根据CID获取房间信息
func (s *rtcRoomService) GetRoomByCid(request *RtcGetRoomByCidRequest) (*RtcResult[*RtcGetRoomResponse], error) {
	pathParams := map[string]string{
		"cid": strconv.FormatInt(request.Cid, 10),
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.GET, GetRoomByCid, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtcGetRoomResponse](apiResponse)
}

// GetRoomByCname 根据CNAME获取房间信息
func (s *rtcRoomService) GetRoomByCname(request *RtcGetRoomByCnameRequest) (*RtcResult[*RtcGetRoomResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.GET, GetRoomByCname, nil, queryString, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtcGetRoomResponse](apiResponse)
}

// ListRoomMembersV2 V2版本列出房间成员
func (s *rtcRoomService) ListRoomMembersV2(request *RtcListRoomMembersRequestV2) (*RtcResult[*RtcListRoomMembersResponse], error) {
	uri := ListMembersV2
	pathParams := map[string]string{
		"cid": strconv.FormatInt(request.Cid, 10),
	}

	if request.Uid != 0 {
		pathParams["uid"] = strconv.FormatUint(request.Uid, 10)
		uri = ListMembersV2WithUid
	}

	queryString := make(map[string]string)
	if request.UserRole != nil {
		queryString["userRole"] = strconv.Itoa(*request.UserRole)
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.GET, uri, pathParams, queryString, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtcListRoomMembersResponse](apiResponse)
}

// ListRoomMembersV3 V3版本列出房间成员
func (s *rtcRoomService) ListRoomMembersV3(request *RtcListRoomMembersRequestV3) (*RtcResult[*RtcListRoomMembersResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	if request.Uid != 0 {
		queryString["uid"] = strconv.FormatUint(request.Uid, 10)
	}

	if request.UserRole != nil {
		queryString["userRole"] = strconv.Itoa(*request.UserRole)
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.GET, ListMembersV3, nil, queryString, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtcListRoomMembersResponse](apiResponse)
}

// AddMemberToKicklistV2 V2版本添加成员到踢出列表
func (s *rtcRoomService) AddMemberToKicklistV2(request *RtcAddMemberToKicklistRequestV2) (*RtcResult[*RtcAddMemberToKicklistResponse], error) {
	uri := AddMemberToKicklistV2
	pathParams := map[string]string{
		"cid": strconv.FormatInt(request.Cid, 10),
		"uid": strconv.FormatUint(request.Uid, 10),
	}

	if request.Duration != nil {
		uri = AddMemberToKicklistV2WithDuration
		pathParams["duration"] = strconv.FormatUint(*request.Duration, 10)
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, uri, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtcAddMemberToKicklistResponse](apiResponse)
}

// AddMemberToKicklistV3 V3版本添加成员到踢出列表
func (s *rtcRoomService) AddMemberToKicklistV3(request *RtcAddMemberToKicklistRequestV3) (*RtcResult[*RtcAddMemberToKicklistResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
		"uid":   strconv.FormatUint(request.Uid, 10),
	}

	if request.Duration != nil {
		queryString["duration"] = strconv.FormatUint(*request.Duration, 10)
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.POST, AddMemberToKicklistV3, nil, queryString, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtcAddMemberToKicklistResponse](apiResponse)
}

// DeleteRoomV2 V2版本删除房间
func (s *rtcRoomService) DeleteRoomV2(request *RtcDeleteRoomRequestV2) (*RtcResult[*RtcDeleteRoomResponse], error) {
	pathParams := map[string]string{
		"cid": strconv.FormatInt(request.Cid, 10),
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.DELETE, DeleteRoomV2, pathParams, nil, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtcDeleteRoomResponse](apiResponse)
}

// DeleteRoomV3 V3版本删除房间
func (s *rtcRoomService) DeleteRoomV3(request *RtcDeleteRoomRequestV3) (*RtcResult[*RtcDeleteRoomResponse], error) {
	queryString := map[string]string{
		"cname": request.Cname,
	}

	apiResponse, err := s.httpClient.ExecuteJson(http.DELETE, DeleteRoomV3, nil, queryString, "")
	if err != nil {
		return nil, err
	}

	return ConvertToRtcResult[*RtcDeleteRoomResponse](apiResponse)
}
