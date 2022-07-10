package moxerr_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/kyledinh/btk-go/pkg/moxerr"
	"github.com/stretchr/testify/assert"
)

func Test_(t *testing.T) {
	t.Parallel()

	type Want struct {
		ErrType        error
		MoxMessage     string
		WrapperMessage string
	}

	tests := []struct {
		name     string
		newError error
		moxError error
		want     Want
	}{
		{
			name:     "Error with ErrCLIAction",
			newError: errors.New("failed to connect to db"),
			moxError: moxerr.ErrCLIAction,
			want: Want{
				ErrType:        moxerr.ErrCLIAction,
				MoxMessage:     "CLI_ACTION_FAILED",
				WrapperMessage: "wrapped message: failed to connect to db",
			},
		},
		{
			name:     "Error with ErrResourceNotFound",
			newError: errors.New("File not found"),
			moxError: moxerr.ErrResourceNotFound,
			want: Want{
				ErrType:        moxerr.ErrResourceNotFound,
				MoxMessage:     "RESOURCE_NOT_FOUND",
				WrapperMessage: "wrapped message: File not found",
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d: %s", i, tt.name), func(t *testing.T) {
			result := moxerr.NewWrappedError(tt.newError.Error(), &tt.moxError)
			assert.Equal(t, tt.want.MoxMessage, error(*result.MoxErr).Error())
			assert.Equal(t, tt.want.WrapperMessage, result.Error())
		})
	}
}
