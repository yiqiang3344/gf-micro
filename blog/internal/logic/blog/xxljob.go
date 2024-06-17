package blog

import (
	"context"
	"fmt"
	"github.com/xxl-job/xxl-job-executor-go"
	"yijunqiang/gf-micro/blog/internal/dao"
	"yijunqiang/gf-micro/blog/internal/model/do"
	"yijunqiang/gf-micro/blog/internal/model/entity"
)

func (s *sBlog) Stats(ctx context.Context, param *xxl.RunReq) (msg string) {
	//统计每个用户的博客数量
	ret, err := dao.Blog.Ctx(ctx).Fields("nickname,count(1) as cnt").Group("nickname").All()
	if err != nil {
		msg = fmt.Sprintf("统计失败:%v", err)
		return
	}
	for _, v := range ret.List() {
		var stats *entity.Stats
		err1 := dao.Stats.Ctx(ctx).Where("nickname", v["nickname"]).Limit(1).Scan(&stats)
		if err1 != nil {
			msg = fmt.Sprintf("统计失败:%v", err)
			return
		}
		if stats != nil {
			_, err1 = dao.Stats.Ctx(ctx).Where("id", stats.Id).Data(do.Stats{
				BlogCnt: v["cnt"],
			}).Update()
		} else {
			_, err1 = dao.Stats.Ctx(ctx).Data(do.Stats{
				Nickname: v["nickname"],
				BlogCnt:  v["cnt"],
			}).Insert()
		}
		if err1 != nil {
			msg = fmt.Sprintf("统计失败:%v", err)
			return
		}
	}
	msg = fmt.Sprintf("统计成功,数量:%d", ret.Len())
	return
}
