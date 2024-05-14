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
