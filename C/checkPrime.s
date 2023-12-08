	.file	"checkPrime.c"
	.text
	.def	scanf;	.scl	3;	.type	32;	.endef
	.seh_proc	scanf
scanf:
	pushq	%rbp
	.seh_pushreg	%rbp
	pushq	%rbx
	.seh_pushreg	%rbx
	subq	$56, %rsp
	.seh_stackalloc	56
	leaq	48(%rsp), %rbp
	.seh_setframe	%rbp, 48
	.seh_endprologue
	movq	%rcx, 32(%rbp)
	movq	%rdx, 40(%rbp)
	movq	%r8, 48(%rbp)
	movq	%r9, 56(%rbp)
	leaq	40(%rbp), %rax
	movq	%rax, -16(%rbp)
	movq	-16(%rbp), %rbx
	movl	$0, %ecx
	movq	__imp___acrt_iob_func(%rip), %rax
	call	*%rax
	movq	%rax, %rcx
	movq	32(%rbp), %rax
	movq	%rbx, %r8
	movq	%rax, %rdx
	call	__mingw_vfscanf
	movl	%eax, -4(%rbp)
	movl	-4(%rbp), %eax
	addq	$56, %rsp
	popq	%rbx
	popq	%rbp
	ret
	.seh_endproc
	.def	printf;	.scl	3;	.type	32;	.endef
	.seh_proc	printf
printf:
	pushq	%rbp
	.seh_pushreg	%rbp
	pushq	%rbx
	.seh_pushreg	%rbx
	subq	$56, %rsp
	.seh_stackalloc	56
	leaq	48(%rsp), %rbp
	.seh_setframe	%rbp, 48
	.seh_endprologue
	movq	%rcx, 32(%rbp)
	movq	%rdx, 40(%rbp)
	movq	%r8, 48(%rbp)
	movq	%r9, 56(%rbp)
	leaq	40(%rbp), %rax
	movq	%rax, -16(%rbp)
	movq	-16(%rbp), %rbx
	movl	$1, %ecx
	movq	__imp___acrt_iob_func(%rip), %rax
	call	*%rax
	movq	%rax, %rcx
	movq	32(%rbp), %rax
	movq	%rbx, %r8
	movq	%rax, %rdx
	call	__mingw_vfprintf
	movl	%eax, -4(%rbp)
	movl	-4(%rbp), %eax
	addq	$56, %rsp
	popq	%rbx
	popq	%rbp
	ret
	.seh_endproc
	.globl	checkPrime
	.def	checkPrime;	.scl	2;	.type	32;	.endef
	.seh_proc	checkPrime
checkPrime:
	pushq	%rbp
	.seh_pushreg	%rbp
	movq	%rsp, %rbp
	.seh_setframe	%rbp, 0
	.seh_endprologue
	movl	%ecx, 16(%rbp)
	cmpl	$2, 16(%rbp)
	jne	.L6
	movl	$1, %eax
	jmp	.L7
.L6:
	cmpl	$3, 16(%rbp)
	jne	.L8
	movl	$1, %eax
	jmp	.L7
.L8:
	cmpl	$5, 16(%rbp)
	jne	.L9
	movl	$1, %eax
	jmp	.L7
.L9:
	cmpl	$7, 16(%rbp)
	jne	.L10
	movl	$1, %eax
	jmp	.L7
.L10:
	cmpl	$11, 16(%rbp)
	jne	.L11
	movl	$1, %eax
	jmp	.L7
.L11:
	cmpl	$13, 16(%rbp)
	jne	.L12
	movl	$1, %eax
	jmp	.L7
.L12:
	cmpl	$17, 16(%rbp)
	jne	.L13
	movl	$1, %eax
	jmp	.L7
.L13:
	cmpl	$19, 16(%rbp)
	jne	.L14
	movl	$1, %eax
	jmp	.L7
.L14:
	cmpl	$23, 16(%rbp)
	jne	.L15
	movl	$1, %eax
	jmp	.L7
.L15:
	cmpl	$29, 16(%rbp)
	jne	.L16
	movl	$1, %eax
	jmp	.L7
.L16:
	cmpl	$31, 16(%rbp)
	jne	.L17
	movl	$1, %eax
	jmp	.L7
.L17:
	cmpl	$37, 16(%rbp)
	jne	.L18
	movl	$1, %eax
	jmp	.L7
