import { useState, useEffect, useRef, useCallback } from 'react'

export function useWebSocket(url) {
    const [isConnected, setIsConnected] = useState(false)
    const [error, setError] = useState(null)
    const socketRef = useRef(null)

    const connect = useCallback(() => {
        try {
            const socket = new WebSocket(url)
            socketRef.current = socket

            socket.onopen = () => {
                console.log('WebSocket connection established')
                setIsConnected(true)
                setError(null)
            }

            socket.onclose = (event) => {
                console.log('WebSocket connection closed:', event)
                setIsConnected(false)
                
                // Attempt to reconnect after a delay
                setTimeout(() => {
                    if (socketRef.current) {
                        connect()
                    }
                }, 3000)
            }

            socket.onerror = (error) => {
                console.error('WebSocket error:', error)
                setError('WebSocket connection error')
            }
        } catch (err) {
            console.error('Error setting up WebSocket:', err)
            setError(err.message)
        }
    }, [url])

    useEffect(() => {
        connect()
        
        return () => {
            if (socketRef.current) {
                socketRef.current.close()
                socketRef.current = null
            }
        }
    }, [connect])

    return { socket: socketRef.current, isConnected, error }
} 