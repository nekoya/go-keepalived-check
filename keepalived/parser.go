//line parser.go.y:2
package keepalived

import __yyfmt__ "fmt"

//line parser.go.y:2
//line parser.go.y:5
type yySymType struct {
	yys        int
	statements []Statement
	statement  Statement
	exprs      []Expression
	expr       Expression
	tok        Token
}

const IDENT = 57346
const NUMBER = 57347
const IP = 57348
const STRING = 57349
const GLOBAL_DEFS = 57350
const NOTIFICATION_EMAIL = 57351
const NOTIFICATION_EMAIL_FROM = 57352
const SMTP_SERVER = 57353
const SMTP_CONNECT_TIMEOUT = 57354
const ROUTER_ID = 57355
const ENABLE_TRAPS = 57356
const VRRP_SYNC_GROUP = 57357
const GROUP = 57358
const NOTIFY_MASTER = 57359
const NOTIFY_BACKUP = 57360
const NOTIFY_FAULT = 57361
const NOTIFY = 57362
const SMTP_ALERT = 57363
const VRRP_INSTANCE = 57364
const STATE = 57365
const MASTER = 57366
const BACKUP = 57367
const INTERFACE = 57368
const GARP_MASTER_DELAY = 57369
const VIRTUAL_ROUTER_ID = 57370
const PRIORITY = 57371
const ADVERT_INT = 57372
const NOPREEMPT = 57373
const AUTHENTICATION = 57374
const AUTH_TYPE = 57375
const AUTH_PASS = 57376
const PASS = 57377
const AH = 57378
const VIRTUAL_IPADDRESS = 57379
const DEV = 57380
const VIRTUAL_SERVER_GROUP = 57381
const FWMARK = 57382
const VIRTUAL_SERVER = 57383
const DELAY_LOOP = 57384
const NAT_MASK = 57385
const SORRY_SERVER = 57386
const LVS_SCHED = 57387
const LB_ALGO = 57388
const RR = 57389
const WRR = 57390
const LC = 57391
const WLC = 57392
const LBLC = 57393
const SH = 57394
const DH = 57395
const LVS_METHOD = 57396
const LB_KIND = 57397
const NAT = 57398
const DR = 57399
const TUN = 57400
const PERSISTENCE_TIMEOUT = 57401
const PERSISTENCE_GRANULARITY = 57402
const PROTOCOL = 57403
const TCP = 57404
const UDP = 57405
const REAL_SERVER = 57406
const WEIGHT = 57407
const INHIBIT_ON_FAILURE = 57408
const HTTP_GET = 57409
const SSL_GET = 57410
const TCP_CHECK = 57411
const MISC_CHECK = 57412
const URL = 57413
const PATH = 57414
const DIGEST = 57415
const STATUS_CODE = 57416
const CONNECT_PORT = 57417
const CONNECT_TIMEOUT = 57418
const NB_GET_RETRY = 57419
const DELAY_BEFORE_RETRY = 57420
const MISC_PATH = 57421
const MISC_TIMEOUT = 57422
const MISC_DYNAMIC = 57423

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"NUMBER",
	"IP",
	"STRING",
	"GLOBAL_DEFS",
	"NOTIFICATION_EMAIL",
	"NOTIFICATION_EMAIL_FROM",
	"SMTP_SERVER",
	"SMTP_CONNECT_TIMEOUT",
	"ROUTER_ID",
	"ENABLE_TRAPS",
	"VRRP_SYNC_GROUP",
	"GROUP",
	"NOTIFY_MASTER",
	"NOTIFY_BACKUP",
	"NOTIFY_FAULT",
	"NOTIFY",
	"SMTP_ALERT",
	"VRRP_INSTANCE",
	"STATE",
	"MASTER",
	"BACKUP",
	"INTERFACE",
	"GARP_MASTER_DELAY",
	"VIRTUAL_ROUTER_ID",
	"PRIORITY",
	"ADVERT_INT",
	"NOPREEMPT",
	"AUTHENTICATION",
	"AUTH_TYPE",
	"AUTH_PASS",
	"PASS",
	"AH",
	"VIRTUAL_IPADDRESS",
	"DEV",
	"VIRTUAL_SERVER_GROUP",
	"FWMARK",
	"VIRTUAL_SERVER",
	"DELAY_LOOP",
	"NAT_MASK",
	"SORRY_SERVER",
	"LVS_SCHED",
	"LB_ALGO",
	"RR",
	"WRR",
	"LC",
	"WLC",
	"LBLC",
	"SH",
	"DH",
	"LVS_METHOD",
	"LB_KIND",
	"NAT",
	"DR",
	"TUN",
	"PERSISTENCE_TIMEOUT",
	"PERSISTENCE_GRANULARITY",
	"PROTOCOL",
	"TCP",
	"UDP",
	"REAL_SERVER",
	"WEIGHT",
	"INHIBIT_ON_FAILURE",
	"HTTP_GET",
	"SSL_GET",
	"TCP_CHECK",
	"MISC_CHECK",
	"URL",
	"PATH",
	"DIGEST",
	"STATUS_CODE",
	"CONNECT_PORT",
	"CONNECT_TIMEOUT",
	"NB_GET_RETRY",
	"DELAY_BEFORE_RETRY",
	"MISC_PATH",
	"MISC_TIMEOUT",
	"MISC_DYNAMIC",
	"'{'",
	"'}'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:377

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 115
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 239

