package base

import "fmt"

const (
	ADD = iota
	MINUS
	MULTI
	DIV
	REMAINDER
	LD
	LDF
	SET
	QUOTE
	DEFUN
	CALL
	IF
	NOT
	AND
	OR
	CASE
	EQ
	ATOMP
	PAIRP
	LISTP
	NILP
	STRINGP
	GT
	GE
	LT
	LE
	PRINT
	PRINC
	JMP
	LJMP
	LABEL
	CONS
	CAR
	CDR
	LAMBDA
	FUNCALL
	APPEND
	PROGN
	LET
	PUSH_LEXICAL_FRAME
	POP_LEXICAL_FRAME
	EXEC_FUNC
	TIME_START
	TIME_END
	FUNCTION
	ERROR
	PROGN_START
	PROGN_END
	FORMAT
	COMMAND
	LENGTH
	MEMBER
	DEFSERVER
	DEFHANDLER
	RUNSERVER
	SETSTATUS
	GETQUERY
	JSON_PARSE
	MAKE_VECTOR
	AREF
	NTH
	ASSOC
	RANDOM
	OPEN
	WRITE
	CLOSE
	READ_FILE
	SPLIT
	MAKE_REQUEST
	DO_REQUEST
	ADD_REQUEST_HEADER
	VECTOR_PUSH
	VECTOR_POP
	VECTOR_LEN
	LIST_FUNCTION
	EVENP
	ODDP
	MAPCAR
	REGEXP_REPLACE_ALL
	REGEXP_MATCH
	SUBSTRING
	MAKE_HASH
	SET_HASH
	GET_HASH
	GLOBAL
	STR_TO_LIST
	INTERN
	TYPE
)

