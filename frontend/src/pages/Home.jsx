import { useState, useEffect } from 'react'
import { useWebSocket } from '../hooks/useWebSocket'
import CreateMessage from '../components/features/messaging/CreateMessage'
import AuthModal from '../components/features/modal/Auth/AuthModal'
import MessageCard from '../components/features/messaging/MessageCard'
import ConnectionStatus from '../components/ConnectionStatus'
import MessageFeed from '../components/features/messaging/MessageFeed'
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
        
    const { 
        socket, 
        isConnected, 
        error: wsError, 
        reconnect, 
        requestInitialMessages 
    } = useWebSocket(wsUrl)

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

    const handleCreateMessage = (content) => {
        if (!user) {
            setPendingMessage(content)
            setShowAuthModal(true)
            return
        }
        
        if (socket && isConnected) {
            const message = {
                type: 'new_message',
                payload: {
                    content,
                    username: user.username
                }
            }
            
            socket.send(JSON.stringify(message))
        } else {
            setError('Not connected to server. Please try again.')
        }
    }

    const handleAuthSuccess = (userData) => {
        setUser(userData)
        setShowAuthModal(false)
        
        if (pendingMessage) {
            handleCreateMessage(pendingMessage)
            setPendingMessage(null)
        }
    }

    const handleConnectionAction = () => {
        if (isConnected) {
            requestInitialMessages()
        } else {
            reconnect()
        }
    }

    return (
        <div className="home">
            <ConnectionStatus 
                isConnected={isConnected} 
                onClick={handleConnectionAction} 
            />
            
            <CreateMessage onSubmit={handleCreateMessage} />
            
            <MessageFeed 
                messages={messages}
                error={error}
                onRetry={() => requestInitialMessages()}
            />

            <AuthModal 
                isOpen={showAuthModal}
                onClose={() => {
                    setShowAuthModal(false)
                    setPendingMessage(null)
                }}
                onSuccess={handleAuthSuccess}
            />
        </div>
    )
} 