.L18:
	cmpl	$41, 16(%rbp)
	jne	.L19
	movl	$1, %eax
	jmp	.L7
.L19:
	cmpl	$43, 16(%rbp)
	jne	.L20
	movl	$1, %eax
	jmp	.L7
.L20:
	cmpl	$47, 16(%rbp)
	jne	.L21
	movl	$1, %eax
	jmp	.L7
.L21:
	cmpl	$53, 16(%rbp)
	jne	.L22
	movl	$1, %eax
	jmp	.L7
.L22:
	cmpl	$59, 16(%rbp)
	jne	.L23
	movl	$1, %eax
	jmp	.L7
.L23:
	cmpl	$61, 16(%rbp)
	jne	.L24
	movl	$1, %eax
	jmp	.L7
.L24:
	cmpl	$67, 16(%rbp)
	jne	.L25
	movl	$1, %eax
	jmp	.L7
.L25:
	cmpl	$71, 16(%rbp)
	jne	.L26
	movl	$1, %eax
	jmp	.L7
.L26:
	cmpl	$73, 16(%rbp)
	jne	.L27
	movl	$1, %eax
	jmp	.L7
.L27:
	cmpl	$79, 16(%rbp)
	jne	.L28
	movl	$1, %eax
	jmp	.L7
.L28:
	cmpl	$83, 16(%rbp)
	jne	.L29
	movl	$1, %eax
	jmp	.L7
.L29:
	cmpl	$89, 16(%rbp)
	jne	.L30
	movl	$1, %eax
	jmp	.L7
.L30:
	cmpl	$97, 16(%rbp)
	jne	.L31
	movl	$1, %eax
	jmp	.L7
.L31:
	cmpl	$101, 16(%rbp)
	jne	.L32
	movl	$1, %eax
	jmp	.L7
.L32:
	cmpl	$103, 16(%rbp)
	jne	.L33
	movl	$1, %eax
	jmp	.L7
.L33:
	cmpl	$107, 16(%rbp)
	jne	.L34
	movl	$1, %eax
	jmp	.L7
.L34:
	cmpl	$109, 16(%rbp)
	jne	.L35
	movl	$1, %eax
	jmp	.L7
.L35:
	cmpl	$113, 16(%rbp)
	jne	.L36
	movl	$1, %eax
	jmp	.L7
.L36:
	cmpl	$127, 16(%rbp)
	jne	.L37
	movl	$1, %eax
	jmp	.L7
.L37:
	cmpl	$131, 16(%rbp)
	jne	.L38
	movl	$1, %eax
	jmp	.L7
.L38:
	cmpl	$137, 16(%rbp)
	jne	.L39
	movl	$1, %eax
	jmp	.L7
.L39:
	cmpl	$139, 16(%rbp)
	jne	.L40
	movl	$1, %eax
	jmp	.L7
.L40:
	cmpl	$149, 16(%rbp)
	jne	.L41
	movl	$1, %eax
	jmp	.L7
.L41:
	cmpl	$151, 16(%rbp)
	jne	.L42
	movl	$1, %eax
	jmp	.L7
.L42:
	cmpl	$157, 16(%rbp)
	jne	.L43
	movl	$1, %eax
	jmp	.L7
.L43:
	cmpl	$163, 16(%rbp)
	jne	.L44
	movl	$1, %eax
	jmp	.L7
.L44:
	cmpl	$167, 16(%rbp)
	jne	.L45
	movl	$1, %eax
	jmp	.L7
.L45:
	cmpl	$173, 16(%rbp)
	jne	.L46
	movl	$1, %eax
	jmp	.L7
.L46:
	cmpl	$179, 16(%rbp)
	jne	.L47
	movl	$1, %eax
	jmp	.L7
.L47:
	cmpl	$181, 16(%rbp)
	jne	.L48
	movl	$1, %eax
	jmp	.L7
.L48:
	cmpl	$191, 16(%rbp)
	jne	.L49
	movl	$1, %eax
	jmp	.L7
.L49:
	cmpl	$193, 16(%rbp)
	jne	.L50
	movl	$1, %eax
	jmp	.L7
.L50:
	cmpl	$197, 16(%rbp)
	jne	.L51
	movl	$1, %eax
	jmp	.L7
.L51:
	cmpl	$199, 16(%rbp)
	jne	.L52
	movl	$1, %eax
	jmp	.L7
.L52:
	cmpl	$211, 16(%rbp)
	jne	.L53
	movl	$1, %eax
	jmp	.L7
