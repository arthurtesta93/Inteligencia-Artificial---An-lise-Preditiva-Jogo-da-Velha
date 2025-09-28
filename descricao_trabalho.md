##PUCRS – Inteligência Artificial - Silvia Moraes - Turma 32

# T1 – Tic Tac Toe com ML

## Enunciado

Neste primeiro trabalho prático da disciplina, você vai construir **um sistema de IA para o jogo
da velha, considerando um tabuleiro clássico 3x3**. O **objetivo da IA** não é ser um dos players,
mas sim **verificar o estado de jogo** (Figura 1).
Figura 1 - Estados de jogo
A seguir serão descritas as etapas do trabalho.

1. **Objetivo:** A IA que você implementará deve receber como entrada o tabuleiro de um
    jogo da velha em andamento e classificar esse estado em:
    ● Tem jogo
    ● Possibilidade de Fim de Jogo
    ● Fim de Jogo
       Figura 2 - Modelo de IA em execução
2.
**2. Dataset:** Acesse o link https://archive.ics.uci.edu/dataset/101/tic+tac+toe+endgame e
    obtenha um dataset que possui instâncias do tabuleiro do jogo da velha. O dataset não
    está pronto para ser usado neste trabalho. Portanto, analise o dataset e realize as
    adequações que julgar pertinente (limpeza, transformação, inclusão de novas
    classes,...). Não esqueça de registrar cada adequação e todos os passos que você
    executou (justificando esses passos) para isso. Procure gerar um dataset balanceado e
    **não use todas as instâncias** , apenas as mais representativas. Use no máximo 250
    instâncias de cada classe, se for possível, para formar o dataset do trabalho.

**3. Solução de IA:** Divida fisicamente o dataset em treino, validação e teste (80, 10, 10,
    respectivamente). Para uma classe de 250 amostras, 200 serão usadas para treino, 25
    para validação e 25 para teste. Precisam ser os mesmos subconjuntos nos
    experimentos, pois testaremos mais de um algoritmo de IA. Construa a sua solução
    testando ao menos 3 algoritmos classificadores: k-NN, MLP e árvores de decisão. Inclua
    um 4º algoritmo nos seus testes a sua escolha. Não esqueça de incluir em seu
    relatório(ppt) uma pequena explicação de como funciona esse 4º algoritmo. Todas as
    decisões referentes a parâmetros desses algoritmos devem ser apresentadas e
    justificadas no seu relatório(ppt) também. No caso da MLP, não esqueça de informar a
    topologia usada. Meça os seus resultados utilizando acurácia, precision, recall e
    F-measure. Busque bons resultados e evite overfitting (procure as melhores
    configurações e parâmetros para os algoritmos). Compare os resultados e escolha o
    melhor algoritmo para o problema. Mostre a sua comparação usando tabelas e
    gráficos. Discuta os resultados, analise erros e confusões. Justifique sua escolha no
    texto do seu relatório (formato ppt).
4. **Front end** : Construa um front end mínimo (não precisa ser gráfico) para o jogo da
    velha, onde dois players possam interagir. Um player deve ser humano e o outro a
    máquina jogando de forma aleatória. No front end, a cada turno (a cada jogada de
    usuário/computador), a solução de IA escolhida por você deve indicar (exibir) um dos
    estados de jogo conforme a Figura 1. Na tela, durante o jogo, deve ser exibido o
    algoritmo de IA que está analisando o tabuleiro e sua predição a cada jogada. Deve
    aparecer também o estado real do jogo (crie um método que verifica o estado atual do
    tabuleiro). Esse método será usado para contabilizar erros e acertos do seu algoritmo
    de IA durante um jogo. Use esse método também para controlar o fim das partidas.
    Não esqueça de durante cada partida, a cada jogada, exibir acertos e erros da solução.
    Meça também a acurácia da solução durante as interações com os usuários. Registre
    isso no seu relatório (ppt).

## Definições e critérios:

● Os grupos podem ser de até 6 alunos. Distribua as atividades entre os integrantes do
grupo de forma que todos trabalhem. Se inscreva no moodle. Alunos que não
formarem grupo, terão grupo definido pelo professor.
● Data de entrega e apresentação: no cronograma disponível no moodle.
● Na data da apresentação, todos os integrantes do grupo devem estar presentes e a
avaliação não é apenas sobre o que foi entregue, mas também sobre o
domínio/conhecimento demonstrado pelos integrantes durante a apresentação.
**Pontuação**
● Dataset .......................................: 1,5 ponto (documentado, mostre por meio de um
gráfico a distribuição de amostras por classe)
● Soluções de IA e documentação..: 4,0 pontos (1,0 por algoritmo e configuração)
● Front End......................................: 1,5 ponto


● Apresentação (ppt).......................................: 3,0 pontos, contendo a seguinte estrutura
(em torno de 10 slides, aproximadamente):
o Capa: título do trabalho e nome dos integrantes
o Dataset: modificações feitas (com justificativa) e distribuição em classes
o Algoritmos e resultados: descrições, configurações usadas/testadas e
resultados. Além de apresentar gráficos por métrica para cada algoritmo.
Consolide os melhores resultados em um único gráfico. Análise dos resultados,
discuta erros e confusões. Justifique e escolha o melhor modelo.
o Front End: comente o desempenho e os resultados do modelo quando foi
usado durante as partidas.
o Considerações finais: Inclua uma conclusão descrevendo dificuldades
encontradas e ganhos obtidos em decorrência da execução desse trabalho. Se
os resultados foram satisfatórios tanto no desenvolvimento quanto no uso do
modelo (Front end) e como melhorá-los. Discuta erros, acertos e confusões.
**o Mencione ao final, que ferramentas de IA você usou e com que propósito.**
● Observações:
o Código incorreto, ausência na apresentação (não justificada), não domínio
durante apresentação e não cumprimento do enunciado provocam decréscimo
na nota.
o Cópia de trabalhos de colegas zeram o trabalho.
o Indique as ferramentas de IA que você usou, especificando onde foram usadas.