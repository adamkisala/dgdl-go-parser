// Generated from DGDL.g4 by ANTLR 4.7.

package parser // DGDL

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 263,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 42, 10, 2, 13, 2, 14, 2, 43, 3,
	2, 3, 2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 2, 3, 2, 7, 2, 55,
	10, 2, 12, 2, 14, 2, 58, 11, 2, 3, 2, 3, 2, 7, 2, 62, 10, 2, 12, 2, 14,
	2, 65, 11, 2, 3, 2, 3, 2, 6, 2, 69, 10, 2, 13, 2, 14, 2, 70, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3,
	86, 10, 3, 12, 3, 14, 3, 89, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 115, 10, 4, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 128, 10, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 139, 10, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 148, 10, 7, 12, 7, 14, 7, 151,
	11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 160, 10, 8, 12,
	8, 14, 8, 163, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 172,
	10, 9, 3, 9, 3, 9, 5, 9, 176, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 189, 10, 11, 12, 11, 14,
	11, 192, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12,
	201, 10, 12, 3, 12, 3, 12, 5, 12, 205, 10, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 7, 13, 214, 10, 13, 12, 13, 14, 13, 217, 11, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 226, 10, 14, 12,
	14, 14, 14, 229, 11, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	7, 15, 238, 10, 15, 12, 15, 14, 15, 241, 11, 15, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 7, 16, 250, 10, 16, 12, 16, 14, 16, 253, 11, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 261, 10, 17, 3, 17, 2,
	2, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 7,
	3, 2, 7, 9, 3, 2, 10, 11, 3, 2, 12, 13, 3, 2, 14, 15, 3, 2, 16, 18, 2,
	267, 2, 34, 3, 2, 2, 2, 4, 74, 3, 2, 2, 2, 6, 101, 3, 2, 2, 2, 8, 118,
	3, 2, 2, 2, 10, 131, 3, 2, 2, 2, 12, 142, 3, 2, 2, 2, 14, 154, 3, 2, 2,
	2, 16, 166, 3, 2, 2, 2, 18, 179, 3, 2, 2, 2, 20, 183, 3, 2, 2, 2, 22, 195,
	3, 2, 2, 2, 24, 208, 3, 2, 2, 2, 26, 220, 3, 2, 2, 2, 28, 232, 3, 2, 2,
	2, 30, 244, 3, 2, 2, 2, 32, 260, 3, 2, 2, 2, 34, 35, 7, 40, 2, 2, 35, 36,
	7, 3, 2, 2, 36, 37, 5, 6, 4, 2, 37, 38, 7, 4, 2, 2, 38, 41, 5, 8, 5, 2,
	39, 40, 7, 4, 2, 2, 40, 42, 5, 10, 6, 2, 41, 39, 3, 2, 2, 2, 42, 43, 3,
	2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 49, 3, 2, 2, 2, 45,
	46, 7, 4, 2, 2, 46, 48, 5, 4, 3, 2, 47, 45, 3, 2, 2, 2, 48, 51, 3, 2, 2,
	2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 56, 3, 2, 2, 2, 51, 49,
	3, 2, 2, 2, 52, 53, 7, 4, 2, 2, 53, 55, 5, 12, 7, 2, 54, 52, 3, 2, 2, 2,
	55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 63, 3,
	2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 60, 7, 4, 2, 2, 60, 62, 5, 14, 8, 2, 61,
	59, 3, 2, 2, 2, 62, 65, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2,
	2, 64, 68, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 67, 7, 4, 2, 2, 67, 69,
	5, 20, 11, 2, 68, 66, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 68, 3, 2, 2,
	2, 70, 71, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 7, 5, 2, 2, 73, 3, 3,
	2, 2, 2, 74, 75, 7, 36, 2, 2, 75, 76, 7, 3, 2, 2, 76, 77, 7, 28, 2, 2,
	77, 78, 7, 6, 2, 2, 78, 79, 7, 40, 2, 2, 79, 80, 7, 4, 2, 2, 80, 81, 7,
	30, 2, 2, 81, 82, 7, 3, 2, 2, 82, 87, 7, 40, 2, 2, 83, 84, 7, 4, 2, 2,
	84, 86, 7, 40, 2, 2, 85, 83, 3, 2, 2, 2, 86, 89, 3, 2, 2, 2, 87, 85, 3,
	2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 90, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 90,
	91, 7, 5, 2, 2, 91, 92, 7, 4, 2, 2, 92, 93, 7, 37, 2, 2, 93, 94, 7, 6,
	2, 2, 94, 95, 9, 2, 2, 2, 95, 96, 7, 4, 2, 2, 96, 97, 7, 39, 2, 2, 97,
	98, 7, 6, 2, 2, 98, 99, 9, 3, 2, 2, 99, 100, 7, 5, 2, 2, 100, 5, 3, 2,
	2, 2, 101, 102, 7, 38, 2, 2, 102, 103, 7, 3, 2, 2, 103, 104, 7, 24, 2,
	2, 104, 105, 7, 6, 2, 2, 105, 106, 9, 4, 2, 2, 106, 107, 7, 4, 2, 2, 107,
	108, 7, 29, 2, 2, 108, 109, 7, 6, 2, 2, 109, 114, 9, 5, 2, 2, 110, 111,
	7, 4, 2, 2, 111, 112, 7, 25, 2, 2, 112, 113, 7, 6, 2, 2, 113, 115, 7, 41,
	2, 2, 114, 110, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2,
	116, 117, 7, 5, 2, 2, 117, 7, 3, 2, 2, 2, 118, 119, 7, 32, 2, 2, 119, 120,
	7, 3, 2, 2, 120, 121, 7, 26, 2, 2, 121, 122, 7, 6, 2, 2, 122, 127, 7, 41,
	2, 2, 123, 124, 7, 4, 2, 2, 124, 125, 7, 25, 2, 2, 125, 126, 7, 6, 2, 2,
	126, 128, 7, 41, 2, 2, 127, 123, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128,
	129, 3, 2, 2, 2, 129, 130, 7, 5, 2, 2, 130, 9, 3, 2, 2, 2, 131, 132, 7,
	31, 2, 2, 132, 133, 7, 3, 2, 2, 133, 134, 7, 28, 2, 2, 134, 135, 7, 6,
	2, 2, 135, 138, 7, 40, 2, 2, 136, 137, 7, 4, 2, 2, 137, 139, 5, 12, 7,
	2, 138, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140,
	141, 7, 5, 2, 2, 141, 11, 3, 2, 2, 2, 142, 143, 7, 34, 2, 2, 143, 144,
	7, 3, 2, 2, 144, 149, 7, 40, 2, 2, 145, 146, 7, 4, 2, 2, 146, 148, 7, 40,
	2, 2, 147, 145, 3, 2, 2, 2, 148, 151, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2,
	149, 150, 3, 2, 2, 2, 150, 152, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 152,
	153, 7, 5, 2, 2, 153, 13, 3, 2, 2, 2, 154, 155, 7, 33, 2, 2, 155, 156,
	7, 3, 2, 2, 156, 161, 5, 16, 9, 2, 157, 158, 7, 4, 2, 2, 158, 160, 5, 16,
	9, 2, 159, 157, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2,
	161, 162, 3, 2, 2, 2, 162, 164, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 164,
	165, 7, 5, 2, 2, 165, 15, 3, 2, 2, 2, 166, 167, 7, 40, 2, 2, 167, 168,
	7, 3, 2, 2, 168, 171, 5, 18, 10, 2, 169, 170, 7, 4, 2, 2, 170, 172, 5,
	26, 14, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 175, 3, 2,
	2, 2, 173, 174, 7, 4, 2, 2, 174, 176, 5, 28, 15, 2, 175, 173, 3, 2, 2,
	2, 175, 176, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 178, 7, 5, 2, 2, 178,
	17, 3, 2, 2, 2, 179, 180, 7, 35, 2, 2, 180, 181, 7, 6, 2, 2, 181, 182,
	9, 6, 2, 2, 182, 19, 3, 2, 2, 2, 183, 184, 7, 27, 2, 2, 184, 185, 7, 3,
	2, 2, 185, 190, 5, 22, 12, 2, 186, 187, 7, 4, 2, 2, 187, 189, 5, 22, 12,
	2, 188, 186, 3, 2, 2, 2, 189, 192, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190,
	191, 3, 2, 2, 2, 191, 193, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 193, 194,
	7, 5, 2, 2, 194, 21, 3, 2, 2, 2, 195, 196, 7, 40, 2, 2, 196, 197, 7, 3,
	2, 2, 197, 200, 5, 24, 13, 2, 198, 199, 7, 4, 2, 2, 199, 201, 5, 26, 14,
	2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2, 202,
	203, 7, 4, 2, 2, 203, 205, 5, 28, 15, 2, 204, 202, 3, 2, 2, 2, 204, 205,
	3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 7, 5, 2, 2, 207, 23, 3, 2,
	2, 2, 208, 209, 7, 22, 2, 2, 209, 210, 7, 3, 2, 2, 210, 215, 7, 40, 2,
	2, 211, 212, 7, 4, 2, 2, 212, 214, 7, 40, 2, 2, 213, 211, 3, 2, 2, 2, 214,
	217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 218,
	3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 219, 7, 5, 2, 2, 219, 25, 3, 2,
	2, 2, 220, 221, 7, 21, 2, 2, 221, 222, 7, 3, 2, 2, 222, 227, 5, 30, 16,
	2, 223, 224, 7, 4, 2, 2, 224, 226, 5, 30, 16, 2, 225, 223, 3, 2, 2, 2,
	226, 229, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228,
	230, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 231, 7, 5, 2, 2, 231, 27, 3,
	2, 2, 2, 232, 233, 7, 23, 2, 2, 233, 234, 7, 3, 2, 2, 234, 239, 5, 30,
	16, 2, 235, 236, 7, 4, 2, 2, 236, 238, 5, 30, 16, 2, 237, 235, 3, 2, 2,
	2, 238, 241, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240,
	242, 3, 2, 2, 2, 241, 239, 3, 2, 2, 2, 242, 243, 7, 5, 2, 2, 243, 29, 3,
	2, 2, 2, 244, 245, 7, 40, 2, 2, 245, 246, 7, 19, 2, 2, 246, 251, 5, 32,
	17, 2, 247, 248, 7, 4, 2, 2, 248, 250, 5, 32, 17, 2, 249, 247, 3, 2, 2,
	2, 250, 253, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252,
	254, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 254, 255, 7, 20, 2, 2, 255, 31,
	3, 2, 2, 2, 256, 257, 7, 40, 2, 2, 257, 258, 7, 6, 2, 2, 258, 261, 7, 40,
	2, 2, 259, 261, 7, 40, 2, 2, 260, 256, 3, 2, 2, 2, 260, 259, 3, 2, 2, 2,
	261, 33, 3, 2, 2, 2, 23, 43, 49, 56, 63, 70, 87, 114, 127, 138, 149, 161,
	171, 175, 190, 200, 204, 215, 227, 239, 251, 260,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "','", "'}'", "':'", "'set'", "'queue'", "'stack'", "'public'",
	"'private'", "'single'", "'multiple'", "'strict'", "'liberal'", "'initial'",
	"'turnwise'", "'movewise'", "'('", "')'", "'conditions'", "'content'",
	"'effects'", "'magnitude'", "'max'", "'min'", "'moves'", "'name'", "'ordering'",
	"'owner'", "'player'", "'players'", "'principles'", "'roles'", "'scope'",
	"'store'", "'structure'", "'turns'", "'visibility'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "CONDITIONS", "CONTENT", "EFFECTS", "MAGNITUDE", "MAX", "MIN", "MOVES",
	"NAME", "ORDERING", "OWNER", "PLAYER", "PLAYERS", "PRINCIPLES", "ROLES",
	"SCOPE", "STORE", "STRUCTURE", "TURNS", "VISIBILITY", "IDENT", "INT", "UCHAR",
	"CHAR", "CHARS", "WS",
}

