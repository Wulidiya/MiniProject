package usecase

import (
	"errors"
	"fmt"
	"miniproject/model"
	"miniproject/database"
)

type TooldetailUsecase interface {
	CreateTooldetail(Tooldetail *model.Tooldetail) error
	GetTooldetail(id uint) (Tool model.Tooldetail, err error)
	GetListTooldetails() (Tools []model.Tooldetail, err error)
	UpdateTooldetail(Tool *model.Tooldetail) (err error)
	DeleteTooldetail(id uint) (err error)
}

type TooldetailUsecase struct {
	ToolRepository database.ToolRepository
}

func NewTooldetailUsecase(TooldetailRepo database.TooldetailRepository) *TooldetailUsecase {
	return &TooldetailUsecase{TooldetailRepository: TooldetailRepo}
}

func (b *TooldetailUsecase) CreateTooldetail(Tooldetail *model.Tooldetail) error {

	// check tool_id cannot be empty
	if Tool.Tool_id == "" {
		return errors.New("tool id cannot be empty")
	}

	// check tool_code
	if tool.tool_code == "" {
		return errors.New("tool creator cannot be empty")
	}

	err := b.tooldetailRepository.Createtooldetail(tooldetail)
	if err != nil {
		return err
	}

	return nil
}

func (b *tooldetailUsecase) Gettooldetail(id uint) (tooldetail model.tooldetail, err error) {
	tooldetail, err = b.tooldetailRepository.Gettooldetail(id)
	if err != nil {
		fmt.Println("Gettooldetail: Error getting tooldetail from database")
		return
	}
	return
}

func (b *tooldetailUsecase) GetListtooldetails() (tooldetails []model.tooldetail, err error) {
	tooldetails, err = b.bookRepository.Gettools()
	if err != nil {
		fmt.Println("GetListtooldetails: Error getting tooldetails from database")
		return
	}
	return
}

func (b *tooldetailUsecase) Updatetooldetail(tooldetail *model.tooldetail) (err error) {
	err = b.toolRepository.Updatetooldetail(tooldetail)
	if err != nil {
		fmt.Println("Updatedetail : Error updating detail, err: ", err)
		return
	}

	return
}

func (b *tooldetailUsecase) Deletetooldetail(id uint) (err error) {
	tooldetail := model.tooldetail{}
	tooldetail.ID = id
	err = b.tooldetailRepository.Deletetooldetail(&tooldetail)
	if err != nil {
		fmt.Println("Deletetooldetail : error deleting tooldetail, err: ", err)
		return
	}

	return
}