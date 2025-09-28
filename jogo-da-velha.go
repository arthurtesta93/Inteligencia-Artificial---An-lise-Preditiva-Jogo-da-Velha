package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// GameState representa os possíveis estados do jogo
type GameState int

const (
	TemJogo GameState = iota
	PossivelFimDeJogo
	FimDeJogo
)

func (gs GameState) String() string {
	switch gs {
	case TemJogo:
		return "Tem Jogo"
	case PossivelFimDeJogo:
		return "Possível Fim de Jogo"
	case FimDeJogo:
		return "Fim de Jogo"
	default:
		return "Estado Desconhecido"
	}
}

// Board representa o tabuleiro 3x3
type Board [9]string

// AIModel representa um modelo de IA disponível
type AIModel struct {
	Name        string  `json:"name"`
	DisplayName string  `json:"display_name"`
	Accuracy    float64 `json:"accuracy"`
	Available   bool    `json:"available"`
}

// PredictionResult representa o resultado de uma predição
type PredictionResult struct {
	Model           string   `json:"model"`
	PredictionInt   int      `json:"prediction_int"`
	PredictionLabel string   `json:"prediction_label"`
	Confidence      *float64 `json:"confidence,omitempty"`
	Error           string   `json:"error,omitempty"`
}

// Game representa o estado do jogo
type Game struct {
	board            Board
	currentPlayer    string
	gameOver         bool
	winner           string
	aiCorrect        int
	aiIncorrect      int
	totalPredictions int
	selectedModel    string
	availableModels  []AIModel
}

// NewGame cria um novo jogo
func NewGame() *Game {
	game := &Game{
		board:         Board{"b", "b", "b", "b", "b", "b", "b", "b", "b"},
		currentPlayer: "x",
		gameOver:      false,
		winner:        "",
		selectedModel: "rule_based", // default to rule-based
	}
	game.initializeModels()
	return game
}

// initializeModels carrega informações dos modelos disponíveis
func (g *Game) initializeModels() {
	// Modelos com suas respectivas acurácias (baseado no seu CSV de resultados)
	g.availableModels = []AIModel{
		{Name: "rule_based", DisplayName: "Classificador Baseado em Regras", Accuracy: 0.0, Available: true},
		{Name: "mlp", DisplayName: "MLP Neural Network", Accuracy: 82.53, Available: false},
		{Name: "random_forest", DisplayName: "Random Forest", Accuracy: 75.96, Available: false},
		{Name: "knn", DisplayName: "k-Nearest Neighbors (k=7)", Accuracy: 72.22, Available: false},
		{Name: "decision_tree", DisplayName: "Decision Tree", Accuracy: 66.68, Available: false},
	}

	// Verificar quais modelos estão disponíveis
	g.checkModelAvailability()
}

// checkModelAvailability verifica se os modelos Python estão disponíveis
func (g *Game) checkModelAvailability() {
	// Try to use virtual environment python first, then fall back to system python3
	pythonCmd := "python3"
	if _, err := os.Stat(".venv/bin/python"); err == nil {
		pythonCmd = ".venv/bin/python"
	}

	cmd := exec.Command(pythonCmd, "model_predictor.py", "info")
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("⚠️  Python models não disponíveis: %v\n", err)
		return
	}

	var info struct {
		AvailableModels []string `json:"available_models"`
	}

	if err := json.Unmarshal(output, &info); err != nil {
		fmt.Printf("⚠️  Erro ao verificar modelos: %v\n", err)
		return
	}

	// Marcar modelos como disponíveis
	for i := range g.availableModels {
		for _, available := range info.AvailableModels {
			if g.availableModels[i].Name == available {
				g.availableModels[i].Available = true
				break
			}
		}
	}
}

