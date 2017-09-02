// Generated from DGDL.g4 by ANTLR 4.7.

package parser // DGDL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DGDLParser.
type DGDLVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DGDLParser#game.
	VisitGame(ctx *GameContext) interface{}

	// Visit a parse tree produced by DGDLParser#store.
	VisitStore(ctx *StoreContext) interface{}

	// Visit a parse tree produced by DGDLParser#turns.
	VisitTurns(ctx *TurnsContext) interface{}

	// Visit a parse tree produced by DGDLParser#players.
	VisitPlayers(ctx *PlayersContext) interface{}

	// Visit a parse tree produced by DGDLParser#player.
	VisitPlayer(ctx *PlayerContext) interface{}

	// Visit a parse tree produced by DGDLParser#roles.
	VisitRoles(ctx *RolesContext) interface{}

	// Visit a parse tree produced by DGDLParser#principles.
	VisitPrinciples(ctx *PrinciplesContext) interface{}

	// Visit a parse tree produced by DGDLParser#principle.
	VisitPrinciple(ctx *PrincipleContext) interface{}

	// Visit a parse tree produced by DGDLParser#scope.
	VisitScope(ctx *ScopeContext) interface{}

	// Visit a parse tree produced by DGDLParser#moves.
	VisitMoves(ctx *MovesContext) interface{}

	// Visit a parse tree produced by DGDLParser#move.
	VisitMove(ctx *MoveContext) interface{}

	// Visit a parse tree produced by DGDLParser#content.
	VisitContent(ctx *ContentContext) interface{}

	// Visit a parse tree produced by DGDLParser#conditions.
	VisitConditions(ctx *ConditionsContext) interface{}

	// Visit a parse tree produced by DGDLParser#effects.
	VisitEffects(ctx *EffectsContext) interface{}

	// Visit a parse tree produced by DGDLParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by DGDLParser#param.
	VisitParam(ctx *ParamContext) interface{}
}
