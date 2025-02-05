import { useState } from 'react'
import './AuthModal.css'

export default function AuthModal({ isOpen, onClose, onSuccess }) {
    const [isLogin, setIsLogin] = useState(true)
    const [formData, setFormData] = useState({
        username: '',
        password: '',
        email: ''
    })
    const [error, setError] = useState(null)
    const [isLoading, setIsLoading] = useState(false)

    const handleSubmit = async (e) => {
        e.preventDefault()
        setError(null)
        setIsLoading(true)

        const endpoint = isLogin ? '/api/auth/login' : '/api/auth/register'
        const requestData = isLogin ? {
            username: formData.username,
            password: formData.password
        } : formData

        try {
            // Log the full request details
            console.log('Making request to:', window.location.origin + endpoint)
            console.log('Request data:', requestData)

            const response = await fetch(endpoint, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json'
                },
                body: JSON.stringify(requestData),
            })

            console.log('Response status:', response.status)
            console.log('Response headers:', Object.fromEntries([...response.headers]))

            // Get the raw response text
            const text = await response.text()
            console.log('Raw response:', text)

            // Try to parse as JSON
            let data
            try {
                data = JSON.parse(text)
                console.log('Parsed response data:', data)
            } catch (err) {
                console.error('Failed to parse response as JSON:', err)
                console.error('Raw response was:', text)
                throw new Error(`Server returned invalid JSON: ${text.substring(0, 100)}...`)
            }

            if (!response.ok) {
                throw new Error(data.error || 'Authentication failed')
            }

            // Store token
            localStorage.setItem('token', response.data.token)
            localStorage.setItem('user', JSON.stringify(data.user))

            onSuccess(data)
            onClose()
        } catch (err) {
            console.error('Full error details:', err)
            setError(err.message)
        } finally {
            setIsLoading(false)
        }
    }

    if (!isOpen) return null

    return (
        <div className="auth-modal-overlay">
            <div className="auth-modal">
                <button className="close-button" onClick={onClose}>×</button>
                <h2>{isLogin ? 'Login' : 'Sign Up'}</h2>
                
                {error && <div className="error-message">{error}</div>}
                
                <form onSubmit={handleSubmit}>
                    <input
                        type="text"
                        placeholder="Username"
                        value={formData.username}
                        onChange={e => setFormData({...formData, username: e.target.value})}
                        required
                        disabled={isLoading}
                    />
                    
                    {!isLogin && (
                        <input
                            type="email"
                            placeholder="Email"
                            value={formData.email}
                            onChange={e => setFormData({...formData, email: e.target.value})}
                            required
                            disabled={isLoading}
                        />
                    )}
                    
                    <input
                        type="password"
                        placeholder="Password"
                        value={formData.password}
                        onChange={e => setFormData({...formData, password: e.target.value})}
                        required
                        disabled={isLoading}
                        minLength={6}
                    />
                    
                    <button 
                        type="submit" 
                        className="submit-button"
                        disabled={isLoading}
                    >
                        {isLoading ? 'Please wait...' : (isLogin ? 'Login' : 'Sign Up')}
                    </button>
                </form>
                
                <button 
                    className="toggle-auth-mode"
                    onClick={() => {
                        setIsLogin(!isLogin)
                        setFormData({
                            username: '',
                            password: '',
                            email: ''
                        })
                        setError(null)
                    }}
                    disabled={isLoading}
                >
                    {isLogin ? 'Need an account? Sign up' : 'Already have an account? Login'}
                </button>
            </div>
        </div>
    )
} 