var ruleNames = []string{
	"game", "store", "turns", "players", "player", "roles", "principles", "principle",
	"scope", "moves", "move", "content", "conditions", "effects", "expr", "param",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type DGDLParser struct {
	*antlr.BaseParser
}

func NewDGDLParser(input antlr.TokenStream) *DGDLParser {
	this := new(DGDLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DGDL.g4"

	return this
}

// DGDLParser tokens.
const (
	DGDLParserEOF        = antlr.TokenEOF
	DGDLParserT__0       = 1
	DGDLParserT__1       = 2
	DGDLParserT__2       = 3
	DGDLParserT__3       = 4
	DGDLParserT__4       = 5
	DGDLParserT__5       = 6
	DGDLParserT__6       = 7
	DGDLParserT__7       = 8
	DGDLParserT__8       = 9
	DGDLParserT__9       = 10
	DGDLParserT__10      = 11
	DGDLParserT__11      = 12
	DGDLParserT__12      = 13
	DGDLParserT__13      = 14
	DGDLParserT__14      = 15
	DGDLParserT__15      = 16
	DGDLParserT__16      = 17
	DGDLParserT__17      = 18
	DGDLParserCONDITIONS = 19
	DGDLParserCONTENT    = 20
	DGDLParserEFFECTS    = 21
	DGDLParserMAGNITUDE  = 22
	DGDLParserMAX        = 23
	DGDLParserMIN        = 24
	DGDLParserMOVES      = 25
	DGDLParserNAME       = 26
	DGDLParserORDERING   = 27
	DGDLParserOWNER      = 28
	DGDLParserPLAYER     = 29
	DGDLParserPLAYERS    = 30
	DGDLParserPRINCIPLES = 31
	DGDLParserROLES      = 32
	DGDLParserSCOPE      = 33
	DGDLParserSTORE      = 34
	DGDLParserSTRUCTURE  = 35
	DGDLParserTURNS      = 36
	DGDLParserVISIBILITY = 37
	DGDLParserIDENT      = 38
	DGDLParserINT        = 39
	DGDLParserUCHAR      = 40
	DGDLParserCHAR       = 41
	DGDLParserCHARS      = 42
	DGDLParserWS         = 43
)

// DGDLParser rules.
const (
	DGDLParserRULE_game       = 0
	DGDLParserRULE_store      = 1
	DGDLParserRULE_turns      = 2
	DGDLParserRULE_players    = 3
	DGDLParserRULE_player     = 4
	DGDLParserRULE_roles      = 5
	DGDLParserRULE_principles = 6
	DGDLParserRULE_principle  = 7
	DGDLParserRULE_scope      = 8
	DGDLParserRULE_moves      = 9
	DGDLParserRULE_move       = 10
	DGDLParserRULE_content    = 11
	DGDLParserRULE_conditions = 12
	DGDLParserRULE_effects    = 13
	DGDLParserRULE_expr       = 14
	DGDLParserRULE_param      = 15
)

// IGameContext is an interface to support dynamic dispatch.
type IGameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGameContext differentiates from other interfaces.
	IsGameContext()
}

type GameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGameContext() *GameContext {
	var p = new(GameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_game
	return p
}

func (*GameContext) IsGameContext() {}

func NewGameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GameContext {
	var p = new(GameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_game

	return p
}

func (s *GameContext) GetParser() antlr.Parser { return s.parser }

func (s *GameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DGDLParserIDENT, 0)
}

func (s *GameContext) Turns() ITurnsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITurnsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITurnsContext)
}

