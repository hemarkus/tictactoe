package tictactoe

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func genericDialog(prompt string, defaultResult interface{}, result interface{}) error {
	resVal := reflect.ValueOf(result)
	// we don't return anything, so result must point to something writeable
	if resVal.Kind() != reflect.Ptr {
		return errors.New("invalid call by value")
	}
	for resVal.Kind() == reflect.Ptr || resVal.Kind() == reflect.Interface {
		resVal = resVal.Elem()
	}
	resKind := resVal.Kind()

	defVal := reflect.ValueOf(defaultResult)
	defaultProvided := reflect.TypeOf(defaultResult) != nil
	if defaultProvided {
		for defVal.Kind() == reflect.Ptr || defVal.Kind() == reflect.Interface {
			defVal = defVal.Elem()
		}
		defKind := defVal.Kind()

		if resKind != defKind {
			return errors.New("inconsistant types provided")
		}
	}

	var response string
	for response == "" {
		fmt.Printf("%s: ", prompt)
		reader := bufio.NewReader(os.Stdin)
		var err error
		response, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Please ... %v\n", err)
			continue
		}
		response = strings.TrimRight(response, "\n")
		if response == "" && defaultProvided {
			resVal.Set(defVal)
			return nil
		}
	}

	switch resKind {
	case reflect.String:
		resVal.SetString(response)
	case reflect.Uint:
		resI, err := strconv.ParseUint(response, 10, 64)
		if err != nil {
			return err
		}
		resVal.SetUint(resI)
	default:
		return fmt.Errorf("failed to read unsupported type %v", resKind)
	}
	return nil
}

func StringDialogDefault(prompt string, defaultResult string) (string, error) {
	var result string
	err := genericDialog(prompt, &defaultResult, &result)
	if err != nil {
		return "", err
	}
	return result, nil
}

func UintDialogDefault(prompt string, defaultResult uint) (uint, error) {
	var result uint
	err := genericDialog(prompt, &defaultResult, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func UintDialog(prompt string) (uint, error) {
	var result uint
	err := genericDialog(prompt, nil, &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func StringOptionDialog(prompt string, options []string) (string, error) {
	var result string
	for {
		newPrompt := fmt.Sprintf("%s [%s]", prompt, strings.Join(options, ", "))
		err := genericDialog(newPrompt, options[0], &result)
		if err != nil {
			return "", err
		}
		for _, o := range options {
			if o == result {
				return o, nil
			}
		}
	}
}
