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
const VRRP_SYNC_GROUP = 57350
const GROUP = 57351
const NOTIFY_MASTER = 57352
const NOTIFY_BACKUP = 57353
const NOTIFY_FAULT = 57354
const NOTIFY = 57355
const SMTP_ALERT = 57356
const VRRP_INSTANCE = 57357
const STATE = 57358
const MASTER = 57359
const BACKUP = 57360
const INTERFACE = 57361
const GARP_MASTER_DELAY = 57362
const VIRTUAL_ROUTER_ID = 57363
const PRIORITY = 57364
const ADVERT_INT = 57365
const AUTHENTICATION = 57366
const AUTH_TYPE = 57367
const AUTH_PASS = 57368
const PASS = 57369
const AH = 57370
const VIRTUAL_IPADDRESS = 57371
const DEV = 57372
const VIRTUAL_SERVER_GROUP = 57373
const VIRTUAL_SERVER = 57374
const DELAY_LOOP = 57375
const SORRY_SERVER = 57376
const LVS_SCHED = 57377
const LB_ALGO = 57378
const RR = 57379
const WRR = 57380
const LC = 57381
const WLC = 57382
const LBLC = 57383
const SH = 57384
const DH = 57385
const LVS_METHOD = 57386
const LB_KIND = 57387
const NAT = 57388
const DR = 57389
const TUN = 57390
const PROTOCOL = 57391
const TCP = 57392
const REAL_SERVER = 57393
const WEIGHT = 57394
const INHIBIT_ON_FAILURE = 57395
const HTTP_GET = 57396
const SSL_GET = 57397
const TCP_CHECK = 57398
const MISC_CHECK = 57399
const URL = 57400
const PATH = 57401
const STATUS_CODE = 57402
const CONNECT_PORT = 57403
const CONNECT_TIMEOUT = 57404
const NB_GET_RETRY = 57405
const DELAY_BEFORE_RETRY = 57406
const MISC_PATH = 57407
const MISC_TIMEOUT = 57408

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"NUMBER",
	"IP",
	"STRING",
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

//line parser.go.y:317

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 94
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 170