func (s *GameContext) Players() IPlayersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlayersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlayersContext)
}

func (s *GameContext) AllPlayer() []IPlayerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPlayerContext)(nil)).Elem())
	var tst = make([]IPlayerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPlayerContext)
		}
	}

	return tst
}

func (s *GameContext) Player(i int) IPlayerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlayerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPlayerContext)
}

func (s *GameContext) AllStore() []IStoreContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStoreContext)(nil)).Elem())
	var tst = make([]IStoreContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStoreContext)
		}
	}

	return tst
}

func (s *GameContext) Store(i int) IStoreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStoreContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStoreContext)
}

func (s *GameContext) AllRoles() []IRolesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRolesContext)(nil)).Elem())
	var tst = make([]IRolesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRolesContext)
		}
	}

	return tst
}

func (s *GameContext) Roles(i int) IRolesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRolesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRolesContext)
}

func (s *GameContext) AllPrinciples() []IPrinciplesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrinciplesContext)(nil)).Elem())
	var tst = make([]IPrinciplesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrinciplesContext)
		}
	}

	return tst
}

func (s *GameContext) Principles(i int) IPrinciplesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrinciplesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrinciplesContext)
}

func (s *GameContext) AllMoves() []IMovesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMovesContext)(nil)).Elem())
	var tst = make([]IMovesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMovesContext)
		}
	}

	return tst
}

