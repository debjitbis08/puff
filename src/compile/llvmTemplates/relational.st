; ***************Relational {{.Instr}} ***************
%a.{{.Ninstr}} = call i64()* @popV()
%b.{{.Ninstr}} = call i64()* @popV()
%cond.{{.Ninstr}} = icmp {{.Ninstr}} i64 %a.{{.Ninstr}}, %b.{{.Ninstr}}
br i1 %cond.{{.Ninstr}}, label %TRUE.{{.Ninstr}}, label %FALSE.{{.Ninstr}}

TRUE.{{.Ninstr}}:
    call void(i64)* @pushV(i64 {{.TrueTag}})
    br label %END.{{.Ninstr}}

FALSE.{{.Ninstr}}:
    call void(i64)* @pushV(i64 {{.FalseTag}})
    br label %END.{{.Ninstr}}

END.{{.Ninstr}}: