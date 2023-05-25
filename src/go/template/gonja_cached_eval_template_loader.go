package template

import (
    "github.com/nikolalohinski/gonja"
    "github.com/nikolalohinski/gonja/exec"
)

type GonjaCachedEvalTemplateLoader struct {
    env *gonja.Environment
}

func NewGonjaCachedEvalTemplateLoader(env *gonja.Environment) *GonjaCachedEvalTemplateLoader {
    return &GonjaCachedEvalTemplateLoader{env}
}

func (loader *GonjaCachedEvalTemplateLoader) GetTemplate(filename string) (*exec.Template, error) {
    tpl, has := loader.env.Cache[filename]
    if has {
        return tpl, nil
    }

    return loader.env.FromFile(filename)
}

func (loader *GonjaCachedEvalTemplateLoader) Path(path string) (string, error) {
    return loader.env.Loader.Path(path)
}
