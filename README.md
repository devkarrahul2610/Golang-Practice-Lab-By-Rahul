# Golang Practice Lab by Rahul 🚀

This branch **functions-methods-05** demonstrates the difference between **functions** and **methods** in Go, along with practical examples of attaching methods to different kinds of types.

---

## 📌 Functions vs Methods in Go

- **Function**  
  A function is a block of code that performs a task.  
  - Defined using the `func` keyword.  
  - Does **not** belong to any type (struct, interface, etc.).  
  - Can only modify what you explicitly pass to it.  

- **Method**  
  A method is a function **with a receiver**.  
  - A receiver appears before the method name and binds it to a type.  
  - Methods “belong” to that type.  
  - Pointer receivers allow methods to modify the original type directly.  

---

## 📘 Topics Covered

1. Declaring variables with `var` vs creating new types with `type`.  
   - `var` → when you want a variable of an existing type.  
   - `type` → when you want a new type (even based on an existing one), usually to add:  
     - Type safety  
     - Methods  
     - Interface implementations  

2. Functions on built-in types.  
3. Methods on custom types (`struct`, `map`, `function type`).  
4. Implementing interfaces.  
5. Type safety by creating named types.  

---

## 📂 Files & Concepts

- **main.go**  
  Entry point. Calls different demos (`Demo1` → `Demo4`).

- **functions.go (Demo1)**  
  - Example of a **normal function** (`UpperCase`).  
  - Demonstrates string conversion using the `strings` package.  

- **methods.go (MethodDemo, InterfaceDemo)**  
  - Shows **type safety** by creating new types (`Kilometer`, `miles`).  
  - Demonstrates attaching a method to a custom type (`Name.NewUpperCase`).  
  - Shows how a type can implement an interface (`fmt.Stringer`).  

- **methods1.go (Demo2)**  
  - Attaching a **method to a map-based custom type** (`marks.add`).  
  - Demonstrates reference type behavior.  
  - Error handling when duplicate keys are inserted.  

- **methods2.go (Demo3)**  
  - Methods on **function types** (`FullName`).  
  - Demonstrates anonymous functions and attaching methods to them.  

- **methods3.go (Demo4)**  
  - Methods on **structs** (`Rectangle.Area`).  
  - Classic OOP-style behavior in Go.  

---

## ▶️ How to Run

```bash
go run main.go
