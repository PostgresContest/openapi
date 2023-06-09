// Code generated by ogen, DO NOT EDIT.

package v1

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}
	if prefix := s.cfg.Prefix; len(prefix) > 0 {
		if strings.HasPrefix(elem, prefix) {
			// Cut prefix from the path.
			elem = strings.TrimPrefix(elem, prefix)
		} else {
			// Prefix doesn't match.
			s.notFound(w, r)
			return
		}
	}
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "a"
				if l := len("a"); len(elem) >= l && elem[0:l] == "a" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 't': // Prefix: "ttempt/"
					if l := len("ttempt/"); len(elem) >= l && elem[0:l] == "ttempt/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "attempt_id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleAttemptAttemptIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				case 'u': // Prefix: "uth/"
					if l := len("uth/"); len(elem) >= l && elem[0:l] == "uth/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'l': // Prefix: "login"
						if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "POST":
								s.handleAuthLoginPostRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
					case 'v': // Prefix: "verify"
						if l := len("verify"); len(elem) >= l && elem[0:l] == "verify" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleAuthVerifyGetRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}
					}
				}
			case 't': // Prefix: "task"
				if l := len("task"); len(elem) >= l && elem[0:l] == "task" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "POST":
						s.handleTaskPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "task_id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleTaskTaskIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/attempt"
						if l := len("/attempt"); len(elem) >= l && elem[0:l] == "/attempt" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "POST":
								s.handleTaskTaskIDAttemptPostRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "POST")
							}

							return
						}
						switch elem[0] {
						case 's': // Prefix: "s"
							if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "GET":
									s.handleTaskTaskIDAttemptsGetRequest([1]string{
										args[0],
									}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "GET")
								}

								return
							}
						}
					}
				case 's': // Prefix: "s"
					if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handleTasksGetRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
				}
			case 'u': // Prefix: "user"
				if l := len("user"); len(elem) >= l && elem[0:l] == "user" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleUserGetRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}
			}
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'a': // Prefix: "a"
				if l := len("a"); len(elem) >= l && elem[0:l] == "a" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 't': // Prefix: "ttempt/"
					if l := len("ttempt/"); len(elem) >= l && elem[0:l] == "ttempt/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "attempt_id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: AttemptAttemptIDGet
							r.name = "AttemptAttemptIDGet"
							r.operationID = ""
							r.pathPattern = "/attempt/{attempt_id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
				case 'u': // Prefix: "uth/"
					if l := len("uth/"); len(elem) >= l && elem[0:l] == "uth/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'l': // Prefix: "login"
						if l := len("login"); len(elem) >= l && elem[0:l] == "login" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								// Leaf: AuthLoginPost
								r.name = "AuthLoginPost"
								r.operationID = ""
								r.pathPattern = "/auth/login"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					case 'v': // Prefix: "verify"
						if l := len("verify"); len(elem) >= l && elem[0:l] == "verify" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "GET":
								// Leaf: AuthVerifyGet
								r.name = "AuthVerifyGet"
								r.operationID = ""
								r.pathPattern = "/auth/verify"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
					}
				}
			case 't': // Prefix: "task"
				if l := len("task"); len(elem) >= l && elem[0:l] == "task" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "POST":
						r.name = "TaskPost"
						r.operationID = ""
						r.pathPattern = "/task"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "task_id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = "TaskTaskIDGet"
							r.operationID = ""
							r.pathPattern = "/task/{task_id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/attempt"
						if l := len("/attempt"); len(elem) >= l && elem[0:l] == "/attempt" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "POST":
								r.name = "TaskTaskIDAttemptPost"
								r.operationID = ""
								r.pathPattern = "/task/{task_id}/attempt"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case 's': // Prefix: "s"
							if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								switch method {
								case "GET":
									// Leaf: TaskTaskIDAttemptsGet
									r.name = "TaskTaskIDAttemptsGet"
									r.operationID = ""
									r.pathPattern = "/task/{task_id}/attempts"
									r.args = args
									r.count = 1
									return r, true
								default:
									return
								}
							}
						}
					}
				case 's': // Prefix: "s"
					if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							// Leaf: TasksGet
							r.name = "TasksGet"
							r.operationID = ""
							r.pathPattern = "/tasks"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
				}
			case 'u': // Prefix: "user"
				if l := len("user"); len(elem) >= l && elem[0:l] == "user" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: UserGet
						r.name = "UserGet"
						r.operationID = ""
						r.pathPattern = "/user"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
			}
		}
	}
	return r, false
}
