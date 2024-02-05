# Tinder Matching Service

可以新增使用者，並且和已有資料配對的服務。

配對條件為：男性只能與比自己矮的女性配對，女性只能和比自己高的男性配對。

每配對成功一次，就會減少一次約會機會，減到0的使用者，會從系統中刪除。

## API Document

[API 文件](https://gitlab.com/way11229/tinder_matching/-/blob/main/doc/service.swagger.json?ref_type=heads)

亦可參考./doc/service.swagger.json

Error List

3: invalid parameters

5: not found

13: internal server error

|  code   | error message  |
|  ----  | ----  |
| 3  | miss required parameters |
| 3  | user id is invalid |
| 3  | user name is invalid |
| 3  | user height is invalid |
| 3  | user gender is invalid |
| 3  | user number of wanted dates is invalid |
| 5  | record not found |
| 13  | internal server error |

## System Design Documentation

1. 該服務使用go 1.21.6
2. 程式架構參考 [go-clean-arch](https://github.com/bxcodec/go-clean-arch)
3. 實現grpc接口，監控9000 port; 並使用[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)實現http接口及Restful API，監控8080 port
4. 使用[go-memdb](https://github.com/hashicorp/go-memdb)存取資料
5. 使用[mockery](https://github.com/vektra/mockery)撰寫測試
6. 使用alpine3.19作為執行容器

### Structured Project Layout

![image](https://github.com/way11229/tinder_matching/blob/main/tinder_matching_struct_project_layout.png)

### API Time Complexity

該服務使用go-memdb存取資料，go-memdb底層使用[radix tree](https://github.com/hashicorp/go-immutable-radix)資料結構，時間複雜度為O(k)。故三隻API的時間複雜度均為O(k)。
