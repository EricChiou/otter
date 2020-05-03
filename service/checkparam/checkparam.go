package checkparam

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
)

// Valid check request parameters
func Valid(ctx *fasthttp.RequestCtx, src interface{}) error {
	if reflect.ValueOf(src).Kind() != reflect.Ptr {
		fmt.Println("need send ptr")
		return errors.New("need send ptr")
	}

	if len(ctx.PostBody()) > 0 {
		if err := json.Unmarshal(ctx.PostBody(), src); err != nil {
			fmt.Println("json parse error")
			return errors.New("json parse error")
		}
	}

	if reflect.ValueOf(src).Elem().Kind() != reflect.Struct {
		fmt.Println("need struct type")
		return errors.New("need struct type")
	}

	return checkStruct(ctx, src)
}

func checkStruct(ctx *fasthttp.RequestCtx, src interface{}) error {
	val := reflect.ValueOf(src).Elem()
	typ := reflect.TypeOf(src).Elem()

	for i := 0; i < val.NumField(); i++ {
		req, _ := strconv.ParseBool(typ.Field(i).Tag.Get("req"))
		paramType := typ.Field(i).Tag.Get("typ")
		json := typ.Field(i).Tag.Get("json")
		key := strings.Split(json, ",")[0]

		switch val.Field(i).Kind() {
		case reflect.String:
			if paramType == "para" {
				valStr := string(ctx.QueryArgs().Peek(key))
				val.Field(i).SetString(valStr)
			}
			if req {
				if err := checkParam(val.Field(i).Interface()); err != nil {
					return err
				}
			}

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if paramType == "para" {
				valStr := string(ctx.QueryArgs().Peek(key))
				r, _ := strconv.ParseInt(valStr, 10, 64)
				val.Field(i).SetInt(r)
			}
			if req {
				if err := checkParam(val.Field(i).Interface()); err != nil {
					return err
				}
			}

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if paramType == "para" {
				valStr := string(ctx.QueryArgs().Peek(key))
				r, _ := strconv.ParseUint(valStr, 10, 64)
				val.Field(i).SetUint(r)
			}
			if req {
				if err := checkParam(val.Field(i).Interface()); err != nil {
					return err
				}
			}

		case reflect.Float32, reflect.Float64:
			if paramType == "para" {
				valStr := string(ctx.QueryArgs().Peek(key))
				r, _ := strconv.ParseFloat(valStr, 64)
				val.Field(i).SetFloat(r)
			}
			if req {
				if err := checkParam(val.Field(i).Interface()); err != nil {
					return err
				}
			}

		case reflect.Complex64, reflect.Complex128:
			if paramType == "para" {
				valStr := string(ctx.QueryArgs().Peek(key))
				r, _ := strconv.ParseFloat(valStr, 64)
				val.Field(i).SetComplex(complex(r, 0))
			}
			if req {
				if err := checkParam(val.Field(i).Interface()); err != nil {
					return err
				}
			}

		case reflect.Bool:
			if paramType == "para" {
				valStr := string(ctx.QueryArgs().Peek(key))
				r, _ := strconv.ParseBool(valStr)
				val.Field(i).SetBool(r)
			}
			if req {
				if err := checkParam(val.Field(i).Interface()); err != nil {
					return err
				}
			}

		case reflect.Slice, reflect.Array:
			if paramType == "para" {
				valStr := string(ctx.QueryArgs().Peek(key))
				split := strings.Split(valStr, ",")
				slice := reflect.MakeSlice(val.Field(i).Type(), len(split), len(split))
				val.Field(i).Set(slice)

				if val.Field(i).Len() > 0 {
					switch val.Field(i).Index(0).Kind() {
					case reflect.String:
						for j := 0; j < val.Field(i).Len(); j++ {
							val.Field(i).Index(j).SetString(split[j])
						}

					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						for j := 0; j < len(split); j++ {
							r, _ := strconv.ParseInt(split[j], 10, 64)
							val.Field(i).Index(j).SetInt(r)
						}

					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						for j := 0; j < len(split); j++ {
							r, _ := strconv.ParseUint(split[j], 10, 64)
							val.Field(i).Index(j).SetUint(r)
						}

					case reflect.Float32, reflect.Float64:
						for j := 0; j < len(split); j++ {
							r, _ := strconv.ParseFloat(split[j], 64)
							val.Field(i).Index(j).SetFloat(r)
						}

					case reflect.Complex64, reflect.Complex128:
						for j := 0; j < len(split); j++ {
							r, _ := strconv.ParseFloat(split[j], 64)
							val.Field(i).Index(j).SetComplex(complex(r, 0))
						}

					case reflect.Bool:
						for j := 0; j < len(split); j++ {
							r, _ := strconv.ParseBool(split[j])
							val.Field(i).Index(j).SetBool(r)
						}

					default:
						return errors.New("param error")
					}
				}
			}

			if req {
				if err := checkParam(val.Field(i).Interface()); err != nil {
					return err
				}
			}

			if val.Field(i).Len() > 0 && val.Field(i).Index(0).Kind() == reflect.Struct {
				for j := 0; j < val.Field(i).Len(); j++ {
					if err := checkStruct(ctx, val.Field(i).Index(j).Addr().Interface()); err != nil {
						return err
					}
				}
			}

		case reflect.Struct:
			if req {
				if err := checkParam(val.Field(i).Interface()); err != nil {
					return err
				}
			}

			if err := checkStruct(ctx, val.Field(i).Addr().Interface()); err != nil {
				return err
			}

		default:
			fmt.Println("src type not support")
			return errors.New("src type not support")
		}
	}

	return nil
}

func checkParam(param interface{}) error {
	switch reflect.TypeOf(param).Kind() {
	case reflect.String:
		if len(param.(string)) == 0 {
			return errors.New("param error")
		}

	case
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		if param == 0 {
			return errors.New("param error")
		}

	case reflect.Bool:
		return nil

	case reflect.Slice, reflect.Array:
		if param == nil {
			return errors.New("param error")
		}

	case reflect.Struct:
		if param == nil {
			return errors.New("param error")
		}

	default:
		return errors.New("param error")
	}

	return nil
}
