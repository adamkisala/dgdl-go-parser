// Generated from DGDL.g4 by ANTLR 4.7.

package parser // DGDL

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DGDLListener is a complete listener for a parse tree produced by DGDLParser.
type DGDLListener interface {
	antlr.ParseTreeListener

	// EnterGame is called when entering the game production.
	EnterGame(c *GameContext)

	// EnterStore is called when entering the store production.
	EnterStore(c *StoreContext)

	// EnterTurns is called when entering the turns production.
	EnterTurns(c *TurnsContext)

	// EnterPlayers is called when entering the players production.
	EnterPlayers(c *PlayersContext)

	// EnterPlayer is called when entering the player production.
	EnterPlayer(c *PlayerContext)

	// EnterRoles is called when entering the roles production.
	EnterRoles(c *RolesContext)

	// EnterPrinciples is called when entering the principles production.
	EnterPrinciples(c *PrinciplesContext)

	// EnterPrinciple is called when entering the principle production.
	EnterPrinciple(c *PrincipleContext)

	// EnterScope is called when entering the scope production.
	EnterScope(c *ScopeContext)

	// EnterMoves is called when entering the moves production.
	EnterMoves(c *MovesContext)

	// EnterMove is called when entering the move production.
	EnterMove(c *MoveContext)

	// EnterContent is called when entering the content production.
	EnterContent(c *ContentContext)

	// EnterConditions is called when entering the conditions production.
	EnterConditions(c *ConditionsContext)

	// EnterEffects is called when entering the effects production.
	EnterEffects(c *EffectsContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// ExitGame is called when exiting the game production.
	ExitGame(c *GameContext)

	// ExitStore is called when exiting the store production.
	ExitStore(c *StoreContext)

	// ExitTurns is called when exiting the turns production.
	ExitTurns(c *TurnsContext)

	// ExitPlayers is called when exiting the players production.
	ExitPlayers(c *PlayersContext)

	// ExitPlayer is called when exiting the player production.
	ExitPlayer(c *PlayerContext)

	// ExitRoles is called when exiting the roles production.
	ExitRoles(c *RolesContext)

	// ExitPrinciples is called when exiting the principles production.
	ExitPrinciples(c *PrinciplesContext)

	// ExitPrinciple is called when exiting the principle production.
	ExitPrinciple(c *PrincipleContext)

	// ExitScope is called when exiting the scope production.
	ExitScope(c *ScopeContext)

	// ExitMoves is called when exiting the moves production.
	ExitMoves(c *MovesContext)

	// ExitMove is called when exiting the move production.
	ExitMove(c *MoveContext)

	// ExitContent is called when exiting the content production.
	ExitContent(c *ContentContext)

	// ExitConditions is called when exiting the conditions production.
	ExitConditions(c *ConditionsContext)

	// ExitEffects is called when exiting the effects production.
	ExitEffects(c *EffectsContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)
}
