/**
 * name: middleware cors
 * author: jie
 * date: 2019-6-20
 * note: 中间件
 */

package cors

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/context"
)

func Mycors() context.Handler {
	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //允许通过的主机名称
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
		//AllowCredentials: true,
	})
	return crs
}

/*AllowAllOrigins:  true,
AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
AllowCredentials: true,*/
