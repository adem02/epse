# EPSE - Express Project Structure Generator 🚀

EPSE (**Express Project Structure Generator**) est un outil CLI (Command-Line Interface) conçu pour simplifier la génération de structures de projet Node.js/Express avec TypeScript. Il aide les développeurs à gagner du temps en créant des templates de projet adaptés à différents besoins.

---

## Fonctionnalités 🌟

- 📂 **Génération rapide** de structures de projet prêtes à l'emploi.
- 🔧 **Personnalisation flexible** avec deux types de structures :
  - **Lite** : Une structure minimale pour des projets simples ou prototypes.
  - **Clean** : Une structure complète et bien organisée, conforme aux principes de **clean architecture**, idéale pour des projets robustes et modulaires.
- ✨ **Interface interactive** pour guider les utilisateurs dans leurs choix.
- ⚙️ **Commandes directes** pour les utilisateurs avancés.

---

## Templates disponibles 📜

### 1. Lite - Minimaliste  
Structure légère pour un démarrage rapide avec Express et TypeScript.

**Commande :**  
```bash
epse generate <project-name> --lite [destination]
```

### 2. Clean - Complète  
Structure basée sur les principes de **clean architecture** et intégrant **TSOA**. Elle est adaptée aux projets conséquents nécessitant une architecture modulaire.

**Commande :**  
```bash
epse generate <project-name> --clean [destination]