package services

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"time"

	"github.com/ZEQUANR/zhulong/ent"
	"github.com/ZEQUANR/zhulong/ent/thesis"
	"github.com/ZEQUANR/zhulong/ent/user"
	"github.com/ZEQUANR/zhulong/model"
	"github.com/ZEQUANR/zhulong/model/api"
)

func CreateThesis(id int, data api.CreateThesis) (*ent.Thesis, error) {
	ctx := context.Background()

	authors, err := json.Marshal(data.Authors)
	if err != nil {
		return nil, fmt.Errorf("func: CreateThesis | index: 0 | err: %w", err)
	}

	teachers, err := json.Marshal(data.Teachers)
	if err != nil {
		return nil, fmt.Errorf("func: CreateThesis | index: 1 | err: %w", err)
	}

	user, err := client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: CreateThesis | index: 2 | err: %w", err)
	}

	arr, err := user.
		QueryThesis().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: CreateThesis | index: 3 | err: %w", err)
	}

	if len(arr) != 0 {
		for _, elem := range arr {
			if elem.FileState == model.FileState.ToBeUploaded {
				result, err := elem.
					Update().
					SetFileState(model.FileState.ToBeUploaded).
					SetChineseTitle(data.ChineseTitle).
					SetEnglishTitle(data.EnglishTitle).
					SetAuthors(string(authors)).
					SetTeachers(string(teachers)).
					SetFirstAdvance(data.FirstAdvance).
					SetSecondAdvance(data.SecondAdvance).
					SetThirdAdvance(data.ThirdAdvance).
					SetDrawback(data.Drawback).
					SetUploaders(user).
					Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("func: CreateThesis | index: 4 | err: %w", err)
				}

				return result, nil
			}
		}
	}

	result, err := client.Thesis.
		Create().
		SetFileState(model.FileState.ToBeUploaded).
		SetChineseTitle(data.ChineseTitle).
		SetEnglishTitle(data.EnglishTitle).
		SetAuthors(string(authors)).
		SetTeachers(string(teachers)).
		SetFirstAdvance(data.FirstAdvance).
		SetSecondAdvance(data.SecondAdvance).
		SetThirdAdvance(data.ThirdAdvance).
		SetDrawback(data.Drawback).
		SetUploaders(user).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: CreateThesis | index: 5 | err: %w", err)
	}

	return result, nil
}

func UploadThesis(userId int, thesisID int, file model.File) (*ent.Thesis, error) {
	ctx := context.Background()

	t, err := client.Thesis.
		Query().
		Where(thesis.ID(thesisID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UploadThesis | index: 0 | err: %w", err)
	}

	user, err := t.
		QueryUploaders().
		Only(ctx)
	if err != nil || user.ID != userId {
		return nil, fmt.Errorf("func: UploadThesis | index: 1 | err: %w", err)
	}

	result, err := t.
		Update().
		SetFileName(file.Name).
		SetFileURL(path.Join(file.Path, file.Name)).
		SetFileState(model.FileState.ToBeReviewed).
		SetUploadTime(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UploadThesis | index: 2 | err: %w", err)
	}

	_, err = RecordThesisUpload(user)
	if err != nil {
		return nil, fmt.Errorf("func: UploadThesis | index: 3 | err: %w", err)
	}

	return result, nil
}

func QueryReviewRecord(userId int) ([]*ent.OperationLog, error) {
	ctx := context.Background()

	user, err := client.User.
		Query().
		Where(user.ID(userId)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryReviewRecord | index: 0 | err: %w", err)
	}

	result, err := user.
		QueryOperatingRecord().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryReviewRecord | index: 1 | err: %w", err)
	}

	return result, nil
}

func QueryToBeReviewedThesisList() ([]api.ToBeReviewedThesisList, error) {
	ctx := context.Background()

	arr, err := client.Thesis.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryToBeReviewedThesisList | index: 0 | err: %w", err)
	}

	var array []api.ToBeReviewedThesisList
	for _, elem := range arr {
		u, err := elem.QueryUploaders().Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("func: QueryToBeReviewedThesisList | index: 1 | err: %w", err)
		}

		var (
			name    string
			number  string
			college string
		)
		if u.Role == model.Admin {
			info, err := u.QueryAdministrators().Only(ctx)
			if err != nil {
				return nil, fmt.Errorf("func: QueryToBeReviewedThesisList | index: 2 | err: %w", err)
			}

			name = info.Name
			number = info.Number
			college = info.College
		}
		if u.Role == model.Student {
			info, err := u.QueryStudents().Only(ctx)
			if err != nil {
				return nil, fmt.Errorf("func: QueryToBeReviewedThesisList | index: 3 | err: %w", err)
			}

			name = info.Name
			number = info.Number
			college = info.College
		}
		if u.Role == model.Teacher {
			info, err := u.QueryTeachers().Only(ctx)
			if err != nil {
				return nil, fmt.Errorf("func: QueryToBeReviewedThesisList | index: 4 | err: %w", err)
			}

			name = info.Name
			number = info.Number
			college = info.College
		}

		array = append(array, api.ToBeReviewedThesisList{
			Id:           elem.ID,
			Name:         name,
			Number:       number,
			College:      college,
			FileState:    elem.FileState,
			ChineseTitle: elem.ChineseTitle,
			EnglishTitle: elem.EnglishTitle,
			UploadTime:   elem.UploadTime,
		})
	}

	return array, nil
}

