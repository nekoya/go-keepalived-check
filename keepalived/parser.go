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
const AUTHENTICATION = 57373
const AUTH_TYPE = 57374
const AUTH_PASS = 57375
const PASS = 57376
const AH = 57377
const VIRTUAL_IPADDRESS = 57378
const DEV = 57379
const VIRTUAL_SERVER_GROUP = 57380
const VIRTUAL_SERVER = 57381
const DELAY_LOOP = 57382
const SORRY_SERVER = 57383
const LVS_SCHED = 57384
const LB_ALGO = 57385
const RR = 57386
const WRR = 57387
const LC = 57388
const WLC = 57389
const LBLC = 57390
const SH = 57391
const DH = 57392
const LVS_METHOD = 57393
const LB_KIND = 57394
const NAT = 57395
const DR = 57396
const TUN = 57397
const PROTOCOL = 57398
const TCP = 57399
const REAL_SERVER = 57400
const WEIGHT = 57401
const INHIBIT_ON_FAILURE = 57402
const HTTP_GET = 57403
const SSL_GET = 57404
const TCP_CHECK = 57405
const MISC_CHECK = 57406
const URL = 57407
const PATH = 57408
const STATUS_CODE = 57409
const CONNECT_PORT = 57410
const CONNECT_TIMEOUT = 57411
const NB_GET_RETRY = 57412
const DELAY_BEFORE_RETRY = 57413
const MISC_PATH = 57414
const MISC_TIMEOUT = 57415

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

//line parser.go.y:336

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 103
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 192

