### 获取 jsapi_ticket 示例
```Go
package main

import (
	"fmt"

	"github.com/xinhuang327/wechat/corp"
	"github.com/xinhuang327/wechat/corp/jssdk"
)

var AccessTokenServer = corp.NewDefaultAccessTokenServer("corpId", "corpSecret", nil)
var CorpClient = corp.NewClient(AccessTokenServer, nil)
var TicketServer = jssdk.NewDefaultTicketServer(CorpClient)

func main() {
	fmt.Println(TicketServer.Ticket())
}
```