.L53:
	cmpl	$223, 16(%rbp)
	jne	.L54
	movl	$1, %eax
	jmp	.L7
.L54:
	cmpl	$227, 16(%rbp)
	jne	.L55
	movl	$1, %eax
	jmp	.L7
.L55:
	cmpl	$229, 16(%rbp)
	jne	.L56
	movl	$1, %eax
	jmp	.L7
.L56:
	cmpl	$233, 16(%rbp)
	jne	.L57
	movl	$1, %eax
	jmp	.L7
.L57:
	cmpl	$239, 16(%rbp)
	jne	.L58
	movl	$1, %eax
	jmp	.L7
.L58:
	cmpl	$241, 16(%rbp)
	jne	.L59
	movl	$1, %eax
	jmp	.L7
.L59:
	cmpl	$251, 16(%rbp)
	jne	.L60
	movl	$1, %eax
	jmp	.L7
.L60:
	cmpl	$257, 16(%rbp)
	jne	.L61
	movl	$1, %eax
	jmp	.L7
.L61:
	cmpl	$263, 16(%rbp)
	jne	.L62
	movl	$1, %eax
	jmp	.L7
.L62:
	cmpl	$269, 16(%rbp)
	jne	.L63
	movl	$1, %eax
	jmp	.L7
.L63:
	cmpl	$271, 16(%rbp)
	jne	.L64
	movl	$1, %eax
	jmp	.L7
.L64:
	cmpl	$277, 16(%rbp)
	jne	.L65
	movl	$1, %eax
	jmp	.L7
.L65:
	cmpl	$281, 16(%rbp)
	jne	.L66
	movl	$1, %eax
	jmp	.L7
.L66:
	cmpl	$283, 16(%rbp)
	jne	.L67
	movl	$1, %eax
	jmp	.L7
.L67:
	cmpl	$293, 16(%rbp)
	jne	.L68
	movl	$1, %eax
	jmp	.L7
.L68:
	cmpl	$307, 16(%rbp)
	jne	.L69
	movl	$1, %eax
	jmp	.L7
.L69:
	cmpl	$311, 16(%rbp)
	jne	.L70
	movl	$1, %eax
	jmp	.L7
.L70:
	cmpl	$313, 16(%rbp)
	jne	.L71
	movl	$1, %eax
	jmp	.L7
.L71:
	cmpl	$317, 16(%rbp)
	jne	.L72
	movl	$1, %eax
	jmp	.L7
.L72:
	cmpl	$331, 16(%rbp)
	jne	.L73
	movl	$1, %eax
	jmp	.L7
.L73:
	cmpl	$337, 16(%rbp)
	jne	.L74
	movl	$1, %eax
	jmp	.L7
.L74:
	cmpl	$347, 16(%rbp)
	jne	.L75
	movl	$1, %eax
	jmp	.L7
.L75:
	cmpl	$349, 16(%rbp)
	jne	.L76
	movl	$1, %eax
	jmp	.L7
.L76:
	cmpl	$353, 16(%rbp)
	jne	.L77
	movl	$1, %eax
	jmp	.L7
.L77:
	cmpl	$359, 16(%rbp)
	jne	.L78
	movl	$1, %eax
	jmp	.L7
.L78:
	cmpl	$367, 16(%rbp)
	jne	.L79
	movl	$1, %eax
	jmp	.L7
.L79:
	cmpl	$373, 16(%rbp)
	jne	.L80
	movl	$1, %eax
	jmp	.L7
.L80:
	cmpl	$379, 16(%rbp)
	jne	.L81
	movl	$1, %eax
	jmp	.L7
.L81:
	cmpl	$383, 16(%rbp)
	jne	.L82
	movl	$1, %eax
	jmp	.L7
.L82:
	cmpl	$389, 16(%rbp)
	jne	.L83
	movl	$1, %eax
	jmp	.L7
.L83:
	cmpl	$397, 16(%rbp)
	jne	.L84
	movl	$1, %eax
	jmp	.L7
.L84:
	cmpl	$401, 16(%rbp)
	jne	.L85
	movl	$1, %eax
	jmp	.L7
.L85:
	cmpl	$409, 16(%rbp)
	jne	.L86
	movl	$1, %eax
	jmp	.L7
.L86:
	cmpl	$419, 16(%rbp)
	jne	.L87
	movl	$1, %eax
	jmp	.L7
.L87:
	cmpl	$421, 16(%rbp)
	jne	.L88
	movl	$1, %eax
	jmp	.L7
.L88:
	cmpl	$431, 16(%rbp)
	jne	.L89
	movl	$1, %eax
	jmp	.L7
.L89:
	cmpl	$433, 16(%rbp)
	jne	.L90
	movl	$1, %eax
	jmp	.L7
.L90:
	cmpl	$439, 16(%rbp)
	jne	.L91
	movl	$1, %eax
	jmp	.L7
.L91:
	cmpl	$443, 16(%rbp)
	jne	.L92
	movl	$1, %eax
	jmp	.L7
.L92:
	cmpl	$449, 16(%rbp)
	jne	.L93
	movl	$1, %eax
	jmp	.L7
.L93:
	cmpl	$457, 16(%rbp)
	jne	.L94
	movl	$1, %eax
	jmp	.L7
.L94:
	cmpl	$461, 16(%rbp)
	jne	.L95
	movl	$1, %eax
	jmp	.L7
.L95:
	cmpl	$463, 16(%rbp)
	jne	.L96
	movl	$1, %eax
	jmp	.L7
.L96:
	cmpl	$467, 16(%rbp)
	jne	.L97
	movl	$1, %eax
	jmp	.L7
.L97:
	cmpl	$479, 16(%rbp)
	jne	.L98
	movl	$1, %eax
	jmp	.L7
.L98:
	cmpl	$487, 16(%rbp)
	jne	.L99
	movl	$1, %eax
	jmp	.L7
.L99:
	cmpl	$491, 16(%rbp)
	jne	.L100
	movl	$1, %eax
	jmp	.L7
.L100:
	cmpl	$499, 16(%rbp)
	jne	.L101
	movl	$1, %eax
	jmp	.L7
.L101:
	cmpl	$503, 16(%rbp)
	jne	.L102
	movl	$1, %eax
	jmp	.L7
.L102:
	cmpl	$509, 16(%rbp)
	jne	.L103
	movl	$1, %eax
	jmp	.L7
.L103:
	cmpl	$521, 16(%rbp)
	jne	.L104
	movl	$1, %eax
	jmp	.L7
.L104:
	cmpl	$523, 16(%rbp)
	jne	.L105
	movl	$1, %eax
	jmp	.L7
.L105:
	cmpl	$541, 16(%rbp)
	jne	.L106
	movl	$1, %eax
	jmp	.L7
.L106:
	cmpl	$547, 16(%rbp)
	jne	.L107
	movl	$1, %eax
	jmp	.L7
.L107:
	cmpl	$557, 16(%rbp)
	jne	.L108
	movl	$1, %eax
	jmp	.L7
.L108:
	cmpl	$563, 16(%rbp)
	jne	.L109
	movl	$1, %eax
	jmp	.L7
.L109:
	cmpl	$569, 16(%rbp)
	jne	.L110
	movl	$1, %eax
	jmp	.L7
.L110:
	cmpl	$571, 16(%rbp)
	jne	.L111
	movl	$1, %eax
	jmp	.L7
.L111:
	cmpl	$577, 16(%rbp)
	jne	.L112
	movl	$1, %eax
	jmp	.L7
.L112:
	cmpl	$587, 16(%rbp)
	jne	.L113
	movl	$1, %eax
	jmp	.L7
.L113:
	cmpl	$593, 16(%rbp)
	jne	.L114
	movl	$1, %eax
	jmp	.L7
.L114:
	cmpl	$599, 16(%rbp)
	jne	.L115
	movl	$1, %eax
	jmp	.L7
.L115:
	cmpl	$601, 16(%rbp)
	jne	.L116
	movl	$1, %eax
	jmp	.L7
.L116:
	cmpl	$607, 16(%rbp)
	jne	.L117
	movl	$1, %eax
	jmp	.L7
.L117:
	cmpl	$613, 16(%rbp)
	jne	.L118
	movl	$1, %eax
	jmp	.L7
.L118:
	cmpl	$617, 16(%rbp)
	jne	.L119
	movl	$1, %eax
	jmp	.L7
.L119:
	cmpl	$619, 16(%rbp)
	jne	.L120
	movl	$1, %eax
	jmp	.L7
.L120:
	cmpl	$631, 16(%rbp)
	jne	.L121
	movl	$1, %eax
	jmp	.L7