var yyAct = [...]int{

	144, 91, 105, 96, 53, 168, 50, 143, 142, 51,
	52, 54, 55, 56, 58, 34, 33, 36, 35, 57,
	82, 88, 83, 84, 32, 37, 141, 39, 176, 177,
	140, 85, 86, 116, 61, 128, 87, 174, 89, 42,
	43, 44, 45, 46, 47, 65, 66, 67, 68, 133,
	134, 135, 136, 137, 138, 80, 119, 120, 48, 21,
	22, 23, 24, 25, 26, 131, 154, 78, 77, 150,
	151, 152, 153, 64, 154, 63, 155, 150, 151, 152,
	153, 31, 162, 163, 148, 160, 113, 30, 104, 109,
	34, 33, 36, 35, 158, 159, 17, 110, 40, 117,
	123, 156, 114, 59, 16, 15, 9, 106, 107, 108,
	97, 98, 99, 100, 101, 102, 103, 3, 125, 126,
	70, 71, 127, 13, 4, 19, 34, 33, 36, 35,
	171, 5, 112, 111, 179, 172, 170, 169, 167, 166,
	92, 165, 145, 164, 139, 122, 121, 6, 7, 95,
	79, 76, 75, 74, 73, 38, 178, 129, 72, 18,
	12, 90, 11, 10, 1, 60, 115, 8, 124, 69,
	29, 93, 161, 157, 175, 149, 132, 81, 118, 49,
	41, 20, 2, 147, 146, 173, 130, 62, 94, 28,
	27, 14,
}
var yyPact = [...]int{

	109, -1000, 109, 32, 159, 158, 156, 107, -1000, -1000,
	31, 30, 22, 155, 50, -1000, -1000, -1000, 13, -1000,
	-1000, 7, 122, 122, 150, 122, -1000, 23, -17, 28,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1, 122, 122, 122, 122, -1000, -1000, -1000,
	96, 154, 149, -1000, 148, 147, 146, -6, -7, -1000,
	-1000, 145, -20, 86, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 144, 66, 66, 54, 54, 40, 127, 126,
	-1000, -1000, 11, 27, 24, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 141, 140, -1000, -1000, -1000, 63, -1000, -1000, 84,
	122, -1000, -39, 153, -1000, -1000, -1000, -1000, -1000, -1000,
	-10, -1000, -1000, 139, -1000, -44, -48, -66, -67, -1000,
	-1000, -1000, -1000, -1000, 9, 1, 26, 10, -1000, -1000,
	138, 136, 134, 133, -69, -1000, -1000, -1000, 132, 131,
	-1000, -1000, 123, 130, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -38, -1000, -1000, 152, 129, -1000, -1000,
}
var yyPgo = [...]int{

	0, 164, 191, 190, 189, 188, 187, 186, 0, 185,
	184, 183, 182, 181, 180, 179, 178, 177, 176, 175,
	174, 173, 172, 75, 171, 170, 1, 169, 168, 166,
	165, 3, 2,
}
var yyR1 = [...]int{

	0, 1, 1, 12, 12, 12, 12, 12, 2, 2,
	13, 13, 13, 13, 13, 13, 3, 3, 14, 14,
	14, 14, 14, 14, 4, 4, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 6, 6, 17, 17, 17,
	17, 17, 17, 17, 17, 31, 31, 31, 31, 31,
	31, 31, 32, 32, 32, 7, 7, 18, 18, 18,
	18, 18, 18, 8, 8, 19, 19, 19, 19, 19,
	10, 10, 21, 21, 11, 11, 22, 22, 9, 9,
	20, 20, 24, 24, 25, 25, 5, 5, 16, 16,
	28, 28, 27, 27, 29, 29, 30, 23, 23, 26,
	26, 26, 26,
}
var yyR2 = [...]int{

	0, 0, 2, 4, 5, 5, 5, 6, 0, 2,
	4, 2, 2, 2, 2, 1, 0, 2, 4, 2,
	2, 2, 2, 1, 0, 2, 2, 2, 2, 1,
	2, 2, 2, 4, 4, 0, 2, 2, 2, 2,
	2, 2, 2, 3, 6, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 0, 2, 2, 1, 4,
	4, 4, 4, 0, 2, 2, 2, 2, 2, 4,
	0, 2, 2, 2, 0, 2, 2, 2, 0, 2,
	2, 2, 0, 2, 0, 2, 0, 2, 2, 2,
	1, 1, 1, 1, 1, 3, 2, 0, 2, 1,
	1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -12, 8, 15, 22, 38, 39, -1, 74,
	4, 4, 4, 16, -2, 74, 74, 74, 4, 75,
	-13, 9, 10, 11, 12, 13, 14, -3, -4, -25,
	74, 74, -26, 5, 4, 7, 6, -26, 5, -26,
	75, -14, 16, 17, 18, 19, 20, 21, 75, -15,
	23, 26, 27, 21, 28, 29, 30, 36, 31, 75,
	-30, 6, -6, -23, 74, -26, -26, -26, -26, -27,
	24, 25, 4, 5, 5, 5, 5, 74, 74, 5,
	75, -17, 40, 42, 43, 51, 52, 56, 41, 58,
	75, -26, -23, -24, -5, 5, -31, 44, 45, 46,
	47, 48, 49, 50, -31, -32, 53, 54, 55, -32,
	57, 6, 6, 75, 75, -29, 6, 75, -16, 32,
	33, 5, 5, 37, -28, 34, 35, -26, 74, 4,
	-7, 75, -18, 59, 60, 61, 62, 63, 64, 5,
	74, 74, 74, 74, -8, -8, -10, -11, 75, -19,
	68, 69, 70, 71, 65, 75, 75, -21, 68, 69,
	75, -22, 72, 73, 5, 5, 5, 5, 74, 5,
	5, 7, 5, -9, 75, -20, 66, 67, 4, 5,
}
var yyDef = [...]int{

	1, -2, 1, 0, 0, 0, 0, 0, 2, 8,
	0, 0, 0, 0, 0, 16, 24, 84, 0, 3,
	9, 0, 0, 0, 0, 0, 15, 0, 0, 0,
	35, 97, 11, 99, 100, 101, 102, 12, 13, 14,
	4, 17, 0, 0, 0, 0, 0, 23, 5, 25,
	0, 0, 0, 29, 0, 0, 0, 0, 0, 6,
	85, 0, 0, 0, 97, 19, 20, 21, 22, 26,
	92, 93, 27, 28, 30, 31, 32, 82, 86, 96,
	7, 36, 0, 0, 0, 0, 0, 0, 0, 0,
	10, 98, 0, 0, 0, 37, 38, 45, 46, 47,
	48, 49, 50, 51, 39, 40, 52, 53, 54, 41,
	42, 0, 0, 18, 33, 83, 94, 34, 87, 0,
	0, 43, 0, 0, 88, 90, 91, 89, 55, 95,
	0, 44, 56, 0, 58, 0, 0, 0, 0, 57,
	63, 63, 70, 74, 0, 0, 0, 0, 59, 64,
	0, 0, 0, 0, 0, 60, 61, 71, 0, 0,
	62, 75, 0, 0, 65, 66, 67, 68, 78, 72,
	73, 76, 77, 0, 69, 79, 0, 0, 80, 81,
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
	3, 3, 3, 74, 3, 75,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
	72, 73,
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
		//line parser.go.y:50
		{
			yyVAL.statements = nil
			if l, ok := yylex.(*Lexer); ok {
				l.statements = yyVAL.statements
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:57
		{
			yyVAL.statements = append([]Statement{yyDollar[1].statement}, yyDollar[2].statements...)
			if l, ok := yylex.(*Lexer); ok {
				l.statements = yyVAL.statements
			}
		}
	case 3:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:65
		{
			yyVAL.statement = &GlobalDefsStmt{stmts: yyDollar[3].statements}
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:69
		{
			yyVAL.statement = &VrrpSyncGroupStmt{group: yyDollar[2].tok.lit, stmts: yyDollar[4].statements}
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:73
		{
			yyVAL.statement = &VrrpInstanceStmt{inside_network: yyDollar[2].tok.lit, stmts: yyDollar[4].statements}
		}
	case 6:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.statement = &VirtualServerGroupStmt{group: yyDollar[2].tok.lit, vips: yyDollar[4].exprs}
		}
	case 7:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:81
		{
			yyVAL.statement = &VirtualServerStmt{group: yyDollar[3].tok.lit, stmts: yyDollar[5].statements}
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:86
		{
			yyVAL.statements = []Statement{}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:87
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 10:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:90
		{
			yyVAL.statement = &GroupStmt{exprs: yyDollar[3].exprs}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:91
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:92
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:93
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:94
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:95
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:98
		{
			yyVAL.statements = []Statement{}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:99
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:102
		{
			yyVAL.statement = &GroupStmt{exprs: yyDollar[3].exprs}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:103
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:104
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:105
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:106
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:107
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:110
		{
			yyVAL.statements = []Statement{}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:111
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:115
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:119
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:123
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:127
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:135
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.statement = &VirtualIpaddressStmt{addresses: yyDollar[3].exprs}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:147
		{
			yyVAL.statement = &AuthenticationStmt{stmts: yyDollar[3].statements}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:152
		{
			yyVAL.statements = []Statement{}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:153
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:157
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:161
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:165
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:169
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:177
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:181
		{
			yyVAL.statement = &SorryServerStmt{ip: &IdentExpr{lit: yyDollar[2].tok.lit}, port: &NumExpr{lit: yyDollar[3].tok.lit}}
		}
	case 44:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:185
		{
			yyVAL.statement = &RealServerStmt{ip: &IdentExpr{lit: yyDollar[2].tok.lit}, port: &NumExpr{lit: yyDollar[3].tok.lit}, stmts: yyDollar[5].statements}
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:190
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:191
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:192
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:194
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:195
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:196
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:199
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:200
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:201
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:204
		{
			yyVAL.statements = []Statement{}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:205
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:209
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:213
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:217
		{
			yyVAL.statement = &HttpGetStmt{stmts: yyDollar[3].statements}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:221
		{
			yyVAL.statement = &SslGetStmt{stmts: yyDollar[3].statements}
		}
	case 61:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:225
		{
			yyVAL.statement = &TcpCheckStmt{stmts: yyDollar[3].statements}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:229
		{
			yyVAL.statement = &MiscCheckStmt{stmts: yyDollar[3].statements}
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:234
		{
			yyVAL.statements = []Statement{}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:235
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:239
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:243
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:247
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:251
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:255
		{
			yyVAL.statement = &UrlStmt{stmts: yyDollar[3].statements}
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:260
		{
			yyVAL.statements = []Statement{}
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:261
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:265
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:269
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.statements = []Statement{}
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:275
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:279
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &StringExpr{lit: yyDollar[2].tok.lit}}
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:283
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 78:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:288
		{
			yyVAL.statements = []Statement{}
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:289
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:292
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:293
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 82:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:296
		{
			yyVAL.exprs = []Expression{}
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:297
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 84:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:300
		{
			yyVAL.exprs = []Expression{}
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:301
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:304
		{
			yyVAL.statements = []Statement{}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:305
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:308
		{
			yyVAL.statement = &AuthTypeStmt{expr: yyDollar[2].expr}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:309
		{
			yyVAL.statement = &AuthPassStmt{expr: yyDollar[2].expr}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:312
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:313
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:316
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:317
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:320
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:321
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit, dev: yyDollar[3].tok.lit}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:324
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit, port: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:327
		{
			yyVAL.exprs = []Expression{}
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:328
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:331
		{
			yyVAL.expr = &NumExpr{lit: yyDollar[1].tok.lit}
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:332
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:333
		{
			yyVAL.expr = &StringExpr{lit: yyDollar[1].tok.lit}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:334
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit}
		}
	}
	goto yystack /* stack new state and value */
}
