# Conway's Game of Life

Conway's Game of Life is a zero-player cellular automaton where cells evolve over discrete time steps according to simple rules. This implementation is built in Go and runs in a terminal using `tview` and `tcell`.

---

## 📹 Demo

A demo video showcasing the simulation is available in the repository.



https://github.com/user-attachments/assets/2aa43c42-4f27-4190-a589-c2da36b8f6bb




---

## 🚀 Features

- Terminal-based UI using **tview** and **tcell**
- Supports different patterns including **random** and **glider gun**
- Uses **Run Length Encoding (RLE)** format for defining patterns
- Real-time updates with **smooth animations**
- Highly configurable grid size

---

## 🛠️ Installation & Usage

### Prerequisites

- Go (1.18+ recommended)

### Clone the Repository

```sh
git clone "https://github.com/Tapasvi48/Convoy-Game-of-Life"

```

### Run the Simulation

```sh
go run main.go
```

By default, the game runs with the **glider gun** pattern on a 15x60 grid.

---

## 🕹️ Controls

- **Run Automatically**: The simulation updates every 20ms.
- **Exit**: Press `Ctrl + C` or close the terminal.

---

## 📝 Rules of the Game

1. Any **alive** cell with **two or three** live neighbors **survives**.
2. Any **alive** cell with **fewer than two** live neighbors **dies** (underpopulation).
3. Any **alive** cell with **more than three** live neighbors **dies** (overpopulation).
4. Any **dead** cell with **exactly three** live neighbors **becomes alive** (reproduction).

---

## 🏗️ Code Structure

- `Game` struct handles the grid and logic.
- `RenderGrid()` updates the UI.
- `Update()` applies the game rules to the grid.
- `LoadRLE()` allows loading patterns in RLE format.
- `main.go` initializes and runs the game.

---

## 📦 Dependencies

- [gdamore/tcell](https://github.com/gdamore/tcell) - Terminal UI Library
- [rivo/tview](https://github.com/rivo/tview) - Rich Terminal UI Framework

Install dependencies using:

```sh
go mod tidy
```

---

## 📜 License

This project is licensed under the MIT License.

---

## 👨‍💻 Author

Developed by **Tapasvi Arora**.

Feel free to contribute, raise issues, or suggest improvements!

Happy Coding! 🚀