func (s *GameContext) Moves(i int) IMovesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMovesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMovesContext)
}

func (s *GameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterGame(s)
	}
}

func (s *GameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitGame(s)
	}
}

func (s *GameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitGame(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Game() (localctx IGameContext) {
	localctx = NewGameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DGDLParserRULE_game)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(DGDLParserIDENT)
	}
	{
		p.SetState(33)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(34)
		p.Turns()
	}
	{
		p.SetState(35)
		p.Match(DGDLParserT__1)
	}
	{
		p.SetState(36)
		p.Players()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(37)
				p.Match(DGDLParserT__1)
			}
			{
				p.SetState(38)
				p.Player()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(43)
				p.Match(DGDLParserT__1)
			}
			{
				p.SetState(44)
				p.Store()
			}

		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(50)
				p.Match(DGDLParserT__1)
			}
			{
				p.SetState(51)
				p.Roles()
			}

		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(57)
				p.Match(DGDLParserT__1)
			}
			{
				p.SetState(58)
				p.Principles()
			}

		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DGDLParserT__1 {
		{
			p.SetState(64)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(65)
			p.Moves()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IStoreContext is an interface to support dynamic dispatch.
type IStoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStoreContext differentiates from other interfaces.
	IsStoreContext()
}

type StoreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStoreContext() *StoreContext {
	var p = new(StoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_store
	return p
}

func (*StoreContext) IsStoreContext() {}

func NewStoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StoreContext {
	var p = new(StoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_store

	return p
}

func (s *StoreContext) GetParser() antlr.Parser { return s.parser }

func (s *StoreContext) STORE() antlr.TerminalNode {
	return s.GetToken(DGDLParserSTORE, 0)
}

func (s *StoreContext) NAME() antlr.TerminalNode {
	return s.GetToken(DGDLParserNAME, 0)
}

func (s *StoreContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(DGDLParserIDENT)
}

func (s *StoreContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(DGDLParserIDENT, i)
}

func (s *StoreContext) OWNER() antlr.TerminalNode {
	return s.GetToken(DGDLParserOWNER, 0)
}

func (s *StoreContext) STRUCTURE() antlr.TerminalNode {
	return s.GetToken(DGDLParserSTRUCTURE, 0)
}

func (s *StoreContext) VISIBILITY() antlr.TerminalNode {
	return s.GetToken(DGDLParserVISIBILITY, 0)
}

func (s *StoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterStore(s)
	}
}

func (s *StoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitStore(s)
	}
}

func (s *StoreContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitStore(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Store() (localctx IStoreContext) {
	localctx = NewStoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DGDLParserRULE_store)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(DGDLParserSTORE)
	}
	{
		p.SetState(73)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(74)
		p.Match(DGDLParserNAME)
	}
	{
		p.SetState(75)
		p.Match(DGDLParserT__3)
	}
	{
		p.SetState(76)
		p.Match(DGDLParserIDENT)
	}
	{
		p.SetState(77)
		p.Match(DGDLParserT__1)
	}
	{
		p.SetState(78)
		p.Match(DGDLParserOWNER)
	}
	{
		p.SetState(79)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(80)
		p.Match(DGDLParserIDENT)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DGDLParserT__1 {
		{
			p.SetState(81)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(82)
			p.Match(DGDLParserIDENT)
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(88)
		p.Match(DGDLParserT__2)
	}
	{
		p.SetState(89)
		p.Match(DGDLParserT__1)
	}
	{
		p.SetState(90)
		p.Match(DGDLParserSTRUCTURE)
	}
	{
		p.SetState(91)
		p.Match(DGDLParserT__3)
	}
	p.SetState(92)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DGDLParserT__4)|(1<<DGDLParserT__5)|(1<<DGDLParserT__6))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(93)
		p.Match(DGDLParserT__1)
	}
	{
		p.SetState(94)
		p.Match(DGDLParserVISIBILITY)
	}
	{
		p.SetState(95)
		p.Match(DGDLParserT__3)
	}
	p.SetState(96)
	_la = p.GetTokenStream().LA(1)

	if !(_la == DGDLParserT__7 || _la == DGDLParserT__8) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(97)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// ITurnsContext is an interface to support dynamic dispatch.
type ITurnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTurnsContext differentiates from other interfaces.
	IsTurnsContext()
}

type TurnsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTurnsContext() *TurnsContext {
	var p = new(TurnsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_turns
	return p
}

func (*TurnsContext) IsTurnsContext() {}

func NewTurnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TurnsContext {
	var p = new(TurnsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_turns

	return p
}

func (s *TurnsContext) GetParser() antlr.Parser { return s.parser }

func (s *TurnsContext) TURNS() antlr.TerminalNode {
	return s.GetToken(DGDLParserTURNS, 0)
}

func (s *TurnsContext) MAGNITUDE() antlr.TerminalNode {
	return s.GetToken(DGDLParserMAGNITUDE, 0)
}

func (s *TurnsContext) ORDERING() antlr.TerminalNode {
	return s.GetToken(DGDLParserORDERING, 0)
}

func (s *TurnsContext) MAX() antlr.TerminalNode {
	return s.GetToken(DGDLParserMAX, 0)
}

func (s *TurnsContext) INT() antlr.TerminalNode {
	return s.GetToken(DGDLParserINT, 0)
}

func (s *TurnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TurnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TurnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterTurns(s)
	}
}

func (s *TurnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitTurns(s)
	}
}

func (s *TurnsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitTurns(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Turns() (localctx ITurnsContext) {
	localctx = NewTurnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DGDLParserRULE_turns)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(DGDLParserTURNS)
	}
	{
		p.SetState(100)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(101)
		p.Match(DGDLParserMAGNITUDE)
	}
	{
		p.SetState(102)
		p.Match(DGDLParserT__3)
	}
	p.SetState(103)
	_la = p.GetTokenStream().LA(1)

	if !(_la == DGDLParserT__9 || _la == DGDLParserT__10) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(104)
		p.Match(DGDLParserT__1)
	}
	{
		p.SetState(105)
		p.Match(DGDLParserORDERING)
	}
	{
		p.SetState(106)
		p.Match(DGDLParserT__3)
	}
	p.SetState(107)
	_la = p.GetTokenStream().LA(1)

	if !(_la == DGDLParserT__11 || _la == DGDLParserT__12) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DGDLParserT__1 {
		{
			p.SetState(108)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(109)
			p.Match(DGDLParserMAX)
		}
		{
			p.SetState(110)
			p.Match(DGDLParserT__3)
		}
		{
			p.SetState(111)
			p.Match(DGDLParserINT)
		}

	}
	{
		p.SetState(114)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IPlayersContext is an interface to support dynamic dispatch.
type IPlayersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlayersContext differentiates from other interfaces.
	IsPlayersContext()
}

type PlayersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlayersContext() *PlayersContext {
	var p = new(PlayersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_players
	return p
}

func (*PlayersContext) IsPlayersContext() {}

func NewPlayersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlayersContext {
	var p = new(PlayersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_players

	return p
}

func (s *PlayersContext) GetParser() antlr.Parser { return s.parser }

func (s *PlayersContext) PLAYERS() antlr.TerminalNode {
	return s.GetToken(DGDLParserPLAYERS, 0)
}

func (s *PlayersContext) MIN() antlr.TerminalNode {
	return s.GetToken(DGDLParserMIN, 0)
}

func (s *PlayersContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(DGDLParserINT)
}

func (s *PlayersContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(DGDLParserINT, i)
}

func (s *PlayersContext) MAX() antlr.TerminalNode {
	return s.GetToken(DGDLParserMAX, 0)
}

func (s *PlayersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlayersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlayersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterPlayers(s)
	}
}

func (s *PlayersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitPlayers(s)
	}
}

func (s *PlayersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitPlayers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Players() (localctx IPlayersContext) {
	localctx = NewPlayersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DGDLParserRULE_players)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(DGDLParserPLAYERS)
	}
	{
		p.SetState(117)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(118)
		p.Match(DGDLParserMIN)
	}
	{
		p.SetState(119)
		p.Match(DGDLParserT__3)
	}
	{
		p.SetState(120)
		p.Match(DGDLParserINT)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DGDLParserT__1 {
		{
			p.SetState(121)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(122)
			p.Match(DGDLParserMAX)
		}
		{
			p.SetState(123)
			p.Match(DGDLParserT__3)
		}
		{
			p.SetState(124)
			p.Match(DGDLParserINT)
		}

	}
	{
		p.SetState(127)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IPlayerContext is an interface to support dynamic dispatch.
type IPlayerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlayerContext differentiates from other interfaces.
	IsPlayerContext()
}

type PlayerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlayerContext() *PlayerContext {
	var p = new(PlayerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_player
	return p
}

func (*PlayerContext) IsPlayerContext() {}

func NewPlayerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlayerContext {
	var p = new(PlayerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_player

	return p
}

func (s *PlayerContext) GetParser() antlr.Parser { return s.parser }

func (s *PlayerContext) PLAYER() antlr.TerminalNode {
	return s.GetToken(DGDLParserPLAYER, 0)
}

func (s *PlayerContext) NAME() antlr.TerminalNode {
	return s.GetToken(DGDLParserNAME, 0)
}

func (s *PlayerContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DGDLParserIDENT, 0)
}

func (s *PlayerContext) Roles() IRolesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRolesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRolesContext)
}

func (s *PlayerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlayerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlayerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterPlayer(s)
	}
}

func (s *PlayerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitPlayer(s)
	}
}

func (s *PlayerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitPlayer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Player() (localctx IPlayerContext) {
	localctx = NewPlayerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DGDLParserRULE_player)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Match(DGDLParserPLAYER)
	}
	{
		p.SetState(130)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(131)
		p.Match(DGDLParserNAME)
	}
	{
		p.SetState(132)
		p.Match(DGDLParserT__3)
	}
	{
		p.SetState(133)
		p.Match(DGDLParserIDENT)
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DGDLParserT__1 {
		{
			p.SetState(134)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(135)
			p.Roles()
		}

	}
	{
		p.SetState(138)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IRolesContext is an interface to support dynamic dispatch.
type IRolesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRolesContext differentiates from other interfaces.
	IsRolesContext()
}

type RolesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRolesContext() *RolesContext {
	var p = new(RolesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_roles
	return p
}

func (*RolesContext) IsRolesContext() {}

func NewRolesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RolesContext {
	var p = new(RolesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_roles

	return p
}

func (s *RolesContext) GetParser() antlr.Parser { return s.parser }

func (s *RolesContext) ROLES() antlr.TerminalNode {
	return s.GetToken(DGDLParserROLES, 0)
}

func (s *RolesContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(DGDLParserIDENT)
}

func (s *RolesContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(DGDLParserIDENT, i)
}

func (s *RolesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RolesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RolesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterRoles(s)
	}
}

func (s *RolesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitRoles(s)
	}
}

func (s *RolesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitRoles(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Roles() (localctx IRolesContext) {
	localctx = NewRolesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DGDLParserRULE_roles)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(DGDLParserROLES)
	}
	{
		p.SetState(141)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(142)
		p.Match(DGDLParserIDENT)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DGDLParserT__1 {
		{
			p.SetState(143)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(144)
			p.Match(DGDLParserIDENT)
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IPrinciplesContext is an interface to support dynamic dispatch.
type IPrinciplesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrinciplesContext differentiates from other interfaces.
	IsPrinciplesContext()
}

type PrinciplesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrinciplesContext() *PrinciplesContext {
	var p = new(PrinciplesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_principles
	return p
}

func (*PrinciplesContext) IsPrinciplesContext() {}

func NewPrinciplesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrinciplesContext {
	var p = new(PrinciplesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_principles

	return p
}

func (s *PrinciplesContext) GetParser() antlr.Parser { return s.parser }

func (s *PrinciplesContext) PRINCIPLES() antlr.TerminalNode {
	return s.GetToken(DGDLParserPRINCIPLES, 0)
}

func (s *PrinciplesContext) AllPrinciple() []IPrincipleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrincipleContext)(nil)).Elem())
	var tst = make([]IPrincipleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrincipleContext)
		}
	}

	return tst
}

