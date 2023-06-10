// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DeleteProjectParams is parameters of deleteProject operation.
type DeleteProjectParams struct {
	XAPIKey   string
	ProjectID int64
}

func unpackDeleteProjectParams(packed middleware.Parameters) (params DeleteProjectParams) {
	{
		key := middleware.ParameterKey{
			Name: "X-Api-Key",
			In:   "header",
		}
		params.XAPIKey = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "projectID",
			In:   "path",
		}
		params.ProjectID = packed[key].(int64)
	}
	return params
}

func decodeDeleteProjectParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteProjectParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: X-Api-Key.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Api-Key",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XAPIKey = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Api-Key",
			In:   "header",
			Err:  err,
		}
	}
	// Decode path: projectID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "projectID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ProjectID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.ProjectID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "projectID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteTaskParams is parameters of deleteTask operation.
type DeleteTaskParams struct {
	XAPIKey   string
	ProjectID int64
	TaskID    int64
}

func unpackDeleteTaskParams(packed middleware.Parameters) (params DeleteTaskParams) {
	{
		key := middleware.ParameterKey{
			Name: "X-Api-Key",
			In:   "header",
		}
		params.XAPIKey = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "projectID",
			In:   "path",
		}
		params.ProjectID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "taskID",
			In:   "path",
		}
		params.TaskID = packed[key].(int64)
	}
	return params
}

func decodeDeleteTaskParams(args [2]string, argsEscaped bool, r *http.Request) (params DeleteTaskParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: X-Api-Key.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Api-Key",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XAPIKey = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Api-Key",
			In:   "header",
			Err:  err,
		}
	}
	// Decode path: projectID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "projectID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ProjectID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.ProjectID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "projectID",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: taskID.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "taskID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.TaskID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.TaskID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "taskID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// GetProjectsParams is parameters of getProjects operation.
type GetProjectsParams struct {
	// リソースの最大取得数を指定する。.
	Limit OptInt
	// リソースの取得開始位置を指定する。.
	Offset OptInt
	// 並び順を指定する。`-`をつければ降順になり、つけなければ昇順となる。.
	Sort    OptGetProjectsSort
	XAPIKey string
}

func unpackGetProjectsParams(packed middleware.Parameters) (params GetProjectsParams) {
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "offset",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Offset = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "sort",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Sort = v.(OptGetProjectsSort)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "X-Api-Key",
			In:   "header",
		}
		params.XAPIKey = packed[key].(string)
	}
	return params
}

func decodeGetProjectsParams(args [0]string, argsEscaped bool, r *http.Request) (params GetProjectsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	h := uri.NewHeaderDecoder(r.Header)
	// Set default value for query: limit.
	{
		val := int(10)
		params.Limit.SetTo(val)
	}
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Limit.Set {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        true,
							Max:           30,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(params.Limit.Value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: offset.
	{
		val := int(0)
		params.Offset.SetTo(val)
	}
	// Decode query: offset.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "offset",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOffsetVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotOffsetVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Offset.SetTo(paramsDotOffsetVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Offset.Set {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           0,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(params.Offset.Value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "offset",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: sort.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "sort",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSortVal GetProjectsSort
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSortVal = GetProjectsSort(c)
					return nil
				}(); err != nil {
					return err
				}
				params.Sort.SetTo(paramsDotSortVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Sort.Set {
					if err := func() error {
						if err := params.Sort.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "sort",
			In:   "query",
			Err:  err,
		}
	}
	// Decode header: X-Api-Key.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Api-Key",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XAPIKey = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Api-Key",
			In:   "header",
			Err:  err,
		}
	}
	return params, nil
}

// GetTasksParams is parameters of getTasks operation.
type GetTasksParams struct {
	// リソースの最大取得数を指定する。.
	Limit OptInt
	// リソースの取得開始位置を指定する。.
	Offset OptInt
	// 並び順を指定する。`-`をつければ降順になり、つけなければ昇順となる。.
	Sort      OptGetTasksSort
	XAPIKey   string
	ProjectID int64
}

func unpackGetTasksParams(packed middleware.Parameters) (params GetTasksParams) {
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "offset",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Offset = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "sort",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Sort = v.(OptGetTasksSort)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "X-Api-Key",
			In:   "header",
		}
		params.XAPIKey = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "projectID",
			In:   "path",
		}
		params.ProjectID = packed[key].(int64)
	}
	return params
}