// DisplayBoard exibe o tabuleiro atual
func (g *Game) DisplayBoard() {
	fmt.Println("\n   Tabuleiro Atual:")
	fmt.Println("   0 | 1 | 2")
	fmt.Println("  -----------")
	fmt.Println("   3 | 4 | 5")
	fmt.Println("  -----------")
	fmt.Println("   6 | 7 | 8")
	fmt.Println()

	fmt.Println("   Jogo:")
	for i := 0; i < 9; i += 3 {
		row := ""
		for j := 0; j < 3; j++ {
			cell := g.board[i+j]
			if cell == "b" {
				cell = " "
			}
			row += fmt.Sprintf(" %s ", cell)
			if j < 2 {
				row += "|"
			}
		}
		fmt.Println("  " + row)
		if i < 6 {
			fmt.Println("  -----------")
		}
	}
	fmt.Println()
}

// GetRealGameState verifica o estado real do jogo
func (g *Game) GetRealGameState() GameState {
	// Verifica se há um vencedor
	winPatterns := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // linhas
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // colunas
		{0, 4, 8}, {2, 4, 6}, // diagonais
	}

	// Verifica vitória
	for _, pattern := range winPatterns {
		if g.board[pattern[0]] != "b" &&
			g.board[pattern[0]] == g.board[pattern[1]] &&
			g.board[pattern[1]] == g.board[pattern[2]] {
			g.winner = g.board[pattern[0]]
			g.gameOver = true
			return FimDeJogo
		}
	}

	// Conta posições vazias
	emptySpaces := 0
	for _, cell := range g.board {
		if cell == "b" {
			emptySpaces++
		}
	}

	// Se não há espaços vazios, é empate
	if emptySpaces == 0 {
		g.gameOver = true
		g.winner = "empate"
		return FimDeJogo
	}

	// Verifica se há possibilidade de vitória em uma jogada
	if g.canWinInOneMove("x") || g.canWinInOneMove("o") {
		return PossivelFimDeJogo
	}

	return TemJogo
}

// canWinInOneMove verifica se um jogador pode vencer em uma jogada
func (g *Game) canWinInOneMove(player string) bool {
	winPatterns := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // linhas
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // colunas
		{0, 4, 8}, {2, 4, 6}, // diagonais
	}

	for _, pattern := range winPatterns {
		count := 0
		emptyPos := -1
		for _, pos := range pattern {
			if g.board[pos] == player {
				count++
			} else if g.board[pos] == "b" {
				emptyPos = pos
			}
		}
		if count == 2 && emptyPos != -1 {
			return true
		}
	}
	return false
}

// PredictGameStateWithAI prediz o estado usando o modelo selecionado
func (g *Game) PredictGameStateWithAI() GameState {
	if g.selectedModel == "rule_based" {
		return g.predictWithRules()
	}

	// Usar modelo treinado via Python
	return g.predictWithTrainedModel()
}

// predictWithRules implementa o classificador baseado em regras original
func (g *Game) predictWithRules() GameState {
	// Conta X's e O's
	xCount := 0
	oCount := 0
	emptyCount := 0

	for _, cell := range g.board {
		switch cell {
		case "x":
			xCount++
		case "o":
			oCount++
		case "b":
			emptyCount++
		}
	}

	// Lógica simplificada baseada em padrões
	if emptyCount <= 2 {
		return FimDeJogo
	} else if emptyCount <= 4 && (xCount >= 2 || oCount >= 2) {
		return PossivelFimDeJogo
	} else {
		return TemJogo
	}
}

// predictWithTrainedModel usa o modelo Python para predição
func (g *Game) predictWithTrainedModel() GameState {
	// Converter board para string formato Python
	boardString := strings.Join(g.board[:], ",")

	// Use virtual environment python if available
	pythonCmd := "python3"
	if _, err := os.Stat(".venv/bin/python"); err == nil {
		pythonCmd = ".venv/bin/python"
	}

	// Chamar script Python
	cmd := exec.Command(pythonCmd, "model_predictor.py", "predict", g.selectedModel, boardString)
	output, err := cmd.Output()

	if err != nil {
		fmt.Printf("⚠️  Erro ao executar modelo %s: %v\n", g.selectedModel, err)
		// Fallback para regras
		return g.predictWithRules()
	}

	// Parse resultado JSON
	var result PredictionResult
	if err := json.Unmarshal(output, &result); err != nil {
		fmt.Printf("⚠️  Erro ao interpretar resultado: %v\n", err)
		return g.predictWithRules()
	}

	if result.Error != "" {
		fmt.Printf("⚠️  Erro no modelo: %s\n", result.Error)
		return g.predictWithRules()
	}

	// Converter resultado para GameState
	switch result.PredictionInt {
	case 0:
		return TemJogo
	case 1:
		return PossivelFimDeJogo
	case 2:
		return FimDeJogo
	default:
		return TemJogo
	}
}

// IsValidMove verifica se a jogada é válida
func (g *Game) IsValidMove(position int) bool {
	return position >= 0 && position <= 8 && g.board[position] == "b"
}

// MakeMove executa uma jogada
func (g *Game) MakeMove(position int) bool {
	if !g.IsValidMove(position) {
		return false
	}
	g.board[position] = g.currentPlayer
	return true
}

// GetComputerMove gera uma jogada aleatória para o computador
func (g *Game) GetComputerMove() int {
	var validMoves []int
	for i, cell := range g.board {
		if cell == "b" {
			validMoves = append(validMoves, i)
		}
	}

	if len(validMoves) == 0 {
		return -1
	}

	return validMoves[rand.Intn(len(validMoves))]
}

// SwitchPlayer alterna entre jogadores
func (g *Game) SwitchPlayer() {
	if g.currentPlayer == "x" {
		g.currentPlayer = "o"
	} else {
		g.currentPlayer = "x"
	}
}

// CompareStates compara predição da IA com estado real
func (g *Game) CompareStates(predicted, real GameState) {
	g.totalPredictions++
	if predicted == real {
		g.aiCorrect++
		fmt.Printf("✓ IA ACERTOU! Predição: %s | Real: %s\n", predicted, real)
	} else {
		g.aiIncorrect++
		fmt.Printf("✗ IA ERROU! Predição: %s | Real: %s\n", predicted, real)
	}
}

// ShowGameStats exibe estatísticas do jogo
func (g *Game) ShowGameStats() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("ESTATÍSTICAS DA PARTIDA")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Printf("Total de predições: %d\n", g.totalPredictions)
	fmt.Printf("Acertos da IA: %d\n", g.aiCorrect)
	fmt.Printf("Erros da IA: %d\n", g.aiIncorrect)

	if g.totalPredictions > 0 {
		accuracy := float64(g.aiCorrect) / float64(g.totalPredictions) * 100
		fmt.Printf("Acurácia da IA: %.2f%%\n", accuracy)
	}
	fmt.Println(strings.Repeat("=", 50))
}

// selectModel permite ao usuário escolher o modelo de IA
func (g *Game) selectModel() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🧠 SELEÇÃO DE MODELO DE IA")
	fmt.Println(strings.Repeat("=", 60))

	// Mostrar modelos disponíveis
	fmt.Println("Modelos disponíveis:")
	for i, model := range g.availableModels {
		status := "❌ Não disponível"
		if model.Available {
			status = "✅ Disponível"
		}
		accuracyStr := ""
		if model.Accuracy > 0 {
			accuracyStr = fmt.Sprintf(" (Acurácia: %.2f%%)", model.Accuracy)
		}
		fmt.Printf("%d. %s %s%s\n", i+1, model.DisplayName, status, accuracyStr)
	}

	fmt.Print("\nEscolha o modelo (número) ou Enter para usar o padrão: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	if input == "" {
		// Usar primeiro modelo disponível
		for _, model := range g.availableModels {
			if model.Available {
				g.selectedModel = model.Name
				break
			}
		}
	} else {
		choice, err := strconv.Atoi(input)
		if err == nil && choice >= 1 && choice <= len(g.availableModels) {
			selectedModel := g.availableModels[choice-1]
			if selectedModel.Available {
				g.selectedModel = selectedModel.Name
			} else {
				fmt.Println("⚠️  Modelo não disponível, usando padrão.")
			}
		}
	}

	// Mostrar modelo selecionado
	for _, model := range g.availableModels {
		if model.Name == g.selectedModel {
			fmt.Printf("\n🎯 Modelo selecionado: %s\n", model.DisplayName)
			break
		}
	}
}

// getSelectedModelDisplayName retorna o nome de exibição do modelo selecionado
func (g *Game) getSelectedModelDisplayName() string {
	for _, model := range g.availableModels {
		if model.Name == g.selectedModel {
			return model.DisplayName
		}
	}
	return "Modelo Desconhecido"
}

// PlayGame executa o loop principal do jogo
func (g *Game) PlayGame() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🎮 BEM-VINDO AO JOGO DA VELHA COM IA! 🎮")
	fmt.Println(strings.Repeat("=", 60))

	// Seleção de modelo
	g.selectModel()

	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("Você é 'X' e o computador é 'O'")
	fmt.Println("Digite um número de 0-8 para fazer sua jogada")
	fmt.Printf("Algoritmo de IA: %s\n", g.getSelectedModelDisplayName())
	fmt.Println(strings.Repeat("=", 60))

	for !g.gameOver {
		g.DisplayBoard()

		if g.currentPlayer == "x" {
			// Turno do jogador humano
			fmt.Printf("Sua vez (X)! Digite a posição (0-8): ")
			scanner.Scan()
			input := scanner.Text()

			position, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("❌ Entrada inválida! Digite um número de 0 a 8.")
				continue
			}

			if !g.MakeMove(position) {
				fmt.Println("❌ Jogada inválida! Posição ocupada ou fora do range.")
				continue
			}
		} else {
			// Turno do computador
			fmt.Println("Turno do computador (O)...")
			time.Sleep(1 * time.Second)

			computerMove := g.GetComputerMove()
			if computerMove == -1 {
				break
			}

			g.MakeMove(computerMove)
			fmt.Printf("🤖 Computador jogou na posição %d\n", computerMove)
		}

		// Análise da IA após cada jogada
		fmt.Println("\n" + strings.Repeat("-", 40))
		fmt.Println("🧠 ANÁLISE DA IA:")
		fmt.Println(strings.Repeat("-", 40))

		realState := g.GetRealGameState()
		predictedState := g.PredictGameStateWithAI()

		fmt.Printf("Algoritmo: %s\n", g.getSelectedModelDisplayName())
		fmt.Printf("Predição da IA: %s\n", predictedState)
		fmt.Printf("Estado Real: %s\n", realState)

		g.CompareStates(predictedState, realState)
		fmt.Printf("Acertos: %d | Erros: %d\n", g.aiCorrect, g.aiIncorrect)
		fmt.Println(strings.Repeat("-", 40))

		if g.gameOver {
			break
		}

		g.SwitchPlayer()
	}

	// Fim do jogo
	g.DisplayBoard()
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🏁 FIM DE JOGO!")
	fmt.Println(strings.Repeat("=", 50))

	switch g.winner {
	case "empate":
		fmt.Println("🤝 EMPATE!")
	case "x":
		fmt.Println("🎉 VOCÊ VENCEU!")
	case "o":
		fmt.Println("🤖 COMPUTADOR VENCEU!")
	}

	g.ShowGameStats()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		game := NewGame()
		game.PlayGame()

		fmt.Println("\nDeseja jogar novamente? (s/n): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		response := strings.ToLower(scanner.Text())

		if response != "s" && response != "sim" {
			fmt.Println("Obrigado por jogar! 👋")
			break
		}

		fmt.Println("\n" + strings.Repeat("=", 60))
		fmt.Println("🔄 NOVA PARTIDA!")
		fmt.Println(strings.Repeat("=", 60))
	}
}
