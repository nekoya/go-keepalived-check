%{
package keepalived
%}

%union{
    statements []Statement
    statement  Statement
    exprs      []Expression
    expr       Expression
    tok        Token
}

%type<statements> statements global_defs_stmts vrrp_sync_group_stmts vrrp_instance_stmts auth_stmts virtual_server_stmts real_server_stmts http_get_stmts url_stmts tcp_check_stmts misc_check_stmts
%type<statement> statement global_defs_stmt vrrp_sync_group_stmt vrrp_instance_stmt auth_stmt virtual_server_stmt real_server_stmt http_get_stmt url_stmt tcp_check_stmt misc_check_stmt
%type<exprs> exprs ipaddresses ip_ports
%type<expr> expr state_type auth_type ipaddress ip_port lb_algo lb_kind

%token<tok> IDENT NUMBER IP STRING

%token<tok> GLOBAL_DEFS
%token<tok> NOTIFICATION_EMAIL NOTIFICATION_EMAIL_FROM SMTP_SERVER SMTP_CONNECT_TIMEOUT ROUTER_ID ENABLE_TRAPS

%token<tok> VRRP_SYNC_GROUP GROUP
%token<tok> NOTIFY_MASTER NOTIFY_BACKUP NOTIFY_FAULT NOTIFY SMTP_ALERT

%token<tok> VRRP_INSTANCE
%token<tok> STATE MASTER BACKUP
%token<tok> INTERFACE GARP_MASTER_DELAY VIRTUAL_ROUTER_ID PRIORITY ADVERT_INT NOPREEMPT
%token<tok> AUTHENTICATION AUTH_TYPE AUTH_PASS PASS AH
%token<tok> VIRTUAL_IPADDRESS DEV

%token<tok> VIRTUAL_SERVER_GROUP VIRTUAL_SERVER
%token<tok> DELAY_LOOP SORRY_SERVER
%token<tok> LVS_SCHED LB_ALGO RR WRR LC WLC LBLC SH DH
%token<tok> LVS_METHOD LB_KIND NAT DR TUN
%token<tok> PERSISTENCE_TIMEOUT PERSISTENCE_GRANULARITY
%token<tok> PROTOCOL TCP

%token<tok> REAL_SERVER
%token<tok> WEIGHT
%token<tok> INHIBIT_ON_FAILURE
%token<tok> HTTP_GET SSL_GET TCP_CHECK MISC_CHECK
%token<tok> URL PATH DIGEST STATUS_CODE
%token<tok> CONNECT_PORT CONNECT_TIMEOUT NB_GET_RETRY DELAY_BEFORE_RETRY
%token<tok> MISC_PATH MISC_TIMEOUT

%%

statements
    :
    {
        $$ = nil
        if l, ok := yylex.(*Lexer); ok {
            l.statements = $$
        }
    }
    | statement statements
    {
        $$ = append([]Statement{$1}, $2...)
        if l, ok := yylex.(*Lexer); ok {
            l.statements = $$
        }
    }
statement
    : GLOBAL_DEFS '{' global_defs_stmts '}'
    {
        $$ = &GlobalDefsStmt{stmts: $3}
    }
    | VRRP_SYNC_GROUP IDENT '{' vrrp_sync_group_stmts '}'
    {
        $$ = &VrrpSyncGroupStmt{group:$2.lit, stmts: $4}
    }
    | VRRP_INSTANCE IDENT '{' vrrp_instance_stmts '}'
    {
        $$ = &VrrpInstanceStmt{inside_network:$2.lit, stmts: $4}
    }
    | VIRTUAL_SERVER_GROUP IDENT '{' ip_ports '}'
    {
        $$ = &VirtualServerGroupStmt{group:$2.lit, vips: $4}
    }
    | VIRTUAL_SERVER GROUP IDENT '{' virtual_server_stmts '}'
    {
        $$ = &VirtualServerStmt{group:$3.lit, stmts: $5}
    }
    | VIRTUAL_SERVER IP NUMBER '{' virtual_server_stmts '}'
    {
        $$ = &VirtualServerStmt{ip: &IdentExpr{lit: $2.lit}, port: &NumExpr{lit: $3.lit}, stmts: $5}
    }

