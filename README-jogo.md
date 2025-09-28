# 🎮 Jogo da Velha com IA - Front End Mínimo

## Descrição

Este é um front-end mínimo em linha de comando para o jogo da velha, implementado em Go, que atende aos requisitos do trabalho de Inteligência Artificial da PUCRS.

## Funcionalidades Implementadas

### ✅ Requisitos Atendidos

1. **Dois jogadores**: Humano (X) vs Computador (O)
2. **Computador joga aleatoriamente**: Seleciona posições válidas aleatoriamente
3. **Análise de IA a cada jogada**: Predição do estado do jogo após cada movimento
4. **Estado real vs Predição**: Comparação e contabilização de acertos/erros
5. **Estatísticas**: Acurácia da IA durante a partida
6. **Interface não-gráfica**: Terminal com interface clara e intuitiva

### 🧠 Estados de Jogo

- **Tem Jogo**: Jogo ainda em andamento
- **Possível Fim de Jogo**: Um jogador pode vencer na próxima jogada
- **Fim de Jogo**: Jogo terminou (vitória ou empate)

### 📊 Informações Exibidas

- Tabuleiro atual com posições numeradas (0-8)
- Algoritmo de IA utilizado
- Predição da IA para o estado atual
- Estado real do jogo
- Comparação (acerto/erro) da predição
- Contador de acertos e erros
- Acurácia final da IA

## Como Executar

### 🚀 Setup Inicial (Primeira vez)

```bash
# 1. Navegar para o diretório do projeto
cd "/Users/arthurtesta/Documents/PUCRS/Inteligencia Artificial"

# 2. Executar setup automático (instala dependências e valida modelos)
python3 setup_models.py

# 3. Executar o jogo
go run jogo-da-velha.go
```

### ⚡ Execução Rápida (Após setup)

```bash
# Executar diretamente
go run jogo-da-velha.go
```

### 🔧 Comandos de Teste

```bash
# Verificar modelos disponíveis
python3 model_predictor.py info

# Testar predição específica
python3 model_predictor.py predict mlp b,x,o,b,x,o,b,b,b
```

## Como Jogar

1. **Início**: Você joga como 'X' e o computador como 'O'
2. **Sua jogada**: Digite um número de 0 a 8 para escolher a posição
3. **Análise**: Após cada jogada, a IA analisa o estado do tabuleiro
4. **Feedback**: Veja se a IA acertou ou errou a predição
5. **Fim**: O jogo termina quando há vitória, empate ou não há mais jogadas

### Layout do Tabuleiro
```
0 | 1 | 2
-----------
3 | 4 | 5
-----------
6 | 7 | 8
```

## Algoritmos de IA Disponíveis

O sistema suporta **5 algoritmos diferentes** com seleção interativa:

### 🧠 Modelos Implementados

1. **Classificador Baseado em Regras** (Fallback)
   - Sempre disponível
   - Heurísticas simples baseadas em contagem de peças

2. **MLP Neural Network** (Acurácia: 82.53%)
   - Rede neural multicamadas
   - Configuração: (64, 32, 16, 8) camadas, alpha=0.0001

3. **Random Forest** (Acurácia: 75.96%)
   - Ensemble de árvores de decisão
   - Configuração: 100 estimators, max_depth=10

4. **k-Nearest Neighbors** (Acurácia: 72.22%)
   - Classificação baseada em vizinhos próximos
   - Configuração: k=7

5. **Decision Tree** (Acurácia: 66.68%)
   - Árvore de decisão única
   - Configuração: max_depth=5, min_samples_leaf=10

### 🔄 Integração Completa Implementada

✅ **Sistema híbrido Go + Python**:
- Front-end em Go para interface e lógica do jogo
- Bridge Python para carregamento e execução dos modelos
- Comunicação via JSON entre os componentes
- Sistema de fallback para tolerância a falhas

✅ **Seleção interativa de modelos**:
- Interface para escolha do algoritmo
- Verificação automática de disponibilidade
- Exibição de acurácias dos modelos

✅ **Arquitetura robusta**:
- Tratamento de erros gracioso
- Fallback para regras se Python falhar
- Validação de dependências automática

## Estatísticas Coletadas

- **Total de predições**: Número total de análises feitas
- **Acertos**: Predições corretas da IA
- **Erros**: Predições incorretas da IA  
- **Acurácia**: Percentual de acertos

## Exemplo de Saída

```
🎮 BEM-VINDO AO JOGO DA VELHA COM IA! 🎮
============================================================
Você é 'X' e o computador é 'O'
Digite um número de 0-8 para fazer sua jogada
Algoritmo de IA: Classificador Baseado em Regras
============================================================

   Tabuleiro Atual:
   0 | 1 | 2
  -----------
   3 | 4 | 5
  -----------
   6 | 7 | 8

   Jogo:
     |   |  
  -----------
     |   |  
  -----------
     |   |  

Sua vez (X)! Digite a posição (0-8): 4

----------------------------------------
🧠 ANÁLISE DA IA:
----------------------------------------
Algoritmo: Classificador Baseado em Regras
Predição da IA: Tem Jogo
Estado Real: Tem Jogo
✓ IA ACERTOU! Predição: Tem Jogo | Real: Tem Jogo
Acertos: 1 | Erros: 0
----------------------------------------
```

## Melhorias Futuras

1. **Integração com modelos treinados**: MLP, Random Forest, k-NN, Decision Tree
2. **Interface gráfica**: Versão web ou desktop
3. **Diferentes níveis de IA**: Fácil, médio, difícil
4. **Histórico de partidas**: Salvar estatísticas de múltiplas partidas
5. **Análise mais detalhada**: Matriz de confusão, métricas por classe

## Estrutura do Código

- `GameState`: Enum para estados do jogo
- `Board`: Array representando o tabuleiro 3x3
- `Game`: Struct principal com lógica do jogo
- `GetRealGameState()`: Verifica estado real do tabuleiro
- `PredictGameStateWithAI()`: Predição da IA (personalizável)
- `CompareStates()`: Compara predição vs realidade
- `PlayGame()`: Loop principal do jogo

---

**Desenvolvido para o Trabalho T1 - PUCRS Inteligência Artificial**