func dumpByteCode(byte_code any) {

	switch byte_code {
	case ADD:
		fmt.Print("ADD\n\n")
	case MINUS:
		fmt.Print("MINUS\n\n")
	case MULTI:
		fmt.Print("MULTI\n\n")
	case DIV:
		fmt.Print("DIV\n\n")
	case REMAINDER:
		fmt.Print("REMAINDER\n\n")
	case LD:
		fmt.Print("LD")
	case LDF:
		fmt.Print("LDF")
	case SET:
		fmt.Print("SET\n\n")
	case QUOTE:
		fmt.Print("QUOTE")
	case DEFUN:
		fmt.Print("DEFUN")
	case CALL:
		fmt.Print("CALL\n\n")
	case IF:
		fmt.Print("IF")
	case CASE:
		fmt.Print("CASE")
	case NOT:
		fmt.Print("NOT\n\n")
	case AND:
		fmt.Print("AND\n\n")
	case OR:
		fmt.Print("OR\n\n")
	case EQ:
		fmt.Print("EQ\n\n")
	case ATOMP:
		fmt.Print("ATOMP\n\n")
	case PAIRP:
		fmt.Print("PAIRP\n\n")
	case LISTP:
		fmt.Print("LISTP\n\n")
	case NILP:
		fmt.Print("NILP\n\n")
	case GT:
		fmt.Print("GT\n\n")
	case GE:
		fmt.Print("GE\n\n")
	case LT:
		fmt.Print("LT\n\n")
	case LE:
		fmt.Print("LE\n\n")
	case PRINT:
		fmt.Print("PRINT\n\n")
	case PRINC:
		fmt.Print("PRINC\n\n")
	case JMP:
		fmt.Print("JMP")
	case LJMP:
		fmt.Print("LJMP")
	case LABEL:
		fmt.Print("LABEL")
	case CONS:
		fmt.Print("CONS\n\n")
	case CAR:
		fmt.Print("CAR")
	case CDR:
		fmt.Print("CDR")
	case LAMBDA:
		fmt.Print("LAMBDA")
	case FUNCALL:
		fmt.Print("FUNCALL\n\n")
	case APPEND:
		fmt.Print("APPEND")
	case PROGN:
		fmt.Print("PROGN")
	case LET:
		fmt.Print("LET")
	case PUSH_LEXICAL_FRAME:
		fmt.Print("PUSH_LEXICAL_FRAME\n\n")
	case POP_LEXICAL_FRAME:
		fmt.Print("POP_LEXICAL_FRAME\n\n")
	case EXEC_FUNC:
		fmt.Print("EXEC_FUNC")
	case TIME_START:
		fmt.Print("TIME_START\n\n")
	case TIME_END:
		fmt.Print("TIME_END\n\n")
	case PROGN_START:
		fmt.Print("PROGN_START\n\n")
	case PROGN_END:
		fmt.Print("PROGN_END\n\n")
	case FUNCTION:
		fmt.Print("FUNCTION\n\n")
	case ERROR:
		fmt.Print("ERROR\n\n")
	case FORMAT:
		fmt.Print("FORMAT")
	case COMMAND:
		fmt.Print("COMMAND")
	case LENGTH:
		fmt.Print("LENGTH\n\n")
	case MEMBER:
		fmt.Print("MEMBER\n\n")
	case DEFSERVER:
		fmt.Print("DEFSERVER\n\n")
	case DEFHANDLER:
		fmt.Print("DEFHANDLER\n\n")
	case RUNSERVER:
		fmt.Print("RUNSERVER\n\n")
	case GETQUERY:
		fmt.Print("GETQUERY\n\n")
	case SETSTATUS:
		fmt.Print("SETSTATUS\n\n")
	case JSON_PARSE:
		fmt.Print("JSON_PARSE\n\n")
	case MAKE_VECTOR:
		fmt.Print("MAKE_VECTOR")
	case MAKE_HASH:
		fmt.Print("MAKE_HASH\n\n")
	case SET_HASH:
		fmt.Print("SET_HASH\n\n")
	case GET_HASH:
		fmt.Print("GET_HASH\n\n")
	case AREF:
		fmt.Print("AREF\n\n")
	case NTH:
		fmt.Print("NTH\n\n")
	case ASSOC:
		fmt.Print("ASSOC\n\n")
	case RANDOM:
		fmt.Print("RANDOM\n\n")
	case OPEN:
		fmt.Print("OPEN\n\n")
	case WRITE:
		fmt.Print("WRITE\n\n")
	case CLOSE:
		fmt.Print("CLOSE\n\n")
	case READ_FILE:
		fmt.Print("READ_FILE\n\n")
	case SPLIT:
		fmt.Print("SPLIT\n\n")
	case MAKE_REQUEST:
		fmt.Print("MAKE_REQUEST")
	case DO_REQUEST:
		fmt.Print("DO_REQUEST\n\n")
	case ADD_REQUEST_HEADER:
		fmt.Print("ADD_REQUEST_HEADER\n\n")
	case VECTOR_PUSH:
		fmt.Print("VECTOR_PUSH\n\n")
	case VECTOR_POP:
		fmt.Print("VECTOR_POP\n\n")
	case VECTOR_LEN:
		fmt.Print("VECTOR_LEN\n\n")
	case LIST_FUNCTION:
		fmt.Print("LIST_FUNCTION")
	case EVENP:
		fmt.Print("EVENP\n\n")
	case ODDP:
		fmt.Print("ODDP\n\n")
	case MAPCAR:
		fmt.Print("MAPCAR")
	case REGEXP_REPLACE_ALL:
		fmt.Print("REGEXP_REPLACE_ALL\n\n")
	case REGEXP_MATCH:
		fmt.Print("REGEXP_MATCH\n\n")
	case SUBSTRING:
		fmt.Print("SUBSTRING\n\n")
	case GLOBAL:
		fmt.Print("GLOBAL\n\n")
	case STR_TO_LIST:
		fmt.Print("STR_TO_LIST\n\n")
	case INTERN:
		fmt.Print("INTERN\n\n")
	case STRINGP:
		fmt.Print("STRINGP\n\n")
	case TYPE:
		fmt.Print("TYPE\n\n")
	default:
		fmt.Print(" ")
		fmt.Println(byte_code)
		fmt.Print("\n")
	}
}
