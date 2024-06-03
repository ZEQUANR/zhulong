package services

import (
	"context"
	"fmt"
	"path"

	"github.com/ZEQUANR/zhulong/ent"
	"github.com/ZEQUANR/zhulong/ent/thesis"
	"github.com/ZEQUANR/zhulong/ent/user"
	"github.com/ZEQUANR/zhulong/model"
)

func UploadReview(userId int, thesisID int, file model.File) (*ent.Reviews, error) {
	ctx := context.Background()

	u, err := client.User.
		Query().
		Where(user.ID(userId)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UploadReview | index: 0 | err: %w", err)
	}

	t, err := client.Thesis.
		Query().
		Where(thesis.ID(thesisID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UploadReview | index: 1 | err: %w", err)
	}

	r, err := client.Reviews.
		Create().
		SetFileName(file.Name).
		SetFileURL(path.Join(file.Path, file.Name)).
		SetUploaders(u).
		SetThesisResult(t).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UploadReview | index: 2 | err: %w", err)
	}

	_, err = t.
		Update().
		SetFileState(model.FileState.ReviewCompleted).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: UploadReview | index: 3 | err: %w", err)
	}

	return r, nil
}
