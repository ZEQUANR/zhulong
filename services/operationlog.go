package services

import (
	"context"
	"fmt"
	"time"

	"github.com/ZEQUANR/zhulong/ent"
	"github.com/ZEQUANR/zhulong/model"
)

func RecordThesisUpload(user *ent.User) (*ent.OperationLog, error) {
	ctx := context.Background()

	var name string

	if user.Role == model.Admin {
		info, err := user.
			QueryAdministrators().
			Only(ctx)

		if err != nil {
			return nil, fmt.Errorf("func: RecordThesisUpload | index: 0 | err: %w", err)
		}

		name = info.Name
	}

	if user.Role == model.Student {
		info, err := user.
			QueryStudents().
			Only(ctx)

		if err != nil {
			return nil, fmt.Errorf("func: RecordThesisUpload | index: 1 | err: %w", err)
		}

		name = info.Name
	}

	if user.Role == model.Teacher {
		info, err := user.
			QueryTeachers().
			Only(ctx)

		if err != nil {
			return nil, fmt.Errorf("func: RecordThesisUpload | index: 2 | err: %w", err)
		}

		name = info.Name
	}

	result, err := client.OperationLog.
		Create().
		SetName(name).
		SetContext(model.OperationLogContext.ThesisUpload).
		SetStatus(model.OperationLogStatus.UnderReview).
		SetTime(time.Now()).
		SetOperator(user).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: RecordThesisUpload | index: 3 | err: %w", err)
	}

	return result, nil
}

func RecordThesisReviewer(user *ent.User, t *ent.Thesis) (*ent.OperationLog, error) {
	ctx := context.Background()

	var name string

	if user.Role == model.Admin {
		info, err := user.
			QueryAdministrators().
			Only(ctx)

		if err != nil {
			return nil, fmt.Errorf("func: RecordThesisReviewer | index: 0 | err: %w", err)
		}

		name = info.Name
	}

	if user.Role == model.Student {
		info, err := user.
			QueryStudents().
			Only(ctx)

		if err != nil {
			return nil, fmt.Errorf("func: RecordThesisReviewer | index: 1 | err: %w", err)
		}

		name = info.Name
	}

	if user.Role == model.Teacher {
		info, err := user.
			QueryTeachers().
			Only(ctx)

		if err != nil {
			return nil, fmt.Errorf("func: RecordThesisReviewer | index: 2 | err: %w", err)
		}

		name = info.Name
	}

	userUp, err := t.
		QueryUploaders().
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: RecordThesisUpload | index: 3 | err: %w", err)
	}

	result, err := client.OperationLog.
		Create().
		SetName(name).
		SetContext(model.OperationLogContext.UploadTheReview).
		SetStatus(model.OperationLogStatus.ReviewCompleted).
		SetTime(time.Now()).
		SetOperator(userUp).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("func: RecordThesisUpload | index: 4 | err: %w", err)
	}

	return result, nil
}
