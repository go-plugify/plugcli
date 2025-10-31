package main

import "context"

var ExportPlugin = Plugin{
	BasePlugin: &BasePlugin{},
}

type BasePlugin struct {
	Name        string
	Description string
	Version     string
	Components  PluginComponents
}

type PluginComponents interface {
	GetLogger() any
	GetUtil() any
	GetComponents() any
}

func (p *BasePlugin) Load(components any) error {
	p.Components = components.(PluginComponents)
	return nil
}

func (p *BasePlugin) Destroy(args any) error {
	return nil
}

type Plugin struct {
	*BasePlugin
}

func (p Plugin) GetName() string {
	return p.Name
}

func (p Plugin) GetDescription() string {
	return p.Description
}

type Logger interface {
	WarnCtx(ctx context.Context, format string, args ...any)
	ErrorCtx(ctx context.Context, format string, args ...any)
	InfoCtx(ctx context.Context, format string, args ...any)
	Warn(format string, args ...any)
	Error(format string, args ...any)
	Info(format string, args ...any)
}

type Util interface {
	GetAttr(obj any, attrName string) any
	CallMethod(obj any, methodName string, args ...any) ([]any, error)
	ConvertTo(src, dist any) error
}

type Components interface {
	Get(name string) any
}

type Component interface {
	Name() string
	Service() any
}

func (p Plugin) Component(name string) any {
	return p.Components.GetComponents().(Components).Get(name).(Component).Service()
}

func (p Plugin) Logger() Logger {
	return p.Components.GetLogger().(Logger)
}

func (p Plugin) GetAttr(obj any, attrName string) any {
	return p.Components.GetUtil().(Util).GetAttr(obj, attrName)
}

func (p Plugin) CallMethod(obj any, methodName string, args ...any) ([]any, error) {
	return p.Components.GetUtil().(Util).CallMethod(obj, methodName, args...)
}

func (p Plugin) CallIgnore(obj any, methodName string, args ...any) []any {
	resp, _ := p.Components.GetUtil().(Util).CallMethod(obj, methodName, args...)
	return resp
}

func (p Plugin) ConvertTo(src, dist any) error {
	return p.Components.GetUtil().(Util).ConvertTo(src, dist)
}

type HttpContext interface {
	Query(key string) string
	JSON(code int, obj any)
}

type HttpRouter interface {
	ReplaceHandler(method, path string, handler func(ctx context.Context)) error
	GetHandler(method, path string) (func(ctx context.Context), error)
	GetHandlerName(method, path string) (string, error)
}