func UpdateAllocationThesis(data api.AllocationThesis) error {
	ctx := context.Background()

	user, err := client.User.
		Query().
		Where(user.ID(data.TeacherID)).
		Only(ctx)
	if err != nil || user.Role == model.Student {
		return fmt.Errorf("func: UpdateAllocationThesis | index: 0 | err: %w", err)
	}

	for _, value := range data.ThesisID {
		result, err := client.Thesis.
			Query().
			Where(thesis.ID(value)).
			Only(ctx)
		if err != nil || result.FileState != model.FileState.ToBeReviewed {
			return fmt.Errorf("func: UpdateAllocationThesis | index: 1 | err: %w", err)
		}
	}

	for _, value := range data.ThesisID {
		_, err := client.Thesis.
			Update().
			Where(thesis.ID(value)).
			SetFileState(model.FileState.UnderReview).
			SetExamine(user).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("func: UpdateAllocationThesis | index: 2 | err: %w", err)
		}
	}

	return nil
}

func QueryUnderReviewThesisList(user *ent.User) ([]api.UnderReviewThesisList, error) {
	ctx := context.Background()

	var arr []api.UnderReviewThesisList
	t, err := user.
		QueryExamineThesis().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryUnderReviewThesisList | index: 0 | err: %w", err)
	}

	for _, elem := range t {
		info, err := elem.QueryExamine().Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("func: QueryUnderReviewThesisList | index: 1 | err: %w", err)
		}

		var (
			name    string
			number  string
			college string
			phone   string
		)

		if info.Role == model.Admin {
			u, err := info.QueryAdministrators().Only(ctx)
			if err != nil {
				return nil, fmt.Errorf("func: QueryUnderReviewThesisList | index: 2 | err: %w", err)
			}

			name = u.Name
			number = u.Number
			college = u.College
			phone = u.Phone
		}

		if info.Role == model.Teacher {
			u, err := info.QueryTeachers().Only(ctx)
			if err != nil {
				return nil, fmt.Errorf("func: QueryUnderReviewThesisList | index: 3 | err: %w", err)
			}

			name = u.Name
			number = u.Number
			college = u.College
			phone = u.Phone
		}

		var authors []string
		err = json.Unmarshal([]byte(elem.Authors), &authors)
		if err != nil {
			return nil, fmt.Errorf("func: QueryUnderReviewThesisList | index: 4 | err: %w", err)
		}

		var teachers []string
		err = json.Unmarshal([]byte(elem.Teachers), &teachers)
		if err != nil {
			return nil, fmt.Errorf("func: QueryUnderReviewThesisList | index: 5 | err: %w", err)
		}

		arr = append(arr, api.UnderReviewThesisList{
			Id:            elem.ID,
			Name:          name,
			Number:        number,
			College:       college,
			Phone:         phone,
			FileName:      elem.FileName,
			FileState:     elem.FileState,
			ChineseTitle:  elem.ChineseTitle,
			EnglishTitle:  elem.EnglishTitle,
			Authors:       authors,
			Teachers:      teachers,
			FirstAdvance:  elem.FirstAdvance,
			SecondAdvance: elem.SecondAdvance,
			ThirdAdvance:  elem.ThirdAdvance,
			Drawback:      elem.Drawback,
			UploadTime:    elem.UploadTime,
			CreateTime:    elem.CreateTime,
		})
	}

	return arr, nil
}

func QueryThesisDownloadPath(id int, data api.DownloadThesis) (*ent.Thesis, error) {
	ctx := context.Background()

	user, err := client.User.
		Query().
		Select().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryThesisDownloadPath | index: 1 | err: %w", err)
	}

	t, err := client.Thesis.
		Query().
		Select().
		Where(thesis.ID(data.ThesisId)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryThesisDownloadPath | index: 2 | err: %w", err)
	}

	qid, err := t.
		QueryExamine().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryThesisDownloadPath | index: 3 | err: %w", err)
	}

	eid, err := t.
		QueryUploaders().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: QueryThesisDownloadPath | index: 4 | err: %w", err)
	}

	if user.Role != model.Admin && (qid.ID != id && eid.ID != id) {
		return nil, fmt.Errorf("func: QueryThesisDownloadPath | index: 5 | err: %w", err)
	}

	return t, nil
}

func UpdateThesisRandomAllocation(data api.RandomAllocationThesis) (*ent.Thesis, error) {
	ctx := context.Background()

	for _, v := range data.ThesisID {
		t, err := client.Thesis.
			Query().
			Select().
			Where(thesis.ID(v)).
			Only(ctx)
		if err != nil || t.FileState != model.FileState.ToBeReviewed {
			return nil, fmt.Errorf("func: UpdateThesisRandomAllocation | index: 0 | err: %w", err)
		}
	}

	return nil, nil
}
