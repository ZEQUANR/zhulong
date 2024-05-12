package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ZEQUANR/zhulong/ent"
	"github.com/ZEQUANR/zhulong/ent/thesis"
	"github.com/ZEQUANR/zhulong/ent/user"
	"github.com/ZEQUANR/zhulong/model"
	"github.com/ZEQUANR/zhulong/model/api"
)

func CreateThesis(id int, createThesis api.CreateThesis) (*ent.Thesis, error) {
	ctx := context.Background()

	user, err := client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: CreateThesis %w", err)
	}

	thesis, err := client.Thesis.
		Create().
		SetThesisTitle(createThesis.ThesisTitle).
		SetUploaders(user).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: CreateThesis %w", err)
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
		return nil, fmt.Errorf("func: UploadThesis %w", err)
	}

	user, err := paper.
		QueryUploaders().
		Only(ctx)
	if err != nil || user.ID != userId {
		return nil, fmt.Errorf("func: UploadThesis %w", err)
	}

	result, err := paper.
		Update().
		SetFileName(file.Name).
		SetFileURL(file.Path).
		SetFileState(0).
		SetUploadTime(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UploadThesis %w", err)
	}

	return result, nil
}
