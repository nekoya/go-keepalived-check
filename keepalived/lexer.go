package keepalived

import (
	"log"
	"regexp"
)

const (
	EOF = -1
	EOL = '\n'
)

var keywords = map[string]int{
	// global_defs
	"global_defs":             GLOBAL_DEFS,
	"notification_email":      NOTIFICATION_EMAIL,
	"notification_email_from": NOTIFICATION_EMAIL_FROM,
	"smtp_server":             SMTP_SERVER,
	"smtp_connect_timeout":    SMTP_CONNECT_TIMEOUT,
	"router_id":               ROUTER_ID,
	"enable_traps":            ENABLE_TRAPS,

	// authentication
	"authentication": AUTHENTICATION,
	"auth_type":      AUTH_TYPE,
	"auth_pass":      AUTH_PASS,
	"PASS":           PASS,
	"AH":             AH,

	// state
	"state":  STATE,
	"MASTER": MASTER,
	"BACKUP": BACKUP,

	// vrrp_sync_group
	"vrrp_sync_group": VRRP_SYNC_GROUP,
	"group":           GROUP,
	"notify_master":   NOTIFY_MASTER,
	"notify_backup":   NOTIFY_BACKUP,
	"notify_fault":    NOTIFY_FAULT,
	"notify":          NOTIFY,
	"smtp_alert":      SMTP_ALERT,

	// vrrp_instance
	"vrrp_instance":     VRRP_INSTANCE,
	"interface":         INTERFACE,
	"garp_master_delay": GARP_MASTER_DELAY,
	"virtual_router_id": VIRTUAL_ROUTER_ID,
	"priority":          PRIORITY,
	"advert_int":        ADVERT_INT,
	"nopreempt":         NOPREEMPT,
	"virtual_ipaddress": VIRTUAL_IPADDRESS,
	"dev":               DEV,

	// virtual_server_group
	"virtual_server_group": VIRTUAL_SERVER_GROUP,
	"fwmark":               FWMARK,

	// virtual_server
	"virtual_server":          VIRTUAL_SERVER,
	"delay_loop":              DELAY_LOOP,
	"nat_mask":                NAT_MASK,
	"sorry_server":            SORRY_SERVER,
	"persistence_timeout":     PERSISTENCE_TIMEOUT,
	"persistence_granularity": PERSISTENCE_GRANULARITY,
	// LVS scheduler
	"lvs_sched": LVS_SCHED,
	"lb_algo":   LB_ALGO,
	"rr":        RR,
	"wrr":       WRR,
	"lc":        LC,
	"wlc":       WLC,
	"lblc":      LBLC,
	"sh":        SH,
	"dh":        DH,
	// LVS forwarding method
	"lvs_method": LVS_METHOD,
	"lb_kind":    LB_KIND,
	"NAT":        NAT,
	"DR":         DR,
	"TUN":        TUN,
	"protocol":   PROTOCOL,
	"TCP":        TCP,
	"UDP":        UDP,
	// real_server
	"real_server":        REAL_SERVER,
	"weight":             WEIGHT,
	"inhibit_on_failure": INHIBIT_ON_FAILURE,
	// HTTP_GET|SSL_GET
	"HTTP_GET":           HTTP_GET,
	"SSL_GET":            SSL_GET,
	"url":                URL,
	"path":               PATH,
	"digest":             DIGEST,
	"status_code":        STATUS_CODE,
	"connect_port":       CONNECT_PORT,
	"connect_timeout":    CONNECT_TIMEOUT,
	"nb_get_retry":       NB_GET_RETRY,
	"delay_before_retry": DELAY_BEFORE_RETRY,
	// TCP_CHECK
	"TCP_CHECK": TCP_CHECK,
	// MISC_CHECK
	"MISC_CHECK":   MISC_CHECK,
	"misc_path":    MISC_PATH,
	"misc_timeout": MISC_TIMEOUT,
	"misc_dynamic": MISC_DYNAMIC,
}

type Position struct {
	Line   int
	Column int
}

type Scanner struct {
	src      []rune
	offset   int
	lineHead int
	line     int
}

func (s *Scanner) Init(src string) {
	s.src = []rune(src)
}

func (s *Scanner) Scan() (tok int, lit string, pos Position) {
retry:
	s.skipWhiteSpace()
	pos = s.position()
	switch ch := s.peek(); {
	case isLetter(ch):
		lit = s.scanIdentifier()
		if keyword, ok := keywords[lit]; ok {
			tok = keyword
		} else {
			tok = IDENT
		}
	case isDigit(ch):
		lit = s.scanIdentifier()
		if matched, _ := regexp.MatchString("^[0-9]+$", lit); matched {
			tok = NUMBER
		} else if matched, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}(|\\-[0-9]{1,3}|/[0-9]{1,2})$", lit); matched {
			tok = IP
		} else {
			tok = IDENT
		}
	default:
		switch ch {
		case -1:
			tok = EOF
		case '#', '!':
			for !isEOL(s.peek()) {
				s.next()
			}
			goto retry
		case '{', '}':
			tok = int(ch)
			lit = string(ch)
		case '"':
			tok = STRING
			lit = s.scanString('"')
		}
		s.next()
	}
	return
}

// ========================================

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '.' || ch == '/' || ch == '@' || ch == '-'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isIp(ch rune) bool {
	return ('0' <= ch && ch <= '9') || ch == '.'
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// isEOL return true if the rune is at end-of-line or end-of-file.
func isEOL(ch rune) bool {
	return ch == '\n' || ch == -1
}

func (s *Scanner) peek() rune {
	if !s.reachEOF() {
		return s.src[s.offset]
	} else {
		return -1
	}
}

func (s *Scanner) next() {
	if !s.reachEOF() {
		if s.peek() == '\n' {
			s.lineHead = s.offset + 1
			s.line++
		}
		s.offset++
	}
}

func (s *Scanner) reachEOF() bool {
	return len(s.src) <= s.offset
}

func (s *Scanner) position() Position {
	return Position{Line: s.line + 1, Column: s.offset - s.lineHead + 1}
}

func (s *Scanner) skipWhiteSpace() {
	for isWhiteSpace(s.peek()) {
		s.next()
	}
}

func (s *Scanner) scanIdentifier() string {
	var ret []rune
	for isLetter(s.peek()) || isDigit(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret)
}

func (s *Scanner) scanNumber() string {
	var ret []rune
	for isIp(s.peek()) {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret)
}

func (s *Scanner) scanString(l rune) string {
	s.next()
	var ret []rune
	for s.peek() != l {
		ret = append(ret, s.peek())
		s.next()
	}
	return string(ret)
}

type Lexer struct {
	s          *Scanner
	recentLit  string
	recentPos  Position
	statements []Statement
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit, pos := l.s.Scan()
	if tok == EOF {
		return 0
	}
	lval.tok = Token{tok: tok, lit: lit, pos: pos}
	l.recentLit = lit
	l.recentPos = pos
	return tok
}

func (l *Lexer) Error(e string) {
	log.Fatalf("Line %d, Column %d: %q %s",
		l.recentPos.Line, l.recentPos.Column, l.recentLit, e)
}

func Parse(s *Scanner) []Statement {
	l := Lexer{s: s}
	if yyParse(&l) != 0 {
		panic("Parse error")
	}
	return l.statements
}
