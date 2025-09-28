#!/usr/bin/env python3
"""
Setup script for Tic-Tac-Toe ML Models Integration
Installs dependencies and validates model files
"""

import sys
import subprocess
import os
from pathlib import Path

def check_python_version():
    """Check if Python version is compatible"""
    if sys.version_info < (3, 7):
        print("❌ Python 3.7 ou superior é necessário")
        print(f"Versão atual: {sys.version}")
        return False
    print(f"✅ Python {sys.version.split()[0]} detectado")
    return True

def install_dependencies():
    """Install required Python packages"""
    packages = [
        'scikit-learn==1.7.1',  # Match the version used to train models
        'pandas>=1.3.0',
        'numpy>=1.21.0'
    ]
    
    print("\n🔧 Instalando dependências Python...")
    
    for package in packages:
        print(f"Instalando {package}...")
        try:
            subprocess.check_call([
                sys.executable, '-m', 'pip', 'install', package
            ], stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
            print(f"✅ {package} instalado com sucesso")
        except subprocess.CalledProcessError:
            print(f"❌ Falha ao instalar {package}")
            return False
    
    return True

def validate_model_files():
    """Check if model files exist"""
    models_dir = Path("data/models")
    required_files = [
        'best_mlp.pkl',
        'best_random_forest.pkl', 
        'best_knn.pkl',
        'best_decision_tree.pkl'
    ]
    
    print(f"\n📁 Verificando arquivos de modelo em {models_dir}...")
    
    if not models_dir.exists():
        print(f"❌ Diretório {models_dir} não encontrado")
        return False
    
    missing_files = []
    for filename in required_files:
        filepath = models_dir / filename
        if filepath.exists():
            size_mb = filepath.stat().st_size / (1024 * 1024)
            print(f"✅ {filename} encontrado ({size_mb:.2f} MB)")
        else:
            print(f"❌ {filename} não encontrado")
            missing_files.append(filename)
    
    if missing_files:
        print(f"\n⚠️  Arquivos de modelo faltando: {', '.join(missing_files)}")
        print("Certifique-se de que os modelos foram treinados e salvos corretamente.")
        return False
    
    return True

def test_model_loading():
    """Test if models can be loaded successfully"""
    print("\n🧪 Testando carregamento de modelos...")
    
    try:
        # Import required modules
        import pickle
        import numpy as np
        from sklearn.base import BaseEstimator
        
        models_dir = Path("data/models")
        model_files = {
            'MLP': 'best_mlp.pkl',
            'Random Forest': 'best_random_forest.pkl',
            'k-NN': 'best_knn.pkl', 
            'Decision Tree': 'best_decision_tree.pkl'
        }
        
        loaded_models = 0
        for model_name, filename in model_files.items():
            filepath = models_dir / filename
            try:
                with open(filepath, 'rb') as f:
                    model = pickle.load(f)
                
                # Verify it's a sklearn model
                if hasattr(model, 'predict'):
                    print(f"✅ {model_name}: Carregado e validado")
                    loaded_models += 1
                else:
                    print(f"❌ {model_name}: Não é um modelo sklearn válido")
                    
            except Exception as e:
                print(f"❌ {model_name}: Erro ao carregar - {e}")
        
        if loaded_models == len(model_files):
            print(f"✅ Todos os {loaded_models} modelos carregados com sucesso!")
            return True
        else:
            print(f"⚠️  Apenas {loaded_models}/{len(model_files)} modelos carregados")
            return False
            
    except ImportError as e:
        print(f"❌ Erro de importação: {e}")
        return False

def test_integration():
    """Test the full integration pipeline"""
    print("\n🔗 Testando integração completa...")
    
    try:
        # Test the model_predictor script
        import subprocess
        import json
        
        # Test info command
        result = subprocess.run([
            sys.executable, 'model_predictor.py', 'info'
        ], capture_output=True, text=True)
        
        if result.returncode != 0:
            print(f"❌ Falha no teste 'info': {result.stderr}")
            return False
        
        try:
            info = json.loads(result.stdout)
            available_models = info.get('available_models', [])
            print(f"✅ Info: {len(available_models)} modelos disponíveis")
        except json.JSONDecodeError:
            print("❌ Resposta JSON inválida do script")
            return False
        
        # Test prediction
        test_board = "b,x,o,b,x,o,b,b,b"
        for model in available_models:
            result = subprocess.run([
                sys.executable, 'model_predictor.py', 'predict', model, test_board
            ], capture_output=True, text=True)
            
            if result.returncode != 0:
                print(f"❌ Falha na predição com {model}: {result.stderr}")
                continue
            
            try:
                prediction = json.loads(result.stdout)
                if 'error' in prediction:
                    print(f"❌ {model}: {prediction['error']}")
                else:
                    print(f"✅ {model}: Predição = {prediction['prediction_label']}")
            except json.JSONDecodeError:
                print(f"❌ {model}: Resposta JSON inválida")
        
        return True
        
    except Exception as e:
        print(f"❌ Erro no teste de integração: {e}")
        return False

def main():
    """Main setup function"""
    print("🚀 SETUP - INTEGRAÇÃO MODELOS ML COM JOGO DA VELHA")
    print("=" * 60)
    
    # Check Python version
    if not check_python_version():
        return False
    
    # Install dependencies
    if not install_dependencies():
        print("\n❌ Falha na instalação das dependências")
        return False
    
    # Validate model files
    if not validate_model_files():
        print("\n❌ Validação dos arquivos de modelo falhou")
        print("Dica: Execute seus notebooks de treinamento primeiro")
        return False
    
    # Test model loading
    if not test_model_loading():
        print("\n❌ Teste de carregamento dos modelos falhou")
        return False
    
    # Test integration
    if not test_integration():
        print("\n❌ Teste de integração falhou")
        return False
    
    print("\n" + "=" * 60)
    print("🎉 SETUP CONCLUÍDO COM SUCESSO!")
    print("=" * 60)
    print("\n📋 Próximos passos:")
    print("1. Execute: go run jogo-da-velha.go")
    print("2. Escolha seu modelo de IA preferido")
    print("3. Jogue e observe as predições!")
    print("\n📚 Consulte INTEGRATION_GUIDE.md para detalhes técnicos")
    
    return True

if __name__ == "__main__":
    success = main()
    sys.exit(0 if success else 1)
