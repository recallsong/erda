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

package dbclient

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/erda-project/erda/apistructs"
)

// 检出 major 下所有版本号, 判断与给定的版本号是否冲突
// ok == true, 表示无冲突
func CheckVersionConflict(orgID uint64, assetID string, major, minor, patch uint64) (result [][3]uint64, ok bool, err error) {
	var (
		versions   []*apistructs.APIAssetVersionsModel
		curVersion = [3]uint64{major, minor, patch}
	)
	if err := DB.Where(map[string]interface{}{
		"org_id":   orgID,
		"asset_id": assetID,
		"major":    major,
	}).Order("major DESC, minor DESC, patch DESC").
		Find(&versions).
		Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, true, nil // 没有任何 version, 则无冲突
		}
		return nil, false, err
	}
	if len(versions) == 0 {
		return nil, true, nil // 没有任何 version, 则无冲突
	}
	for _, m := range versions {
		v := [3]uint64{m.Major, m.Minor, m.Patch}
		result = append(result, v)
		if curVersion == v {
			return result, false, nil // 有冲突
		}
	}
	return result, true, nil
}

func PagingAPIAssetVersions(req *apistructs.PagingAPIAssetVersionsReq) (uint64,
	[]*apistructs.APIAssetVersionsModel,
	error) {
	var (
		total    uint64
		versions []*apistructs.APIAssetVersionsModel
	)

	sq := Sq()

	if req.OrgID != 0 {
		sq = sq.Where("org_id = ?", req.OrgID)
	}
	sq = sq.Where("asset_id = ?", req.URIParams.AssetID)
	if req.QueryParams.Major != nil {
		sq = sq.Where("major = ?", *req.QueryParams.Major)
		if req.QueryParams.Minor != nil {
			sq = sq.Where("minor = ?", *req.QueryParams.Minor)
		}
	}

	if err := sq.Offset((req.QueryParams.PageNo - 1) * req.QueryParams.PageSize).Limit(req.QueryParams.PageSize).
		Order("id DESC").Find(&versions).
		Offset(0).Limit(-1).Count(&total).
		Error; err != nil {
		return 0, nil, err
	}

	return total, versions, nil
}

func GetAPIAssetVersion(req *apistructs.GetAPIAssetVersionReq) (*apistructs.APIAssetVersionsModel, error) {
	var version apistructs.APIAssetVersionsModel
	if err := Sq().First(&version, map[string]interface{}{
		"id":       req.URIParams.VersionID,
		"asset_id": req.URIParams.AssetID,
		"org_id":   req.OrgID,
	}).Error; err != nil {
		return nil, err
	}
	return &version, nil
}

// 根据给定主键(id)删除 APIAssetVersion 表的记录
func DeleteAPIAssetVersion(orgID uint64, assetID string, versionID uint64, cascade bool) error {
	tx := Tx()
	defer tx.RollbackUnlessCommitted()

	if cascade {
		if err := tx.Where(map[string]interface{}{
			"org_id":     orgID,
			"asset_id":   assetID,
			"version_id": versionID,
		}).Delete(APIAssetVersionSpecsModel{}).
			Error; err != nil {
			return err
		}
	}

	if err := tx.Where(map[string]interface{}{
		"org_id":   orgID,
		"asset_id": assetID,
		"id":       versionID,
	}).Delete(new(apistructs.APIAssetVersionsModel)).
		Error; err != nil {
		return err
	}

	if d := tx.Delete(new(apistructs.APIOAS3IndexModel), map[string]interface{}{"version_id": versionID}); d.Error != nil {
		return d.Error
	}

	if d := tx.Delete(new(apistructs.APIOAS3FragmentModel), map[string]interface{}{"version_id": versionID}); d.Error != nil {
		return d.Error
	}

	tx.Commit()

	return nil
}

func GenSemVer(orgID uint64, assetID, swaggerVersion string, major, minor, patch *uint64) error {
	if major == nil || minor == nil || patch == nil {
		return errors.New("semVer is nil")
	}

	// 查找指定 swagger_version 的一个 version model
	var (
		swaggerVersionModel apistructs.APIAssetVersionsModel
		swaggerVerExists    bool
	)
	switch err := DB.First(&swaggerVersionModel, map[string]interface{}{
		"org_id":          orgID,
		"asset_id":        assetID,
		"swagger_version": swaggerVersion,
	}).Error; {
	case err == nil:
		swaggerVerExists = true
	case gorm.IsRecordNotFoundError(err):
	default:
		return err
	}

	switch {
	// 如果请求参数中 major 缺省, 则补全 major 并检查版本冲突; 如果不冲突, 则 minor, patch 按用户填写, 如果版本冲突, 令 minor 为最新版本 minor, patch 为最新
	case *major == 0:
		if swaggerVerExists {
			*major = swaggerVersionModel.Major
		} else {
			latestVersion, err := QueryAPILatestVersion(orgID, assetID)
			if err == nil {
				*major = latestVersion.Major + 1
			} else if gorm.IsRecordNotFoundError(err) {
				*major = 1
			} else {
				return err
			}
		}
		versions, ok, err := CheckVersionConflict(orgID, assetID, *major, *minor, *patch)
		if err != nil {
			return err
		}
		if !ok && len(versions) > 0 {
			*minor = versions[0][1]
			*patch = versions[0][2] + 1
		}

	// 如果请求参数中的 major 正确, 则检查版本冲突; 如果不冲突, 按用户填写继续流程, 如果冲突则报错
	case *major == swaggerVersionModel.Major:
		_, ok, err := CheckVersionConflict(orgID, assetID, *major, *minor, *patch)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("request version(major.minor.path) already exists")
		}

	// 如果存在这样的 swaggerVersion,
	case swaggerVerExists:
		return errors.Errorf("输入的版本号 %v.%v.%v 与 API 描述文档中 info/version 对应的 %v.*.* 不匹配, 请修改主版本号或更换文档",
			*major, *minor, *patch, swaggerVersionModel.Major)

	default:
		// 查找特定 major 的 version model
		var exVersion apistructs.APIAssetVersionsModel
		switch err := DB.First(&exVersion, map[string]interface{}{
			"org_id":   orgID,
			"asset_id": assetID,
			"major":    major,
		}).Error; {
		case err == nil:
			// 不存在给定的 swaggerVersion, 但存在给定的 majorVersion, 说明 swaggerVersion 和 major 冲突
			return errors.Errorf("输入的语义化版本号 %v.*.* 已存在于名称为 %s 的版本下, 请修改主版本号或更换文档",
				*major, exVersion.SwaggerVersion)

		case gorm.IsRecordNotFoundError(err):
			// 给定的 swaggerVersion 和 majorVersion 都不存在, 则正常创建

		default:
			return err
		}
	}

	return nil
}
