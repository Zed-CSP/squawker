import { useState } from 'react'
import './App.css'

function App() {
  const [message, setMessage] = useState('')
  const [error, setError] = useState('')

  const testApiConnection = async () => {
    try {
      setError('')
      setMessage('Loading...')
      
      const response = await fetch('/api/messages')
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`)
      }
      
      const data = await response.json()
      setMessage(JSON.stringify(data, null, 2))
    } catch (err) {
      setError(`Error: ${err.message}`)
      setMessage('')
    }
  }

  return (
    <div className="App">
      <h1>Squawker</h1>
      <button onClick={testApiConnection}>
        Test API Connection
      </button>
      
      {message && (
        <pre style={{ textAlign: 'left', marginTop: '20px' }}>
          {message}
        </pre>
      )}
      
      {error && (
        <div style={{ color: 'red', marginTop: '20px' }}>
          {error}
        </div>
      )}
    </div>
  )
}

export default App
