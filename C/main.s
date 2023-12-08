	.file	"main.c"
	.text
	.def	fprintf;	.scl	3;	.type	32;	.endef
	.seh_proc	fprintf
fprintf:
	pushq	%rbp
	.seh_pushreg	%rbp
	movq	%rsp, %rbp
	.seh_setframe	%rbp, 0
	subq	$48, %rsp
	.seh_stackalloc	48
	.seh_endprologue
	movq	%rcx, 16(%rbp)
	movq	%rdx, 24(%rbp)
	movq	%r8, 32(%rbp)
	movq	%r9, 40(%rbp)
	leaq	32(%rbp), %rax
	movq	%rax, -16(%rbp)
	movq	-16(%rbp), %rcx
	movq	24(%rbp), %rdx
	movq	16(%rbp), %rax
	movq	%rcx, %r8
	movq	%rax, %rcx
	call	__mingw_vfprintf
	movl	%eax, -4(%rbp)
	movl	-4(%rbp), %eax
	addq	$48, %rsp
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
	.globl	isPrime
	.def	isPrime;	.scl	2;	.type	32;	.endef
	.seh_proc	isPrime
isPrime:
	pushq	%rbp
	.seh_pushreg	%rbp
	movq	%rsp, %rbp
	.seh_setframe	%rbp, 0
	subq	$16, %rsp
	.seh_stackalloc	16
	.seh_endprologue
	movl	%ecx, 16(%rbp)
	cmpl	$1, 16(%rbp)
	jg	.L6
	movl	$0, %eax
	jmp	.L7
.L6:
	movl	$2, -4(%rbp)
	jmp	.L8
.L10:
	movl	16(%rbp), %eax
	cltd
	idivl	-4(%rbp)
	movl	%edx, %eax
	testl	%eax, %eax
	jne	.L9
	movl	$0, %eax
	jmp	.L7
.L9:
	addl	$1, -4(%rbp)
.L8:
	movl	-4(%rbp), %eax
	imull	%eax, %eax
	cmpl	%eax, 16(%rbp)
	jge	.L10
	movl	$1, %eax
.L7:
	addq	$16, %rsp
	popq	%rbp
	ret
	.seh_endproc
	.def	__main;	.scl	2;	.type	32;	.endef
	.section .rdata,"dr"
.LC0:
	.ascii "w\0"
.LC1:
	.ascii "checkPrime.c\0"
	.align 8
.LC2:
	.ascii "Error opening file for writing.\12\0"
.LC3:
	.ascii "#ifndef CHECK_PRIME_H\12\0"
.LC4:
	.ascii "#define CHECK_PRIME_H\12\12\0"
.LC5:
	.ascii "#include <stdio.h>\12\0"
.LC6:
	.ascii "#include <stdbool.h>\12\12\0"
.LC7:
	.ascii "bool checkPrime(int num) {\12\0"
.LC8:
	.ascii "    if (num == %d) {\12\0"
.LC9:
	.ascii "        return true;\12\0"
.LC10:
	.ascii "    }\12\0"
.LC11:
	.ascii "    return false;\12}\12\12\0"
.LC12:
	.ascii "int main() {\12\0"
.LC13:
	.ascii "    int num;\12\0"
	.align 8
.LC14:
	.ascii "    printf(\"Enter a number: \");\12\0"
.LC15:
	.ascii "    scanf(\"%%d\", &num);\12\12\0"
.LC16:
	.ascii "    if (checkPrime(num)) {\12\0"
	.align 8
.LC17:
	.ascii "        printf(\"%%d is a prime number.\\n\", num);\12\0"
.LC18:
	.ascii "    } else {\12\0"
	.align 8
.LC19:
	.ascii "        printf(\"%%d is not a prime number.\\n\", num);\12\0"
.LC20:
	.ascii "    }\12\12\0"
.LC21:
	.ascii "    return 0;\12\0"
.LC22:
	.ascii "}\12\0"
.LC23:
	.ascii "#endif // CHECK_PRIME_H\12\0"
.LC24:
	.ascii "File written successfully.\12\0"
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
	movq	%rax, %rdx
	leaq	.LC1(%rip), %rax
	movq	%rax, %rcx
	call	fopen
	movq	%rax, -16(%rbp)
	cmpq	$0, -16(%rbp)
	jne	.L12
	movl	$2, %ecx
	movq	__imp___acrt_iob_func(%rip), %rax
	call	*%rax
	movq	%rax, %rcx
	leaq	.LC2(%rip), %rax
	movq	%rax, %rdx
	call	fprintf
	movl	$1, %eax
	jmp	.L13
.L12:
	movq	-16(%rbp), %rax
	leaq	.LC3(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC4(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC5(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC6(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC7(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movl	$1, -4(%rbp)
	jmp	.L14
.L16:
	movl	-4(%rbp), %eax
	movl	%eax, %ecx
	call	isPrime
	testb	%al, %al
	je	.L15
	movl	-4(%rbp), %edx
	movq	-16(%rbp), %rax
	movl	%edx, %r8d
	leaq	.LC8(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC9(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC10(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
.L15:
	addl	$1, -4(%rbp)
.L14:
	cmpl	$1000, -4(%rbp)
	jle	.L16
	movq	-16(%rbp), %rax
	leaq	.LC11(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC12(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC13(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC14(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC15(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC16(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC17(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC18(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC19(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC20(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC21(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC22(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	leaq	.LC23(%rip), %rdx
	movq	%rax, %rcx
	call	fprintf
	movq	-16(%rbp), %rax
	movq	%rax, %rcx
	call	fclose
	leaq	.LC24(%rip), %rax
	movq	%rax, %rcx
	call	printf
	movl	$0, %eax
.L13:
	addq	$48, %rsp
	popq	%rbp
	ret
	.seh_endproc
	.ident	"GCC: (MinGW-W64 x86_64-msvcrt-posix-seh, built by Brecht Sanders) 13.1.0"
	.def	__mingw_vfprintf;	.scl	2;	.type	32;	.endef
	.def	fopen;	.scl	2;	.type	32;	.endef
	.def	fclose;	.scl	2;	.type	32;	.endef
