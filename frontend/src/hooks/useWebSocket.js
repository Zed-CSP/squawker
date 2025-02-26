import { useState, useEffect, useRef, useCallback } from 'react'

export function useWebSocket(url) {
    const [isConnected, setIsConnected] = useState(false)
    const [error, setError] = useState(null)
    const socketRef = useRef(null)
    const reconnectTimeoutRef = useRef(null)
    const initialRequestSentRef = useRef(false)

    const connect = useCallback(() => {
        try {
            // Close existing connection if any
            if (socketRef.current && socketRef.current.readyState !== WebSocket.CLOSED) {
                socketRef.current.close()
            }

            console.log('Attempting to connect to WebSocket at:', url)
            const socket = new WebSocket(url)
            socketRef.current = socket
            initialRequestSentRef.current = false

            socket.onopen = () => {
                console.log('WebSocket connection established')
                setIsConnected(true)
                setError(null)
                
                // Request initial messages immediately after connection
                socket.send(JSON.stringify({ type: 'get_initial_messages' }))
                initialRequestSentRef.current = true
                console.log('Sent initial messages request')
                
                // Clear any pending reconnect timeouts
                if (reconnectTimeoutRef.current) {
                    clearTimeout(reconnectTimeoutRef.current)
                    reconnectTimeoutRef.current = null
                }
            }

            socket.onclose = (event) => {
                console.log('WebSocket connection closed:', event)
                setIsConnected(false)
                initialRequestSentRef.current = false
                
                // Always attempt to reconnect
                console.log('Attempting to reconnect...')
                reconnectTimeoutRef.current = setTimeout(() => {
                    connect()
                }, 2000)
            }

            socket.onerror = (error) => {
                console.error('WebSocket error:', error)
                setError('WebSocket connection error')
            }
        } catch (err) {
            console.error('Error setting up WebSocket:', err)
            setError(err.message)
            
            // Try to reconnect after error
            reconnectTimeoutRef.current = setTimeout(() => {
                connect()
            }, 2000)
        }
    }, [url])

    useEffect(() => {
        connect()
        
        // Ping the server every 15 seconds to keep the connection alive
        const pingInterval = setInterval(() => {
            if (socketRef.current && socketRef.current.readyState === WebSocket.OPEN) {
                // Send a ping message
                socketRef.current.send(JSON.stringify({ type: 'ping' }))
                
                // If initial messages request wasn't sent, try again
                if (!initialRequestSentRef.current) {
                    socketRef.current.send(JSON.stringify({ type: 'get_initial_messages' }))
                    initialRequestSentRef.current = true
                    console.log('Retry: Sent initial messages request')
                }
            }
        }, 15000)
        
        return () => {
            clearInterval(pingInterval)
            
            if (reconnectTimeoutRef.current) {
                clearTimeout(reconnectTimeoutRef.current)
            }
            
            if (socketRef.current) {
                socketRef.current.close()
                socketRef.current = null
            }
        }
    }, [connect])

    // Add a function to manually reconnect
    const reconnect = useCallback(() => {
        console.log('Manually reconnecting WebSocket...')
        connect()
    }, [connect])

    // Add a function to manually request initial messages
    const requestInitialMessages = useCallback(() => {
        if (socketRef.current && socketRef.current.readyState === WebSocket.OPEN) {
            socketRef.current.send(JSON.stringify({ type: 'get_initial_messages' }))
            initialRequestSentRef.current = true
            console.log('Manually requested initial messages')
        }
    }, [])

    return { 
        socket: socketRef.current, 
        isConnected, 
        error,
        reconnect,
        requestInitialMessages
    }
} 