global_defs_stmts
    : { $$ = []Statement{} }
    | global_defs_stmts global_defs_stmt { $$ = append($1, $2) }

global_defs_stmt
    : NOTIFICATION_EMAIL '{' exprs '}' { $$ = &GroupStmt{exprs: $3} }
    | NOTIFICATION_EMAIL_FROM expr { $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2} }
    | SMTP_SERVER expr { $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2} }
    | SMTP_CONNECT_TIMEOUT NUMBER { $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}} }
    | ROUTER_ID expr { $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2} }
    | ENABLE_TRAPS { $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}} }

vrrp_sync_group_stmts
    : { $$ = []Statement{} }
    | vrrp_sync_group_stmts vrrp_sync_group_stmt { $$ = append($1, $2) }

vrrp_sync_group_stmt
    : GROUP '{' exprs '}' { $$ = &GroupStmt{exprs: $3} }
    | NOTIFY_MASTER expr { $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2} }
    | NOTIFY_BACKUP expr { $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2} }
    | NOTIFY_FAULT  expr { $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2} }
    | NOTIFY        expr { $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2} }
    | SMTP_ALERT { $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}} }

vrrp_instance_stmts
    : { $$ = []Statement{} }
    | vrrp_instance_stmts vrrp_instance_stmt { $$ = append($1, $2) }

vrrp_instance_stmt
    : STATE state_type
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2}
    }
    | INTERFACE IDENT
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &IdentExpr{lit: $2.lit}}
    }
    | GARP_MASTER_DELAY NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | NOPREEMPT
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}}
    }
    | SMTP_ALERT
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}}
    }
    | VIRTUAL_ROUTER_ID NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | PRIORITY NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | ADVERT_INT NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | VIRTUAL_IPADDRESS '{' ipaddresses '}'
    {
        $$ = &VirtualIpaddressStmt{addresses: $3}
    }
    | AUTHENTICATION '{' auth_stmts '}'
    {
        $$ = &AuthenticationStmt{stmts: $3}
    }

virtual_server_stmts
    : { $$ = []Statement{} }
    | virtual_server_stmts virtual_server_stmt { $$ = append($1, $2) }

virtual_server_stmt
    : DELAY_LOOP NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | LVS_SCHED lb_algo
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2}
    }
    | LB_ALGO lb_algo
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2}
    }
    | LVS_METHOD lb_kind
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2}
    }
    | LB_KIND lb_kind
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: $2}
    }
    | PERSISTENCE_TIMEOUT NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | PERSISTENCE_GRANULARITY IDENT
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &IdentExpr{lit: $2.lit}}
    }
    | PROTOCOL TCP
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &IdentExpr{lit: $2.lit}}
    }
    | SORRY_SERVER IP NUMBER
    {
        $$ = &SorryServerStmt{ip: &IdentExpr{lit: $2.lit}, port: &NumExpr{lit: $3.lit}}
    }
    | REAL_SERVER IP NUMBER '{' real_server_stmts '}'
    {
        $$ = &RealServerStmt{ip: &IdentExpr{lit: $2.lit}, port: &NumExpr{lit: $3.lit}, stmts: $5}
    }

lb_algo
    : RR   { $$ = &IdentExpr{lit: $1.lit} }
    | WRR  { $$ = &IdentExpr{lit: $1.lit} }
    | LC   { $$ = &IdentExpr{lit: $1.lit} }
    | WLC  { $$ = &IdentExpr{lit: $1.lit} }
    | LBLC { $$ = &IdentExpr{lit: $1.lit} }
    | SH   { $$ = &IdentExpr{lit: $1.lit} }
    | DH   { $$ = &IdentExpr{lit: $1.lit} }

lb_kind
    : NAT { $$ = &IdentExpr{lit: $1.lit} }
    | DR  { $$ = &IdentExpr{lit: $1.lit} }
    | TUN { $$ = &IdentExpr{lit: $1.lit} }

