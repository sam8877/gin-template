package model

import "net/http"

type CommonRes struct {
	Code int    `json:"code"` // http返回码
	Msg  string `json:"msg"`  // 提示信息
	Data any    `json:"data"` // 数据
}

func BuildSuccess(DataIn any) *CommonRes {
	return &CommonRes{
		Code: http.StatusOK,
		Msg:  "success",
		Data: DataIn,
	}
}

func BuildInternalErr(MsgIn string) *CommonRes {
	return &CommonRes{
		Code: http.StatusInternalServerError,
		Msg:  MsgIn,
	}
}

func BuildBadReq(MsgIn string) *CommonRes {
	return &CommonRes{
		Code: http.StatusBadRequest,
		Msg:  MsgIn,
	}
}
