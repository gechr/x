//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package terminal

type responseState uint8

const (
	responseGround responseState = iota
	responseEscape
	responseCSI
	responseDA1
	responseOSC
	responseOSCEscape
	responseString
)

type queryResponse struct {
	state      responseState
	payload    [maxBackgroundResponse]byte
	length     int
	overflow   bool
	background string
}

func (r *queryResponse) consume(b byte) bool {
	if b == '\x18' || b == '\x1a' {
		r.state = responseGround
		return false
	}

	switch r.state {
	case responseGround:
		if b == '\x1b' {
			r.state = responseEscape
		}
	case responseEscape:
		r.escape(b)
	case responseCSI:
		r.state = responseGround
		switch b {
		case '?':
			r.state = responseDA1
		case '\x1b':
			r.state = responseEscape
		}
	case responseDA1:
		switch {
		case b == 'c':
			r.state = responseGround
			return true
		case b >= '0' && b <= '9', b == ';':
		case b == '\x1b':
			r.state = responseEscape
		default:
			r.state = responseGround
		}
	case responseOSC:
		switch b {
		case '\a':
			r.finishOSC()
		case '\x1b':
			r.state = responseOSCEscape
		default:
			if r.length < len(r.payload) {
				r.payload[r.length] = b
				r.length++
			} else {
				r.overflow = true
			}
		}
	case responseOSCEscape:
		if b == '\\' {
			r.finishOSC()
		} else {
			r.escape(b)
		}
	case responseString:
		if b == '\x1b' {
			r.state = responseEscape
		}
	}
	return false
}

func (r *queryResponse) escape(b byte) {
	r.state = responseGround
	switch b {
	case '\x1b':
		r.state = responseEscape
	case '[':
		r.state = responseCSI
	case ']':
		r.state = responseOSC
		r.length = 0
		r.overflow = false
	case 'P', 'X', '^', '_':
		r.state = responseString
	}
}

func (r *queryResponse) finishOSC() {
	r.state = responseGround
	if r.overflow {
		return
	}
	response := "\x1b]" + string(r.payload[:r.length]) + "\a"
	if _, _, _, ok := parseBackgroundResponse(response); ok {
		r.background = response
	}
}