func (s *PrinciplesContext) Principle(i int) IPrincipleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrincipleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrincipleContext)
}

func (s *PrinciplesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrinciplesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrinciplesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterPrinciples(s)
	}
}

func (s *PrinciplesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitPrinciples(s)
	}
}

func (s *PrinciplesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitPrinciples(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Principles() (localctx IPrinciplesContext) {
	localctx = NewPrinciplesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DGDLParserRULE_principles)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(DGDLParserPRINCIPLES)
	}
	{
		p.SetState(153)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(154)
		p.Principle()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DGDLParserT__1 {
		{
			p.SetState(155)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(156)
			p.Principle()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(162)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IPrincipleContext is an interface to support dynamic dispatch.
type IPrincipleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrincipleContext differentiates from other interfaces.
	IsPrincipleContext()
}

type PrincipleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrincipleContext() *PrincipleContext {
	var p = new(PrincipleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_principle
	return p
}

func (*PrincipleContext) IsPrincipleContext() {}

func NewPrincipleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrincipleContext {
	var p = new(PrincipleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_principle

	return p
}

func (s *PrincipleContext) GetParser() antlr.Parser { return s.parser }

func (s *PrincipleContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DGDLParserIDENT, 0)
}

func (s *PrincipleContext) Scope() IScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *PrincipleContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *PrincipleContext) Effects() IEffectsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEffectsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEffectsContext)
}

func (s *PrincipleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrincipleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrincipleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterPrinciple(s)
	}
}

func (s *PrincipleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitPrinciple(s)
	}
}

func (s *PrincipleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitPrinciple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Principle() (localctx IPrincipleContext) {
	localctx = NewPrincipleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DGDLParserRULE_principle)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(DGDLParserIDENT)
	}
	{
		p.SetState(165)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(166)
		p.Scope()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(167)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(168)
			p.Conditions()
		}

	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DGDLParserT__1 {
		{
			p.SetState(171)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(172)
			p.Effects()
		}

	}
	{
		p.SetState(175)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_scope
	return p
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) SCOPE() antlr.TerminalNode {
	return s.GetToken(DGDLParserSCOPE, 0)
}

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterScope(s)
	}
}

func (s *ScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitScope(s)
	}
}

func (s *ScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Scope() (localctx IScopeContext) {
	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DGDLParserRULE_scope)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(DGDLParserSCOPE)
	}
	{
		p.SetState(178)
		p.Match(DGDLParserT__3)
	}
	p.SetState(179)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DGDLParserT__13)|(1<<DGDLParserT__14)|(1<<DGDLParserT__15))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IMovesContext is an interface to support dynamic dispatch.
type IMovesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMovesContext differentiates from other interfaces.
	IsMovesContext()
}

type MovesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMovesContext() *MovesContext {
	var p = new(MovesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_moves
	return p
}

func (*MovesContext) IsMovesContext() {}

func NewMovesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MovesContext {
	var p = new(MovesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_moves

	return p
}

func (s *MovesContext) GetParser() antlr.Parser { return s.parser }

func (s *MovesContext) MOVES() antlr.TerminalNode {
	return s.GetToken(DGDLParserMOVES, 0)
}

func (s *MovesContext) AllMove() []IMoveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMoveContext)(nil)).Elem())
	var tst = make([]IMoveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMoveContext)
		}
	}

	return tst
}

func (s *MovesContext) Move(i int) IMoveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMoveContext)
}

func (s *MovesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MovesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MovesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterMoves(s)
	}
}

func (s *MovesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitMoves(s)
	}
}

func (s *MovesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitMoves(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Moves() (localctx IMovesContext) {
	localctx = NewMovesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DGDLParserRULE_moves)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(DGDLParserMOVES)
	}
	{
		p.SetState(182)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(183)
		p.Move()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DGDLParserT__1 {
		{
			p.SetState(184)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(185)
			p.Move()
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(191)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IMoveContext is an interface to support dynamic dispatch.
type IMoveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoveContext differentiates from other interfaces.
	IsMoveContext()
}

type MoveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoveContext() *MoveContext {
	var p = new(MoveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_move
	return p
}

func (*MoveContext) IsMoveContext() {}

func NewMoveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoveContext {
	var p = new(MoveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_move

	return p
}

func (s *MoveContext) GetParser() antlr.Parser { return s.parser }

func (s *MoveContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DGDLParserIDENT, 0)
}

func (s *MoveContext) Content() IContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContentContext)
}

func (s *MoveContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *MoveContext) Effects() IEffectsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEffectsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEffectsContext)
}

