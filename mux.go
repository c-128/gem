package gem

import "strings"

func NewMux() *Mux {
	mux := &Mux{
		children: make(map[string]*Mux),
		handler:  nil,
		dynamic:  false,
	}
	return mux
}

type Mux struct {
	children map[string]*Mux
	handler  Handler
	dynamic  bool
}

func (mux *Mux) Handle(path string, handler Handler) {
	segments := mux.splitPath(path)

	mux.addRoute(segments, handler)
}

func (mux *Mux) Handler(ctx *Ctx) error {
	url := ctx.URL()
	path := url.Path
	segments := mux.splitPath(path)

	route := mux.search(segments, ctx.params)
	if route == nil {
		return NotFoundErr
	}

	if route.handler == nil {
		return NotFoundErr
	}

	err := route.handler(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (mux *Mux) addRoute(segments []string, handler Handler) {
	if len(segments) == 0 {
		mux.handler = handler
		return
	}

	component := segments[0]
	dynamic := false
	if strings.HasPrefix(component, ":") {
		component = strings.TrimPrefix(component, ":")
		dynamic = true
	}

	child, found := mux.children[component]
	if !found {
		child = NewMux()
		child.dynamic = dynamic
		mux.children[component] = child
	}

	child.addRoute(segments[1:], handler)
}

func (mux *Mux) search(segments []string, params map[string]string) *Mux {
	if len(segments) == 0 {
		return mux
	}

	component := segments[0]
	for name, child := range mux.children {
		if child.dynamic {
			params[name] = component
			return child.search(segments[1:], params)
		}

		if name == component {
			return child.search(segments[1:], params)
		}
	}

	return nil
}

func (mux *Mux) splitPath(path string) []string {
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	if path == "" {
		return make([]string, 0)
	}

	return strings.Split(path, "/")
}
