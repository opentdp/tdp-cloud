package member

import "tdp-cloud/core/dborm"

type CreateTATParam struct {
	UserId      uint   `json:"userId"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

func CreateTAT(post *CreateTATParam) error {

	result := dborm.Db.Create(&dborm.TAT{
		UserId:      post.UserId,
		Name:        post.Name,
		Description: post.Description,
		Content:     post.Content,
	})

	return result.Error

}

type UpdateTATParam struct {
	ID          uint   `json:"id"  binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

func UpdateTAT(post *UpdateTATParam) error {
	result := dborm.Db.Model(&dborm.TAT{}).Where("id = ?", post.ID).Updates(dborm.TAT{Name: post.Name, Content: post.Content, Description: post.Description})
	return result.Error
}

func ListTAT(uid uint) ([]*dborm.TAT, error) {
	var tats []*dborm.TAT
	result := dborm.Db.Find(&tats, "user_id = ?", uid)
	return tats, result.Error
}

func InfoTAT(id int) (dborm.TAT, error) {
	var tat dborm.TAT
	result := dborm.Db.Find(&tat, "id = ?", id)
	return tat, result.Error
}

func DeleteTAT(id int) error {
	result := dborm.Db.Delete(&dborm.TAT{}, id)
	return result.Error
}