real_server_stmts
    : { $$ = []Statement{} }
    | real_server_stmts real_server_stmt { $$ = append($1, $2) }

real_server_stmt
    : WEIGHT NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | INHIBIT_ON_FAILURE
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}}
    }
    | HTTP_GET '{' http_get_stmts '}'
    {
        $$ = &HttpGetStmt{stmts: $3}
    }
    | SSL_GET '{' http_get_stmts '}'
    {
        $$ = &SslGetStmt{stmts: $3}
    }
    | TCP_CHECK '{' tcp_check_stmts '}'
    {
        $$ = &TcpCheckStmt{stmts: $3}
    }
    | MISC_CHECK '{' misc_check_stmts '}'
    {
        $$ = &MiscCheckStmt{stmts: $3}
    }

http_get_stmts
    : { $$ = []Statement{} }
    | http_get_stmts http_get_stmt { $$ = append($1, $2) }

http_get_stmt
    : CONNECT_PORT NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | CONNECT_TIMEOUT NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | NB_GET_RETRY NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | DELAY_BEFORE_RETRY NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | URL '{' url_stmts '}'
    {
        $$ = &UrlStmt{stmts: $3}
    }

tcp_check_stmts
    : { $$ = []Statement{} }
    | tcp_check_stmts tcp_check_stmt { $$ = append($1, $2) }

tcp_check_stmt
    : CONNECT_PORT NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }
    | CONNECT_TIMEOUT NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }

misc_check_stmts
    : { $$ = []Statement{} }
    | misc_check_stmts misc_check_stmt { $$ = append($1, $2) }

misc_check_stmt
    : MISC_PATH STRING
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &StringExpr{lit: $2.lit}}
    }
    | MISC_TIMEOUT NUMBER
    {
        $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &NumExpr{lit: $2.lit}}
    }

url_stmts
    : { $$ = []Statement{} }
    | url_stmts url_stmt { $$ = append($1, $2) }

url_stmt
    : PATH   IDENT { $$ = &ExprStmt{key: &IdentExpr{lit: $1.lit}, value: &IdentExpr{lit: $2.lit}} }
    | DIGEST IDENT { $$ = &ExprStmt{key: &IdentExpr{$1.lit}, value: &NumExpr{lit: $2.lit}} }
    | STATUS_CODE NUMBER { $$ = &ExprStmt{key: &IdentExpr{$1.lit}, value: &NumExpr{lit: $2.lit}} }

ipaddresses
    : { $$ = []Expression{} }
    | ipaddresses ipaddress { $$ = append($1, $2) }

ip_ports
    : { $$ = []Expression{} }
    | ip_ports ip_port { $$ = append($1, $2) }

auth_stmts
    : { $$ = []Statement{} }
    | auth_stmts auth_stmt { $$ = append($1, $2) }

auth_stmt
    : AUTH_TYPE auth_type { $$ = &AuthTypeStmt{expr: $2} }
    | AUTH_PASS expr { $$ = &AuthPassStmt{expr: $2} }

auth_type
    : PASS { $$ = &IdentExpr{lit: $1.lit} }
    | AH { $$ = &IdentExpr{lit: $1.lit} }

state_type
    : MASTER { $$ = &IdentExpr{lit: $1.lit} }
    | BACKUP { $$ = &IdentExpr{lit: $1.lit} }

ipaddress
    : IP { $$ = &IpExpr{lit: $1.lit} }
    | IP DEV IDENT { $$ = &IpExpr{lit: $1.lit, dev: $3.lit} }

ip_port
    : IP NUMBER { $$ = &IpExpr{lit: $1.lit, port: &NumExpr{lit: $2.lit}} }

exprs
    : { $$ = []Expression{} }
    | exprs expr { $$ = append($1, $2) }

expr
    : NUMBER { $$ = &NumExpr{lit: $1.lit} }
    | IDENT { $$ = &IdentExpr{lit: $1.lit} }
    | STRING { $$ = &StringExpr{lit: $1.lit} }
    | IP { $$ = &IpExpr{lit: $1.lit} }

%%
