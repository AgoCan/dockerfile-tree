package service

import (
	"backend/models"
	"backend/serializer"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/jmoiron/sqlx"
)

/*
构建镜像：
	1. 根据层级的ID进行构建，那先后顺序怎么解决？第一步先查库？
	2. 获取到Dockerfile，并且生成记录
	3. 如果成功，记录recored和全部的dockerfile。其中dockerfile直接变成一个记录即可
	4. 记录每一层，如果是最后一层记录，标识，方便查询
	5. 构建的时候是后台进行，必须马上返回。也就是发起的时候应该怎么记录
*/

// BuildJob 构建镜像
type BuildJob struct {
	LevelIDs []int `json:"level_ids" binding:"required"`
	UseProxy bool  `json:"use_proxy" binding:"required"`
}

// build 构建镜像
func (b *BuildJob) build() {
	// 构建是串行的
}

// build 推送镜像
func (b *BuildJob) push(imageName string, authConfig types.AuthConfig) (reader io.ReadCloser, err error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return nil, err
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	reader, err = cli.ImagePush(ctx, imageName, types.ImagePushOptions{RegistryAuth: authStr})
	return reader, err
}

// log 记录构建日志
func (b *BuildJob) log() {
	// 写在文件里面
}

// GetLevels 获取dockerfile
func (b *BuildJob) getLevels() (levelList []models.LevelDockerfile, err error) {
	// 获取level
	sqlStr := `select 
		level.id "level.id",
		level.updated_at "level.updated_at",level.name "level.name",level.order_id "level.order_id",
		level.parent_id "level.parent_id",dockerfile.dockerfile "dockerfile.dockerfile"
	from 
		level INNER JOIN dockerfile ON dockerfile.level_id = level.id
	where 
		level.id in (?); `
	query, args, err := sqlx.In(sqlStr, b.LevelIDs)
	if err != nil {
		return levelList, err
	}
	rows, err := models.DB.Queryx(query, args...)
	// orderList := []map[int]int{}
	LevelIDList := map[int][]int{}
	for rows.Next() {
		var (
			LevelID         int
			LevelUpdateTime time.Time
			LevelName       string
			LevelOrderID    int
			LevelParentID   int
			Dockerfile      string
		)
		if rows.Scan(&LevelID, &LevelUpdateTime, &LevelName, &LevelOrderID, &LevelParentID, &Dockerfile) != nil {
			break
		}
		// 如果ID为0，代表是顶层，可以报错了。顶层是没有dockerfile的
		// 不为0的时候，应该获取到顶层的name，做镜像名使用

		if LevelParentID == 0 {

		} else {
			// select * from level AS l1 INNER JOIN level AS l2 ON l2.parent_id = l1.id where l2.id in (LevelIDList)
			// 返回等长的切片，并且已经排完序
			// 轮询结构体，然后根据排序 [1,4,2,3,5]
			// 替换结构体顺序
			// 再次轮询结构体进行构建，并且结构体应该有对应的名称，制作镜像名和ENV
			// framework-tensorflow1.13-openvino,再次判断 上层的
			if _, ok := LevelIDList[LevelOrderID]; ok {
				LevelIDList[LevelOrderID] = append(LevelIDList[LevelOrderID], LevelID)
			} else {
				LevelIDList[LevelOrderID] = []int{LevelID} // 对id进行排序
			}
			levelList = append(levelList, models.LevelDockerfile{
				ID:         LevelID,
				UpdatedAt:  LevelUpdateTime,
				Name:       LevelName,
				OrderID:    LevelOrderID,
				ParentID:   LevelParentID,
				Dockerfile: Dockerfile,
			})
		}

	}
	return levelList, nil
}

// getResources 创建数据库记录
func (b *BuildJob) getResources() {
	// 获取资源，然后构建的时候进行排除不替换的镜像。并且只有加在资源表的才保证不替换
}

// make  制作level_task日志
func (b *BuildJob) makeLevelTask() {

}

// getProxy 制作环境变量，主要是获取proxy
func (b *BuildJob) getProxy() (buildArgs map[string]string, err error) {
	var (
		httpProxyStr  string
		httpsProxyStr string
		noProxyStr    string
	)
	err = models.DB.Get(&httpProxyStr, "select config_value from config where config_key = http_proxy")
	if err != nil {
		return nil, err
	}
	err = models.DB.Get(&httpsProxyStr, "select config_value from config where config_key = https_proxy")
	if err != nil {
		return nil, err
	}
	err = models.DB.Get(&noProxyStr, "select config_value from config where config_key = no_proxy")
	if err != nil {
		return nil, err
	}
	buildArgs["http_proxy"] = httpProxyStr
	buildArgs["https_proxy"] = httpsProxyStr
	buildArgs["no_proxy"] = noProxyStr
	return buildArgs, nil

}

// getRegistryConfig 制作docker认证信息
func (b *BuildJob) getRegistryConfig() (authConfig types.AuthConfig, err error) {
	var (
		username string
		password string
	)
	err = models.DB.Get(&username, "select config_value from config where config_key = registry_username")
	if err != nil {
		return authConfig, err
	}
	err = models.DB.Get(&username, "select config_value from config where config_key = registry_password")
	if err != nil {
		return authConfig, err
	}
	authConfig.Username = username
	authConfig.Password = password
	return authConfig, nil
}

// makeImageInfo 制作镜像名
func (b *BuildJob) makeImageInfo() {
	// 镜像名按照首层的首字母+版本号 ccudnn6.5.4-p36-o4.1.1-v-i-ftf1.13-l-c:latest
	// 制作镜像信息,添加到镜像env当中      cuda-cudnn6.5.4-python36-openvino4.1.1-video-ice-frameworktensorflow1.13-lab-coding-env

}

// record 记录构建记录
func (b *BuildJob) record() {

}

// Create 创建数据库记录
func (b *BuildJob) Create() serializer.Response {
	res, err := b.getLevels()
	fmt.Println(res, err)
	return serializer.Success("ok")
}
