import { useState } from 'react'
import Home from './pages/Home'
import './App.css'

function App() {
    return (
        <div className="App">
            <header className="app-header">
                <h1>Squawker</h1>
            </header>
            <main>
                <Home />
            </main>
        </div>
    )
}

export default App
