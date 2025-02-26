import React from 'react'
import './MessageList.css'

function formatDate(dateString) {
    try {
        const date = new Date(dateString)
        return date.toLocaleString()
    } catch (error) {
        console.error('Error formatting date:', error, dateString)
        return 'Just now'
    }
}

export default function MessageList({ messages }) {
    return (
        <div className="message-list">
            {messages.length === 0 ? (
                <div className="no-messages">No messages yet. Be the first to post!</div>
            ) : (
                messages.map((message) => (
                    <div key={message.id} className="message">
                        <div className="message-header">
                            <span className="username">@{message.username}</span>
                            <span className="timestamp">{formatDate(message.createdAt)}</span>
                        </div>
                        <div className="message-content">{message.content}</div>
                    </div>
                ))
            )}
        </div>
    )
} 