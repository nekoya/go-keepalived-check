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
const VRRP_INSTANCE = 57350
const STATE = 57351
const MASTER = 57352
const BACKUP = 57353
const INTERFACE = 57354
const GARP_MASTER_DELAY = 57355
const SMTP_ALERT = 57356
const VIRTUAL_ROUTER_ID = 57357
const PRIORITY = 57358
const ADVERT_INT = 57359
const AUTHENTICATION = 57360
const AUTH_TYPE = 57361
const AUTH_PASS = 57362
const PASS = 57363
const AH = 57364
const VIRTUAL_IPADDRESS = 57365
const DEV = 57366
const VIRTUAL_SERVER_GROUP = 57367
const VIRTUAL_SERVER = 57368
const GROUP = 57369
const DELAY_LOOP = 57370
const LVS_SCHED = 57371
const LB_ALGO = 57372
const RR = 57373
const WRR = 57374
const LC = 57375
const WLC = 57376
const LBLC = 57377
const SH = 57378
const DH = 57379
const LVS_METHOD = 57380
const LB_KIND = 57381
const NAT = 57382
const DR = 57383
const TUN = 57384
const PROTOCOL = 57385
const TCP = 57386
const REAL_SERVER = 57387
const WEIGHT = 57388
const INHIBIT_ON_FAILURE = 57389
const HTTP_GET = 57390
const SSL_GET = 57391
const TCP_CHECK = 57392
const MISC_CHECK = 57393
const URL = 57394
const PATH = 57395
const STATUS_CODE = 57396
const CONNECT_PORT = 57397
const CONNECT_TIMEOUT = 57398
const NB_GET_RETRY = 57399
const DELAY_BEFORE_RETRY = 57400
const MISC_PATH = 57401
const MISC_TIMEOUT = 57402

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"NUMBER",
	"IP",
	"STRING",
	"VRRP_INSTANCE",
	"STATE",
	"MASTER",
	"BACKUP",
	"INTERFACE",
	"GARP_MASTER_DELAY",
	"SMTP_ALERT",
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
	"GROUP",
	"DELAY_LOOP",
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

//line parser.go.y:289

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 81
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 142

