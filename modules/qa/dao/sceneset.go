// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/pkg/database/dbengine"
)

type SceneSet struct {
	dbengine.BaseModel
	Name        string
	Description string
	SpaceID     uint64
	PreID       uint64
	CreatorID   string
	UpdaterID   string
}

// Test TableName
func (SceneSet) TableName() string {
	return "dice_autotest_scene_set"
}

// Create Scene Set
func (client *DBClient) CreateSceneSet(sceneSet *SceneSet) error {
	return client.Create(sceneSet).Error
}

func (client *DBClient) CountSceneSetByName(name string, spaceId uint64) (int, error) {
	var res int
	if err := client.Model(&SceneSet{}).Where("`space_id` = ? and name = ?", spaceId, name).Count(&res).Error; err != nil {
		return -1, err
	}
	return res, nil
}

// Get Sceneset by id
func (client *DBClient) GetSceneSet(id uint64) (*SceneSet, error) {
	var res SceneSet
	if err := client.First(&res, id).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

//	Get Scenesets by spaceID
func (client *DBClient) SceneSetsBySpaceID(spaceID uint64) ([]SceneSet, error) {
	var res []SceneSet
	if err := client.Where("`space_id` = ?", spaceID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// Update Sceneset
func (client *DBClient) UpdateSceneSet(sceneSet *SceneSet) (*SceneSet, error) {
	err := client.Save(sceneSet).Error
	return sceneSet, err
}

// Delete Sceneset
func (client *DBClient) DeleteSceneSet(sceneSet *SceneSet, scenes []uint64) error {
	return client.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(AutoTestScene{}).Where(scenes).Delete(AutoTestScene{}).Error; err != nil {
			return err
		}
		if err := tx.Where(AutoTestSceneInput{}).Where("scene_id IN (?)", scenes).Delete(AutoTestSceneInput{}).Error; err != nil {
			return err
		}
		if err := tx.Where(AutoTestSceneOutput{}).Where("scene_id IN (?)", scenes).Delete(AutoTestSceneOutput{}).Error; err != nil {
			return err
		}
		if err := tx.Where(AutoTestSceneStep{}).Where("scene_id IN (?)", scenes).Delete(AutoTestSceneStep{}).Error; err != nil {
			return err
		}

		var next SceneSet
		if err := tx.Where("pre_id = ?", sceneSet.ID).Find(&next).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return client.Delete(sceneSet).Error
			}
			return err
		}
		next.PreID = sceneSet.PreID
		if err := tx.Save(&next).Error; err != nil {
			return err
		}
		return client.Delete(sceneSet).Error
	})
}

func (client *DBClient) MoveSceneSet(req apistructs.SceneSetRequest) error {
	return client.Transaction(func(tx *gorm.DB) error {
		var sceneSet, next SceneSet
		if err := tx.Where("id = ?", req.SetID).Find(&sceneSet).Error; err != nil {
			return err
		}

		if err := tx.Where("pre_id = ?", req.SetID).Find(&next).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				goto LABEL1
			}
			return err
		}
		next.PreID = sceneSet.PreID
		if err := tx.Save(&next).Error; err != nil {
			return err
		}
	LABEL1:
		var target SceneSet
		if err := tx.Where("id = ?", req.DropKey).Find(&target).Error; err != nil {
			return err
		}
		if req.Position == 1 {
			sceneSet.PreID = target.ID
			var next SceneSet
			if err := tx.Where("pre_id = ?", req.DropKey).Find(&next).Error; err != nil {
				if gorm.IsRecordNotFoundError(err) {
					goto LABEL2
				}
				return err
			}
			next.PreID = sceneSet.ID
			if err := tx.Save(&next).Error; err != nil {
				return err
			}
			goto LABEL2
		}
		sceneSet.PreID = target.PreID
		target.PreID = sceneSet.ID

		if err := tx.Save(&target).Error; err != nil {
			return err
		}
	LABEL2:
		if err := tx.Save(&sceneSet).Error; err != nil {
			return err
		}
		return nil
	})
}

func (client *DBClient) FindByPreId(id uint64) (*SceneSet, error) {
	var res SceneSet
	if err := client.Where("pre_id = ?", id).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (client *DBClient) GetSceneSetByPreID(preID uint64) (*SceneSet, error) {
	var res SceneSet
	if err := client.Where("pre_id = ?", preID).Find(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

// CheckSceneSetIsExists check the SceneSet is exists
func (client *DBClient) CheckSceneSetIsExists(setID uint64) error {
	var count int64
	if err := client.Model(&SceneSet{}).Where("id = ?", setID).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.Errorf("SceneSet is not exist: %d", setID)
	}

	return nil
}