.L121:
	cmpl	$641, 16(%rbp)
	jne	.L122
	movl	$1, %eax
	jmp	.L7
.L122:
	cmpl	$643, 16(%rbp)
	jne	.L123
	movl	$1, %eax
	jmp	.L7
.L123:
	cmpl	$647, 16(%rbp)
	jne	.L124
	movl	$1, %eax
	jmp	.L7
.L124:
	cmpl	$653, 16(%rbp)
	jne	.L125
	movl	$1, %eax
	jmp	.L7
.L125:
	cmpl	$659, 16(%rbp)
	jne	.L126
	movl	$1, %eax
	jmp	.L7
.L126:
	cmpl	$661, 16(%rbp)
	jne	.L127
	movl	$1, %eax
	jmp	.L7
.L127:
	cmpl	$673, 16(%rbp)
	jne	.L128
	movl	$1, %eax
	jmp	.L7
.L128:
	cmpl	$677, 16(%rbp)
	jne	.L129
	movl	$1, %eax
	jmp	.L7
.L129:
	cmpl	$683, 16(%rbp)
	jne	.L130
	movl	$1, %eax
	jmp	.L7
.L130:
	cmpl	$691, 16(%rbp)
	jne	.L131
	movl	$1, %eax
	jmp	.L7
.L131:
	cmpl	$701, 16(%rbp)
	jne	.L132
	movl	$1, %eax
	jmp	.L7
.L132:
	cmpl	$709, 16(%rbp)
	jne	.L133
	movl	$1, %eax
	jmp	.L7
.L133:
	cmpl	$719, 16(%rbp)
	jne	.L134
	movl	$1, %eax
	jmp	.L7
.L134:
	cmpl	$727, 16(%rbp)
	jne	.L135
	movl	$1, %eax
	jmp	.L7
.L135:
	cmpl	$733, 16(%rbp)
	jne	.L136
	movl	$1, %eax
	jmp	.L7
.L136:
	cmpl	$739, 16(%rbp)
	jne	.L137
	movl	$1, %eax
	jmp	.L7
.L137:
	cmpl	$743, 16(%rbp)
	jne	.L138
	movl	$1, %eax
	jmp	.L7
.L138:
	cmpl	$751, 16(%rbp)
	jne	.L139
	movl	$1, %eax
	jmp	.L7
.L139:
	cmpl	$757, 16(%rbp)
	jne	.L140
	movl	$1, %eax
	jmp	.L7
.L140:
	cmpl	$761, 16(%rbp)
	jne	.L141
	movl	$1, %eax
	jmp	.L7
.L141:
	cmpl	$769, 16(%rbp)
	jne	.L142
	movl	$1, %eax
	jmp	.L7
.L142:
	cmpl	$773, 16(%rbp)
	jne	.L143
	movl	$1, %eax
	jmp	.L7
.L143:
	cmpl	$787, 16(%rbp)
	jne	.L144
	movl	$1, %eax
	jmp	.L7
.L144:
	cmpl	$797, 16(%rbp)
	jne	.L145
	movl	$1, %eax
	jmp	.L7
.L145:
	cmpl	$809, 16(%rbp)
	jne	.L146
	movl	$1, %eax
	jmp	.L7
.L146:
	cmpl	$811, 16(%rbp)
	jne	.L147
	movl	$1, %eax
	jmp	.L7
.L147:
	cmpl	$821, 16(%rbp)
	jne	.L148
	movl	$1, %eax
	jmp	.L7
.L148:
	cmpl	$823, 16(%rbp)
	jne	.L149
	movl	$1, %eax
	jmp	.L7
.L149:
	cmpl	$827, 16(%rbp)
	jne	.L150
	movl	$1, %eax
	jmp	.L7
.L150:
	cmpl	$829, 16(%rbp)
	jne	.L151
	movl	$1, %eax
	jmp	.L7
.L151:
	cmpl	$839, 16(%rbp)
	jne	.L152
	movl	$1, %eax
	jmp	.L7
.L152:
	cmpl	$853, 16(%rbp)
	jne	.L153
	movl	$1, %eax
	jmp	.L7
.L153:
	cmpl	$857, 16(%rbp)
	jne	.L154
	movl	$1, %eax
	jmp	.L7
.L154:
	cmpl	$859, 16(%rbp)
	jne	.L155
	movl	$1, %eax
	jmp	.L7