func decodeGetTasksParams(args [1]string, argsEscaped bool, r *http.Request) (params GetTasksParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	h := uri.NewHeaderDecoder(r.Header)
	// Set default value for query: limit.
	{
		val := int(10)
		params.Limit.SetTo(val)
	}
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Limit.Set {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           1,
							MaxSet:        true,
							Max:           30,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(params.Limit.Value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	// Set default value for query: offset.
	{
		val := int(0)
		params.Offset.SetTo(val)
	}
	// Decode query: offset.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "offset",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotOffsetVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotOffsetVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Offset.SetTo(paramsDotOffsetVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Offset.Set {
					if err := func() error {
						if err := (validate.Int{
							MinSet:        true,
							Min:           0,
							MaxSet:        false,
							Max:           0,
							MinExclusive:  false,
							MaxExclusive:  false,
							MultipleOfSet: false,
							MultipleOf:    0,
						}).Validate(int64(params.Offset.Value)); err != nil {
							return errors.Wrap(err, "int")
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "offset",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: sort.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "sort",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSortVal GetTasksSort
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSortVal = GetTasksSort(c)
					return nil
				}(); err != nil {
					return err
				}
				params.Sort.SetTo(paramsDotSortVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if params.Sort.Set {
					if err := func() error {
						if err := params.Sort.Value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "sort",
			In:   "query",
			Err:  err,
		}
	}
	// Decode header: X-Api-Key.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Api-Key",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XAPIKey = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Api-Key",
			In:   "header",
			Err:  err,
		}
	}
	// Decode path: projectID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "projectID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ProjectID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.ProjectID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "projectID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PatchProjectParams is parameters of patchProject operation.
type PatchProjectParams struct {
	XAPIKey   string
	ProjectID int64
}

func unpackPatchProjectParams(packed middleware.Parameters) (params PatchProjectParams) {
	{
		key := middleware.ParameterKey{
			Name: "X-Api-Key",
			In:   "header",
		}
		params.XAPIKey = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "projectID",
			In:   "path",
		}
		params.ProjectID = packed[key].(int64)
	}
	return params
}

func decodePatchProjectParams(args [1]string, argsEscaped bool, r *http.Request) (params PatchProjectParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: X-Api-Key.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Api-Key",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XAPIKey = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Api-Key",
			In:   "header",
			Err:  err,
		}
	}
	// Decode path: projectID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "projectID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ProjectID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.ProjectID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "projectID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PatchTaskParams is parameters of patchTask operation.
type PatchTaskParams struct {
	XAPIKey   string
	ProjectID int64
	TaskID    int64
}

func unpackPatchTaskParams(packed middleware.Parameters) (params PatchTaskParams) {
	{
		key := middleware.ParameterKey{
			Name: "X-Api-Key",
			In:   "header",
		}
		params.XAPIKey = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "projectID",
			In:   "path",
		}
		params.ProjectID = packed[key].(int64)
	}
	{
		key := middleware.ParameterKey{
			Name: "taskID",
			In:   "path",
		}
		params.TaskID = packed[key].(int64)
	}
	return params
}

func decodePatchTaskParams(args [2]string, argsEscaped bool, r *http.Request) (params PatchTaskParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: X-Api-Key.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Api-Key",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XAPIKey = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Api-Key",
			In:   "header",
			Err:  err,
		}
	}
	// Decode path: projectID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "projectID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ProjectID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.ProjectID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "projectID",
			In:   "path",
			Err:  err,
		}
	}
	// Decode path: taskID.
	if err := func() error {
		param := args[1]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[1])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "taskID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.TaskID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.TaskID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "taskID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// PostProjectsParams is parameters of PostProjects operation.
type PostProjectsParams struct {
	XAPIKey string
}

func unpackPostProjectsParams(packed middleware.Parameters) (params PostProjectsParams) {
	{
		key := middleware.ParameterKey{
			Name: "X-Api-Key",
			In:   "header",
		}
		params.XAPIKey = packed[key].(string)
	}
	return params
}

func decodePostProjectsParams(args [0]string, argsEscaped bool, r *http.Request) (params PostProjectsParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: X-Api-Key.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Api-Key",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XAPIKey = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Api-Key",
			In:   "header",
			Err:  err,
		}
	}
	return params, nil
}

// PostTasksParams is parameters of PostTasks operation.
type PostTasksParams struct {
	XAPIKey   string
	ProjectID int64
}

func unpackPostTasksParams(packed middleware.Parameters) (params PostTasksParams) {
	{
		key := middleware.ParameterKey{
			Name: "X-Api-Key",
			In:   "header",
		}
		params.XAPIKey = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "projectID",
			In:   "path",
		}
		params.ProjectID = packed[key].(int64)
	}
	return params
}

func decodePostTasksParams(args [1]string, argsEscaped bool, r *http.Request) (params PostTasksParams, _ error) {
	h := uri.NewHeaderDecoder(r.Header)
	// Decode header: X-Api-Key.
	if err := func() error {
		cfg := uri.HeaderParameterDecodingConfig{
			Name:    "X-Api-Key",
			Explode: false,
		}
		if err := h.HasParam(cfg); err == nil {
			if err := h.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.XAPIKey = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "X-Api-Key",
			In:   "header",
			Err:  err,
		}
	}
	// Decode path: projectID.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "projectID",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToInt64(val)
				if err != nil {
					return err
				}

				params.ProjectID = c
				return nil
			}(); err != nil {
				return err
			}
			if err := func() error {
				if err := (validate.Int{
					MinSet:        true,
					Min:           1,
					MaxSet:        false,
					Max:           0,
					MinExclusive:  false,
					MaxExclusive:  false,
					MultipleOfSet: false,
					MultipleOf:    0,
				}).Validate(int64(params.ProjectID)); err != nil {
					return errors.Wrap(err, "int")
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "projectID",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
