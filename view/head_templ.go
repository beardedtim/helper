// Code generated by templ@v0.2.364 DO NOT EDIT.

package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"strings"
)

func favIcons(pageData *PageData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(strings.Join([]string{
			pageData.AssetURL,
			"apple-touch-icon.png",
		}, "/")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><link rel=\"icon\" type=\"image/png\" sizes=\"32x32\" href=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(strings.Join([]string{
			pageData.AssetURL,
			"favicon-32x32.png",
		}, "/")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><link rel=\"icon\" type=\"image/png\" sizes=\"16x16\" href=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(strings.Join([]string{
			pageData.AssetURL,
			"favicon-16x16.png",
		}, "/")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><link rel=\"manifest\" href=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(strings.Join([]string{
			pageData.AssetURL,
			"site.webmanifest",
		}, "/")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func meta() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_2 := templ.GetChildren(ctx)
		if var_2 == nil {
			var_2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func globalHeadScripts(pageData *PageData) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_3 := templ.GetChildren(ctx)
		if var_3 == nil {
			var_3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<script src=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(strings.Join([]string{
			pageData.AssetURL,
			"js",
			"global.js",
		}, "/")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var_4 := ``
		_, err = templBuffer.WriteString(var_4)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func globalStyleSheets(pageData *PageData, customSheets []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_5 := templ.GetChildren(ctx)
		if var_5 == nil {
			var_5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<link rel=\"stylesheet\" href=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(strings.Join([]string{
			pageData.AssetURL,
			"css",
			"global.css",
		}, "/")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		for _, sheetName := range customSheets {
			_, err = templBuffer.WriteString("<link rel=\"stylesheet\" href=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString(strings.Join([]string{
				pageData.AssetURL,
				"css",
				sheetName,
			}, "/")))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\">")
			if err != nil {
				return err
			}
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func Head(pageData *PageData, customSheets []string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_6 := templ.GetChildren(ctx)
		if var_6 == nil {
			var_6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<head>")
		if err != nil {
			return err
		}
		err = meta().Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = favIcons(pageData).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = globalHeadScripts(pageData).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = globalStyleSheets(pageData, customSheets).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</head>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
