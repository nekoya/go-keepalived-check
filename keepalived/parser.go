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
const VIRTUAL_SERVER = 57382
const DELAY_LOOP = 57383
const SORRY_SERVER = 57384
const LVS_SCHED = 57385
const LB_ALGO = 57386
const RR = 57387
const WRR = 57388
const LC = 57389
const WLC = 57390
const LBLC = 57391
const SH = 57392
const DH = 57393
const LVS_METHOD = 57394
const LB_KIND = 57395
const NAT = 57396
const DR = 57397
const TUN = 57398
const PERSISTENCE_TIMEOUT = 57399
const PERSISTENCE_GRANULARITY = 57400
const PROTOCOL = 57401
const TCP = 57402
const REAL_SERVER = 57403
const WEIGHT = 57404
const INHIBIT_ON_FAILURE = 57405
const HTTP_GET = 57406
const SSL_GET = 57407
const TCP_CHECK = 57408
const MISC_CHECK = 57409
const URL = 57410
const PATH = 57411
const DIGEST = 57412
const STATUS_CODE = 57413
const CONNECT_PORT = 57414
const CONNECT_TIMEOUT = 57415
const NB_GET_RETRY = 57416
const DELAY_BEFORE_RETRY = 57417
const MISC_PATH = 57418
const MISC_TIMEOUT = 57419

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
	"VIRTUAL_SERVER",
	"DELAY_LOOP",
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
	"'{'",
	"'}'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:354

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 108
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 214

