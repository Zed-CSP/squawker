import { useState, useEffect } from 'react'
import Home from './pages/Home'
import './App.css'
import GridHorizon from './components/GridHorizon'
import DigitalRain from './components/DigitalRain'

function App() {
    const [user, setUser] = useState(null)

    useEffect(() => {
        const storedUser = localStorage.getItem('user')
        if (storedUser) {
            setUser(JSON.parse(storedUser))
        }
    }, [])

    const handleLogout = () => {
        localStorage.removeItem('token')
        localStorage.removeItem('user')
        setUser(null)
    }

    return (
        <div className="App">
            <GridHorizon />
            <DigitalRain />
            <header className="app-header">
                <h1>Squawker</h1>
                {user && (
                    <div className="user-status">
                        <span>@{user.username}</span>
                        <button onClick={handleLogout} className="logout-button">Logout</button>
                    </div>
                )}
            </header>
            <main>
                <Home user={user} setUser={setUser} />
            </main>
        </div>
    )
}

export default App
