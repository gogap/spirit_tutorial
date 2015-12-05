package component

import (
	"encoding/base64"
	"github.com/gogap/errors"
	"github.com/gogap/spirit"
	"strings"

	"github.com/gogap/spirit_tutorial/todo/controllers"
)

var _ spirit.Component = new(AuthComponent)

const (
	taskAuthComponentURN = "urn:spirit_tutorial:component:todo:auth"
)

var (
	ErrAuthorizationHeaderNotExist = errors.TN(TodoErrorNamespace, 30001, "header of Authorization not exist!")
	ErrAuthorizationHeaderFmtError = errors.TN(TodoErrorNamespace, 30002, "header of Authorization format error.")
)

type AuthComponent struct {
	name string
	*controllers.Auth
}

func init() {
	spirit.RegisterComponent(taskAuthComponentURN, NewAuthComponent)
}

func NewAuthComponent(name string, options spirit.Map) (component spirit.Component, err error) {
	component = &AuthComponent{
		name: name,
		Auth: controllers.NewAuth(),
	}
	return
}

func (p *AuthComponent) Name() string {
	return p.name
}

func (p *AuthComponent) URN() string {
	return taskAuthComponentURN
}

func (p *AuthComponent) Labels() spirit.Labels {
	return spirit.Labels{
		"version": "0.0.1",
	}
}

func (p *AuthComponent) CheckAuth(payload spirit.Payload) (result interface{}, err error) {

	if result, err = payload.GetData(); err != nil {
		return
	}

	var contexts interface{}
	var headerExist bool
	if contexts, headerExist = payload.GetContext(CtxHttpHeaders); !headerExist {
		err = ErrAuthorizationHeaderNotExist.New()
		return
	}

	if headers, ok := contexts.(map[string]interface{}); !ok {
		err = ErrAuthorizationHeaderNotExist.New()
		return
	} else if v, exist := headers["Authorization"]; !exist {
		err = ErrAuthorizationHeaderNotExist.New()
		return
	} else if authVal, ok := v.(string); !ok {
		err = ErrAuthorizationHeaderNotExist.New()
		return
	} else {
		authVal = strings.TrimLeft(authVal, "Basic ")

		var val []byte
		if val, err = base64.StdEncoding.DecodeString(authVal); err != nil {
			return
		}

		account := strings.Split(string(val), ":")

		if len(account) != 2 {
			err = ErrAuthorizationHeaderFmtError.New()
			return
		}

		if err = p.Auth.CheckAuth(account[0], account[1]); err != nil {
			return
		}
	}

	return
}
