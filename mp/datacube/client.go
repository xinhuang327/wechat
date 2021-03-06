// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package datacube

import (
	"net/http"

	"github.com/xinhuang327/wechat/mp"
)

type Client struct {
	*mp.Client
}

// 兼容保留, 建議實際項目全局維護一個 *mp.Client
func NewClient(srv mp.AccessTokenServer, clt *http.Client) Client {
	return Client{
		Client: mp.NewClient(srv, clt),
	}
}
