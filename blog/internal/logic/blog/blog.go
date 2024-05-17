package blog

import (
	"context"
	"yijunqiang/gf-micro/blog/internal/model/entity"

	"yijunqiang/gf-micro/blog/api/pbentity"
	"yijunqiang/gf-micro/blog/internal/dao"
	"yijunqiang/gf-micro/blog/internal/model/do"
	"yijunqiang/gf-micro/blog/internal/service"
)

type (
	sBlog struct{}
)

func init() {
	service.RegisterBlog(&sBlog{})
}

func (s *sBlog) Create(ctx context.Context, title string, content string, nickname string) (*entity.Blog, error) {
	blog := &entity.Blog{
		Title:    title,
		Content:  content,
		Nickname: nickname,
	}
	_, err := dao.Blog.Ctx(ctx).Data(blog).Insert()
	return blog, err
}

func (s *sBlog) Edit(ctx context.Context, title string, content string, nickname string) (err error) {
	_, err = dao.Blog.Ctx(ctx).Where(do.Blog{
		Nickname: nickname,
	}).Data(do.Blog{
		Title:   title,
		Content: content,
	}).Update()
	if err != nil {
		return
	}
	return
}

func (s *sBlog) GetById(ctx context.Context, id uint64) (*pbentity.Blog, error) {
	var blog *pbentity.Blog
	err := dao.Blog.Ctx(ctx).Where(do.Blog{
		Id: id,
	}).Scan(&blog)
	return blog, err
}

func (s *sBlog) GetList(ctx context.Context) (list []*pbentity.Blog, err error) {
	list = []*pbentity.Blog{}
	err = dao.Blog.Ctx(ctx).Scan(&list)
	if err != nil {
		return
	}
	return
}

func (s *sBlog) Delete(ctx context.Context, id uint64) (err error) {
	_, err = dao.Blog.Ctx(ctx).Where(do.Blog{
		Id: id,
	}).Delete()
	if err != nil {
		return
	}
	return
}

func (s *sBlog) BatDelete(ctx context.Context, ids []uint64) (batNo string, err error) {
	_, err = dao.Blog.Ctx(ctx).WhereIn("id", ids).Delete()
	if err != nil {
		return
	}
	batNo = "bat12345"
	return
}

func (s *sBlog) GetBatDeleteStatus(ctx context.Context, batNo string) (status string, err error) {
	status = "pending"
	return
}
