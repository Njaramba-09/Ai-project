# Beginner’s Toolkit for Go (Golang)  
*Moringa AI Capstone Project*

---

## 1. Title & Objective
**Title:** Beginner’s Toolkit for Go (Golang)  

**Objective:**  
This toolkit is designed to help beginners learn the basics of the Go programming language using Generative AI prompts. The goal is to guide learners through installing Go, creating a simple web service that serves static HTML/CSS content, and documenting the learning process with AI assistance. By the end, learners will have confidence in running a Go backend and connecting it with a frontend.

---

## 2. Quick Summary of the Technology
**What is Go?**  
Go (Golang) is an open-source programming language created by Google. It is known for simplicity, speed, and strong performance.  

**Where is it used?**  
- Web servers and APIs  
- Cloud applications and distributed systems  
- Backend systems for high-performance applications  

**Real-world Example:**  
Uber uses Go to power high-performance backend services that handle millions of requests per second.

---

## 3. System Requirements
To follow this toolkit, you need:  
- **Operating System:** Windows, macOS, or Linux  
- **Go Compiler:** Official Go installer from [https://go.dev/dl](https://go.dev/dl)  
- **Text Editor:** VS Code (recommended)  
- **Terminal:** Command Prompt, PowerShell, or macOS/Linux Terminal  
- **Browser:** For viewing the frontend  

---

## 4. Installation & Setup Instructions

### **Install Go**
Follow the official instructions for your OS (Windows/macOS/Linux). Verify installation:

```bash
go version
```

### **Project Structure**
```
bree-AI/
│
├─ Backend/
│   ├─ main.go
│   └─ go_toolkit.md
│
└─ Frontend/
    ├─ index.html
    └─ style.css
```

### **Run the Web Service**
1. Open a terminal in the `Backend` folder.
2. Run the Go server:

```bash
go run main.go
```

3. Open a browser and navigate to [http://localhost:8080](http://localhost:8080) to see your frontend served by Go.

---

## 5. Minimal Working Example

### **Backend (main.go)**
```go
package main

import (
    "net/http"
    "log"
)

func main() {
    // Serve static files in Frontend folder under /static/
    fs := http.FileServer(http.Dir("../Frontend"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Serve index.html at root URL
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "../Frontend/index.html")
    })

    log.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### **Frontend (index.html)**
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Go Beginner Toolkit</title>
    <link rel="stylesheet" href="/static/style.css">
</head>
<body>
    <header>
        <h1>Welcome to the Go Beginner Toolkit!</h1>
        <p>Learn the basics of Go programming through fun facts and examples.</p>
    </header>

    <section class="cards-container">
        <div class="card">
            <h2>Fact 1</h2>
            <p>Go was created at Google in 2007 to simplify software development.</p>
        </div>
        <div class="card">
            <h2>Fact 2</h2>
            <p>Go is statically typed and compiled, making it fast and efficient.</p>
        </div>
        <div class="card">
            <h2>Fact 3</h2>
            <p>Popular companies using Go include Uber, Dropbox, and SoundCloud.</p>
        </div>
        <div class="card">
            <h2>Fact 4</h2>
            <p>Go has built-in support for concurrency with goroutines.</p>
        </div>
    </section>

    <footer>
        <p>Beginner’s Toolkit © 2025</p>
    </footer>
</body>
</html>
```

### **CSS (style.css)**
```css
body {
    font-family: Arial, sans-serif;
    background: #f0f8ff;
    margin: 0;
    padding: 0;
}

header {
    background: #4caf50;
    color: white;
    text-align: center;
    padding: 30px 10px;
}

.cards-container {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    margin: 20px;
    gap: 20px;
}

.card {
    background: white;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0,0,0,0.2);
    padding: 20px;
    width: 250px;
    transition: transform 0.2s;
}

.card:hover {
    transform: translateY(-5px);
}

.card h2 {
    color: #4caf50;
}

footer {
    background: #333;
    color: white;
    text-align: center;
    padding: 10px 0;
    margin-top: 40px;
}
```

---

## 6. AI Prompt Journal (Learning Experiences)

**Prompt 1:** "Explain Go programming language in simple terms for a beginner."  
**AI Summary:** Go is a modern, fast, and simple language suitable for web servers, APIs, and cloud apps.  
**Reflection:** Helped me understand Go’s purpose and uses.

**Prompt 2:** "Give me a minimal Go example that serves a static HTML page."  
**AI Summary:** Provided `main.go` to serve HTML/CSS content with Go.  
**Reflection:** Learned how to serve frontend assets from a backend.

**Prompt 3:** "How do I handle CSS and other static files in Go?"  
**AI Summary:** Explained `http.FileServer` and `http.StripPrefix`.  
**Reflection:** Resolved CSS not loading issue.

**Prompt 4:** "List common beginner errors when serving static files in Go."  
**AI Summary:** Errors like wrong paths or missing `StripPrefix`.  
**Reflection:** Learned to troubleshoot backend/frontend integration.

---

## 7. Common Issues & Fixes

- **CSS not loading** – Wrong path.  
  **Fix:** Use `/static/style.css` in HTML and `http.StripPrefix` in Go.

- **404 on page** – Backend not serving the route correctly.  
  **Fix:** Use `http.HandleFunc("/", ...)` for the root page.

- **Go runtime errors** – typos like stray `r` in `func main() { r ... }`  
  **Fix:** Remove stray characters and check syntax.

---

## 8. References

- [Official Go Documentation](https://go.dev/doc/)  
- [Go by Example](https://gobyexample.com/)  
- [W3Schools Go Tutorial](https://www.w3schools.com/go/)  
- [DigitalOcean Go Articles](https://www.digitalocean.com/community/tags/go)
