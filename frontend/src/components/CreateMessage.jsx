import { useState } from 'react'
import './CreateMessage.css'

export default function CreateMessage({ onSubmit }) {
    const [message, setMessage] = useState('')

    const handleSubmit = async (e) => {
        e.preventDefault()
        console.log('CreateMessage submit triggered') // Debug log
        
        if (message.trim()) {
            await onSubmit(message)
            setMessage('')
        }
    }

    return (
        <form onSubmit={handleSubmit} className="create-message">
            <input
                type="text"
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                placeholder="Type your message..."
                className="message-input"
            />
            <button type="submit" className="send-button">Send</button>
        </form>
    )
} 