package tasksManager

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Task struct représente une tâche avec ses propriétés
type Task struct {
	ID          int       `json:"id"`          // Identifiant de la tâche
	Description string    `json:"description"` // Description de la tâche
	Completed   bool      `json:"completed"`   // Statut de complétion de la tâche
	DueDate     time.Time `json:"due_date"`    // Date d'échéance de la tâche
	Subtasks    []Task    `json:"subtasks"`    // Liste des sous-tâches
}

// TaskList struct contient une liste de tâches
type TaskList struct {
	Tasks      []Task `json:"tasks"` // Slice de tâches
	nextTaskID int    // Prochain ID de tâche à utiliser
}

var tasks TaskList // Variable globale pour stocker les tâches

// TaskManager gère l'interface utilisateur en ligne de commande du gestionnaire de tâches
func TaskManager() {
	loadTasks() // Chargement des tâches depuis le fichier JSON au démarrage

	reader := bufio.NewReader(os.Stdin) // Création d'un lecteur pour lire depuis l'entrée standard

	// Affichage du menu des commandes disponibles
	fmt.Println("Task Manager CLI")
	fmt.Println("----------------")
	fmt.Println("Commands:")
	fmt.Println("1. 'add <description>' - Add a new task")
	fmt.Println("2. 'list' - List all tasks")
	fmt.Println("3. 'complete <taskID>' - Mark a task as completed")
	fmt.Println("4. 'delete <taskID>' - Delete a task")
	fmt.Println("5. 'addsubtask <parentID> <description>' - Add a subtask to a task")
	fmt.Println("6. 'listsubtasks <taskID>' - List subtasks of a task")
	fmt.Println("7. 'completesubtask <taskID> <subtaskID>' - Mark a subtask as completed")
	fmt.Println("8. 'deletesubtask <taskID> <subtaskID>' - Delete a subtask")
	fmt.Println("9. 'listcompleted' - List all completed tasks with subtasks")
	fmt.Println("10. 'listincomplete' - List all incomplete tasks with subtasks")
	fmt.Println("11. 'exit' - Exit the program")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n') // Lecture de l'entrée utilisateur
		input = strings.TrimSpace(input)    // Suppression des espaces autour de l'entrée
		args := strings.Fields(input)       // Découpage de l'entrée en mots séparés

		if len(args) == 0 {
			continue // Si aucune commande n'est entrée, recommence la boucle
		}

		command := args[0] // La première partie de l'entrée est la commande

		// Switch pour traiter différentes commandes
		switch command {
		case "add":
			if len(args) < 2 {
				fmt.Println("Please provide a description for the task.")
				continue
			}
			description := strings.Join(args[1:], " ") // La description est le reste de l'entrée après "add"
			addTask(description)                       // Appel de la fonction pour ajouter une tâche
		case "list":
			listTasksWithSubtasks(tasks.Tasks) // Appel de la fonction pour lister toutes les tâches avec sous-tâches
		case "complete":
			if len(args) < 2 {
				fmt.Println("Please provide a task ID to complete.")
				continue
			}
			taskID, err := strconv.Atoi(args[1]) // Conversion de l'ID de tâche en entier
			if err != nil {
				fmt.Println("Invalid task ID. Please provide a valid integer.")
				continue
			}
			completeTask(taskID) // Appel de la fonction pour marquer une tâche comme complétée
		case "delete":
			if len(args) < 2 {
				fmt.Println("Please provide a task ID to delete.")
				continue
			}
			taskID, err := strconv.Atoi(args[1]) // Conversion de l'ID de tâche en entier
			if err != nil {
				fmt.Println("Invalid task ID. Please provide a valid integer.")
				continue
			}
			deleteTask(taskID) // Appel de la fonction pour supprimer une tâche
		case "addsubtask":
			if len(args) < 3 {
				fmt.Println("Please provide a parent task ID and description for the subtask.")
				continue
			}
			parentID, err := strconv.Atoi(args[1]) // Conversion de l'ID de tâche parent en entier
			if err != nil {
				fmt.Println("Invalid parent task ID. Please provide a valid integer.")
				continue
			}
			description := strings.Join(args[2:], " ") // La description est le reste de l'entrée après "add subtask <parentID>"
			addSubtask(parentID, description)          // Appel de la fonction pour ajouter une sous-tâche
		case "listsubtasks":
			if len(args) < 2 {
				fmt.Println("Please provide a task ID to list subtasks.")
				continue
			}
			taskID, err := strconv.Atoi(args[1]) // Conversion de l'ID de tâche en entier
			if err != nil {
				fmt.Println("Invalid task ID. Please provide a valid integer.")
				continue
			}
			listSubtasks(taskID) // Appel de la fonction pour lister les sous-tâches d'une tâche spécifique
		case "completesubtask":
			if len(args) < 3 {
				fmt.Println("Please provide both task ID and subtask ID to complete a subtask.")
				continue
			}
			taskID, err := strconv.Atoi(args[1])    // Conversion de l'ID de tâche en entier
			subtaskID, err := strconv.Atoi(args[2]) // Conversion de l'ID de sous-tâche en entier
			if err != nil {
				fmt.Println("Invalid ID(s). Please provide valid integers.")
				continue
			}
			completeSubtask(taskID, subtaskID) // Appel de la fonction pour marquer une sous-tâche comme complétée
		case "deletesubtask":
			if len(args) < 3 {
				fmt.Println("Please provide both task ID and subtask ID to delete a subtask.")
				continue
			}
			taskID, err := strconv.Atoi(args[1])    // Conversion de l'ID de tâche en entier
			subtaskID, err := strconv.Atoi(args[2]) // Conversion de l'ID de sous-tâche en entier
			if err != nil {
				fmt.Println("Invalid ID(s). Please provide valid integers.")
				continue
			}
			deleteSubtask(taskID, subtaskID) // Appel de la fonction pour supprimer une sous-tâche
		case "listcompleted":
			listCompletedTasksWithSubtasks(tasks.Tasks) // Appel de la fonction pour lister toutes les tâches complétées avec sous-tâches
		case "listincomplete":
			listIncompleteTasksWithSubtasks(tasks.Tasks) // Appel de la fonction pour lister toutes les tâches incomplètes avec sous-tâches
		case "exit":
			saveTasks() // Sauvegarde des tâches dans le fichier JSON avant de quitter
			fmt.Println("Exiting Task Manager CLI.")
			return // Sortie de la fonction TaskManager et du programme
		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}

// loadTasks charge les tâches depuis le fichier JSON
func loadTasks() {
	file, err := os.Open("tasks.json") // Ouverture du fichier des tâches en lecture
	if err != nil {
		return // Si le fichier n'existe pas ou s'il y a une erreur, retourne sans rien faire
	}
	defer file.Close() // Fermeture du fichier à la fin de la fonction

	decoder := json.NewDecoder(file) // Création d'un décodeur JSON pour lire le fichier
	err = decoder.Decode(&tasks)     // Décodage des données JSON dans la variable tasks
	if err != nil {
		fmt.Println("Error decoding tasks:", err) // Affichage de l'erreur en cas d'échec du décodage
	}
	// Trouver le plus grand ID existant pour initialiser nextTaskID
	for _, task := range tasks.Tasks {
		if task.ID >= tasks.nextTaskID {
			tasks.nextTaskID = task.ID + 1
		}
	}
}

// saveTasks sauvegarde les tâches dans le fichier JSON
func saveTasks() {
	file, err := os.Create("tasks.json") // Création ou ouverture du fichier des tâches en écriture
	if err != nil {
		fmt.Println("Error creating tasks file:", err) // Affichage de l'erreur en cas d'échec de création du fichier
		return
	}
	defer file.Close() // Fermeture du fichier à la fin de la fonction

	encoder := json.NewEncoder(file) // Création d'un encodeur JSON pour écrire dans le fichier
	err = encoder.Encode(tasks)      // Encodage des données de tasks dans le fichier JSON
	if err != nil {
		fmt.Println("Error encoding tasks:", err) // Affichage de l'erreur en cas d'échec de l'encodage
	}
}

// addTask ajoute une nouvelle tâche avec la description fournie
func addTask(description string) {
	newTask := Task{
		ID:          tasks.nextTaskID,
		Description: description,
		Completed:   false,
		DueDate:     time.Time{}, // Date d'échéance par défaut (valeur zéro pour time.Time)
	}
	tasks.Tasks = append(tasks.Tasks, newTask) // Ajout de la nouvelle tâche à la liste des tâches
	tasks.nextTaskID++                         // Incrément de nextTaskID pour la prochaine tâche
	fmt.Println("Task added successfully.")    // Confirmation de l'ajout de la tâche
}

// listTasksWithSubtasks affiche toutes les tâches avec leurs sous-tâches
func listTasksWithSubtasks(taskList []Task) {
	for _, task := range taskList {
		fmt.Printf("[%d]%s %s\n", task.ID, getStatusSymbol(task.Completed), task.Description) // Affichage de l'ID, du statut et de la description de la tâche
		listSubtasks(task.ID)                                                                 // Affichage des sous-tâches de la tâche actuelle
	}
}

// listCompletedTasksWithSubtasks affiche toutes les tâches complétées avec leurs sous-tâches
func listCompletedTasksWithSubtasks(taskList []Task) {
	for _, task := range taskList {
		if task.Completed {
			fmt.Printf("[%d]%s %s\n", task.ID, getStatusSymbol(task.Completed), task.Description) // Affichage de l'ID, du statut et de la description de la tâche
			listSubtasks(task.ID)                                                                 // Affichage des sous-tâches de la tâche actuelle
		}
	}
}

// listIncompleteTasksWithSubtasks affiche toutes les tâches incomplètes avec leurs sous-tâches
func listIncompleteTasksWithSubtasks(taskList []Task) {
	for _, task := range taskList {
		if !task.Completed {
			fmt.Printf("[%d]%s %s\n", task.ID, getStatusSymbol(task.Completed), task.Description) // Affichage de l'ID, du statut et de la description de la tâche
			listSubtasks(task.ID)                                                                 // Affichage des sous-tâches de la tâche actuelle
		}
	}
}

// listSubtasks affiche les sous-tâches de la tâche spécifiée par son ID, indentées
func listSubtasks(taskID int) {
	for _, task := range tasks.Tasks {
		if task.ID == taskID {
			for _, subtask := range task.Subtasks {
				fmt.Printf("\t[%d]%s %s\n", subtask.ID, getStatusSymbol(subtask.Completed), subtask.Description) // Affichage indenté de l'ID, du statut et de la description de la sous-tâche
			}
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", taskID) // Message si aucune tâche ne correspond à l'ID fourni
}

// completeTask marque une tâche spécifique comme complétée en fonction de son ID
func completeTask(taskID int) {
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == taskID {
			tasks.Tasks[i].Completed = true                      // Marque la tâche comme complétée si elle correspond à l'ID fourni
			fmt.Printf("Task %d marked as completed.\n", taskID) // Confirmation de la complétion de la tâche
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", taskID) // Message si aucune tâche ne correspond à l'ID fourni
}

// deleteTask supprime une tâche spécifique de la liste en fonction de son ID
func deleteTask(taskID int) {
	found := false
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == taskID {
			// Déplacer les tâches suivantes vers l'avant pour combler le trou
			copy(tasks.Tasks[i:], tasks.Tasks[i+1:])
			tasks.Tasks = tasks.Tasks[:len(tasks.Tasks)-1]

			// Réorganiser les IDs pour refléter la suppression
			for j := i; j < len(tasks.Tasks); j++ {
				tasks.Tasks[j].ID = j + 1
			}

			fmt.Printf("Task %d deleted.\n", taskID) // Confirmation de la suppression de la tâche
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found.\n", taskID) // Message si aucune tâche ne correspond à l'ID fourni
	}
}

// addSubtask ajoute une nouvelle sous-tâche à une tâche spécifique identifiée par son ID parent
func addSubtask(parentID int, description string) {
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == parentID {
			newSubtask := Task{
				ID:          len(tasks.Tasks[i].Subtasks) + 1, // ID unique pour la nouvelle sous-tâche
				Description: description,
				Completed:   false,
				DueDate:     time.Time{}, // Date d'échéance par défaut (valeur zéro pour time.Time)
			}
			tasks.Tasks[i].Subtasks = append(tasks.Tasks[i].Subtasks, newSubtask) // Ajout de la nouvelle sous-tâche à la liste des sous-tâches de la tâche parent
			tasks.Tasks[i].Completed = false                                      // Réinitialise le statut de la tâche parent à non terminé
			fmt.Printf("Subtask added successfully to task %d.\n", parentID)      // Confirmation de l'ajout de la sous-tâche
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", parentID) // Message si aucune tâche ne correspond à l'ID fourni
}

// completeSubtask marque une sous-tâche spécifique comme complétée en fonction de l'ID de la tâche parente et de l'ID de la sous-tâche
func completeSubtask(taskID, subtaskID int) {
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == taskID {
			for j := range tasks.Tasks[i].Subtasks {
				if tasks.Tasks[i].Subtasks[j].ID == subtaskID {
					tasks.Tasks[i].Subtasks[j].Completed = true                                   // Marque la sous-tâche comme complétée si elle correspond à l'ID fourni
					fmt.Printf("Subtask %d of task %d marked as completed.\n", subtaskID, taskID) // Confirmation de la complétion de la sous-tâche

					// Vérifier si toutes les sous-tâches sont complétées
					allSubtasksCompleted := true
					for _, subtask := range tasks.Tasks[i].Subtasks {
						if !subtask.Completed {
							allSubtasksCompleted = false
							break
						}
					}
					// Si toutes les sous-tâches sont complétées, marquer la tâche principale comme complétée
					if allSubtasksCompleted {
						tasks.Tasks[i].Completed = true
					}
					return
				}
			}
			fmt.Printf("Subtask with ID %d not found for task %d.\n", subtaskID, taskID) // Message si aucune sous-tâche ne correspond à l'ID fourni
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", taskID) // Message si aucune tâche ne correspond à l'ID fourni
}

// deleteSubtask supprime une sous-tâche spécifique d'une tâche spécifique en fonction de l'ID de la tâche parente et de l'ID de la sous-tâche
func deleteSubtask(taskID, subtaskID int) {
	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == taskID {
			for j := range tasks.Tasks[i].Subtasks {
				if tasks.Tasks[i].Subtasks[j].ID == subtaskID {
					// Déplacer les sous-tâches suivantes vers l'avant pour combler le trou
					copy(tasks.Tasks[i].Subtasks[j:], tasks.Tasks[i].Subtasks[j+1:])
					tasks.Tasks[i].Subtasks = tasks.Tasks[i].Subtasks[:len(tasks.Tasks[i].Subtasks)-1]

					fmt.Printf("Subtask %d of task %d deleted.\n", subtaskID, taskID) // Confirmation de la suppression de la sous-tâche
					return
				}
			}
			fmt.Printf("Subtask with ID %d not found for task %d.\n", subtaskID, taskID) // Message si aucune sous-tâche ne correspond à l'ID fourni
			return
		}
	}
	fmt.Printf("Task with ID %d not found.\n", taskID) // Message si aucune tâche ne correspond à l'ID fourni
}

// getStatusSymbol retourne le symbole correspondant à l'état de complétion de la tâche
func getStatusSymbol(completed bool) string {
	if completed {
		return " ✓"
	}
	return " x"
}
