# dgdl-go-parser

Handy small package with [antlr4](http://www.antlr.org/index.html) ([licence](http://www.antlr.org/license.html)) generated `go` parser, lexer, listener and visitor for [DGDL](http://siwells.github.io/DGDL/) grammar, that allows to override either listener or visitor and apply code that would generate dynamic dialouge games.

To implement your own listener create your own struct and assign it's address to `DGDLListener`:

```go
type MyListener struct {}

var _ parser.DGDLListener = &MyListener{}
```

Now implement all the interface methods of `ParseTreeListener` that `DGDLListener` inherits:

```go
// VisitTerminal is called when a terminal node is visited.
func (ml *MyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (ml *MyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (ml *MyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (ml *MyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}
```

Now implement `DGDLListener` interface methods - these that are called on entering and exiting every grammar rule:

```go
// EnterGame is called when production game is entered.
func (ml *MyListener) EnterGame(ctx *parser.GameContext) {}

// ExitGame is called when production game is exited.
func (ml *MyListener) ExitGame(ctx *parser.GameContext) {}

// EnterStore is called when production store is entered.
func (ml *MyListener) EnterStore(ctx *parser.StoreContext) {}

// ExitStore is called when production store is exited.
func (ml *MyListener) ExitStore(ctx *parser.StoreContext) {}

// EnterTurns is called when production turns is entered.
func (ml *MyListener) EnterTurns(ctx *parser.TurnsContext) {}

// ExitTurns is called when production turns is exited.
func (ml *MyListener) ExitTurns(ctx *parser.TurnsContext) {}

// EnterPlayers is called when production players is entered.
func (ml *MyListener) EnterPlayers(ctx *parser.PlayersContext) {}

// ExitPlayers is called when production players is exited.
func (ml *MyListener) ExitPlayers(ctx *parser.PlayersContext) {}

// EnterPlayer is called when production player is entered.
func (ml *MyListener) EnterPlayer(ctx *parser.PlayerContext) {}

// ExitPlayer is called when production player is exited.
func (ml *MyListener) ExitPlayer(ctx *parser.PlayerContext) {}

// EnterRoles is called when production roles is entered.
func (ml *MyListener) EnterRoles(ctx *parser.RolesContext) {}

// ExitRoles is called when production roles is exited.
func (ml *MyListener) ExitRoles(ctx *parser.RolesContext) {}

// EnterPrinciples is called when production principles is entered.
func (ml *MyListener) EnterPrinciples(ctx *parser.PrinciplesContext) {}

// ExitPrinciples is called when production principles is exited.
func (ml *MyListener) ExitPrinciples(ctx *parser.PrinciplesContext) {}

// EnterPrinciple is called when production principle is entered.
func (ml *MyListener) EnterPrinciple(ctx *parser.PrincipleContext) {}

// ExitPrinciple is called when production principle is exited.
func (ml *MyListener) ExitPrinciple(ctx *parser.PrincipleContext) {}

// EnterScope is called when production scope is entered.
func (ml *MyListener) EnterScope(ctx *parser.ScopeContext) {}

// ExitScope is called when production scope is exited.
func (ml *MyListener) ExitScope(ctx *parser.ScopeContext) {}

// EnterMoves is called when production moves is entered.
func (ml *MyListener) EnterMoves(ctx *parser.MovesContext) {}

// ExitMoves is called when production moves is exited.
func (ml *MyListener) ExitMoves(ctx *parser.MovesContext) {}

// EnterMove is called when production move is entered.
func (ml *MyListener) EnterMove(ctx *parser.MoveContext) {}

// ExitMove is called when production move is exited.
func (ml *MyListener) ExitMove(ctx *parser.MoveContext) {}

// EnterContent is called when production content is entered.
func (ml *MyListener) EnterContent(ctx *parser.ContentContext) {}

// ExitContent is called when production content is exited.
func (ml *MyListener) ExitContent(ctx *parser.ContentContext) {}

// EnterConditions is called when production conditions is entered.
func (ml *MyListener) EnterConditions(ctx *parser.ConditionsContext) {}

// ExitConditions is called when production conditions is exited.
func (ml *MyListener) ExitConditions(ctx *parser.ConditionsContext) {}

// EnterEffects is called when production effects is entered.
func (ml *MyListener) EnterEffects(ctx *parser.EffectsContext) {}

// ExitEffects is called when production effects is exited.
func (ml *MyListener) ExitEffects(ctx *parser.EffectsContext) {}

// EnterExpr is called when production expr is entered.
func (ml *MyListener) EnterExpr(ctx *parser.ExprContext) {}

// ExitExpr is called when production expr is exited.
func (ml *MyListener) ExitExpr(ctx *parser.ExprContext) {}

// EnterParam is called when production param is entered.
func (ml *MyListener) EnterParam(ctx *parser.ParamContext) {}

// ExitParam is called when production param is exited.
func (ml *MyListener) ExitParam(ctx *parser.ParamContext) {}
```

Do whatever logic is required in these concrete functions to generate dialogue game.

Example running `main.go` with dialogue game description file passed as arg:

```go
func main()  {
	input := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewDGDLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer,0)
	p := parser.NewDGDLParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Game()
	walker := antlr.NewParseTreeWalker()
	walker.Walk(&MyListener{}, tree)
}
```
