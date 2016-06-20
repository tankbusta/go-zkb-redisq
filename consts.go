package zkbq

import "fmt"

var (
	// ZkillboardServer is the full address of the ZKB RedisQ server
	ZkillboardServer = "https://redisq.zkillboard.com/listen.php"
	// GoZKBPollerVer is the float version of the go-zkb code
	GoZKBPollerVer = 1.0
	// GoZKBStringVer is the string version of the go-zkb code
	GoZKBStringVer = fmt.Sprintf("go-zkb/%f", GoZKBPollerVer)
)
