// Code generated by gowebx, DO AVOID EDIT.
package media

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/ixugo/goweb/pkg/hook"
	"github.com/ixugo/goweb/pkg/orm"
	"github.com/ixugo/goweb/pkg/web"
	"github.com/jinzhu/copier"

	"github.com/gowvp/gb28181/internal/core/bz"
)

// StreamPushStorer Instantiation interface
type StreamPushStorer interface {
	Find(context.Context, *[]*StreamPush, orm.Pager, ...orm.QueryOption) (int64, error)
	Get(context.Context, *StreamPush, ...orm.QueryOption) error
	Add(context.Context, *StreamPush) error
	Edit(context.Context, *StreamPush, func(*StreamPush), ...orm.QueryOption) error
	Del(context.Context, *StreamPush, ...orm.QueryOption) error
}

// FindStreamPush Paginated search
func (c Core) FindStreamPush(ctx context.Context, in *FindStreamPushInput) ([]*StreamPush, int64, error) {
	items := make([]*StreamPush, 0)
	args := make([]orm.QueryOption, 0, 2)
	args = append(args, orm.OrderBy("created_at DESC"))
	if in.Status != "" {
		args = append(args, orm.Where("status=?", in.Status))
	}
	if in.Key != "" {
		args = append(args, orm.Where("id=? OR app LIKE ? OR stream LIKE ?", in.Key, "%"+in.Key+"%", "%"+in.Key+"%"))
	}

	total, err := c.store.StreamPush().Find(ctx, &items, in, args...)
	if err != nil {
		return nil, 0, web.ErrDB.Withf(`Find err[%s]`, err.Error())
	}

	return items, total, nil
}

// GetStreamPush Query a single object
func (c Core) GetStreamPush(ctx context.Context, id string) (*StreamPush, error) {
	var out StreamPush
	if err := c.store.StreamPush().Get(ctx, &out, orm.Where("id=?", id)); err != nil {
		if orm.IsErrRecordNotFound(err) {
			return nil, web.ErrNotFound.Withf(`Get err[%s]`, err.Error())
		}
		return nil, web.ErrDB.Withf(`Get err[%s]`, err.Error())
	}
	return &out, nil
}

func (c Core) GetStreamPushByAppStream(ctx context.Context, app, stream string) (*StreamPush, error) {
	var out StreamPush
	if err := c.store.StreamPush().Get(ctx, &out, orm.Where("app=? AND stream=?", app, stream)); err != nil {
		if orm.IsErrRecordNotFound(err) {
			return nil, web.ErrNotFound.Withf(`Get err[%s]`, err.Error())
		}
		return nil, web.ErrDB.Withf(`Get err[%s]`, err.Error())
	}
	return &out, nil
}

// AddStreamPush Insert into database
func (c Core) AddStreamPush(ctx context.Context, in *AddStreamPushInput) (*StreamPush, error) {
	var out StreamPush
	if err := copier.Copy(&out, in); err != nil {
		slog.Error("Copy", "err", err)
	}

	if in.App == "rtp" {
		return nil, web.ErrBadRequest.With("请更换 app 参数")
	}
	out.ID = c.uniqueID.UniqueID(bz.IDPrefixRTMP)
	if err := c.store.StreamPush().Add(ctx, &out); err != nil {
		if orm.IsDuplicatedKey(err) {
			return nil, web.ErrDB.Msg("stream 重复，请勿重复添加")
		}
		return nil, web.ErrDB.Withf(`Add err[%s]`, err.Error())
	}
	return &out, nil
}

// EditStreamPush Update object information
func (c Core) EditStreamPush(ctx context.Context, in *EditStreamPushInput, id string) (*StreamPush, error) {
	var out StreamPush
	if err := c.store.StreamPush().Edit(ctx, &out, func(b *StreamPush) {
		if err := copier.Copy(b, in); err != nil {
			slog.Error("Copy", "err", err)
		}
	}, orm.Where("id=?", id)); err != nil {
		return nil, web.ErrDB.Withf(`Edit err[%s]`, err.Error())
	}
	return &out, nil
}

// DelStreamPush Delete object
func (c *Core) DelStreamPush(ctx context.Context, id string) (*StreamPush, error) {
	// 检查数据库
	// 如果是推流中，需要先让 sms 停止推流
	// TODO: 待实现国标相关，删除国标相关数据
	var out StreamPush
	if err := c.store.StreamPush().Del(ctx, &out, orm.Where("id=?", id)); err != nil {
		return nil, web.ErrDB.Withf(`Del err[%s]`, err.Error())
	}
	return &out, nil
}

type PublishInput struct {
	App           string
	Stream        string
	MediaServerID string
	Sign          string
	Secret        string
	Session       string
}

// Publish 由于 hook 的函数，无需 web.error 封装
func (c *Core) Publish(ctx context.Context, in PublishInput) error {
	var streamPushInput = AddStreamPushInput{
		App:            in.App,
		Stream:         in.Stream,
		IsAuthDisabled: false,
	}
	c.AddStreamPush(ctx, &streamPushInput)
	result, err := c.GetStreamPushByAppStream(ctx, in.App, in.Stream)
	if err != nil {
		return err
	}
	if !result.IsAuthDisabled {
		if s := hook.MD5(in.Session + in.Secret); s != in.Sign {
			slog.Info("推流鉴权失败", "got", in.Sign, "expect", s)
			return fmt.Errorf("Unauthorized, rtmp secret error, got[%s]", in.Sign)
		}
	}

	var s StreamPush
	return c.store.StreamPush().Edit(ctx, &s, func(b *StreamPush) {
		b.MediaServerID = in.MediaServerID
		b.Status = StatusPushing
		now := orm.Now()
		b.PushedAt = &now
		b.Session = in.Session
	}, orm.Where("id=?", result.ID))

}

func (c *Core) UnPublish(ctx context.Context, app, stream string) error {
	var s StreamPush
	return c.store.StreamPush().Edit(ctx, &s, func(b *StreamPush) {
		b.Status = StatusStopped
		now := orm.Now()
		b.StoppedAt = &now
		b.Session = ""
	}, orm.Where("app = ? AND stream=?", app, stream))
}

type OnPlayInput struct {
	App     string
	Stream  string
	Session string
}

func (c *Core) OnPlay(ctx context.Context, in OnPlayInput) error {
	result, err := c.GetStreamPushByAppStream(ctx, in.App, in.Stream)
	if err != nil {
		return err
	}
	if result.IsAuthDisabled {
		return nil
	}
	if in.Session != result.Session {
		slog.Info("拉流鉴权失败", "got", in.Session, "expect", result.Session)
		return fmt.Errorf("Unauthorized, session error, got[%s]", in.Session)
	}
	return nil
}
