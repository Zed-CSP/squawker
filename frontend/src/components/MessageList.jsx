import { useState, useEffect } from 'react'
import './MessageList.css'

export default function MessageList() {
    const [messages, setMessages] = useState([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(null)

    useEffect(() => {
        fetchMessages()
    }, [])

    const fetchMessages = async () => {
        try {
            const response = await fetch('/api/messages')
            if (!response.ok) {
                throw new Error('Failed to fetch messages')
            }
            const data = await response.json()
            setMessages(data)
            setLoading(false)
        } catch (err) {
            setError(err.message)
            setLoading(false)
        }
    }

    if (loading) return <div className="loading">Loading messages...</div>
    if (error) return <div className="error">{error}</div>

    return (
        <div className="message-list">
            {messages.map((message) => (
                <div key={message.id} className="message-card">
                    <div className="message-header">
                        <span className="username">@{message.username}</span>
                        <span className="date">
                            {new Date(message.createdAt).toLocaleDateString()}
                        </span>
                    </div>
                    <p className="content">{message.content}</p>
                </div>
            ))}
        </div>
    )
} 