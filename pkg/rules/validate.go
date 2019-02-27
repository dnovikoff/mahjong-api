package rules

import (
	"fmt"
	"reflect"
	"strings"

	"go.uber.org/multierr"

	base_proto "github.com/dnovikoff/mahjong-api/genproto/base"
	rules_proto "github.com/dnovikoff/mahjong-api/genproto/rules"
)

func validateEnum(value interface{}, possbile map[int32]string) error {
	return validateEnumImpl(value, possbile, false)
}

func validateEnumImpl(value interface{}, possbile map[int32]string, allowEmpty bool) error {
	typeName := reflect.TypeOf(value).Name()
	intVal := int32(reflect.ValueOf(value).Int())
	if intVal == 0 && !allowEmpty {
		return fmt.Errorf("value of %v should be set", typeName)
	}
	_, found := possbile[intVal]
	if !found {
		return fmt.Errorf("value of %v=%v is unknown", typeName, intVal)
	}
	return nil
}

func validateMoneyValue(caption string, val int64) error {
	if val%100 == 0 {
		return nil
	}
	return fmt.Errorf(
		"incorrect value of %v=%v. Should be dividable by 100",
		caption, val)
}

func validateMoneySlice(caption string, expectedLen int64, value reflect.Value) error {
	ln := value.Len()
	if ln != int(expectedLen) {
		return fmt.Errorf(
			"expected %v len to be %v. Got %v",
			caption, expectedLen, ln)
	}
	var err error
	for i := 0; i < ln; i++ {
		v := value.Index(i).Int()
		err = multierr.Append(err, validateMoneyValue(caption, v))
	}
	return err
}

func validateMoney(expectedLen int64, value interface{}) error {
	return validateMoneyReflect(expectedLen, reflect.ValueOf(value))
}

func validateMoneyReflect(expectedLen int64, val reflect.Value) error {
	tp := val.Type()
	for tp.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
		tp = tp.Elem()
	}
	if tp.Kind() != reflect.Struct {
		return nil
	}
	var err error
	for i := 0; i < tp.NumField(); i++ {
		ft := tp.Field(i)
		fv := val.Field(i)
		if fv.Kind() == reflect.Ptr {
			err = multierr.Append(err, validateMoneyReflect(expectedLen, fv))
		} else if fv.Kind() == reflect.Slice {
			for i := 0; i < fv.Len(); i++ {
				v := fv.Index(i)
				err = multierr.Append(err, validateMoneyReflect(expectedLen, v))
			}
		}
		if !strings.HasSuffix(ft.Name, "Money") {
			continue
		}
		if fv.Kind() == reflect.Slice {
			err = multierr.Append(err, validateMoneySlice(ft.Name, expectedLen, fv))
		} else {
			intVal := fv.Int()
			err = multierr.Append(err, validateMoneyValue(ft.Name, intVal))
		}
	}
	return err
}

func Validate(rules *rules_proto.Ruleset) error {
	err := ValidateGame(rules.GetGame())
	if err != nil {
		return err
	}
	err = validateMoney(rules.GetGame().GetNumberOfPlayers(), rules)
	if err != nil {
		return err
	}
	return validateEnumImpl(rules.GetYaku().GetRenhou(), base_proto.Limit_name, true)
}

func ValidateGame(rules *rules_proto.Game) error {
	var err error
	add := func(x error) {
		if x == nil {
			return
		}
		err = multierr.Append(err, x)
	}
	addf := func(format string, args ...interface{}) {
		add(fmt.Errorf(format, args...))
	}
	add(validateEnum(rules.GetAgariYame(), rules_proto.AgariYame_name))
	add(validateEnum(rules.GetChiShift(), rules_proto.Shifting_name))
	add(validateEnum(rules.GetAtodzuke(), rules_proto.Atodzuke_name))
	add(validateEnum(rules.GetKanDoraOpen(), rules_proto.KanDoraOpen_name))
	add(validateEnum(rules.GetLastWind(), base_proto.Wind_name))
	add(validateEnum(rules.GetMaxLastWind(), base_proto.Wind_name))
	num := rules.GetNumberOfPlayers()
	add(validateMoney(num, rules))
	switch num {
	case 2, 3, 4:
	default:
		addf("possble value for number_of_players are: 2,3,4. Got %v", num)
	}
	if rules.GetUmaShare() {
		checkDiv := func(caption string, values ...int64) {
			for k := int64(2); k <= num; k++ {
				for _, value := range values {
					if value%k != 0 {
						addf("value of %v=%v should be divideble by %v", caption, value, k)
					}
				}
			}
		}
		checkDiv("oka", rules.GetOkaMoney())
		// checkDiv("uma", rules.GetUma().GetDefaultMoney()...)
		// checkDiv("umap1", rules.GetUma().GetMinus1Money()...)
		// checkDiv("umam1", rules.GetUma().GetPlus1Money()...)
	}
	return err
}