.L155:
	cmpl	$863, 16(%rbp)
	jne	.L156
	movl	$1, %eax
	jmp	.L7
.L156:
	cmpl	$877, 16(%rbp)
	jne	.L157
	movl	$1, %eax
	jmp	.L7
.L157:
	cmpl	$881, 16(%rbp)
	jne	.L158
	movl	$1, %eax
	jmp	.L7
.L158:
	cmpl	$883, 16(%rbp)
	jne	.L159
	movl	$1, %eax
	jmp	.L7
.L159:
	cmpl	$887, 16(%rbp)
	jne	.L160
	movl	$1, %eax
	jmp	.L7
.L160:
	cmpl	$907, 16(%rbp)
	jne	.L161
	movl	$1, %eax
	jmp	.L7
.L161:
	cmpl	$911, 16(%rbp)
	jne	.L162
	movl	$1, %eax
	jmp	.L7
.L162:
	cmpl	$919, 16(%rbp)
	jne	.L163
	movl	$1, %eax
	jmp	.L7
.L163:
	cmpl	$929, 16(%rbp)
	jne	.L164
	movl	$1, %eax
	jmp	.L7
.L164:
	cmpl	$937, 16(%rbp)
	jne	.L165
	movl	$1, %eax
	jmp	.L7
.L165:
	cmpl	$941, 16(%rbp)
	jne	.L166
	movl	$1, %eax
	jmp	.L7
.L166:
	cmpl	$947, 16(%rbp)
	jne	.L167
	movl	$1, %eax
	jmp	.L7
.L167:
	cmpl	$953, 16(%rbp)
	jne	.L168
	movl	$1, %eax
	jmp	.L7
.L168:
	cmpl	$967, 16(%rbp)
	jne	.L169
	movl	$1, %eax
	jmp	.L7
.L169:
	cmpl	$971, 16(%rbp)
	jne	.L170
	movl	$1, %eax
	jmp	.L7
.L170:
	cmpl	$977, 16(%rbp)
	jne	.L171
	movl	$1, %eax
	jmp	.L7
.L171:
	cmpl	$983, 16(%rbp)
	jne	.L172
	movl	$1, %eax
	jmp	.L7
.L172:
	cmpl	$991, 16(%rbp)
	jne	.L173
	movl	$1, %eax
	jmp	.L7
.L173:
	cmpl	$997, 16(%rbp)
	jne	.L174
	movl	$1, %eax
	jmp	.L7
.L174:
	movl	$0, %eax
.L7:
	popq	%rbp
	ret
	.seh_endproc
	.def	__main;	.scl	2;	.type	32;	.endef
	.section .rdata,"dr"
.LC0:
	.ascii "Enter a number: \0"
.LC1:
	.ascii "%d\0"
.LC2:
	.ascii "%d is a prime number.\12\0"
.LC3:
	.ascii "%d is not a prime number.\12\0"
	.text
	.globl	main
	.def	main;	.scl	2;	.type	32;	.endef
	.seh_proc	main
main:
	pushq	%rbp
	.seh_pushreg	%rbp
	movq	%rsp, %rbp
	.seh_setframe	%rbp, 0
	subq	$48, %rsp
	.seh_stackalloc	48
	.seh_endprologue
	call	__main
	leaq	.LC0(%rip), %rax
	movq	%rax, %rcx
	call	printf
	leaq	-4(%rbp), %rax
	movq	%rax, %rdx
	leaq	.LC1(%rip), %rax
	movq	%rax, %rcx
	call	scanf
	movl	-4(%rbp), %eax
	movl	%eax, %ecx
	call	checkPrime
	testb	%al, %al
	je	.L176
	movl	-4(%rbp), %eax
	movl	%eax, %edx
	leaq	.LC2(%rip), %rax
	movq	%rax, %rcx
	call	printf
	jmp	.L177
.L176:
	movl	-4(%rbp), %eax
	movl	%eax, %edx
	leaq	.LC3(%rip), %rax
	movq	%rax, %rcx
	call	printf
.L177:
	movl	$0, %eax
	addq	$48, %rsp
	popq	%rbp
	ret
	.seh_endproc
	.ident	"GCC: (MinGW-W64 x86_64-msvcrt-posix-seh, built by Brecht Sanders) 13.1.0"
	.def	__mingw_vfscanf;	.scl	2;	.type	32;	.endef
	.def	__mingw_vfprintf;	.scl	2;	.type	32;	.endef
