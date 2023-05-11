package usecase

import (
	"errors"
	"fmt"
	"miniproject/model"
	"miniproject/repository/database"
)

type ReturnUsecase interface {
	CreateReturn(blog *model.Return) error
	GetReturn(id uint) (Return model.Return, err error)
	GetListReturn() (Returns []model.Return, err error)
	UpdateReturn(Return *model.Return) (err error)
	DeleteReturn(id uint) (err error)
}

type ReturnsUsecase struct {
	blogRepository database.ReturnRepository
}

func NewBlogUsecase(blogRepo database.ReturnRepository) *ReturnUsecase {
	return &ReturnUsecase{blogRepository: blogRepo}
}

func (b *ReturnUsecase) CreateBlog(blog *model.Return) error {

	// check toolname cannot be empty
	if Return.ToolName == "" {
		return errors.New("Return toolname cannot be empty")
	}

	// check sizevolume
	if Return.SizeVolume == "" {
		return errors.New("Return content cannot be empty")
	}

	err := b.ReturnRepository.CreateBlog(blog)
	if err != nil {
		return err
	}

	return nil
}

func (b *ReturnUsecase) GetBlog(id uint) (blog model.Return, err error) {
	blog, err = b.ReturnRepository.GetBlog(id)
	if err != nil {
		fmt.Println("GetBlog: Error getting blog from database")
		return
	}
	return
}

func (b *ReturnUsecase) GetListBlogs() (blogs []model.Return, err error) {
	blogs, err = b.ReturnRepository.GetBlogs()
	if err != nil {
		fmt.Println("GetListBlogs: Error getting blogs from database")
		return
	}
	return
}

func (b *ReturnUsecase) UpdateBlog(blog *model.Return) (err error) {
	err = b.ReturnRepository.UpdateBlog(blog)
	if err != nil {
		fmt.Println("UpdateBlog : Error updating blog, err: ", err)
		return
	}

	return
}

func (b *ReturnUsecase) DeleteBlog(id uint) (err error) {
	blog := model.Return{}
	blog.ID = id
	err = b.ReturnRepository.DeleteBlog(&blog)
	if err != nil {
		fmt.Println("DeleteBlog : error deleting blog, err: ", err)
		return
	}

	return
}
