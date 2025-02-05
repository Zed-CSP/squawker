import { useState, useEffect } from 'react'
import { useWebSocket } from '../hooks/useWebSocket'
import CreateMessage from '../components/CreateMessage'
import AuthModal from '../components/AuthModal'
import './Home.css'

export default function Home({ user, setUser }) {
    const [messages, setMessages] = useState([])
    const [error, setError] = useState(null)
    const [showAuthModal, setShowAuthModal] = useState(false)
    const [pendingMessage, setPendingMessage] = useState(null)
    
    // Update WebSocket URL to match Vite's proxy
    const wsUrl = import.meta.env.DEV 
        ? `ws://localhost:8080/ws`  // Development
        : `ws://${window.location.hostname}/ws`  // Production
        
    const { socket, isConnected, error: wsError } = useWebSocket(wsUrl)

    useEffect(() => {
        if (socket) {
            socket.onmessage = (event) => {
                const data = JSON.parse(event.data)
                
                switch (data.type) {
                    case 'new_message':
                        setMessages(prev => [data.payload, ...prev])
                        break
                    case 'initial_messages':
                        setMessages(data.payload)
                        break
                    case 'error':
                        setError(data.payload)
                        break
                }
            }
        }
    }, [socket])

    // Add logging for debugging
    useEffect(() => {
        console.log('WebSocket connection status:', isConnected)
        if (wsError) {
            console.error('WebSocket error:', wsError)
        }
    }, [isConnected, wsError])

    const handleCreateMessage = async (content) => {
        const token = localStorage.getItem('token')
        console.log('Token being sent:', token) // Debug log
        
        if (!token || !user) {
            setPendingMessage(content)
            setShowAuthModal(true)
            return
        }

        try {
            const response = await fetch('/api/messages', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({ content })
            })

            if (response.status === 401) {
                // Try to verify token first
                const verifyResponse = await fetch('/api/auth/verify', {
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                })

                if (verifyResponse.status === 401) {
                    // Only clear if token is actually invalid
                    localStorage.removeItem('token')
                    localStorage.removeItem('user')
                    setUser(null)
                    setShowAuthModal(true)
                    setPendingMessage(content)
                }
                return
            }

            if (!response.ok) {
                throw new Error('Failed to create message')
            }

            // Clear pending message on success
            setPendingMessage(null)
        } catch (err) {
            setError(err.message)
        }
    }

    // Add useEffect to monitor showAuthModal changes
    useEffect(() => {
        console.log('showAuthModal changed to:', showAuthModal)
    }, [showAuthModal])

    const handleAuthSuccess = async (data) => {
        localStorage.setItem('token', data.token)
        localStorage.setItem('user', JSON.stringify(data.user))
        setUser(data.user)
        
        if (pendingMessage) {
            await handleCreateMessage(pendingMessage)
            setPendingMessage(null)
        }
        
        setShowAuthModal(false)
    }

    return (
        <div className="home">
            <div className="connection-status" style={{
                padding: '10px',
                margin: '10px 0',
                borderRadius: '4px',
                backgroundColor: isConnected ? 'rgba(0, 255, 0, 0.1)' : 'rgba(255, 0, 0, 0.1)',
                border: `1px solid ${isConnected ? 'var(--neon-blue)' : 'var(--neon-pink)'}`,
                color: isConnected ? 'var(--neon-blue)' : 'var(--neon-pink)',
                textAlign: 'center',
                fontFamily: 'monospace'
            }}>
                {isConnected ? (
                    <span>LIVE CONNECTION ESTABLISHED ▇ </span>
                ) : (
                    <span>CONNECTING TO SERVER... ▇</span>
                )}
            </div>
            
            <div className="feed-container">
                <CreateMessage onSubmit={handleCreateMessage} />
                
                <div className="messages-container">
                    {error && (
                        <div className="error">
                            {error}
                            <button onClick={() => {
                                // Implement the logic to try again
                            }}>Try Again</button>
                        </div>
                    )}
                    
                    {!error && (
                        <>
                            {messages.length === 0 ? (
                                <div className="no-messages">
                                    No messages yet. Be the first to post!
                                </div>
                            ) : (
                                <div className="message-list">
                                    {messages.map((message) => (
                                        <div key={message.id} className="message-card">
                                            <div className="message-header">
                                                <div className="user-info">
                                                    <span className="username">@{message.username}</span>
                                                    <span className="time">
                                                        {new Date(message.createdAt).toLocaleTimeString()}
                                                    </span>
                                                </div>
                                                <span className="date">
                                                    {new Date(message.createdAt).toLocaleDateString()}
                                                </span>
                                            </div>
                                            <p className="content">{message.content}</p>
                                        </div>
                                    ))}
                                </div>
                            )}
                        </>
                    )}
                </div>
            </div>

            {/* Add debug element */}
            <div style={{ display: 'none' }}>
                Modal should be: {showAuthModal ? 'visible' : 'hidden'}
            </div>

            <AuthModal 
                isOpen={showAuthModal}
                onClose={() => {
                    console.log('Modal close triggered') // Debug log
                    setShowAuthModal(false)
                    setPendingMessage(null)
                }}
                onSuccess={handleAuthSuccess}
            />
        </div>
    )
} 