var yyAct = [...]int{

	126, 44, 33, 86, 30, 150, 77, 31, 32, 34,
	35, 36, 38, 65, 71, 66, 67, 37, 115, 116,
	117, 118, 119, 120, 68, 69, 49, 50, 51, 70,
	98, 72, 125, 41, 113, 22, 23, 24, 25, 26,
	27, 136, 101, 102, 132, 133, 134, 135, 63, 136,
	91, 137, 132, 133, 134, 135, 28, 158, 159, 130,
	124, 46, 45, 48, 47, 105, 156, 144, 145, 11,
	142, 140, 141, 90, 85, 95, 123, 3, 138, 122,
	110, 61, 60, 43, 4, 99, 19, 14, 13, 12,
	107, 108, 96, 153, 20, 39, 87, 88, 89, 93,
	5, 6, 53, 54, 109, 78, 79, 80, 81, 82,
	83, 84, 46, 45, 48, 47, 92, 161, 154, 152,
	151, 149, 148, 147, 127, 94, 146, 121, 104, 103,
	76, 62, 59, 58, 57, 56, 160, 111, 55, 15,
	10, 9, 8, 1, 40, 97, 7, 106, 52, 18,
	74, 73, 143, 139, 157, 131, 114, 64, 100, 29,
	21, 2, 129, 128, 155, 112, 42, 75, 17, 16,
}
var yyPact = [...]int{

	69, -1000, 69, 138, 137, 136, 60, -1000, 22, 21,
	20, 135, -1000, -1000, -1000, 19, 26, -12, 27, -1000,
	-1000, -1000, 16, 108, 108, 108, 108, -1000, -1000, -1000,
	85, 134, 130, -1000, 129, 128, 127, 15, 14, -1000,
	-1000, 126, -20, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 125, 68, 68, 50, 50,
	0, 110, 93, 57, 24, 17, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 124, 123, -1000, -1000, -1000, -1000, 35, -1000,
	-1000, 63, 108, -1000, 13, 133, -1000, -1000, -1000, -1000,
	-1000, -1000, -34, -1000, -1000, 122, -1000, 12, 9, -7,
	-35, -1000, -1000, -1000, -1000, -1000, -9, -17, 10, 2,
	-1000, -1000, 121, 118, 117, 116, -62, -1000, -1000, -1000,
	115, 114, -1000, -1000, 86, 113, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -2, -1000, -1000, 132, 112,
	-1000, -1000,
}
var yyPgo = [...]int{

	0, 143, 169, 168, 167, 166, 165, 0, 164, 163,
	162, 161, 160, 159, 158, 157, 156, 155, 154, 153,
	152, 151, 150, 149, 1, 148, 147, 145, 144, 6,
	3,
}
var yyR1 = [...]int{

	0, 1, 1, 11, 11, 11, 11, 2, 2, 12,
	12, 12, 12, 12, 12, 3, 3, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 5, 5, 15, 15,
	15, 15, 15, 15, 15, 15, 29, 29, 29, 29,
	29, 29, 29, 30, 30, 30, 6, 6, 16, 16,
	16, 16, 16, 16, 7, 7, 17, 17, 17, 17,
	17, 9, 9, 19, 19, 10, 10, 20, 20, 8,
	8, 18, 18, 22, 22, 23, 23, 4, 4, 14,
	14, 26, 26, 25, 25, 27, 27, 28, 21, 21,
	24, 24, 24, 24,
}
var yyR2 = [...]int{

	0, 0, 2, 5, 5, 5, 6, 0, 2, 4,
	2, 2, 2, 2, 1, 0, 2, 2, 2, 2,
	1, 2, 2, 2, 4, 4, 0, 2, 2, 2,
	2, 2, 2, 2, 3, 6, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 0, 2, 2, 1,
	4, 4, 4, 4, 0, 2, 2, 2, 2, 2,
	4, 0, 2, 2, 2, 0, 2, 2, 2, 0,
	2, 2, 2, 0, 2, 0, 2, 0, 2, 2,
	2, 1, 1, 1, 1, 1, 3, 2, 0, 2,
	1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -1, -11, 8, 15, 31, 32, -1, 4, 4,
	4, 9, 67, 67, 67, 4, -2, -3, -23, 67,
	68, -12, 9, 10, 11, 12, 13, 14, 68, -13,
	16, 19, 20, 14, 21, 22, 23, 29, 24, 68,
	-28, 6, -5, 67, -24, 5, 4, 7, 6, -24,
	-24, -24, -25, 17, 18, 4, 5, 5, 5, 5,
	67, 67, 5, 68, -15, 33, 35, 36, 44, 45,
	49, 34, 51, -21, -22, -4, 5, -29, 37, 38,
	39, 40, 41, 42, 43, -29, -30, 46, 47, 48,
	-30, 50, 6, 6, 68, -24, 68, -27, 6, 68,
	-14, 25, 26, 5, 5, 30, -26, 27, 28, -24,
	67, 4, -6, 68, -16, 52, 53, 54, 55, 56,
	57, 5, 67, 67, 67, 67, -7, -7, -9, -10,
	68, -17, 61, 62, 63, 64, 58, 68, 68, -19,
	61, 62, 68, -20, 65, 66, 5, 5, 5, 5,
	67, 5, 5, 7, 5, -8, 68, -18, 59, 60,
	4, 5,
}
var yyDef = [...]int{

	1, -2, 1, 0, 0, 0, 0, 2, 0, 0,
	0, 0, 7, 15, 75, 0, 0, 0, 0, 26,
	3, 8, 0, 0, 0, 0, 0, 14, 4, 16,
	0, 0, 0, 20, 0, 0, 0, 0, 0, 5,
	76, 0, 0, 88, 10, 90, 91, 92, 93, 11,
	12, 13, 17, 83, 84, 18, 19, 21, 22, 23,
	73, 77, 87, 6, 27, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 28, 29, 36, 37,
	38, 39, 40, 41, 42, 30, 31, 43, 44, 45,
	32, 33, 0, 0, 9, 89, 24, 74, 85, 25,
	78, 0, 0, 34, 0, 0, 79, 81, 82, 80,
	46, 86, 0, 35, 47, 0, 49, 0, 0, 0,
	0, 48, 54, 54, 61, 65, 0, 0, 0, 0,
	50, 55, 0, 0, 0, 0, 0, 51, 52, 62,
	0, 0, 53, 66, 0, 0, 56, 57, 58, 59,
	69, 63, 64, 67, 68, 0, 60, 70, 0, 0,
	71, 72,
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
	3, 3, 3, 67, 3, 68,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66,
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
		//line parser.go.y:47
		{
			yyVAL.statements = nil
			if l, ok := yylex.(*Lexer); ok {
				l.statements = yyVAL.statements
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:54
		{
			yyVAL.statements = append([]Statement{yyDollar[1].statement}, yyDollar[2].statements...)
			if l, ok := yylex.(*Lexer); ok {
				l.statements = yyVAL.statements
			}
		}
	case 3:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:62
		{
			yyVAL.statement = &VrrpSyncGroupStmt{group: yyDollar[2].tok.lit, stmts: yyDollar[4].statements}
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:66
		{
			yyVAL.statement = &VrrpInstanceStmt{inside_network: yyDollar[2].tok.lit, stmts: yyDollar[4].statements}
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:70
		{
			yyVAL.statement = &VirtualServerGroupStmt{group: yyDollar[2].tok.lit, vips: yyDollar[4].exprs}
		}
	case 6:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:74
		{
			yyVAL.statement = &VirtualServerStmt{group: yyDollar[3].tok.lit, stmts: yyDollar[5].statements}
		}
	case 7:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:79
		{
			yyVAL.statements = []Statement{}
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:80
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:83
		{
			yyVAL.statement = &GroupStmt{exprs: yyDollar[3].exprs}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:84
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:85
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:86
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:87
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:88
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:91
		{
			yyVAL.statements = []Statement{}
		}
	case 16:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:92
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:96
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:100
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:104
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:108
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:112
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:116
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:120
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:124
		{
			yyVAL.statement = &VirtualIpaddressStmt{addresses: yyDollar[3].exprs}
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:128
		{
			yyVAL.statement = &AuthenticationStmt{stmts: yyDollar[3].statements}
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:133
		{
			yyVAL.statements = []Statement{}
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:134
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:138
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:142
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:146
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:150
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:154
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:162
		{
			yyVAL.statement = &SorryServerStmt{ip: &IdentExpr{lit: yyDollar[2].tok.lit}, port: &NumExpr{lit: yyDollar[3].tok.lit}}
		}
	case 35:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:166
		{
			yyVAL.statement = &RealServerStmt{ip: &IdentExpr{lit: yyDollar[2].tok.lit}, port: &NumExpr{lit: yyDollar[3].tok.lit}, stmts: yyDollar[5].statements}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:171
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:172
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:174
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:175
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:176
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:177
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:180
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:181
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:182
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:185
		{
			yyVAL.statements = []Statement{}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:186
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:190
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:194
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:198
		{
			yyVAL.statement = &HttpGetStmt{stmts: yyDollar[3].statements}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:202
		{
			yyVAL.statement = &SslGetStmt{stmts: yyDollar[3].statements}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:206
		{
			yyVAL.statement = &TcpCheckStmt{stmts: yyDollar[3].statements}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:210
		{
			yyVAL.statement = &MiscCheckStmt{stmts: yyDollar[3].statements}
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:215
		{
			yyVAL.statements = []Statement{}
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:216
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:220
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:224
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:228
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:232
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:236
		{
			yyVAL.statement = &UrlStmt{stmts: yyDollar[3].statements}
		}
	case 61:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:241
		{
			yyVAL.statements = []Statement{}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:242
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:246
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:250
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:255
		{
			yyVAL.statements = []Statement{}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:256
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:260
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &StringExpr{lit: yyDollar[2].tok.lit}}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:264
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 69:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:269
		{
			yyVAL.statements = []Statement{}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:270
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:273
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 73:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:277
		{
			yyVAL.exprs = []Expression{}
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:278
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 75:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:281
		{
			yyVAL.exprs = []Expression{}
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:282
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 77:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:285
		{
			yyVAL.statements = []Statement{}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:286
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:289
		{
			yyVAL.statement = &AuthTypeStmt{expr: yyDollar[2].expr}
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:290
		{
			yyVAL.statement = &AuthPassStmt{expr: yyDollar[2].expr}
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:293
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:294
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:297
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:298
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:301
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit}
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:302
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit, dev: yyDollar[3].tok.lit}
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:305
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit, port: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 88:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:308
		{
			yyVAL.exprs = []Expression{}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:309
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:312
		{
			yyVAL.expr = &NumExpr{lit: yyDollar[1].tok.lit}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:313
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:314
		{
			yyVAL.expr = &StringExpr{lit: yyDollar[1].tok.lit}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:315
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit}
		}
	}
	goto yystack /* stack new state and value */
}