var yyAct = [...]int{

	154, 99, 113, 104, 87, 95, 88, 89, 143, 144,
	145, 146, 147, 148, 178, 90, 91, 153, 126, 152,
	92, 93, 94, 151, 96, 141, 35, 40, 150, 42,
	138, 83, 45, 46, 47, 48, 49, 50, 37, 36,
	39, 38, 97, 87, 95, 88, 89, 82, 70, 71,
	72, 73, 65, 69, 90, 91, 186, 187, 188, 92,
	93, 94, 57, 96, 53, 120, 184, 54, 55, 58,
	59, 60, 56, 62, 37, 36, 39, 38, 61, 172,
	173, 85, 170, 23, 24, 25, 26, 27, 28, 129,
	130, 124, 164, 112, 117, 43, 160, 161, 162, 163,
	164, 168, 169, 165, 160, 161, 162, 163, 166, 34,
	68, 158, 3, 123, 33, 32, 18, 17, 16, 4,
	51, 9, 114, 115, 116, 63, 5, 133, 135, 136,
	66, 181, 137, 75, 76, 127, 105, 106, 107, 108,
	109, 110, 111, 6, 7, 14, 122, 121, 191, 98,
	182, 180, 155, 21, 179, 13, 37, 36, 39, 38,
	177, 176, 175, 174, 67, 149, 132, 131, 118, 103,
	84, 81, 80, 79, 78, 41, 20, 190, 189, 139,
	100, 119, 77, 19, 12, 11, 10, 1, 64, 125,
	8, 134, 74, 31, 101, 171, 167, 185, 159, 142,
	86, 128, 52, 44, 22, 2, 157, 156, 183, 140,
	102, 30, 29, 15,
}
var yyPact = [...]int{

	104, -1000, 104, 43, 182, 181, 180, 139, -1000, -1000,
	40, 39, 38, 179, 171, 74, -1000, -1000, -1000, 37,
	36, -1000, -1000, 31, 152, 152, 170, 152, -1000, 16,
	41, 46, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -25, 152, 152, 152, 152,
	-1000, -1000, -1000, 109, 178, 169, -1000, -1000, 168, 167,
	166, -31, -47, -1000, -1000, 165, 2, -37, 70, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 164, 91, 91,
	68, 68, 163, 177, 5, 141, 140, -1000, -1000, -1000,
	34, 12, 56, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 162, 161, -1000, -1000, -1000, 89, -1000, -1000, 93,
	152, -1000, -48, 175, -1000, -1000, -1000, -1000, -1000, -1000,
	-54, -1000, -1000, 160, -1000, -50, -55, -59, -61, -1000,
	-1000, -1000, -1000, -1000, 32, 24, 29, 3, -1000, -1000,
	158, 157, 156, 155, -64, -1000, -1000, -1000, 149, 146,
	-1000, -1000, 124, 145, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -13, -1000, -1000, 174, 173, 143, -1000,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 187, 213, 212, 211, 210, 130, 209, 0, 208,
	207, 206, 205, 204, 203, 202, 201, 200, 199, 198,
	197, 196, 195, 110, 194, 193, 1, 192, 191, 189,
	188, 3, 2,
}
var yyR1 = [...]int{

	0, 1, 1, 12, 12, 12, 12, 12, 12, 2,
	2, 13, 13, 13, 13, 13, 13, 3, 3, 14,
	14, 14, 14, 14, 14, 4, 4, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 15, 6, 6, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 31,
	31, 31, 31, 31, 31, 31, 32, 32, 32, 7,
	7, 18, 18, 18, 18, 18, 18, 8, 8, 19,
	19, 19, 19, 19, 10, 10, 21, 21, 11, 11,
	22, 22, 9, 9, 20, 20, 20, 24, 24, 25,
	25, 5, 5, 16, 16, 28, 28, 27, 27, 29,
	29, 30, 23, 23, 26, 26, 26, 26,
}
var yyR2 = [...]int{

	0, 0, 2, 4, 5, 5, 5, 6, 6, 0,
	2, 4, 2, 2, 2, 2, 1, 0, 2, 4,
	2, 2, 2, 2, 1, 0, 2, 2, 2, 2,
	1, 1, 2, 2, 2, 4, 4, 0, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 6, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
	2, 2, 1, 4, 4, 4, 4, 0, 2, 2,
	2, 2, 2, 4, 0, 2, 2, 2, 0, 2,
	2, 2, 0, 2, 2, 2, 2, 0, 2, 0,
	2, 0, 2, 2, 2, 1, 1, 1, 1, 1,
	3, 2, 0, 2, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -12, 8, 15, 22, 39, 40, -1, 78,
	4, 4, 4, 16, 6, -2, 78, 78, 78, 4,
	5, 79, -13, 9, 10, 11, 12, 13, 14, -3,
	-4, -25, 78, 78, 78, -26, 5, 4, 7, 6,
	-26, 5, -26, 79, -14, 16, 17, 18, 19, 20,
	21, 79, -15, 23, 26, 27, 31, 21, 28, 29,
	30, 37, 32, 79, -30, 6, -6, -6, -23, 78,
	-26, -26, -26, -26, -27, 24, 25, 4, 5, 5,
	5, 5, 78, 78, 5, 79, -17, 41, 43, 44,
	52, 53, 57, 58, 59, 42, 61, 79, 79, -26,
	-23, -24, -5, 5, -31, 45, 46, 47, 48, 49,
	50, 51, -31, -32, 54, 55, 56, -32, 5, 4,
	60, 6, 6, 79, 79, -29, 6, 79, -16, 33,
	34, 5, 5, 38, -28, 35, 36, -26, 78, 4,
	-7, 79, -18, 62, 63, 64, 65, 66, 67, 5,
	78, 78, 78, 78, -8, -8, -10, -11, 79, -19,
	72, 73, 74, 75, 68, 79, 79, -21, 72, 73,
	79, -22, 76, 77, 5, 5, 5, 5, 78, 5,
	5, 7, 5, -9, 79, -20, 69, 70, 71, 4,
	4, 5,
}
var yyDef = [...]int{

	1, -2, 1, 0, 0, 0, 0, 0, 2, 9,
	0, 0, 0, 0, 0, 0, 17, 25, 89, 0,
	0, 3, 10, 0, 0, 0, 0, 0, 16, 0,
	0, 0, 37, 37, 102, 12, 104, 105, 106, 107,
	13, 14, 15, 4, 18, 0, 0, 0, 0, 0,
	24, 5, 26, 0, 0, 0, 30, 31, 0, 0,
	0, 0, 0, 6, 90, 0, 0, 0, 0, 102,
	20, 21, 22, 23, 27, 97, 98, 28, 29, 32,
	33, 34, 87, 91, 101, 7, 38, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 8, 11, 103,
	0, 0, 0, 39, 40, 49, 50, 51, 52, 53,
	54, 55, 41, 42, 56, 57, 58, 43, 44, 45,
	46, 0, 0, 19, 35, 88, 99, 36, 92, 0,
	0, 47, 0, 0, 93, 95, 96, 94, 59, 100,
	0, 48, 60, 0, 62, 0, 0, 0, 0, 61,
	67, 67, 74, 78, 0, 0, 0, 0, 63, 68,
	0, 0, 0, 0, 0, 64, 65, 75, 0, 0,
	66, 79, 0, 0, 69, 70, 71, 72, 82, 76,
	77, 80, 81, 0, 73, 83, 0, 0, 0, 84,
	85, 86,
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
	3, 3, 3, 78, 3, 79,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73, 74, 75, 76, 77,
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
		//line parser.go.y:51
		{
			yyVAL.statements = nil
			if l, ok := yylex.(*Lexer); ok {
				l.statements = yyVAL.statements
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:58
		{
			yyVAL.statements = append([]Statement{yyDollar[1].statement}, yyDollar[2].statements...)
			if l, ok := yylex.(*Lexer); ok {
				l.statements = yyVAL.statements
			}
		}
	case 3:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:66
		{
			yyVAL.statement = &GlobalDefsStmt{stmts: yyDollar[3].statements}
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:70
		{
			yyVAL.statement = &VrrpSyncGroupStmt{group: yyDollar[2].tok.lit, stmts: yyDollar[4].statements}
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:74
		{
			yyVAL.statement = &VrrpInstanceStmt{inside_network: yyDollar[2].tok.lit, stmts: yyDollar[4].statements}
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:78
		{
			yyVAL.statement = &VirtualServerGroupStmt{group: yyDollar[2].tok.lit, vips: yyDollar[4].exprs}
		}
	case 7:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:82
		{
			yyVAL.statement = &VirtualServerStmt{group: yyDollar[3].tok.lit, stmts: yyDollar[5].statements}
		}
	case 8:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:86
		{
			yyVAL.statement = &VirtualServerStmt{ip: &IdentExpr{lit: yyDollar[2].tok.lit}, port: &NumExpr{lit: yyDollar[3].tok.lit}, stmts: yyDollar[5].statements}
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:91
		{
			yyVAL.statements = []Statement{}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:92
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:95
		{
			yyVAL.statement = &GroupStmt{exprs: yyDollar[3].exprs}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:96
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:97
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:98
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:99
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:100
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:103
		{
			yyVAL.statements = []Statement{}
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:104
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:107
		{
			yyVAL.statement = &GroupStmt{exprs: yyDollar[3].exprs}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:108
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:109
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:110
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:111
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:112
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:115
		{
			yyVAL.statements = []Statement{}
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:116
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:120
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:124
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:128
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:132
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:136
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:140
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:144
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:148
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:152
		{
			yyVAL.statement = &VirtualIpaddressStmt{addresses: yyDollar[3].exprs}
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:156
		{
			yyVAL.statement = &AuthenticationStmt{stmts: yyDollar[3].statements}
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:161
		{
			yyVAL.statements = []Statement{}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:162
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:166
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:170
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:174
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:178
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:182
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:186
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:190
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:194
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:198
		{
			yyVAL.statement = &SorryServerStmt{ip: &IdentExpr{lit: yyDollar[2].tok.lit}, port: &NumExpr{lit: yyDollar[3].tok.lit}}
		}
	case 48:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:202
		{
			yyVAL.statement = &RealServerStmt{ip: &IdentExpr{lit: yyDollar[2].tok.lit}, port: &NumExpr{lit: yyDollar[3].tok.lit}, stmts: yyDollar[5].statements}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:207
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:208
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:209
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:210
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:211
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:212
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:213
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:216
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:217
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:218
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:221
		{
			yyVAL.statements = []Statement{}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:222
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:226
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:230
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 63:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:234
		{
			yyVAL.statement = &HttpGetStmt{stmts: yyDollar[3].statements}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:238
		{
			yyVAL.statement = &SslGetStmt{stmts: yyDollar[3].statements}
		}
	case 65:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:242
		{
			yyVAL.statement = &TcpCheckStmt{stmts: yyDollar[3].statements}
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:246
		{
			yyVAL.statement = &MiscCheckStmt{stmts: yyDollar[3].statements}
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:251
		{
			yyVAL.statements = []Statement{}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:252
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:256
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:260
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:264
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:268
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:272
		{
			yyVAL.statement = &UrlStmt{stmts: yyDollar[3].statements}
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:277
		{
			yyVAL.statements = []Statement{}
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:278
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:282
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:286
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 78:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:291
		{
			yyVAL.statements = []Statement{}
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:292
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:296
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &StringExpr{lit: yyDollar[2].tok.lit}}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:300
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 82:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:305
		{
			yyVAL.statements = []Statement{}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:306
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:309
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:310
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:311
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 87:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:314
		{
			yyVAL.exprs = []Expression{}
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:315
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 89:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:318
		{
			yyVAL.exprs = []Expression{}
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:319
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:322
		{
			yyVAL.statements = []Statement{}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:323
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:326
		{
			yyVAL.statement = &AuthTypeStmt{expr: yyDollar[2].expr}
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:327
		{
			yyVAL.statement = &AuthPassStmt{expr: yyDollar[2].expr}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:330
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:331
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:334
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:335
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:338
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit}
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:339
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit, dev: yyDollar[3].tok.lit}
		}
	case 101:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:342
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit, port: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 102:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:345
		{
			yyVAL.exprs = []Expression{}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:346
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:349
		{
			yyVAL.expr = &NumExpr{lit: yyDollar[1].tok.lit}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:350
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:351
		{
			yyVAL.expr = &StringExpr{lit: yyDollar[1].tok.lit}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:352
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit}
		}
	}
	goto yystack /* stack new state and value */
}
