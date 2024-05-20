package services

import (
	"context"
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

	user, err := client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: CreateThesis | index: 0 | err: %w", err)
	}

	thesis, err := client.Thesis.
		Create().
		SetChineseTitle(data.ChineseTitle).
		SetEnglishTitle(data.EnglishTitle).
		SetAuthors(data.Authors).
		SetTeachers(data.Teachers).
		SetFirstAdvance(data.FirstAdvance).
		SetSecondAdvance(data.SecondAdvance).
		SetThirdAdvance(data.ThirdAdvance).
		SetDrawback(data.Drawback).
		SetUploaders(user).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: CreateThesis | index: 1 | err: %w", err)
	}

	return thesis, nil
}

func UploadThesis(userId int, thesisID int, file model.File) (*ent.Thesis, error) {
	ctx := context.Background()

	paper, err := client.Thesis.
		Query().
		Where(thesis.ID(thesisID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UploadThesis | index: 0 | err: %w", err)
	}

	user, err := paper.
		QueryUploaders().
		Only(ctx)
	if err != nil || user.ID != userId {
		return nil, fmt.Errorf("func: UploadThesis | index: 1 | err: %w", err)
	}

	result, err := paper.
		Update().
		SetFileName(file.Name).
		SetFileURL(path.Join(file.Path, file.Name)).
		SetFileState(model.FileState.ToBeReviewed).
		SetUploadTime(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UploadThesis | index: 2 | err: %w", err)
	}

	return result, nil
}

func QueryToBeReviewedThesisList() ([]api.ToBeReviewedThesisList, error) {
	ctx := context.Background()

	var result []api.ToBeReviewedThesisList
	err := client.Thesis.
		Query().
		Select(
			thesis.FieldID,
			thesis.FieldFileName,
			thesis.FieldFileState,
			thesis.FieldUploadTime,
			thesis.FieldChineseTitle,
			thesis.FieldEnglishTitle,
			thesis.FieldAuthors,
			thesis.FieldTeachers,
			thesis.FieldFirstAdvance,
			thesis.FieldSecondAdvance,
			thesis.FieldThirdAdvance,
			thesis.FieldDrawback,
		).
		Scan(ctx, &result)
	if err != nil {
		return nil, fmt.Errorf("func: QueryToBeReviewedThesisList | index: 0 | err: %w", err)
	}

	return result, nil
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

func QueryUnderReviewThesisList(id int) ([]api.ToBeReviewedThesisList, error) {
	ctx := context.Background()

	user, err := client.User.
		Query().
		Select().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil || user.Role == model.Student {
		return nil, fmt.Errorf("func: QueryUnderReviewThesisList | index: 0 | err: %w", err)
	}

	var result []api.ToBeReviewedThesisList
	err = user.
		QueryExamineThesis().
		Select(
			thesis.FieldID,
			thesis.FieldFileName,
			thesis.FieldFileState,
			thesis.FieldUploadTime,
			thesis.FieldChineseTitle,
			thesis.FieldEnglishTitle,
			thesis.FieldAuthors,
			thesis.FieldTeachers,
			thesis.FieldFirstAdvance,
			thesis.FieldSecondAdvance,
			thesis.FieldThirdAdvance,
			thesis.FieldDrawback,
		).
		Scan(ctx, &result)
	if err != nil {
		return nil, fmt.Errorf("func: QueryUnderReviewThesisList | index: 1 | err: %w", err)
	}

	return result, nil
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
