# cli-contact

Un petit CRM en ligne de commande développé en Go pour gérer des contacts (ID, Nom, Email).  
Permet d’ajouter, lister, supprimer et mettre à jour des contacts directement depuis le terminal.  

## Fonctionnalités

- Ajouter un contact
- Lister tous les contacts
- Supprimer un contact par ID
- Mettre à jour un contact

## Installation

1. Cloner le repository :

```bash
git clone https://github.com/DuarteJules/mini_crm.git
cd mini_crm
go mod tidy
```

## Utilisation
### CLI
utiliser les commandes suivante
```bash
go run main.go add --nom "Jules" --email "jules@example.com"
go run main.go list
go run main.go update --id 1 --nom "Jules Duarte"
go run main.go delete --id 1
```