func (s *MoveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterMove(s)
	}
}

func (s *MoveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitMove(s)
	}
}

func (s *MoveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitMove(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Move() (localctx IMoveContext) {
	localctx = NewMoveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DGDLParserRULE_move)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(DGDLParserIDENT)
	}
	{
		p.SetState(194)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(195)
		p.Content()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(196)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(197)
			p.Conditions()
		}

	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DGDLParserT__1 {
		{
			p.SetState(200)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(201)
			p.Effects()
		}

	}
	{
		p.SetState(204)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IContentContext is an interface to support dynamic dispatch.
type IContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContentContext differentiates from other interfaces.
	IsContentContext()
}

type ContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentContext() *ContentContext {
	var p = new(ContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_content
	return p
}

func (*ContentContext) IsContentContext() {}

func NewContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentContext {
	var p = new(ContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_content

	return p
}

func (s *ContentContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentContext) CONTENT() antlr.TerminalNode {
	return s.GetToken(DGDLParserCONTENT, 0)
}

func (s *ContentContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(DGDLParserIDENT)
}

func (s *ContentContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(DGDLParserIDENT, i)
}

func (s *ContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterContent(s)
	}
}

func (s *ContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitContent(s)
	}
}

func (s *ContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Content() (localctx IContentContext) {
	localctx = NewContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DGDLParserRULE_content)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(DGDLParserCONTENT)
	}
	{
		p.SetState(207)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(208)
		p.Match(DGDLParserIDENT)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DGDLParserT__1 {
		{
			p.SetState(209)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(210)
			p.Match(DGDLParserIDENT)
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_conditions
	return p
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) CONDITIONS() antlr.TerminalNode {
	return s.GetToken(DGDLParserCONDITIONS, 0)
}

func (s *ConditionsContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ConditionsContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitConditions(s)
	}
}

func (s *ConditionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitConditions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Conditions() (localctx IConditionsContext) {
	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DGDLParserRULE_conditions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(DGDLParserCONDITIONS)
	}
	{
		p.SetState(219)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(220)
		p.Expr()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DGDLParserT__1 {
		{
			p.SetState(221)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(222)
			p.Expr()
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(228)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IEffectsContext is an interface to support dynamic dispatch.
type IEffectsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEffectsContext differentiates from other interfaces.
	IsEffectsContext()
}

type EffectsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEffectsContext() *EffectsContext {
	var p = new(EffectsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_effects
	return p
}

func (*EffectsContext) IsEffectsContext() {}

func NewEffectsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EffectsContext {
	var p = new(EffectsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_effects

	return p
}

func (s *EffectsContext) GetParser() antlr.Parser { return s.parser }

func (s *EffectsContext) EFFECTS() antlr.TerminalNode {
	return s.GetToken(DGDLParserEFFECTS, 0)
}

func (s *EffectsContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EffectsContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EffectsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EffectsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EffectsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterEffects(s)
	}
}

func (s *EffectsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitEffects(s)
	}
}

func (s *EffectsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitEffects(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Effects() (localctx IEffectsContext) {
	localctx = NewEffectsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DGDLParserRULE_effects)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(DGDLParserEFFECTS)
	}
	{
		p.SetState(231)
		p.Match(DGDLParserT__0)
	}
	{
		p.SetState(232)
		p.Expr()
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DGDLParserT__1 {
		{
			p.SetState(233)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(234)
			p.Expr()
		}

		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(240)
		p.Match(DGDLParserT__2)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) IDENT() antlr.TerminalNode {
	return s.GetToken(DGDLParserIDENT, 0)
}

func (s *ExprContext) AllParam() []IParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamContext)(nil)).Elem())
	var tst = make([]IParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamContext)
		}
	}

	return tst
}

func (s *ExprContext) Param(i int) IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DGDLParserRULE_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(DGDLParserIDENT)
	}
	{
		p.SetState(243)
		p.Match(DGDLParserT__16)
	}
	{
		p.SetState(244)
		p.Param()
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DGDLParserT__1 {
		{
			p.SetState(245)
			p.Match(DGDLParserT__1)
		}
		{
			p.SetState(246)
			p.Param()
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(252)
		p.Match(DGDLParserT__17)
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DGDLParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DGDLParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(DGDLParserIDENT)
}

func (s *ParamContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(DGDLParserIDENT, i)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DGDLListener); ok {
		listenerT.ExitParam(s)
	}
}

func (s *ParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DGDLVisitor:
		return t.VisitParam(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DGDLParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DGDLParserRULE_param)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)
			p.Match(DGDLParserIDENT)
		}
		{
			p.SetState(255)
			p.Match(DGDLParserT__3)
		}
		{
			p.SetState(256)
			p.Match(DGDLParserIDENT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
			p.Match(DGDLParserIDENT)
		}

	}

	return localctx
}