var yyAct = [...]int{

	108, 166, 122, 113, 94, 102, 103, 95, 96, 199,
	200, 201, 40, 39, 42, 41, 97, 98, 141, 142,
	197, 99, 100, 101, 191, 165, 104, 38, 43, 164,
	45, 163, 60, 162, 56, 150, 89, 57, 58, 61,
	62, 63, 59, 65, 145, 106, 88, 75, 64, 37,
	76, 77, 78, 79, 94, 102, 103, 95, 96, 155,
	156, 157, 158, 159, 160, 138, 97, 98, 139, 36,
	35, 99, 100, 101, 34, 19, 104, 153, 18, 17,
	94, 102, 103, 95, 96, 48, 49, 50, 51, 52,
	53, 135, 97, 98, 54, 105, 9, 99, 100, 101,
	121, 126, 104, 132, 25, 26, 27, 28, 29, 30,
	176, 70, 180, 181, 172, 173, 174, 175, 74, 176,
	178, 92, 177, 172, 173, 174, 175, 40, 39, 42,
	41, 170, 184, 185, 186, 134, 182, 130, 131, 123,
	124, 125, 136, 149, 71, 69, 147, 148, 81, 82,
	133, 204, 46, 114, 115, 116, 117, 118, 119, 120,
	3, 14, 195, 193, 192, 167, 190, 4, 189, 188,
	187, 13, 161, 144, 5, 143, 127, 112, 23, 91,
	72, 73, 90, 87, 86, 194, 85, 84, 66, 44,
	22, 6, 21, 7, 109, 15, 40, 39, 42, 41,
	203, 202, 151, 128, 83, 20, 107, 12, 11, 10,
	1, 129, 68, 8, 137, 146, 80, 110, 183, 179,
	198, 171, 154, 93, 67, 140, 55, 47, 24, 2,
	169, 168, 196, 152, 33, 111, 32, 31, 16,
}
var yyPact = [...]int{

	152, -1000, 152, 14, 205, 204, 203, 155, -1000, -1000,
	-3, -4, -7, 201, 187, 185, 95, -1000, -1000, -1000,
	-8, -12, -13, -1000, -1000, -33, 192, 192, 184, 192,
	-1000, 69, 11, 105, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -35, 192,
	192, 192, 192, -1000, -1000, -1000, 124, 200, 182, -1000,
	-1000, 181, 179, 178, -36, -46, -1000, -1000, -1000, 177,
	174, 38, 12, -38, 123, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 172, 106, 106, 83, 83, 171,
	199, 75, 192, 144, 129, -1000, -1000, -1000, -1000, 8,
	59, -15, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 170, 168, -1000, -1000, -1000, 6, -1000,
	-1000, 111, 192, -1000, -47, 198, -1000, -1000, -1000, -1000,
	-1000, -1000, -6, -1000, -1000, 167, -1000, -49, -51, -53,
	-57, -1000, -1000, -1000, -1000, -1000, 48, 39, 37, 53,
	-1000, -1000, 165, 164, 163, 161, -58, -1000, -1000, -1000,
	159, 158, -1000, -1000, 192, 157, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -63, -1000, -1000, 197,
	196, 146, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 210, 238, 237, 236, 235, 234, 144, 233, 1,
	232, 231, 230, 229, 228, 227, 226, 225, 224, 223,
	222, 221, 220, 219, 218, 118, 217, 0, 216, 215,
	214, 212, 211, 3, 2,
}
var yyR1 = [...]int{

	0, 1, 1, 13, 13, 13, 13, 13, 13, 13,
	2, 2, 14, 14, 14, 14, 14, 14, 3, 3,
	15, 15, 15, 15, 15, 15, 4, 4, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 6, 6,
	18, 18, 7, 7, 19, 19, 19, 19, 19, 19,
	19, 19, 19, 19, 19, 32, 32, 33, 33, 33,
	33, 33, 33, 33, 34, 34, 34, 8, 8, 20,
	20, 20, 20, 20, 20, 9, 9, 21, 21, 21,
	21, 21, 11, 11, 23, 23, 12, 12, 24, 24,
	24, 10, 10, 22, 22, 22, 26, 26, 5, 5,
	17, 17, 29, 29, 28, 28, 30, 30, 31, 25,
	25, 27, 27, 27, 27,
}
var yyR2 = [...]int{

	0, 0, 2, 4, 5, 5, 5, 6, 6, 6,
	0, 2, 4, 2, 2, 2, 2, 1, 0, 2,
	4, 2, 2, 2, 2, 1, 0, 2, 2, 2,
	2, 1, 1, 2, 2, 2, 4, 4, 0, 2,
	1, 2, 0, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 6, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 0, 2, 2,
	1, 4, 4, 4, 4, 0, 2, 2, 2, 2,
	2, 4, 0, 2, 2, 2, 0, 2, 2, 2,
	1, 0, 2, 2, 2, 2, 0, 2, 0, 2,
	2, 2, 1, 1, 1, 1, 1, 3, 2, 0,
	2, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -13, 8, 15, 22, 39, 41, -1, 82,
	4, 4, 4, 16, 6, 40, -2, 82, 82, 82,
	4, 5, 5, 83, -14, 9, 10, 11, 12, 13,
	14, -3, -4, -6, 82, 82, 82, 82, -27, 5,
	4, 7, 6, -27, 5, -27, 83, -15, 16, 17,
	18, 19, 20, 21, 83, -16, 23, 26, 27, 31,
	21, 28, 29, 30, 37, 32, 83, -18, -31, 40,
	6, -7, -7, -7, -25, 82, -27, -27, -27, -27,
	-28, 24, 25, 4, 5, 5, 5, 5, 82, 82,
	5, 5, 83, -19, 42, 45, 46, 54, 55, 59,
	60, 61, 43, 44, 64, 83, 83, 83, -27, -25,
	-26, -5, 5, -33, 47, 48, 49, 50, 51, 52,
	53, -33, -34, 56, 57, 58, -34, 5, 4, -32,
	62, 63, -27, 6, 6, 83, 83, -30, 6, 83,
	-17, 33, 34, 5, 5, 38, -29, 35, 36, -27,
	82, 4, -8, 83, -20, 65, 66, 67, 68, 69,
	70, 5, 82, 82, 82, 82, -9, -9, -11, -12,
	83, -21, 75, 76, 77, 78, 71, 83, 83, -23,
	75, 76, 83, -24, 79, 80, 81, 5, 5, 5,
	5, 82, 5, 5, -27, 5, -10, 83, -22, 72,
	73, 74, 4, 4, 5,
}
var yyDef = [...]int{

	1, -2, 1, 0, 0, 0, 0, 0, 2, 10,
	0, 0, 0, 0, 0, 0, 0, 18, 26, 38,
	0, 0, 0, 3, 11, 0, 0, 0, 0, 0,
	17, 0, 0, 0, 42, 42, 42, 109, 13, 111,
	112, 113, 114, 14, 15, 16, 4, 19, 0, 0,
	0, 0, 0, 25, 5, 27, 0, 0, 0, 31,
	32, 0, 0, 0, 0, 0, 6, 39, 40, 0,
	0, 0, 0, 0, 0, 109, 21, 22, 23, 24,
	28, 104, 105, 29, 30, 33, 34, 35, 96, 98,
	41, 108, 7, 43, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 8, 9, 12, 110, 0,
	0, 0, 44, 45, 57, 58, 59, 60, 61, 62,
	63, 46, 47, 64, 65, 66, 48, 49, 50, 51,
	55, 56, 52, 0, 0, 20, 36, 97, 106, 37,
	99, 0, 0, 53, 0, 0, 100, 102, 103, 101,
	67, 107, 0, 54, 68, 0, 70, 0, 0, 0,
	0, 69, 75, 75, 82, 86, 0, 0, 0, 0,
	71, 76, 0, 0, 0, 0, 0, 72, 73, 83,
	0, 0, 74, 87, 0, 0, 90, 77, 78, 79,
	80, 91, 84, 85, 88, 89, 0, 81, 92, 0,
	0, 0, 93, 94, 95,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 82, 3, 83,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yychar = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:54
		{
			yyVAL.statements = nil
			if l, ok := yylex.(*Lexer); ok {
				l.statements = yyVAL.statements
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:61
		{
			yyVAL.statements = append([]Statement{yyDollar[1].statement}, yyDollar[2].statements...)
			if l, ok := yylex.(*Lexer); ok {
				l.statements = yyVAL.statements
			}
		}
	case 3:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:69
		{
			yyVAL.statement = &GlobalDefsStmt{stmts: yyDollar[3].statements}
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:73
		{
			yyVAL.statement = &VrrpSyncGroupStmt{group: yyDollar[2].tok.lit, stmts: yyDollar[4].statements}
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.statement = &VrrpInstanceStmt{inside_network: yyDollar[2].tok.lit, stmts: yyDollar[4].statements}
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:81
		{
			yyVAL.statement = &VirtualServerGroupStmt{group: yyDollar[2].tok.lit, stmts: yyDollar[4].statements}
		}
	case 7:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:85
		{
			yyVAL.statement = &VirtualServerStmt{group: yyDollar[3].tok.lit, stmts: yyDollar[5].statements}
		}
	case 8:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:89
		{
			yyVAL.statement = &VirtualServerStmt{ip: &IdentExpr{lit: yyDollar[2].tok.lit}, port: &NumExpr{lit: yyDollar[3].tok.lit}, stmts: yyDollar[5].statements}
		}
	case 9:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:93
		{
			yyVAL.statement = &VirtualServerStmt{fwmark: &NumExpr{lit: yyDollar[3].tok.lit}, stmts: yyDollar[5].statements}
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:98
		{
			yyVAL.statements = []Statement{}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:99
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 12:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.statement = &GroupStmt{exprs: yyDollar[3].exprs}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:103
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:104
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:105
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:106
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:107
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 18:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:110
		{
			yyVAL.statements = []Statement{}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:111
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:114
		{
			yyVAL.statement = &GroupStmt{exprs: yyDollar[3].exprs}
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:115
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:116
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:117
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:118
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:119
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:122
		{
			yyVAL.statements = []Statement{}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:123
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:127
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:135
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:151
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:155
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:159
		{
			yyVAL.statement = &VirtualIpaddressStmt{addresses: yyDollar[3].exprs}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:163
		{
			yyVAL.statement = &AuthenticationStmt{stmts: yyDollar[3].statements}
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:168
		{
			yyVAL.statements = []Statement{}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:169
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:172
		{
			yyVAL.statement = &ExprStmt{key: yyDollar[1].expr}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:176
		{
			yyVAL.statements = []Statement{}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:177
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:181
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:185
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:189
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:197
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:201
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:205
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:209
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:213
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:217
		{
			yyVAL.statement = &SorryServerStmt{ip: &IdentExpr{lit: yyDollar[2].tok.lit}, port: &NumExpr{lit: yyDollar[3].tok.lit}}
		}
	case 54:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:221
		{
			yyVAL.statement = &RealServerStmt{ip: &IdentExpr{lit: yyDollar[2].tok.lit}, port: &NumExpr{lit: yyDollar[3].tok.lit}, stmts: yyDollar[5].statements}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:226
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:227
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:230
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:231
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:232
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:233
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:234
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:235
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:236
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:239
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:240
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:241
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:244
		{
			yyVAL.statements = []Statement{}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:245
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:249
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:253
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 71:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:257
		{
			yyVAL.statement = &HttpGetStmt{stmts: yyDollar[3].statements}
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:261
		{
			yyVAL.statement = &SslGetStmt{stmts: yyDollar[3].statements}
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:265
		{
			yyVAL.statement = &TcpCheckStmt{stmts: yyDollar[3].statements}
		}
	case 74:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:269
		{
			yyVAL.statement = &MiscCheckStmt{stmts: yyDollar[3].statements}
		}
	case 75:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.statements = []Statement{}
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:275
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:279
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:283
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:287
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:291
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 81:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:295
		{
			yyVAL.statement = &UrlStmt{stmts: yyDollar[3].statements}
		}
	case 82:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:300
		{
			yyVAL.statements = []Statement{}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:301
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:305
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:309
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:314
		{
			yyVAL.statements = []Statement{}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:315
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:319
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:323
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:327
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:332
		{
			yyVAL.statements = []Statement{}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:333
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:336
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:337
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:338
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 96:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:341
		{
			yyVAL.exprs = []Expression{}
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:342
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:345
		{
			yyVAL.statements = []Statement{}
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:346
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:349
		{
			yyVAL.statement = &AuthTypeStmt{expr: yyDollar[2].expr}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:350
		{
			yyVAL.statement = &AuthPassStmt{expr: yyDollar[2].expr}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:353
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:354
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:357
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:358
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:361
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:362
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit, dev: yyDollar[3].tok.lit}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:365
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit, port: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 109:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:368
		{
			yyVAL.exprs = []Expression{}
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:369
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:372
		{
			yyVAL.expr = &NumExpr{lit: yyDollar[1].tok.lit}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:373
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:374
		{
			yyVAL.expr = &StringExpr{lit: yyDollar[1].tok.lit}
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:375
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit}
		}
	}
	goto yystack /* stack new state and value */
}
