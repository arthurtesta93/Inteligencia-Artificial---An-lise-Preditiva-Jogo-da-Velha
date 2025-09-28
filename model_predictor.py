#!/usr/bin/env python3
import json
import sys
from pathlib import Path
import joblib
import numpy as np
import pandas as pd

# Pasta dos modelos (mesmo nível do jogo)
MODELS_DIR = Path("/Users/arthurtesta/Documents/PUCRS/Inteligencia Artificial/data/models")

# Mapeia "apelidos" usados no jogo -> arquivos .pkl
# Ajuste os nomes se seus arquivos tiverem outros prefixos/sufixos
MODEL_FILES = {
    "mlp":            ["best_mlp.pkl", "mlp.pkl"],
    "random_forest":  ["best_random_forest.pkl", "random_forest.pkl", "rf.pkl"],
    "knn":            ["best_knn.pkl", "knn.pkl"],
    "decision_tree":  ["best_decision_tree.pkl", "decision_tree.pkl", "dt.pkl"],
}

# Rótulos compatíveis com o jogo
LABELS = {
    0: "Tem Jogo",
    1: "Possível Fim de Jogo",
    2: "Fim de Jogo",
}

# Cache simples para não reabrir modelos a cada chamada
_model_cache = {}

def _load_first_existing(paths):
    for p in paths:
        full = MODELS_DIR / p
        if full.exists():
            return joblib.load(full)
    return None

def load_model(alias):
    if alias in _model_cache:
        return _model_cache[alias]
    files = MODEL_FILES.get(alias, [])
    model = _load_first_existing(files)
    if model is not None:
        _model_cache[alias] = model
    return model

def available_models():
    avail = []
    for alias, files in MODEL_FILES.items():
        if _load_first_existing(files) is not None:
            avail.append(alias)
    return avail

def parse_board(board_csv):
    """
    board_csv: string "x,o,b,..." (9 itens)
    Retorna DataFrame com colunas 0..8 (int) e uma única linha com strings 'x','o','b'
    Isso casa com o pipeline salvo (OneHot nas colunas 0..8).
    """
    parts = [s.strip() for s in board_csv.split(",")]
    if len(parts) != 9:
        raise ValueError("Board deve ter 9 posições separadas por vírgula ('x','o','b').")
    # DataFrame 1x9 com colunas int 0..8
    data = {i: parts[i] for i in range(9)}
    X = pd.DataFrame([data])
    return X

def predict(alias, board_csv):
    model = load_model(alias)
    if model is None:
        return {"model": alias, "error": f"modelo '{alias}' não encontrado em ./models"}

    X = parse_board(board_csv)

    # Alguns pipelines podem esperar colunas inteiras 0..8; garantimos:
    X.columns = [int(c) for c in X.columns]

    # Predição
    y_pred = model.predict(X)
    pred_int = int(y_pred[0])
    pred_label = LABELS.get(pred_int, "Desconhecido")

    # Confiança (se o modelo expõe predict_proba)
    conf = None
    if hasattr(model, "predict_proba"):
        try:
            proba = model.predict_proba(X)[0]
            # Usa a probabilidade da classe prevista
            conf = float(proba[pred_int]) if pred_int < len(proba) else None
        except Exception:
            conf = None
    else:
        # Pode ser um Pipeline; tente acessar o estimador final
        try:
            clf = getattr(model, "named_steps", {}).get("clf", None)
            if clf is not None and hasattr(clf, "predict_proba"):
                proba = clf.predict_proba(model.named_steps.get("prep", lambda v: v).transform(X))[0] \
                        if "prep" in model.named_steps else clf.predict_proba(X)[0]
                conf = float(proba[pred_int]) if pred_int < len(proba) else None
        except Exception:
            conf = None

    return {
        "model": alias,
        "prediction_int": pred_int,
        "prediction_label": pred_label,
        "confidence": conf,
    }

def main():
    """
    Uso:
      python3 model_predictor.py info
      python3 model_predictor.py predict <alias> "x,o,b,x,..."   # 9 valores
    """
    if len(sys.argv) < 2:
        print(json.dumps({"error": "uso: info | predict <alias> <board_csv>"}))
        return

    cmd = sys.argv[1]

    if cmd == "info":
        print(json.dumps({"available_models": available_models()}))
        return

    if cmd == "predict":
        if len(sys.argv) < 4:
            print(json.dumps({"error": "uso: predict <alias> <board_csv>"}))
            return
        alias = sys.argv[2]
        board_csv = sys.argv[3]
        try:
            result = predict(alias, board_csv)
            print(json.dumps(result))
        except Exception as e:
            print(json.dumps({"model": alias, "error": str(e)}))
        return

    print(json.dumps({"error": f"comando desconhecido: {cmd}"}))

if __name__ == "__main__":
    main()