var yyAct = [...]int{

	102, 18, 72, 29, 19, 20, 21, 22, 23, 24,
	26, 44, 45, 46, 112, 25, 126, 108, 109, 110,
	111, 47, 48, 63, 113, 101, 49, 100, 50, 99,
	91, 92, 93, 94, 95, 96, 98, 112, 75, 76,
	108, 109, 110, 111, 54, 42, 89, 106, 134, 135,
	120, 121, 68, 118, 16, 116, 117, 132, 70, 27,
	86, 9, 114, 40, 39, 15, 11, 10, 64, 65,
	66, 3, 67, 55, 56, 57, 58, 59, 60, 61,
	78, 73, 80, 81, 32, 33, 129, 69, 4, 5,
	137, 62, 84, 83, 130, 85, 128, 127, 125, 124,
	103, 123, 122, 97, 77, 53, 41, 38, 37, 36,
	35, 136, 87, 34, 12, 8, 7, 1, 28, 71,
	6, 79, 31, 82, 14, 51, 119, 115, 133, 107,
	90, 43, 74, 17, 2, 105, 104, 131, 88, 30,
	52, 13,
}
var yyPact = [...]int{

	63, -1000, 63, 112, 111, 34, -1000, 6, 5, 110,
	-1000, -1000, 4, -8, -3, -1000, -1000, -1000, 74, 109,
	105, -1000, 104, 103, 102, 3, 2, -1000, -1000, 101,
	-17, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 100, 42, 42, 28, 28, 8,
	81, -4, 19, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 99,
	-1000, -1000, 56, -1000, -1000, 61, 88, -1, 108, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -16, -1000,
	-1000, 98, -1000, -25, -32, -34, -36, -1000, -1000, -1000,
	-1000, -1000, -15, -38, 0, -9, -1000, -1000, 97, 96,
	94, 93, -45, -1000, -1000, -1000, 92, 91, -1000, -1000,
	79, 89, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -5, -1000, -1000, 107, 85, -1000, -1000,
}
var yyPgo = [...]int{

	0, 117, 141, 140, 139, 138, 0, 137, 136, 135,
	134, 133, 132, 131, 130, 129, 128, 127, 126, 125,
	124, 123, 122, 121, 119, 118, 44, 23,
}
var yyR1 = [...]int{

	0, 1, 1, 10, 10, 10, 2, 2, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 4, 4, 13,
	13, 13, 13, 13, 13, 13, 26, 26, 26, 26,
	26, 26, 26, 27, 27, 27, 5, 5, 14, 14,
	14, 14, 14, 14, 6, 6, 15, 15, 15, 15,
	15, 8, 8, 17, 17, 9, 9, 18, 18, 7,
	7, 16, 16, 19, 19, 20, 20, 3, 3, 12,
	12, 23, 23, 22, 22, 24, 24, 25, 21, 21,
	21,
}
var yyR2 = [...]int{

	0, 0, 2, 5, 5, 6, 0, 2, 2, 2,
	2, 1, 2, 2, 2, 4, 4, 0, 2, 2,
	2, 2, 2, 2, 2, 6, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 0, 2, 2, 1,
	4, 4, 4, 4, 0, 2, 2, 2, 2, 2,
	4, 0, 2, 2, 2, 0, 2, 2, 2, 0,
	2, 2, 2, 0, 2, 0, 2, 0, 2, 2,
	2, 1, 1, 1, 1, 1, 3, 2, 1, 1,
	1,
}
var yyChk = [...]int{

	-1000, -1, -10, 8, 25, 26, -1, 4, 4, 27,
	61, 61, 4, -2, -20, 61, 62, -11, 9, 12,
	13, 14, 15, 16, 17, 23, 18, 62, -25, 6,
	-4, -22, 10, 11, 4, 5, 5, 5, 5, 61,
	61, 5, 62, -13, 28, 29, 30, 38, 39, 43,
	45, -19, -3, 5, -26, 31, 32, 33, 34, 35,
	36, 37, -26, -27, 40, 41, 42, -27, 44, 6,
	62, -24, 6, 62, -12, 19, 20, 5, 24, -23,
	21, 22, -21, 5, 4, 7, 61, 4, -5, 62,
	-14, 46, 47, 48, 49, 50, 51, 5, 61, 61,
	61, 61, -6, -6, -8, -9, 62, -15, 55, 56,
	57, 58, 52, 62, 62, -17, 55, 56, 62, -18,
	59, 60, 5, 5, 5, 5, 61, 5, 5, 7,
	5, -7, 62, -16, 53, 54, 4, 5,
}
var yyDef = [...]int{

	1, -2, 1, 0, 0, 0, 2, 0, 0, 0,
	6, 65, 0, 0, 0, 17, 3, 7, 0, 0,
	0, 11, 0, 0, 0, 0, 0, 4, 66, 0,
	0, 8, 73, 74, 9, 10, 12, 13, 14, 63,
	67, 77, 5, 18, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 19, 20, 26, 27, 28, 29, 30,
	31, 32, 21, 22, 33, 34, 35, 23, 24, 0,
	15, 64, 75, 16, 68, 0, 0, 0, 0, 69,
	71, 72, 70, 78, 79, 80, 36, 76, 0, 25,
	37, 0, 39, 0, 0, 0, 0, 38, 44, 44,
	51, 55, 0, 0, 0, 0, 40, 45, 0, 0,
	0, 0, 0, 41, 42, 52, 0, 0, 43, 56,
	0, 0, 46, 47, 48, 49, 59, 53, 54, 57,
	58, 0, 50, 60, 0, 0, 61, 62,
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
	3, 3, 3, 61, 3, 62,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60,
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
		//line parser.go.y:44
		{
			yyVAL.statements = nil
			if l, ok := yylex.(*Lexer); ok {
				l.statements = yyVAL.statements
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:51
		{
			yyVAL.statements = append([]Statement{yyDollar[1].statement}, yyDollar[2].statements...)
			if l, ok := yylex.(*Lexer); ok {
				l.statements = yyVAL.statements
			}
		}
	case 3:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:59
		{
			yyVAL.statement = &VrrpInstanceStmt{inside_network: yyDollar[2].tok.lit, stmts: yyDollar[4].statements}
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:63
		{
			yyVAL.statement = &VirtualServerGroupStmt{group: yyDollar[2].tok.lit, vips: yyDollar[4].exprs}
		}
	case 5:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:67
		{
			yyVAL.statement = &VirtualServerStmt{group: yyDollar[3].tok.lit, stmts: yyDollar[5].statements}
		}
	case 6:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:72
		{
			yyVAL.statements = []Statement{}
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:73
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 8:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:77
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:81
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:85
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:89
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:93
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:97
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:101
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:105
		{
			yyVAL.statement = &VirtualIpaddressStmt{addresses: yyDollar[3].exprs}
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:109
		{
			yyVAL.statement = &AuthenticationStmt{stmts: yyDollar[3].statements}
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:114
		{
			yyVAL.statements = []Statement{}
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:115
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:119
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:123
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:127
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:131
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:135
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: yyDollar[2].expr}
		}
	case 24:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:139
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:143
		{
			yyVAL.statement = &RealServerStmt{ip: &IdentExpr{lit: yyDollar[2].tok.lit}, port: &NumExpr{lit: yyDollar[3].tok.lit}, stmts: yyDollar[5].statements}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:148
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:149
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:150
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:151
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:152
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:153
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:154
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:157
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:158
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:159
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:162
		{
			yyVAL.statements = []Statement{}
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:163
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:167
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:171
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:175
		{
			yyVAL.statement = &HttpGetStmt{stmts: yyDollar[3].statements}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:179
		{
			yyVAL.statement = &SslGetStmt{stmts: yyDollar[3].statements}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:183
		{
			yyVAL.statement = &TcpCheckStmt{stmts: yyDollar[3].statements}
		}
	case 43:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:187
		{
			yyVAL.statement = &MiscCheckStmt{stmts: yyDollar[3].statements}
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:192
		{
			yyVAL.statements = []Statement{}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:197
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:201
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:205
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:209
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:213
		{
			yyVAL.statement = &UrlStmt{stmts: yyDollar[3].statements}
		}
	case 51:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:218
		{
			yyVAL.statements = []Statement{}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:219
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:223
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:227
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:232
		{
			yyVAL.statements = []Statement{}
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:233
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:237
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &StringExpr{lit: yyDollar[2].tok.lit}}
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:241
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 59:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:246
		{
			yyVAL.statements = []Statement{}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:247
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:250
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{lit: yyDollar[1].tok.lit}, value: &IdentExpr{lit: yyDollar[2].tok.lit}}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:251
		{
			yyVAL.statement = &ExprStmt{key: &IdentExpr{yyDollar[1].tok.lit}, value: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:254
		{
			yyVAL.exprs = []Expression{}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:255
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:258
		{
			yyVAL.exprs = []Expression{}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:259
		{
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[2].expr)
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.go.y:262
		{
			yyVAL.statements = []Statement{}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:263
		{
			yyVAL.statements = append(yyDollar[1].statements, yyDollar[2].statement)
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:266
		{
			yyVAL.statement = &AuthTypeStmt{expr: yyDollar[2].expr}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:267
		{
			yyVAL.statement = &AuthPassStmt{expr: yyDollar[2].expr}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:270
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:271
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:274
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:275
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:278
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:279
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit, dev: yyDollar[3].tok.lit}
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:282
		{
			yyVAL.expr = &IpExpr{lit: yyDollar[1].tok.lit, port: &NumExpr{lit: yyDollar[2].tok.lit}}
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:285
		{
			yyVAL.expr = &NumExpr{lit: yyDollar[1].tok.lit}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:286
		{
			yyVAL.expr = &IdentExpr{lit: yyDollar[1].tok.lit}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:287
		{
			yyVAL.expr = &StringExpr{lit: yyDollar[1].tok.lit}
		}
	}
	goto yystack /* stack new state and value */
}
