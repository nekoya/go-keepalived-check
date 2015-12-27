package keepalived

import (
	"fmt"
	"strings"
)

type Block interface{}
type Statement interface {
	String() string
}
type Expression interface {
	expr() string
}

type ExprStmt struct {
	key   Expression
	value Expression
}

type GlobalDefsStmt struct {
	stmts []Statement
}

type VrrpSyncGroupStmt struct {
	group string
	stmts []Statement
}

type VrrpInstanceStmt struct {
	inside_network string
	stmts          []Statement
}

type VirtualIpaddressStmt struct {
	addresses []Expression
}

type VirtualServerGroupStmt struct {
	group string
	vips  []Expression
}

type VirtualServerStmt struct {
	group string
	ip    *IdentExpr
	port  *NumExpr
	stmts []Statement
}

type SorryServerStmt struct {
	ip   *IdentExpr
	port *NumExpr
}

type RealServerStmt struct {
	ip    *IdentExpr
	port  *NumExpr
	stmts []Statement
}

type HttpGetStmt struct {
	stmts []Statement
}

type SslGetStmt struct {
	stmts []Statement
}

type TcpCheckStmt struct {
	stmts []Statement
}

type MiscCheckStmt struct {
	stmts []Statement
}

type UrlStmt struct {
	stmts []Statement
}

type AuthenticationStmt struct {
	stmts []Statement
}

type AuthTypeStmt struct {
	expr Expression
}

type AuthPassStmt struct {
	expr Expression
}

type GroupStmt struct {
	exprs []Expression
}

func (x *GlobalDefsStmt) String() string {
	var ret []string
	ret = append(ret, "global_defs")
	for _, v := range x.stmts {
		ret = append(ret, "- "+v.String())
	}
	return strings.Join(ret[:], "\n")
}

func (x *VrrpSyncGroupStmt) String() string {
	var ret []string
	ret = append(ret, "vrrp_sync_group "+x.group)
	for _, v := range x.stmts {
		ret = append(ret, "- "+v.String())
	}
	return strings.Join(ret[:], "\n")
}

func (x *VrrpInstanceStmt) String() string {
	var ret []string
	ret = append(ret, "vrrp_instance "+x.inside_network)
	for _, v := range x.stmts {
		ret = append(ret, "- "+v.String())
	}
	return strings.Join(ret[:], "\n")
}

func (x *VirtualIpaddressStmt) String() string {
	var ret []string
	ret = append(ret, "virtual_ipaddress ")
	for _, v := range x.addresses {
		ret = append(ret, v.expr())
	}
	return strings.Join(ret[:], " - ")
}

func (x *VirtualServerGroupStmt) String() string {
	var ret []string
	ret = append(ret, "virtual_server_group "+x.group)
	for _, v := range x.vips {
		ret = append(ret, v.expr())
	}
	return strings.Join(ret[:], " - ")
}

func (x *VirtualServerStmt) String() string {
	var ret []string
	ret = append(ret, "virtual_server group "+x.group)
	for _, v := range x.stmts {
		ret = append(ret, "- "+v.String())
	}
	return strings.Join(ret[:], "\n")
}

func (x *SorryServerStmt) String() string {
	return "sorry_server " + x.ip.expr() + ":" + x.port.expr()
}

func (x *RealServerStmt) String() string {
	var ret []string
	ret = append(ret, "real_server "+x.ip.expr()+":"+x.port.expr())
	for _, v := range x.stmts {
		ret = append(ret, "- "+v.String())
	}
	return strings.Join(ret[:], "\n")
}

func (x *HttpGetStmt) String() string {
	var ret []string
	ret = append(ret, "HTTP_GET")
	for _, v := range x.stmts {
		ret = append(ret, "- "+v.String())
	}
	return strings.Join(ret[:], "\n")
}

func (x *SslGetStmt) String() string {
	var ret []string
	ret = append(ret, "SSL_GET")
	for _, v := range x.stmts {
		ret = append(ret, "- "+v.String())
	}
	return strings.Join(ret[:], "\n")
}

func (x *TcpCheckStmt) String() string {
	var ret []string
	ret = append(ret, "TCP_CHECK")
	for _, v := range x.stmts {
		ret = append(ret, "- "+v.String())
	}
	return strings.Join(ret[:], "\n")
}

func (x *MiscCheckStmt) String() string {
	var ret []string
	ret = append(ret, "MISC_CHECK")
	for _, v := range x.stmts {
		ret = append(ret, "- "+v.String())
	}
	return strings.Join(ret[:], "\n")
}

func (x *UrlStmt) String() string {
	var ret []string
	ret = append(ret, "url")
	for _, v := range x.stmts {
		ret = append(ret, "- "+v.String())
	}
	return strings.Join(ret[:], "\n")
}

func (x *GroupStmt) String() string {
	var ret []string
	ret = append(ret, "group")
	for _, v := range x.exprs {
		ret = append(ret, v.expr())
	}
	return strings.Join(ret[:], " - ")
}

func (x *ExprStmt) String() string {
	return fmt.Sprintf("%s : %s", x.key, x.value)
}

func (x *AuthenticationStmt) String() string {
	var ret []string
	ret = append(ret, "AUTH")
	for _, v := range x.stmts {
		ret = append(ret, v.String())
	}
	return strings.Join(ret[:], " - ")
}
func (x *AuthTypeStmt) String() string {
	return fmt.Sprintf("AuthType: %s", x.expr.expr())
}
func (x *AuthPassStmt) String() string {
	return fmt.Sprintf("AuthPass: %s", x.expr.expr())
}

type NumExpr struct {
	lit string
}

type IdentExpr struct {
	lit string
}

type StringExpr struct {
	lit string
}

type IpExpr struct {
	lit  string
	dev  string
	port *NumExpr
}

func (x *NumExpr) expr() string {
	return x.lit
}
func (x *IdentExpr) expr() string {
	return x.lit
}
func (x *StringExpr) expr() string {
	return x.lit
}
func (x *IpExpr) expr() string {
	return x.lit
}

type Token struct {
	tok int
	lit string
	pos Position
}
