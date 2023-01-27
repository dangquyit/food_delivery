package uploadmodel

import "food_delivery/common"

const EntityName = "Upload"

type Upload struct {
	common.SQLModel
	common.Image
}

func (Upload) TableName() string {
	return "uploads"
}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"file is not image",
		"ErrFileIsNotImage",
	)
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save uploaded file",
		"ErrCannotSaveFile",
	)
}
