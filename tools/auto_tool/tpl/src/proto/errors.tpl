syntax = "proto3";

package im_error_proto;

option go_package = "{{.Name}}/src/common/commonproto/im_error_proto";

enum ErrCode {
 success = 0; // 成功

    common_unexpected_err = 1; // 未预期的错误

    // 数据库
	db_read_err = 2;
	db_write_err = 3;
}
// 返回客户端错误
message CommonError {
    ErrCode Code = 1;   // 错误码
    // 往返消息id
    uint32 ReqMsgID = 2;
    uint32 ResMsgID = 3;
    // 往返消息数据
    string ReqPayload = 4;
    string ResPayload = 5;
}